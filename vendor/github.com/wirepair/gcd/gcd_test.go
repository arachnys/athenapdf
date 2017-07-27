package gcd

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"runtime"
	"testing"
	"time"

	"github.com/wirepair/gcd/gcdapi"
)

var (
	debugger                 *Gcd
	testListener             net.Listener
	testSkipNetworkIntercept bool
	testPath                 string
	testDir                  string
	testPort                 string
	testServerAddr           string
)

func init() {
	switch runtime.GOOS {
	case "windows":
		flag.StringVar(&testPath, "chrome", "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "path to chrome")
		flag.StringVar(&testDir, "dir", "C:\\temp\\", "user directory")
	case "darwin":
		flag.StringVar(&testPath, "chrome", "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "path to chrome")
		flag.StringVar(&testDir, "dir", "/tmp/", "user directory")
	case "linux":
		flag.StringVar(&testPath, "chrome", "/usr/bin/chromium-browser", "path to chrome")
		flag.StringVar(&testDir, "dir", "/tmp/", "user directory")
	}
	flag.StringVar(&testPort, "port", "9222", "Debugger port")

	// TODO: remove this once mainline chrome supports it.
	flag.BoolVar(&testSkipNetworkIntercept, "intercept", true, "set to false to test network intercept, will fail if browser does not support yet.")

}

func TestMain(m *testing.M) {
	flag.Parse()
	testServer()

	ret := m.Run()
	testCleanUp()
	os.Exit(ret)
}

func testCleanUp() {
	testListener.Close()
}

func TestGetPages(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	targets, _ := debugger.GetTargets()
	if len(targets) <= 0 {
		t.Fatalf("invalid number of targets, got: %d\n", len(targets))
	}
	t.Logf("page: %s\n", targets[0].Target.Url)
}

func TestEnv(t *testing.T) {
	var ok bool
	debugger = NewChromeDebugger()
	debugger.AddEnvironmentVars([]string{"hello=youze", "zoinks=scoob"})
	debugger.StartProcess(testPath, testRandomTempDir(t), testRandomPort(t))
	defer debugger.ExitProcess()

	t.Logf("%#v\n", debugger.chromeCmd.Env)
	for _, v := range debugger.chromeCmd.Env {
		if v == "hello=youze" {
			ok = true
		}
	}
	if !ok {
		t.Fatalf("error finding our environment vars in chrome process")
	}
}

func TestProcessKilled(t *testing.T) {
	testDefaultStartup(t)
	doneCh := make(chan struct{})
	shutdown := time.NewTimer(time.Second * 4)
	timeout := time.NewTimer(time.Second * 10)
	terminatedHandler := func(reason string) {
		t.Logf("reason: %s\n", reason)
		doneCh <- struct{}{}
	}
	debugger.SetTerminationHandler(terminatedHandler)
	for {
		select {
		case <-doneCh:
			goto DONE
		case <-shutdown.C:
			debugger.ExitProcess()
		case <-timeout.C:
			t.Fatalf("timed out waiting for termination")
		}
	}
DONE:
}

func TestTargetCrashed(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	doneCh := make(chan struct{})
	go testTimeoutListener(t, doneCh, 5, "timed out waiting for crashed to be handled")

	targetCrashedFn := func(targ *ChromeTarget, payload []byte) {
		t.Logf("reason: %s\n", string(payload))
		close(doneCh)
	}

	tab, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error creating new tab")
	}

	tab.Subscribe("Inspector.targetCrashed", targetCrashedFn)
	_, err = tab.Page.Navigate("chrome://crash", "", "typed")
	if err == nil {
		t.Fatalf("Navigation should have failed")
	}

	<-doneCh
}

func TestEvents(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	console := target.Console

	doneCh := make(chan struct{}, 1)
	target.Subscribe("Console.messageAdded", func(target *ChromeTarget, v []byte) {
		target.Unsubscribe("Console.messageAdded")
		msg := &gcdapi.ConsoleMessageAddedEvent{}
		err := json.Unmarshal(v, msg)
		if err != nil {
			t.Fatalf("error unmarshalling event data: %v\n", err)
		}
		t.Logf("METHOD: %s\n", msg.Method)
		eventData := msg.Params.Message
		t.Logf("Got event: %v\n", eventData)
		close(doneCh)
	})

	_, err = console.Enable()
	if err != nil {
		t.Fatalf("error sending enable: %s\n", err)
	}

	if _, err := target.Page.Navigate(testServerAddr+"console_log.html", "", "typed"); err != nil {
		t.Fatalf("error attempting to navigate: %s\n", err)
	}

	go testTimeoutListener(t, doneCh, 5, "console message")

	<-doneCh
}

func TestEvaluate(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()
	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}

	doneCh := make(chan struct{}, 1)
	target.Subscribe("Runtime.executionContextCreated", func(target *ChromeTarget, v []byte) {
		//target.Unsubscribe("Console.messageAdded")
		msg := &gcdapi.RuntimeExecutionContextCreatedEvent{}
		err := json.Unmarshal(v, msg)
		if err != nil {
			t.Fatalf("error unmarshalling event data: %v\n", err)
		}

		if msg.Params.Context.Origin != testServerAddr[:len(testServerAddr)-1] {
			return
		}
		scriptSource := "document.location.href"
		objectGroup := "gcdtest"
		awaitPromise := false
		includeCommandLineAPI := true
		contextId := msg.Params.Context.Id
		silent := true
		returnByValue := false
		generatePreview := true
		userGestures := true
		rro, exception, err := target.Runtime.Evaluate(scriptSource, objectGroup, includeCommandLineAPI, silent, contextId, returnByValue, generatePreview, userGestures, awaitPromise)
		if err != nil {
			t.Fatalf("error evaulating: %s %s\n", err, exception)
		}

		if val, ok := rro.Value.(string); ok {
			if val != testServerAddr {
				t.Fatalf("invalid location returned expected %s got %s\n", testServerAddr, val)
			}
		} else {
			t.Fatalf("error rro.Value was not a string")
		}
		close(doneCh)
	})
	target.Runtime.Enable()
	target.Page.Navigate(testServerAddr, "", "typed")
	<-doneCh
}

