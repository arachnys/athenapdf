// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOMDebugger functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Object event listener.
type DOMDebuggerEventListener struct {
	Type            string               `json:"type"`                      // `EventListener`'s type.
	UseCapture      bool                 `json:"useCapture"`                // `EventListener`'s useCapture.
	Passive         bool                 `json:"passive"`                   // `EventListener`'s passive flag.
	Once            bool                 `json:"once"`                      // `EventListener`'s once flag.
	ScriptId        string               `json:"scriptId"`                  // Script id of the handler code.
	LineNumber      int                  `json:"lineNumber"`                // Line number in the script (0-based).
	ColumnNumber    int                  `json:"columnNumber"`              // Column number in the script (0-based).
	Handler         *RuntimeRemoteObject `json:"handler,omitempty"`         // Event handler function value.
	OriginalHandler *RuntimeRemoteObject `json:"originalHandler,omitempty"` // Event original handler function value.
	BackendNodeId   int                  `json:"backendNodeId,omitempty"`   // Node the listener is added to (if any).
}

type DOMDebugger struct {
	target gcdmessage.ChromeTargeter
}

func NewDOMDebugger(target gcdmessage.ChromeTargeter) *DOMDebugger {
	c := &DOMDebugger{target: target}
	return c
}

type DOMDebuggerGetEventListenersParams struct {
	// Identifier of the object to return listeners for.
	ObjectId string `json:"objectId"`
	// The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
	Depth int `json:"depth,omitempty"`
	// Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled.
	Pierce bool `json:"pierce,omitempty"`
}

// GetEventListenersWithParams - Returns event listeners of the given object.
// Returns -  listeners - Array of relevant listeners.
func (c *DOMDebugger) GetEventListenersWithParams(v *DOMDebuggerGetEventListenersParams) ([]*DOMDebuggerEventListener, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.getEventListeners", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Listeners []*DOMDebuggerEventListener
		}
	}

	if resp == nil {
		return nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, err
	}

	return chromeData.Result.Listeners, nil
}

// GetEventListeners - Returns event listeners of the given object.
// objectId - Identifier of the object to return listeners for.
// depth - The maximum depth at which Node children should be retrieved, defaults to 1. Use -1 for the entire subtree or provide an integer larger than 0.
// pierce - Whether or not iframes and shadow roots should be traversed when returning the subtree (default is false). Reports listeners for all contexts if pierce is enabled.
// Returns -  listeners - Array of relevant listeners.
func (c *DOMDebugger) GetEventListeners(objectId string, depth int, pierce bool) ([]*DOMDebuggerEventListener, error) {
	var v DOMDebuggerGetEventListenersParams
	v.ObjectId = objectId
	v.Depth = depth
	v.Pierce = pierce
	return c.GetEventListenersWithParams(&v)
}

type DOMDebuggerRemoveDOMBreakpointParams struct {
	// Identifier of the node to remove breakpoint from.
	NodeId int `json:"nodeId"`
	// Type of the breakpoint to remove. enum values: subtree-modified, attribute-modified, node-removed
	TheType string `json:"type"`
}

// RemoveDOMBreakpointWithParams - Removes DOM breakpoint that was set using `setDOMBreakpoint`.
func (c *DOMDebugger) RemoveDOMBreakpointWithParams(v *DOMDebuggerRemoveDOMBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.removeDOMBreakpoint", Params: v})
}

// RemoveDOMBreakpoint - Removes DOM breakpoint that was set using `setDOMBreakpoint`.
// nodeId - Identifier of the node to remove breakpoint from.
// type - Type of the breakpoint to remove. enum values: subtree-modified, attribute-modified, node-removed
func (c *DOMDebugger) RemoveDOMBreakpoint(nodeId int, theType string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDebuggerRemoveDOMBreakpointParams
	v.NodeId = nodeId
	v.TheType = theType
	return c.RemoveDOMBreakpointWithParams(&v)
}

type DOMDebuggerRemoveEventListenerBreakpointParams struct {
	// Event name.
	EventName string `json:"eventName"`
	// EventTarget interface name.
	TargetName string `json:"targetName,omitempty"`
}

// RemoveEventListenerBreakpointWithParams - Removes breakpoint on particular DOM event.
func (c *DOMDebugger) RemoveEventListenerBreakpointWithParams(v *DOMDebuggerRemoveEventListenerBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.removeEventListenerBreakpoint", Params: v})
}

// RemoveEventListenerBreakpoint - Removes breakpoint on particular DOM event.
// eventName - Event name.
// targetName - EventTarget interface name.
func (c *DOMDebugger) RemoveEventListenerBreakpoint(eventName string, targetName string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDebuggerRemoveEventListenerBreakpointParams
	v.EventName = eventName
	v.TargetName = targetName
	return c.RemoveEventListenerBreakpointWithParams(&v)
}

