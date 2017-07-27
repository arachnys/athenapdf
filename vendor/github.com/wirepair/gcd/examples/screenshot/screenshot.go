package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"github.com/wirepair/gcd"
	"log"
	"net/url"
	"os"
	"runtime"
	"sync"
	"time"
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
		targets[i].Subscribe("Page.loadEventFired", PageLoaded)
		page.Navigate(urls[i], "")
	}
	wg.Wait()
	for i := 0; i < numTabs; i++ {
		TakeScreenShot(targets[i])
	}
}

func PageLoaded(target *gcd.ChromeTarget, event []byte) {
	wg.Done()
}

func TakeScreenShot(target *gcd.ChromeTarget) {
	dom := target.DOM
	page := target.Page
	doc, err := dom.GetDocument(-1, true)
	if err != nil {
		fmt.Errorf("error getting doc: %s\n", err)
		return
	}
	debugger.ActivateTab(target)
	time.Sleep(1 * time.Second) // give it a sec to paint
	u, urlErr := url.Parse(doc.DocumentURL)
	if urlErr != nil {
		fmt.Errorf("error parsing url: %s\n", urlErr)
		return
	}
	fmt.Printf("Taking screen shot of: %s\n", u.Host)
	img, errCap := page.CaptureScreenshot("png", 0, false)
	if errCap != nil {
		fmt.Errorf("error getting doc: %s\n", errCap)
		return
	}
	imgBytes, errDecode := base64.StdEncoding.DecodeString(img)
	if errDecode != nil {
		fmt.Errorf("error decoding image: %s\n", errDecode)
		return
	}
	f, errFile := os.Create(u.Host + ".png")
	defer f.Close()
	if errFile != nil {
		fmt.Errorf("error creating image file: %s\n", errFile)
		return
	}
	f.Write(imgBytes)
}
