package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"runtime"
	"sync"
	"time"

	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
)

const (
	numTabs = 5
)

var debugger *gcd.Gcd
var wg sync.WaitGroup

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
	var err error
	urls := []string{"http://www.google.com", "http://www.veracode.com", "http://www.microsoft.com", "http://bbc.co.uk", "http://www.reddit.com/r/golang"}

	flag.Parse()

	debugger = gcd.NewChromeDebugger()
	debugger.StartProcess(path, dir, port)
	defer debugger.ExitProcess()

	targets := make([]*gcd.ChromeTarget, numTabs)

	for i := 0; i < numTabs; i++ {
		wg.Add(1)
		targets[i], err = debugger.NewTab()
		if err != nil {
			log.Fatalf("error getting targets")
		}
		page := targets[i].Page
		page.Enable()
		targets[i].Subscribe("Page.loadEventFired", pageLoaded)
		// navigate
		navigateParams := &gcdapi.PageNavigateParams{Url: urls[i]}
		_, _, _, err := page.NavigateWithParams(navigateParams)
		if err != nil {
			log.Fatalf("error: %s\n", err)
		}
	}
	wg.Wait()
	for i := 0; i < numTabs; i++ {
		takeScreenShot(targets[i])
	}
}

func pageLoaded(target *gcd.ChromeTarget, event []byte) {
	target.Unsubscribe("Page.loadEventFired")
	wg.Done()
}

func takeScreenShot(target *gcd.ChromeTarget) {
	dom := target.DOM
	page := target.Page
	doc, err := dom.GetDocument(-1, true)
	if err != nil {
		fmt.Printf("error getting doc: %s\n", err)
		return
	}

	debugger.ActivateTab(target)
	time.Sleep(1 * time.Second) // give it a sec to paint
	u, urlErr := url.Parse(doc.DocumentURL)
	if urlErr != nil {
		fmt.Printf("error parsing url: %s\n", urlErr)
		return
	}

	fmt.Printf("Taking screen shot of: %s\n", u.Host)
	screenShotParams := &gcdapi.PageCaptureScreenshotParams{Format: "png", FromSurface: true}
	img, errCap := page.CaptureScreenshotWithParams(screenShotParams)
	if errCap != nil {
		fmt.Printf("error getting doc: %s\n", errCap)
		return
	}

	imgBytes, errDecode := base64.StdEncoding.DecodeString(img)
	if errDecode != nil {
		fmt.Printf("error decoding image: %s\n", errDecode)
		return
	}

	f, errFile := os.Create(u.Host + ".png")
	defer f.Close()
	if errFile != nil {
		fmt.Printf("error creating image file: %s\n", errFile)
		return
	}
	f.Write(imgBytes)
	debugger.CloseTab(target)
}