type DOMDebuggerRemoveInstrumentationBreakpointParams struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// RemoveInstrumentationBreakpointWithParams - Removes breakpoint on particular native event.
func (c *DOMDebugger) RemoveInstrumentationBreakpointWithParams(v *DOMDebuggerRemoveInstrumentationBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.removeInstrumentationBreakpoint", Params: v})
}

// RemoveInstrumentationBreakpoint - Removes breakpoint on particular native event.
// eventName - Instrumentation name to stop on.
func (c *DOMDebugger) RemoveInstrumentationBreakpoint(eventName string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDebuggerRemoveInstrumentationBreakpointParams
	v.EventName = eventName
	return c.RemoveInstrumentationBreakpointWithParams(&v)
}

type DOMDebuggerRemoveXHRBreakpointParams struct {
	// Resource URL substring.
	Url string `json:"url"`
}

// RemoveXHRBreakpointWithParams - Removes breakpoint from XMLHttpRequest.
func (c *DOMDebugger) RemoveXHRBreakpointWithParams(v *DOMDebuggerRemoveXHRBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.removeXHRBreakpoint", Params: v})
}

// RemoveXHRBreakpoint - Removes breakpoint from XMLHttpRequest.
// url - Resource URL substring.
func (c *DOMDebugger) RemoveXHRBreakpoint(url string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDebuggerRemoveXHRBreakpointParams
	v.Url = url
	return c.RemoveXHRBreakpointWithParams(&v)
}

type DOMDebuggerSetDOMBreakpointParams struct {
	// Identifier of the node to set breakpoint on.
	NodeId int `json:"nodeId"`
	// Type of the operation to stop upon. enum values: subtree-modified, attribute-modified, node-removed
	TheType string `json:"type"`
}

// SetDOMBreakpointWithParams - Sets breakpoint on particular operation with DOM.
func (c *DOMDebugger) SetDOMBreakpointWithParams(v *DOMDebuggerSetDOMBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.setDOMBreakpoint", Params: v})
}

// SetDOMBreakpoint - Sets breakpoint on particular operation with DOM.
// nodeId - Identifier of the node to set breakpoint on.
// type - Type of the operation to stop upon. enum values: subtree-modified, attribute-modified, node-removed
func (c *DOMDebugger) SetDOMBreakpoint(nodeId int, theType string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDebuggerSetDOMBreakpointParams
	v.NodeId = nodeId
	v.TheType = theType
	return c.SetDOMBreakpointWithParams(&v)
}

type DOMDebuggerSetEventListenerBreakpointParams struct {
	// DOM Event name to stop on (any DOM event will do).
	EventName string `json:"eventName"`
	// EventTarget interface name to stop on. If equal to `"*"` or not provided, will stop on any EventTarget.
	TargetName string `json:"targetName,omitempty"`
}

// SetEventListenerBreakpointWithParams - Sets breakpoint on particular DOM event.
func (c *DOMDebugger) SetEventListenerBreakpointWithParams(v *DOMDebuggerSetEventListenerBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.setEventListenerBreakpoint", Params: v})
}

// SetEventListenerBreakpoint - Sets breakpoint on particular DOM event.
// eventName - DOM Event name to stop on (any DOM event will do).
// targetName - EventTarget interface name to stop on. If equal to `"*"` or not provided, will stop on any EventTarget.
func (c *DOMDebugger) SetEventListenerBreakpoint(eventName string, targetName string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDebuggerSetEventListenerBreakpointParams
	v.EventName = eventName
	v.TargetName = targetName
	return c.SetEventListenerBreakpointWithParams(&v)
}

type DOMDebuggerSetInstrumentationBreakpointParams struct {
	// Instrumentation name to stop on.
	EventName string `json:"eventName"`
}

// SetInstrumentationBreakpointWithParams - Sets breakpoint on particular native event.
func (c *DOMDebugger) SetInstrumentationBreakpointWithParams(v *DOMDebuggerSetInstrumentationBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.setInstrumentationBreakpoint", Params: v})
}

// SetInstrumentationBreakpoint - Sets breakpoint on particular native event.
// eventName - Instrumentation name to stop on.
func (c *DOMDebugger) SetInstrumentationBreakpoint(eventName string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDebuggerSetInstrumentationBreakpointParams
	v.EventName = eventName
	return c.SetInstrumentationBreakpointWithParams(&v)
}

type DOMDebuggerSetXHRBreakpointParams struct {
	// Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
	Url string `json:"url"`
}

// SetXHRBreakpointWithParams - Sets breakpoint on XMLHttpRequest.
func (c *DOMDebugger) SetXHRBreakpointWithParams(v *DOMDebuggerSetXHRBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMDebugger.setXHRBreakpoint", Params: v})
}

// SetXHRBreakpoint - Sets breakpoint on XMLHttpRequest.
// url - Resource URL substring. All XHRs having this substring in the URL will get stopped upon.
func (c *DOMDebugger) SetXHRBreakpoint(url string) (*gcdmessage.ChromeResponse, error) {
	var v DOMDebuggerSetXHRBreakpointParams
	v.Url = url
	return c.SetXHRBreakpointWithParams(&v)
}
