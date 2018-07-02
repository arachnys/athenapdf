package runner

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/arachnys/athenapdf/pkg/proto"
	"github.com/arachnys/athenapdf/pkg/runner/plugin"

	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
)

var (
	defaultTimeout = time.Minute
)

type ExitFunc func() error

type Runner struct {
	Debug   bool
	DryRun  bool
	Timeout time.Duration

	Server *url.URL
	Proxy  *url.URL

	Plugins struct {
		Builtin []string
		Custom  []string
	}

	Target *gcd.ChromeTarget

	exited   bool
	exitedMu sync.RWMutex
}

func (r *Runner) AutoTarget() (ExitFunc, error) {
	var exitFunc ExitFunc = func() error { return nil }

	c, clientExit, err := startCDP(r.Server, r.Proxy)
	if err != nil {
		return exitFunc, err
	}

	t, targetExit, err := startTarget(c)
	if err != nil {
		return exitFunc, err
	}

	exitFunc = func() error {
		r.exitedMu.Lock()
		defer r.exitedMu.Unlock()

		if r.exited {
			return nil
		}

		var exit ExitFunc
		if r.Server == nil {
			exit = clientExit
		} else {
			exit = targetExit
		}

		if err := exit(); err != nil {
			return errors.WithStack(err)
		}

		r.exited = true
		return nil
	}

	t.Debug(r.Debug)
	t.DebugEvents(r.Debug)

	r.Target = t

	return exitFunc, nil
}

func (r *Runner) Convert(req *proto.Conversion) ([]byte, error) {
	var (
		requestID string
		err       error
	)

	if r.Target == nil {
		return nil, ErrInvalidTarget
	}

	if r.Timeout == 0 {
		r.Timeout = defaultTimeout
	}

	options := req.GetOptions()

	if err := setOptions(r.Target, options); err != nil {
		return nil, err
	}

	if err := setHeaders(r.Target, req.GetHeaders()); err != nil {
		return nil, err
	}

	if err := setCookies(r.Target, req.GetCookies()); err != nil {
		return nil, err
	}

	if err := setInsecure(r.Target, options["insecure"].GetBoolValue(), r.Debug); err != nil {
		return nil, err
	}

	pageCh := make(chan struct {
		success bool
		err     error
	}, 1)

	// Add page load event listener
	r.Target.Subscribe("Page.loadEventFired", func(_ *gcd.ChromeTarget, _ []byte) {
		pageCh <- struct {
			success bool
			err     error
		}{true, nil}
	})

	// Detect errors in page load
	r.Target.Subscribe("Network.loadingFailed", func(_ *gcd.ChromeTarget, b []byte) {
		var v gcdapi.NetworkLoadingFailedEvent
		if err := json.Unmarshal(b, &v); err != nil {
			if r.Debug {
				log.Println(err)
			}
			return
		}
		if (requestID != "" && v.Params.RequestId == requestID) ||
			(requestID == "" && v.Params.Type == "Document") {
			pageCh <- struct {
				success bool
				err     error
			}{false, errors.New(v.Params.ErrorText)}
		}
	})

	// Load scripts concurrently
	loadedScripts := make(chan string, 1)
	go func() {
		pr, err := plugin.Get(r.Plugins.Builtin, r.Plugins.Custom)
		if err != nil {
			log.Fatalln(err)
		}

		for _, p := range pr {
			var buf bytes.Buffer
			if _, err := buf.ReadFrom(p); err != nil {
				log.Fatalln(err)
			}
			loadedScripts <- buf.String()
		}

		close(loadedScripts)
	}()

	requestID, _, _, err = r.Target.Page.Navigate(req.GetUri(), "", "", "")
	if err != nil {
		return nil, errors.WithStack(err)
	}

	select {
	case <-time.After(r.Timeout):
		return nil, errors.New("timed out waiting for the page to load")
	case page := <-pageCh:
		if !page.success {
			if page.err == nil {
				page.err = errors.New("unknown error, content-type might not be supported (e.g. `application/octet-stream`)")
			}
			return nil, errors.WithMessage(page.err, "failed to load the page")
		}
	}

	// Execute scripts
	for script := range loadedScripts {
		evaluateParams := gcdapi.RuntimeEvaluateParams{
			Expression:   script,
			AwaitPromise: true,
		}
		if _, _, err := r.Target.Runtime.EvaluateWithParams(&evaluateParams); err != nil {
			return nil, errors.WithStack(err)
		}
	}

	// Generate PDF, and convert output to bytes from base64 string
	dimensions := req.GetDimensions()
	pdfParams := gcdapi.PagePrintToPDFParams{
		MarginBottom:      dimensions.GetMarginBottom(),
		MarginLeft:        dimensions.GetMarginLeft(),
		MarginRight:       dimensions.GetMarginRight(),
		MarginTop:         dimensions.GetMarginTop(),
		PaperHeight:       dimensions.GetPageHeight(),
		PaperWidth:        dimensions.GetPageWidth(),
		Scale:             options["scale"].GetDoubleValue(),
		Landscape:         dimensions.GetOrientation() == proto.Conversion_Dimensions_LANDSCAPE,
		PrintBackground:   !options["no_background"].GetBoolValue(),
		PreferCSSPageSize: options["preferCssPageSize"].GetBoolValue(),
	}
	base64String, err := r.Target.Page.PrintToPDFWithParams(&pdfParams)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	b, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return b, nil
}