func TestSimpleReturn(t *testing.T) {
	var ret bool
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	network := target.Network
	if _, err := network.Enable(-1, -1); err != nil {
		t.Fatalf("error enabling network")
	}
	ret, err = network.CanClearBrowserCache()
	if err != nil {
		t.Fatalf("error getting response to clearing browser cache: %s\n", err)
	}
	if !ret {
		t.Fatalf("we should have got true for can clear browser cache\n")
	}
}

func TestSimpleReturnWithParams(t *testing.T) {
	var ret bool
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	network := target.Network

	networkParams := &gcdapi.NetworkEnableParams{
		MaxTotalBufferSize:    -1,
		MaxResourceBufferSize: -1,
	}

	if _, err := network.EnableWithParams(networkParams); err != nil {
		t.Fatalf("error enabling network")
	}
	ret, err = network.CanClearBrowserCache()
	if err != nil {
		t.Fatalf("error getting response to clearing browser cache: %s\n", err)
	}
	if !ret {
		t.Fatalf("we should have got true for can clear browser cache\n")
	}
}

// tests getting a complex object back from inside a fired event that we subscribed to.
func TestComplexReturn(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	doneCh := make(chan struct{}, 1)
	go testTimeoutListener(t, doneCh, 7, "waiting for page load to get cookies")
	target, err := debugger.NewTab()

	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}
	if _, err := target.Network.Enable(-1, -1); err != nil {
		t.Fatalf("error enabling network %s\n", err)
	}

	if _, err := target.Page.Enable(); err != nil {
		t.Fatalf("error enabling page: %s\n", err)
	}

	target.Subscribe("Page.loadEventFired", func(target *ChromeTarget, payload []byte) {
		var ok bool
		t.Logf("page load event fired\n")
		cookies, err := target.Network.GetCookies([]string{testServerAddr})
		if err != nil {
			t.Fatalf("error getting cookies!")
		}
		for _, v := range cookies {
			t.Logf("got cookies: %#v\n", v)
			if v.Name == "HEYA" {
				ok = true
				break
			}
		}
		if !ok {
			t.Fatalf("error finding our cookie value!")
		}
		close(doneCh)
	})

	_, err = target.Page.Navigate(testServerAddr+"cookie.html", "", "typed")
	if err != nil {
		t.Fatalf("error navigating to cookie page: %s\n", err)
	}

	t.Logf("waiting for loadEventFired")
	<-doneCh
}

func TestConnectToInstance(t *testing.T) {
	testDefaultStartup(t)
	defer debugger.ExitProcess()

	doneCh := make(chan struct{})

	go testTimeoutListener(t, doneCh, 15, "timed out waiting for remote connection")
	go func() {
		remoteDebugger := NewChromeDebugger()
		remoteDebugger.ConnectToInstance(debugger.host, debugger.port)

		_, err := remoteDebugger.NewTab()
		if err != nil {
			t.Fatalf("error creating new tab")
		}

		targets, error := remoteDebugger.GetTargets()
		if error != nil {
			t.Fatalf("cannot get targets: %s \n", error)
		}
		if len(targets) <= 0 {
			t.Fatalf("invalid number of targets, got: %d\n", len(targets))
		}
		for _, target := range targets {
			t.Logf("page: %s\n", target.Target.Url)
		}
		close(doneCh)
	}()
	<-doneCh
}

