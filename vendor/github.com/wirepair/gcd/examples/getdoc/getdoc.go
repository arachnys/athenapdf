package main

import (
	"flag"
	"github.com/wirepair/gcd"
	"log"
	"runtime"
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
	flag.Parse()
	debugger := gcd.NewChromeDebugger()
	debugger.StartProcess(path, dir, port)
	defer debugger.ExitProcess()
	target, err := debugger.NewTab()
	if err != nil {
		log.Fatalf("error getting new tab: %s\n", err)
	}
	dom := target.DOM
	r, err := dom.GetDocument(-1, true)
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	log.Printf("NodeID: %d Node Name: %s, URL: %s\n", r.NodeId, r.NodeName, r.DocumentURL)
}
