package runner

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"net/url"
	"sync"
	"time"

	"github.com/arachnys/athenapdf/pkg/proto"

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

	JSPlugins []string

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
			return err
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
	if r.Target == nil {
		return nil, ErrInvalidTarget
	}

	if r.Timeout == 0 {
		r.Timeout = defaultTimeout
	}

	if err := setOptions(r.Target, req.GetOptions()); err != nil {
		return nil, err
	}

	if err := setHeaders(r.Target, req.GetHeaders()); err != nil {
		return nil, err
	}

	if err := setCookies(r.Target, req.GetCookies()); err != nil {
		return nil, err
	}

	fid, err := r.Target.Page.Navigate(req.GetUri(), "", "")
	if err != nil {
		return nil, err
	}

	// Add page load event listener
	pageReady := make(chan bool, 1)
	r.Target.Subscribe("Page.loadEventFired", func(_ *gcd.ChromeTarget, _ []byte) {
		pageReady <- true
	})

	// Detect errors in page load
	pageFailed := make(chan string, 1)
	r.Target.Subscribe("Network.loadingFailed", func(_ *gcd.ChromeTarget, b []byte) {
		var v gcdapi.NetworkLoadingFailedEvent
		if err := json.Unmarshal(b, &v); err != nil {
			if r.Debug {
				log.Println(err)
			}
			return
		}
		if v.Params.RequestId == fid {
			pageFailed <- v.Params.ErrorText
		}
	})

	// Load scripts concurrently
	loadedPlugins := make(chan string, 1)
	go func() {
		p, err := GetJsPlugins(r.JSPlugins...)
		if err != nil {
			log.Fatalln(err)
		}
		loadedPlugins <- p
	}()

	select {
	case <-time.After(r.Timeout):
		return nil, fmt.Errorf("timeout waiting for the page to load")
	case errorText := <-pageFailed:
		if errorText == "" {
			errorText = "unknown error, content-type might not be supported (e.g. `application/octet-stream`)"
		}
		return nil, fmt.Errorf("failed to load the page: %s", errorText)
	case <-pageReady:
	}

	// Execute scripts
	evaluateParams := gcdapi.RuntimeEvaluateParams{Expression: <-loadedPlugins}
	if _, _, err := r.Target.Runtime.EvaluateWithParams(&evaluateParams); err != nil {
		return nil, err
	}

	// Generate PDF, and convert output to bytes from base64 string
	dimensions := req.GetDimensions()
	options := req.GetOptions()
	pdfParams := gcdapi.PagePrintToPDFParams{
		MarginBottom:    dimensions.GetMarginBottom(),
		MarginLeft:      dimensions.GetMarginLeft(),
		MarginRight:     dimensions.GetMarginRight(),
		MarginTop:       dimensions.GetMarginTop(),
		PaperHeight:     dimensions.GetPageHeight(),
		PaperWidth:      dimensions.GetPageWidth(),
		Scale:           options["scale"].GetDoubleValue(),
		Landscape:       dimensions.GetOrientation() == proto.Conversion_Dimensions_LANDSCAPE,
		PrintBackground: !options["no_background"].GetBoolValue(),
	}
	base64String, err := r.Target.Page.PrintToPDFWithParams(&pdfParams)
	if err != nil {
		return nil, err
	}

	b, err := base64.StdEncoding.DecodeString(base64String)
	if err != nil {
		return nil, err
	}

	return b, nil
}
