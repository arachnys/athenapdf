package main

import (
	"flag"
	"log"
	"runtime"
	"sync"

	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
)

var path string
var dir string
var port string

func init() {
	switch runtime.GOOS {
	case "windows":
		flag.StringVar(&path, "chrome", "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "path to chrome")
		flag.StringVar(&dir, "dir", "C:\\temp\\", "user directory")
	case "darwin":
		flag.StringVar(&path, "chrome", "/Applications/Google Chrome.app/Contents/MacOS/Google Chrome", "path to chrome")
		flag.StringVar(&dir, "dir", "/tmp/", "user directory")
	case "linux":
		flag.StringVar(&path, "chrome", "/usr/bin/chromium-browser", "path to chrome")
		flag.StringVar(&dir, "dir", "/tmp/", "user directory")
	}

	flag.StringVar(&port, "port", "9222", "Debugger port")
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	flag.Parse()
	debugger := gcd.NewChromeDebugger()

	// start process, specify a tmp profile path so we get a fresh profiled browser
	// set port 9222 as the debug port
	debugger.StartProcess(path, dir, port)
	defer debugger.ExitProcess() // exit when done

	target, err := debugger.NewTab()
	if err != nil {
		log.Fatalf("error opening new tab: %s\n", err)
	}

	//subscribe to page load
	target.Subscribe("Page.loadEventFired", func(targ *gcd.ChromeTarget, v []byte) {
		doc, err := target.DOM.GetDocument(-1, true)
		if err == nil {
			log.Printf("%s\n", doc.DocumentURL)
		}
		wg.Done() // page loaded, we can exit now
		// if you wanted to inspect the full response data, you could do that here
	})

	// get the Page API and enable it
	if _, err := target.Page.Enable(); err != nil {
		log.Fatalf("error getting page: %s\n", err)
	}

	navigateParams := &gcdapi.PageNavigateParams{Url: "http://www.veracode.com"}
	ret, _, _, err := target.Page.NavigateWithParams(navigateParams) // navigate
	if err != nil {
		log.Fatalf("Error navigating: %s\n", err)
	}

	log.Printf("ret: %#v\n", ret)
	wg.Wait() // wait for page load
	debugger.CloseTab(target)
}
