/*
The MIT License (MIT)

Copyright (c) 2017 isaac dawson

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

package gcd

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdapi"
	"github.com/wirepair/gcd/gcdmessage"
	"golang.org/x/net/websocket"
	"io"
	"log"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

// Defines the 'tab' or target for this chrome instance, can be multiple and background processes
// are included (not just visual tabs)
type TargetInfo struct {
	Description          string `json:"description"`
	DevtoolsFrontendUrl  string `json:"devtoolsFrontendUrl"`
	FaviconUrl           string `json:"faviconUrl"`
	Id                   string `json:"id"`
	Title                string `json:"title"`
	Type                 string `json:"type"`
	Url                  string `json:"url"`
	WebSocketDebuggerUrl string `json:"webSocketDebuggerUrl"`
}

// Our Chrome Target (Tab/Process). Messages are returned to callers via non-buffered channels. Helpfully,
// the remote debugger service uses id's so we can correlate which request should match which response.
// We use a map to store the id of the request which contains a reference to a gcdmessage.Message that holds the
// reply channel for the ChromeTarget to return the response to.
// Events are handled by mapping the method name to a function which takes a target and byte output.
// For now, callers will need to unmarshall the types themselves.
type ChromeTarget struct {
	replyLock       sync.RWMutex                           // lock for dispatching responses
	replyDispatcher map[int64]chan *gcdmessage.Message     // Replies to synch methods using a non-buffered channel
	eventLock       sync.RWMutex                           // lock for dispatching events
	eventDispatcher map[string]func(*ChromeTarget, []byte) // calls the function when events match the subscribed method
	conn            *websocket.Conn                        // the connection to the chrome debugger service for this tab/process

	// Chrome Debugger Domains
	Accessibility     *gcdapi.Accessibility
	Animation         *gcdapi.Animation
	ApplicationCache  *gcdapi.ApplicationCache // application cache API
	Browser           *gcdapi.Browser
	CacheStorage      *gcdapi.CacheStorage
	Console           *gcdapi.Console           // console API
	CSS               *gcdapi.CSS               // CSS API
	Database          *gcdapi.Database          // Database API
	Debugger          *gcdapi.Debugger          // JS Debugger API
	DeviceOrientation *gcdapi.DeviceOrientation // Device Orientation API
	DOMDebugger       *gcdapi.DOMDebugger       // DOM Debugger API
	DOM               *gcdapi.DOM               // DOM API
	DOMSnapshot       *gcdapi.DOMSnapshot
	DOMStorage        *gcdapi.DOMStorage // DOM Storage API
	Emulation         *gcdapi.Emulation
	HeapProfiler      *gcdapi.HeapProfiler // HeapProfiler API
	IndexedDB         *gcdapi.IndexedDB    // IndexedDB API
	Input             *gcdapi.Input        // Why am i doing this, it's obvious what they are, I quit.
	Inspector         *gcdapi.Inspector
	IO                *gcdapi.IO
	LayerTree         *gcdapi.LayerTree
	Log               *gcdapi.Log
	Memory            *gcdapi.Memory
	Network           *gcdapi.Network
	Overlay           *gcdapi.Overlay
	Page              *gcdapi.Page
	Profiler          *gcdapi.Profiler
	Runtime           *gcdapi.Runtime
	Schema            *gcdapi.Schema
	Security          *gcdapi.Security
	ServiceWorker     *gcdapi.ServiceWorker
	Storage           *gcdapi.Storage
	SystemInfo        *gcdapi.SystemInfo
	TargetApi         *gcdapi.Target // buh name collision
	Tracing           *gcdapi.Tracing
	Tethering         *gcdapi.Tethering

	Target      *TargetInfo              // The target information see, TargetInfo
	sendCh      chan *gcdmessage.Message // The channel used for API components to send back to use
	doneCh      chan struct{}            // we be donez.
	apiTimeout  time.Duration            // A customizable timeout for waiting on Chrome to respond to us
	sendId      int64                    // An Id which is atomically incremented per request.
	debugEvents bool                     // flag for spitting out event data as a string which we have not subscribed to.
	debug       bool                     // flag for printing internal debugging messages
	stopped     bool                     // we are/have shutdown
}

// Creates a new Chrome Target by connecting to the service given the URL taken from initial connection.
func openChromeTarget(addr string, target *TargetInfo) (*ChromeTarget, error) {
	conn, err := wsConnection(addr, target.WebSocketDebuggerUrl)
	if err != nil {
		return nil, err
	}
	replier := make(map[int64]chan *gcdmessage.Message)
	var replyLock sync.RWMutex
	var eventLock sync.RWMutex
	eventer := make(map[string]func(*ChromeTarget, []byte))
	sendCh := make(chan *gcdmessage.Message)
	doneCh := make(chan struct{})
	chromeTarget := &ChromeTarget{conn: conn, eventLock: eventLock, replyLock: replyLock, Target: target, sendCh: sendCh, replyDispatcher: replier, eventDispatcher: eventer, doneCh: doneCh, sendId: 0}
	chromeTarget.apiTimeout = 120 * time.Second // default 120 seconds to wait for chrome to respond to us
	chromeTarget.Init()
	chromeTarget.listen()
	return chromeTarget, nil
}

// Initialize all api objects
func (c *ChromeTarget) Init() {
	c.Accessibility = gcdapi.NewAccessibility(c)
	c.Animation = gcdapi.NewAnimation(c)
	c.ApplicationCache = gcdapi.NewApplicationCache(c)
	c.Browser = gcdapi.NewBrowser(c)
	c.CacheStorage = gcdapi.NewCacheStorage(c)
	c.Console = gcdapi.NewConsole(c)
	c.CSS = gcdapi.NewCSS(c)
	c.Database = gcdapi.NewDatabase(c)
	c.Debugger = gcdapi.NewDebugger(c)
	c.DeviceOrientation = gcdapi.NewDeviceOrientation(c)
	c.DOMDebugger = gcdapi.NewDOMDebugger(c)
	c.DOM = gcdapi.NewDOM(c)
	c.DOMSnapshot = gcdapi.NewDOMSnapshot(c)
	c.DOMStorage = gcdapi.NewDOMStorage(c)
	c.Emulation = gcdapi.NewEmulation(c)
	c.HeapProfiler = gcdapi.NewHeapProfiler(c)
	c.IndexedDB = gcdapi.NewIndexedDB(c)
	c.Input = gcdapi.NewInput(c)
	c.Inspector = gcdapi.NewInspector(c)
	c.IO = gcdapi.NewIO(c)
	c.LayerTree = gcdapi.NewLayerTree(c)
	c.Memory = gcdapi.NewMemory(c)
	c.Log = gcdapi.NewLog(c)
	c.Network = gcdapi.NewNetwork(c)
	c.Overlay = gcdapi.NewOverlay(c)
	c.Page = gcdapi.NewPage(c)
	c.Profiler = gcdapi.NewProfiler(c)
	c.Runtime = gcdapi.NewRuntime(c)
	c.Schema = gcdapi.NewSchema(c)
	c.Security = gcdapi.NewSecurity(c)
	c.SystemInfo = gcdapi.NewSystemInfo(c)
	c.ServiceWorker = gcdapi.NewServiceWorker(c)
	c.TargetApi = gcdapi.NewTarget(c)
	c.Tracing = gcdapi.NewTracing(c)
	c.Tethering = gcdapi.NewTethering(c)
}

// clean up this target
func (c *ChromeTarget) shutdown() {
	if c.stopped == true {
		return
	}
	c.stopped = true

	// close websocket read/write goroutines
	close(c.doneCh)

	// close websocket connection
	c.conn.Close()
}

// A timeout for how long we should wait before giving up gcdmessages.
// In the highly unusuable (but it has occurred) event that chrome
// does not respond to one of our messages, we should be able to return
// from gcdmessage functions.
func (c *ChromeTarget) SetApiTimeout(timeout time.Duration) {
	c.apiTimeout = timeout
}

// Used by gcdmessage.SendCustomReturn and gcdmessage.SendDefaultRequest
// to timeout an API call if chrome hasn't responded to us in apiTimeout
// time.
func (c *ChromeTarget) GetApiTimeout() time.Duration {
	return c.apiTimeout
}

/// Subscribes Events, you must know the method name, such as Page.loadFiredEvent, and bind a function
// which takes a ChromeTarget (us) and the raw JSON byte data for that event.
func (c *ChromeTarget) Subscribe(method string, callback func(*ChromeTarget, []byte)) {
	c.eventLock.Lock()
	c.eventDispatcher[method] = callback
	c.eventLock.Unlock()
}

// Unsubscribes the handler for no longer recieving events.
func (c *ChromeTarget) Unsubscribe(method string) {
	c.eventLock.Lock()
	delete(c.eventDispatcher, method)
	c.eventLock.Unlock()
}

// Whether to print out raw JSON event data when not Subscribed.
func (c *ChromeTarget) DebugEvents(debug bool) {
	c.debugEvents = debug
}

func (c *ChromeTarget) Debug(debug bool) {
	c.debug = debug
}

// Listens for API components wanting to send, and recv'ing data from the Chrome Debugger Service
func (c *ChromeTarget) listen() {
	go c.listenRead()
	go c.listenWrite()
}

// Listens for API components wishing to send requests to the Chrome Debugger Service
func (c *ChromeTarget) listenWrite() {
	for {
		select {
		// send message to the browser debugger client
		case msg := <-c.sendCh:
			c.replyLock.Lock()
			c.replyDispatcher[msg.Id] = msg.ReplyCh
			c.replyLock.Unlock()

			c.debugf("%d sending to chrome. %s\n", msg.Id, msg.Data)

			err := websocket.Message.Send(c.conn, string(msg.Data))
			if err != nil {
				c.debugf("error sending message: %s\n", err)
				return
			}
		// receive done from listenRead
		case <-c.doneCh:
			return
		}
	}
}

// Listens for responses coming in from the Chrome Debugger Service.
func (c *ChromeTarget) listenRead() {
	for {
		select {
		// receive done from listenWrite
		case <-c.doneCh:
			return
		// read data from websocket connection
		default:
			var msg string
			err := websocket.Message.Receive(c.conn, &msg)
			if err == io.EOF {
				c.debugf("error io.EOF in websocket read")
				return
			} else if err != nil {
				c.debugf("error in ws read: %s\n", err)
			} else {
				go c.dispatchResponse([]byte(msg))
			}
		}
	}
}

type responseHeader struct {
	Id     int64  `json:"id"`
	Method string `json:"method"`
}

// dispatchResponse takes in the json message and extracts
// the id or method fields to dispatch either responses or events
// to listeners.
func (c *ChromeTarget) dispatchResponse(msg []byte) {
	f := &responseHeader{}
	err := json.Unmarshal(msg, f)
	if err != nil {
		c.debugf("error reading response data from chrome target: %v\n", err)
	}

	c.replyLock.Lock()

	if r, ok := c.replyDispatcher[f.Id]; ok {
		delete(c.replyDispatcher, f.Id)
		c.replyLock.Unlock()
		c.debugf("%d dispatching\n", f.Id)
		c.dispatchWithTimeout(r, f.Id, msg)
		return
	}
	c.replyLock.Unlock()

	c.checkTargetDisconnected(f.Method)

	c.eventLock.RLock()
	if r, ok := c.eventDispatcher[f.Method]; ok {
		c.eventLock.RUnlock()
		c.debugf("dispatching %s event: %s\n", f.Method, string(msg))
		go r(c, msg)
		return

	}
	c.eventLock.RUnlock()

	if c.debugEvents == true {
		log.Printf("\n\nno event recv bound for: %s\n", f.Method)
		log.Printf("data: %s\n\n", string(msg))
	}
}

func (c *ChromeTarget) dispatchWithTimeout(r chan<- *gcdmessage.Message, id int64, msg []byte) {
	timeout := time.NewTimer(c.GetApiTimeout())
	defer timeout.Stop()

	select {
	case r <- &gcdmessage.Message{Id: id, Data: msg}:
		timeout.Stop()
	case <-timeout.C:
		c.debugf("timed out dispatching request id: %d of msg: %s\n", id, msg)
		close(r)
		return
	}
	return
}

// check target detached/crashed
// close any replier channels that are open
// dispatch detached / crashed event as usual
func (c *ChromeTarget) checkTargetDisconnected(method string) {
	switch method {
	case "Inspector.targetCrashed", "Inspector.detached":
		c.replyLock.Lock()
		for _, replyCh := range c.replyDispatcher {
			close(replyCh)
		}
		// empty out the dispatcher
		c.replyDispatcher = make(map[int64]chan *gcdmessage.Message)
		c.replyLock.Unlock()
	}
}

// Connects to the tab/process for sending/recv'ing debug events
func wsConnection(addr, url string) (*websocket.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}

	config, errConfig := websocket.NewConfig(url, "http://localhost")
	if errConfig != nil {
		return nil, errConfig
	}
	client, errWS := websocket.NewClient(config, conn)
	if errWS != nil {
		return nil, errWS
	}
	return client, nil
}

// gcdmessage.ChromeTargeter interface methods

// Increments the Id so we can synchronize our request/responses internally
func (c *ChromeTarget) GetId() int64 {
	return atomic.AddInt64(&c.sendId, 1)
}

// The channel used for API components to send back to use
func (c *ChromeTarget) GetSendCh() chan *gcdmessage.Message {
	return c.sendCh
}

// The channel used to signal any pending SendDefaultRequest and SendCustomReturn
// that we are exiting so we don't block goroutines from exiting.
func (c *ChromeTarget) GetDoneCh() chan struct{} {
	return c.doneCh
}

func (c *ChromeTarget) debugf(format string, args ...interface{}) {
	if c.debug {
		log.Printf(format, args...)
	}
}