func TestLocalExtension(t *testing.T) {
	testExtensionStartup(t)
	debugger.ConnectToInstance(debugger.host, debugger.port)
	defer debugger.ExitProcess()

	doneCh := make(chan struct{})

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error creating new tab")
	}

	if _, err := target.Page.Enable(); err != nil {
		t.Fatalf("error enabling page: %s\n", err)
	}

	target.Subscribe("Page.loadEventFired", func(target *ChromeTarget, payload []byte) {
		t.Logf("page load event fired\n")
		close(doneCh)
	})

	if _, err := target.Network.Enable(-1, -1); err != nil {
		t.Fatalf("error enabling network: %s\n", err)
	}

	params := &gcdapi.PageNavigateParams{Url: "http://www.google.com"}
	_, err = target.Page.NavigateWithParams(params)
	if err != nil {
		t.Fatalf("error navigating: %s\n", err)
	}

	go testTimeoutListener(t, doneCh, 15, "timed out waiting for remote connection")
	<-doneCh
}

func TestNetworkIntercept(t *testing.T) {
	if testSkipNetworkIntercept {
		return
	}

	testDefaultStartup(t)
	defer debugger.ExitProcess()

	doneCh := make(chan struct{})

	target, err := debugger.NewTab()
	if err != nil {
		t.Fatalf("error getting new tab: %s\n", err)
	}

	//target.Debug(true)
	//target.DebugEvents(true)

	go testTimeoutListener(t, doneCh, 5, "timed out waiting for requestIntercepted")
	network := target.Network

	networkParams := &gcdapi.NetworkEnableParams{
		MaxTotalBufferSize:    -1,
		MaxResourceBufferSize: -1,
	}

	if _, err := network.EnableWithParams(networkParams); err != nil {
		t.Fatalf("error enabling network")
	}

	if _, err := network.SetRequestInterceptionEnabled(true); err != nil {
		t.Fatalf("unable to set interception enabled: %s\n", err)
	}

	target.Subscribe("Network.requestIntercepted", func(target *ChromeTarget, payload []byte) {
		fmt.Printf("requestIntercepted event fired %s\n", string(payload))
		t.Logf("requestIntercepted event fired %s\n", string(payload))
		close(doneCh)
	})

	params := &gcdapi.PageNavigateParams{Url: "http://www.example.com"}
	_, err = target.Page.NavigateWithParams(params)
	if err != nil {
		t.Fatalf("error navigating: %s\n", err)
	}

	<-doneCh
}

// UTILITY FUNCTIONS
func testExtensionStartup(t *testing.T) {
	debugger = NewChromeDebugger()
	sep := string(os.PathSeparator)

	path, err := os.Getwd()
	if err != nil {
		t.Fatalf("error getting working directory: %s\n", err)
	}

	extensionPath := "--load-extension=" + path + sep + "testdata" + sep + "extension" + sep
	debugger.AddFlags([]string{extensionPath})
	debugger.StartProcess(testPath, testRandomTempDir(t), testRandomPort(t))
}

func testDefaultStartup(t *testing.T) {
	debugger = NewChromeDebugger()
	debugger.StartProcess(testPath, testRandomTempDir(t), testRandomPort(t))
}

func testServer() {
	testListener, _ = net.Listen("tcp", ":0")
	_, testServerPort, _ := net.SplitHostPort(testListener.Addr().String())
	testServerAddr = fmt.Sprintf("http://localhost:%s/", testServerPort)
	go http.Serve(testListener, http.FileServer(http.Dir("testdata/")))
}

func testTimeoutListener(t *testing.T, closeCh chan struct{}, seconds time.Duration, message string) {
	timeout := time.NewTimer(seconds * time.Second)
	for {
		select {
		case <-closeCh:
			timeout.Stop()
			return
		case <-timeout.C:
			close(closeCh)
			t.Fatalf("timed out waiting for %s", message)
			return
		}
	}
}

func testRandomPort(t *testing.T) string {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		t.Fatal(err)
	}
	_, randPort, _ := net.SplitHostPort(l.Addr().String())
	l.Close()
	return randPort
}

func testRandomTempDir(t *testing.T) string {
	dir, err := ioutil.TempDir(testDir, "gcd")
	if err != nil {
		t.Errorf("error creating temp dir: %s\n", err)
	}
	return dir
}
