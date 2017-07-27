/*
Example of using DebugEvents(true) to listen to various events being emitted by
the Remote Chrome Debugger
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/wirepair/gcd"
	"github.com/wirepair/gcd/gcdapi"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"time"
)

var (
	testListener   net.Listener
	testPath       string
	testDir        string
	testPort       string
	testServerAddr string
)

var testStartupFlags = []string{"--disable-new-tab-first-run", "--no-first-run", "--disable-popup-blocking"}

func init() {
	flag.StringVar(&testPath, "chrome", "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "path to Xvfb")
	flag.StringVar(&testDir, "dir", "C:\\temp\\", "user directory")
	flag.StringVar(&testPort, "port", "9222", "Debugger port")
}

func main() {
	navigatedCh := make(chan struct{})
	debugger := startGcd()
	defer debugger.ExitProcess()

	target := startTarget(debugger)

	// subscribe to page loaded event
	target.Subscribe("Page.loadEventFired", func(targ *gcd.ChromeTarget, payload []byte) {
		navigatedCh <- struct{}{}
	})
	// navigate
	_, err := target.Page.Navigate(testServerAddr + "top.html")
	if err != nil {
		log.Fatalf("error: %s\n", err)
	}

	// wait for navigation to finish
	<-navigatedCh
	// get the document node
	doc, err := target.DOM.GetDocument()
	if err != nil {
		log.Fatal(err)
	}

	// request child nodes, this will come in as DOM.setChildNode events
	for i := 0; i < 3; i++ {
		log.Printf("requesting child nodes...")
		_, err = target.DOM.RequestChildNodes(doc.NodeId, -1)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(1 * time.Second)
	}

	target.Subscribe("DOM.setChildNodes", func(targ *gcd.ChromeTarget, payload []byte) {
		setNodes := &gcdapi.DOMSetChildNodesEvent{}
		err := json.Unmarshal(payload, setNodes)
		if err != nil {
			log.Fatalf("error decoding setChildNodes")
		}
		for _, x := range setNodes.Params.Nodes {
			if x.ContentDocument != nil {
				checkContentDocument(targ, x)
			}
			for _, v := range x.Children {
				if v.ContentDocument != nil {
					checkContentDocument(targ, v)
				}
			}
		}

	})

	// wait for redirect
	//time.Sleep(5 * time.Second)

	// get iframe node id
	iframe, err := target.DOM.QuerySelector(doc.NodeId, "#iframe")
	if err != nil {
		log.Fatalf("error looking for frame")
	}

	log.Printf("iframe %d\n", iframe)
	time.Sleep(10 * time.Second)
	debugger.ExitProcess()
	os.RemoveAll(testDir) // remove new profile junk
}

func checkContentDocument(targ *gcd.ChromeTarget, v *gcdapi.DOMNode) {
	if v.ContentDocument != nil {
		iframeDocId := v.ContentDocument.NodeId
		targ.DOM.RequestChildNodes(iframeDocId, -1)
		//nodes, _ := targ.DOM.QuerySelectorAll(iframeDocId, "div")
		log.Printf("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!got iframe nodes.\n")
	} else {
		log.Printf("got non iframe\n")
	}
}

func startGcd() *gcd.Gcd {
	testDir = testRandomDir()
	testPort = testRandomPort()
	testServer() // start test web server
	debugger := gcd.NewChromeDebugger()
	debugger.AddFlags(testStartupFlags)
	debugger.StartProcess(testPath, testDir, testPort)
	return debugger
}

func startTarget(debugger *gcd.Gcd) *gcd.ChromeTarget {
	target, err := debugger.NewTab()
	if err != nil {
		log.Fatalf("error getting new tab: %s\n", err)
	}
	target.DebugEvents(true)
	target.DOM.Enable()
	target.Console.Enable()
	target.Page.Enable()
	//target.Debugger.Enable()
	return target

}

func testServer() {
	testListener, _ = net.Listen("tcp", ":0")
	_, testServerPort, _ := net.SplitHostPort(testListener.Addr().String())
	testServerAddr = fmt.Sprintf("http://localhost:%s/", testServerPort)
	go http.Serve(testListener, http.FileServer(http.Dir("webdata/")))
}

func testRandomPort() string {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatal(err)
	}
	_, randPort, _ := net.SplitHostPort(l.Addr().String())
	l.Close()
	return randPort
}

func testRandomDir() string {
	dir, err := ioutil.TempDir(testDir, "autogcd")
	if err != nil {
		log.Fatalf("error getting temp dir: %s\n", err)
	}
	return dir
}
