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
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var GCDVERSION = "1.2017.7.8.0"

// When we get an error reading the body from the debugger api endpoint
type GcdBodyReadErr struct {
	Message string
}

func (g *GcdBodyReadErr) Error() string {
	return "error reading response body: " + g.Message
}

// Failure to unmarshal the JSON response from debugger API
type GcdDecodingErr struct {
	Message string
}

func (g *GcdDecodingErr) Error() string {
	return "error decoding inspectable page: " + g.Message
}

type TerminatedHandler func(reason string)

// The Google Chrome Debugger
type Gcd struct {
	timeout           time.Duration // how much time to wait for debugger port to open up
	chromeProcess     *os.Process
	chromeCmd         *exec.Cmd
	terminatedHandler TerminatedHandler
	port              string
	host              string
	addr              string
	readyCh           chan struct{}
	apiEndpoint       string
	flags             []string
	env               []string
}

// Give it a friendly name.
func NewChromeDebugger() *Gcd {
	c := &Gcd{}
	c.timeout = 15
	c.host = "localhost"
	c.readyCh = make(chan struct{})
	c.terminatedHandler = nil
	c.flags = make([]string, 0)
	c.env = make([]string, 0)
	return c
}

// Pass a handler to be notified when the chrome process exits.
func (c *Gcd) SetTerminationHandler(handler TerminatedHandler) {
	c.terminatedHandler = handler
}

// Set the timeout for how long we should wait for debug port to become available.
func (c *Gcd) SetTimeout(timeout time.Duration) {
	c.timeout = timeout
}

// Allows caller to add additional startup flags to the chrome process
func (c *Gcd) AddFlags(flags []string) {
	c.flags = append(c.flags, flags...)
}

// Add environment variables for the chrome process, useful for Xvfb etc.
func (c *Gcd) AddEnvironmentVars(vars []string) {
	c.env = append(c.env, vars...)
}

// Starts the process
// exePath - the path to the executable
// userDir - the user directory to start from so we get a fresh profile
// port - The port to listen on.
func (c *Gcd) StartProcess(exePath, userDir, port string) {
	c.port = port
	c.addr = fmt.Sprintf("%s:%s", c.host, c.port)
	c.apiEndpoint = fmt.Sprintf("http://%s/json", c.addr)
	// profile directory
	c.flags = append(c.flags, fmt.Sprintf("--user-data-dir=%s", userDir))
	// debug port to use
	c.flags = append(c.flags, fmt.Sprintf("--remote-debugging-port=%s", port))
	// bypass first run check
	c.flags = append(c.flags, "--no-first-run")
	// bypass default browser check
	c.flags = append(c.flags, "--no-default-browser-check")

	c.chromeCmd = exec.Command(exePath, c.flags...)
	// add custom environment variables.
	c.chromeCmd.Env = os.Environ()
	c.chromeCmd.Env = append(c.chromeCmd.Env, c.env...)
	go func() {
		err := c.chromeCmd.Start()
		if err != nil {
			log.Fatalf("error starting chrome process: %s", err)
		}
		c.chromeProcess = c.chromeCmd.Process
		err = c.chromeCmd.Wait()

		closeMessage := "exited"
		if err != nil {
			closeMessage = err.Error()
		}
		if c.terminatedHandler != nil {
			c.terminatedHandler(closeMessage)
		}
	}()
	go c.probeDebugPort()
	<-c.readyCh
}

// Kills the process
func (c *Gcd) ExitProcess() error {
	return c.chromeProcess.Kill()
}

// ConnectToInstance connects to a running chrome instance without starting a local process
// Host - The host destination.
// Port - The port to listen on.
func (c *Gcd) ConnectToInstance(host string, port string) {
	c.host = host
	c.port = port
	c.addr = fmt.Sprintf("%s:%s", c.host, c.port)
	c.apiEndpoint = fmt.Sprintf("http://%s/json", c.addr)
	go c.probeDebugPort()
	<-c.readyCh
}

// Gets the primary tabs/processes to work with. Each will have their own references
// to the underlying API components (such as Page, Debugger, DOM etc).
func (c *Gcd) GetTargets() ([]*ChromeTarget, error) {
	empty := make(map[string]struct{}, 0)
	return c.GetNewTargets(empty)
}

// Gets a list of current tabs and creates new chrome targets returning a list
// provided they weren't in the knownIds list. Note it is an error to attempt
// to create a new chrome target from one that already exists.
func (c *Gcd) GetNewTargets(knownIds map[string]struct{}) ([]*ChromeTarget, error) {
	resp, err := http.Get(c.apiEndpoint)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, &GcdBodyReadErr{Message: errRead.Error()}
	}

	targets := make([]*TargetInfo, 0)
	err = json.Unmarshal(body, &targets)
	if err != nil {
		return nil, &GcdDecodingErr{Message: err.Error()}
	}

	connectableTargets := make([]*TargetInfo, 0)
	for _, v := range targets {
		if v.WebSocketDebuggerUrl != "" {
			connectableTargets = append(connectableTargets, v)
		}
	}

	chromeTargets := make([]*ChromeTarget, 0)
	for _, v := range connectableTargets {
		if _, ok := knownIds[v.Id]; !ok {
			target, err := openChromeTarget(c.addr, v)
			if err != nil {
				return nil, err
			}
			chromeTargets = append(chromeTargets, target)
		}
	}
	return chromeTargets, nil
}

// Create a new empty tab, returns the chrome target.
func (c *Gcd) NewTab() (*ChromeTarget, error) {
	resp, err := http.Get(c.apiEndpoint + "/new")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, errRead := ioutil.ReadAll(resp.Body)
	if errRead != nil {
		return nil, &GcdBodyReadErr{Message: errRead.Error()}
	}

	tabTarget := &TargetInfo{}
	err = json.Unmarshal(body, &tabTarget)
	if err != nil {
		return nil, &GcdDecodingErr{Message: err.Error()}
	}
	return openChromeTarget(c.addr, tabTarget)
}

// Closes the tab
func (c *Gcd) CloseTab(target *ChromeTarget) error {
	target.shutdown() // close WS connection first
	resp, err := http.Get(fmt.Sprintf("%s/close/%s", c.apiEndpoint, target.Target.Id))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, errRead := ioutil.ReadAll(resp.Body)
	return errRead
}

// Activates (focus) the tab.
func (c *Gcd) ActivateTab(target *ChromeTarget) error {
	resp, err := http.Get(fmt.Sprintf("%s/activate/%s", c.apiEndpoint, target.Target.Id))
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, errRead := ioutil.ReadAll(resp.Body)
	return errRead
}

// probes the debugger report and signals when it's available.
func (c *Gcd) probeDebugPort() {
	ticker := time.NewTicker(time.Millisecond * 100)
	timeoutTicker := time.NewTicker(time.Second * c.timeout)

	defer func() {
		ticker.Stop()
		timeoutTicker.Stop()
	}()

	for {
		select {
		case <-ticker.C:
			resp, err := http.Get(c.apiEndpoint)
			if err != nil {
				continue
			}
			defer resp.Body.Close()
			c.readyCh <- struct{}{}
			return
		case <-timeoutTicker.C:
			log.Fatalf("Unable to contact debugger at %s after %d seconds, gave up", c.apiEndpoint, c.timeout)
		}
	}
}
