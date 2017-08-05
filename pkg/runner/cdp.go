package runner

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"log"
	"net/url"

	"github.com/arachnys/athenapdf/pkg/proto"

	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
)

const tempDir = "athenapdf-runner"

var defaultFlags = []string{
	"--allow-running-insecure-content",
	"--disable-extensions",
	"--disable-gpu",
	"--disable-new-tab-first-run",
	"--disable-notifications",
	"--headless",
	"--ignore-certificate-errors",
	"--no-default-browser-check",
	"--no-first-run",
	"--no-sandbox",
	"--no-startup-window",
	"--no-zygote",
}

var ErrInvalidTarget = errors.New("invalid / uninitialised target")

func startCDP(server *url.URL, proxy *url.URL) (*gcd.Gcd, ExitFunc, error) {
	var exitFunc ExitFunc = func() error { return nil }

	// Get a random port to avoid conflicting instances
	randomPort, err := getRandomPort()
	if err != nil {
		return nil, exitFunc, err
	}

	// Create a random directory for user data
	randomDir, err := ioutil.TempDir("", tempDir)
	if err != nil {
		return nil, exitFunc, errors.WithStack(err)
	}

	client := gcd.NewChromeDebugger()

	if server == nil {
		// Create a new CDP process if no instance specified
		client.AddFlags(defaultFlags)
		if proxy != nil {
			client.AddFlags([]string{"--proxy-server=" + proxy.String()})
		}
		client.StartProcess(getChromePath(), randomDir, randomPort)
		exitFunc = func() error { return client.ExitProcess() }
	} else {
		// Connect to an existing instance if specified
		client.ConnectToInstance(server.Hostname(), server.Port())
	}

	return client, exitFunc, nil
}

func startTarget(client *gcd.Gcd) (*gcd.ChromeTarget, ExitFunc, error) {
	var exitFunc ExitFunc = func() error { return nil }

	t, err := client.NewTab()
	if err != nil {
		return nil, exitFunc, errors.WithStack(err)
	}

	t.CSS.Enable()
	t.DOM.Enable()
	t.Log.Enable()
	t.Network.Enable(-1, -1)
	t.Page.Enable()
	t.Runtime.Enable()
	t.Security.Enable()

	exitFunc = func() error { return client.CloseTab(t) }

	return t, exitFunc, nil
}

func setOptions(t *gcd.ChromeTarget, options map[string]*proto.Option) error {
	if t == nil {
		return ErrInvalidTarget
	}

	if _, err := t.Network.SetUserAgentOverride(options["user_agent"].GetStringValue()); err != nil {
		return errors.WithStack(err)
	}

	if _, err := t.Emulation.SetEmulatedMedia(options["media_type"].GetStringValue()); err != nil {
		return errors.WithStack(err)
	}

	if _, err := t.Network.SetCacheDisabled(options["no_cache"].GetBoolValue()); err != nil {
		return errors.WithStack(err)
	}

	if _, err := t.Emulation.SetScriptExecutionDisabled(options["no_javascript"].GetBoolValue()); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func setHeaders(t *gcd.ChromeTarget, headers []*proto.Header) error {
	if t == nil {
		return ErrInvalidTarget
	}

	h := make(map[string]interface{}, len(headers))
	for _, header := range headers {
		h[header.GetName()] = header.GetValue()
	}

	if _, err := t.Network.SetExtraHTTPHeaders(h); err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func setCookies(t *gcd.ChromeTarget, cookies []*proto.Cookie) error {
	if t == nil {
		return ErrInvalidTarget
	}

	for _, cookie := range cookies {
		c := gcdapi.NetworkSetCookieParams{
			Name:  cookie.GetName(),
			Value: cookie.GetValue(),
			Url:   cookie.GetUrl(),
		}
		if _, err := t.Network.SetCookieWithParams(&c); err != nil {
			return errors.WithStack(err)
		}
	}

	return nil
}

func setInsecure(t *gcd.ChromeTarget, insecure bool, debug bool) error {
	if !insecure {
		return nil
	}

	if _, err := t.Security.SetOverrideCertificateErrors(true); err != nil {
		return errors.WithStack(err)
	}

	t.Subscribe("Security.certificateError", func(_ *gcd.ChromeTarget, b []byte) {
		var v gcdapi.SecurityCertificateErrorEvent
		if err := json.Unmarshal(b, &v); err != nil {
			if debug {
				log.Println(err)
			}
			return
		}
		t.Security.HandleCertificateError(v.Params.EventId, "continue")
	})

	return nil
}
