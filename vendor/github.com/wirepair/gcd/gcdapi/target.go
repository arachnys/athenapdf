// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Target functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// No Description.
type TargetTargetInfo struct {
	TargetId string `json:"targetId"` //
	Type     string `json:"type"`     //
	Title    string `json:"title"`    //
	Url      string `json:"url"`      //
	Attached bool   `json:"attached"` // Whether the target has an attached client.
}

// No Description.
type TargetRemoteLocation struct {
	Host string `json:"host"` //
	Port int    `json:"port"` //
}

// Issued when a possible inspection target is created.
type TargetTargetCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetInfo *TargetTargetInfo `json:"targetInfo"` //
	} `json:"Params,omitempty"`
}

// Issued when some information about a target has changed. This only happens between <code>targetCreated</code> and <code>targetDestroyed</code>.
type TargetTargetInfoChangedEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetInfo *TargetTargetInfo `json:"targetInfo"` //
	} `json:"Params,omitempty"`
}

// Issued when a target is destroyed.
type TargetTargetDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetId string `json:"targetId"` //
	} `json:"Params,omitempty"`
}

// Issued when attached to target because of auto-attach or <code>attachToTarget</code> command.
type TargetAttachedToTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetInfo         *TargetTargetInfo `json:"targetInfo"`         //
		WaitingForDebugger bool              `json:"waitingForDebugger"` //
	} `json:"Params,omitempty"`
}

// Issued when detached from target for any reason (including <code>detachFromTarget</code> command).
type TargetDetachedFromTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetId string `json:"targetId"` //
	} `json:"Params,omitempty"`
}

// Notifies about new protocol message from attached target.
type TargetReceivedMessageFromTargetEvent struct {
	Method string `json:"method"`
	Params struct {
		TargetId string `json:"targetId"` //
		Message  string `json:"message"`  //
	} `json:"Params,omitempty"`
}

type Target struct {
	target gcdmessage.ChromeTargeter
}

func NewTarget(target gcdmessage.ChromeTargeter) *Target {
	c := &Target{target: target}
	return c
}

type TargetSetDiscoverTargetsParams struct {
	// Whether to discover available targets.
	Discover bool `json:"discover"`
}

// SetDiscoverTargetsWithParams - Controls whether to discover available targets and notify via <code>targetCreated/targetInfoChanged/targetDestroyed</code> events.
func (c *Target) SetDiscoverTargetsWithParams(v *TargetSetDiscoverTargetsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setDiscoverTargets", Params: v})
}

// SetDiscoverTargets - Controls whether to discover available targets and notify via <code>targetCreated/targetInfoChanged/targetDestroyed</code> events.
// discover - Whether to discover available targets.
func (c *Target) SetDiscoverTargets(discover bool) (*gcdmessage.ChromeResponse, error) {
	var v TargetSetDiscoverTargetsParams
	v.Discover = discover
	return c.SetDiscoverTargetsWithParams(&v)
}

type TargetSetAutoAttachParams struct {
	// Whether to auto-attach to related targets.
	AutoAttach bool `json:"autoAttach"`
	// Whether to pause new targets when attaching to them. Use <code>Runtime.runIfWaitingForDebugger</code> to run paused targets.
	WaitForDebuggerOnStart bool `json:"waitForDebuggerOnStart"`
}

// SetAutoAttachWithParams - Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
func (c *Target) SetAutoAttachWithParams(v *TargetSetAutoAttachParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setAutoAttach", Params: v})
}

// SetAutoAttach - Controls whether to automatically attach to new targets which are considered to be related to this one. When turned on, attaches to all existing related targets as well. When turned off, automatically detaches from all currently attached targets.
// autoAttach - Whether to auto-attach to related targets.
// waitForDebuggerOnStart - Whether to pause new targets when attaching to them. Use <code>Runtime.runIfWaitingForDebugger</code> to run paused targets.
func (c *Target) SetAutoAttach(autoAttach bool, waitForDebuggerOnStart bool) (*gcdmessage.ChromeResponse, error) {
	var v TargetSetAutoAttachParams
	v.AutoAttach = autoAttach
	v.WaitForDebuggerOnStart = waitForDebuggerOnStart
	return c.SetAutoAttachWithParams(&v)
}

type TargetSetAttachToFramesParams struct {
	// Whether to attach to frames.
	Value bool `json:"value"`
}

// SetAttachToFramesWithParams -
func (c *Target) SetAttachToFramesWithParams(v *TargetSetAttachToFramesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setAttachToFrames", Params: v})
}

// SetAttachToFrames -
// value - Whether to attach to frames.
func (c *Target) SetAttachToFrames(value bool) (*gcdmessage.ChromeResponse, error) {
	var v TargetSetAttachToFramesParams
	v.Value = value
	return c.SetAttachToFramesWithParams(&v)
}

type TargetSetRemoteLocationsParams struct {
	// List of remote locations.
	Locations []*TargetRemoteLocation `json:"locations"`
}

// SetRemoteLocationsWithParams - Enables target discovery for the specified locations, when <code>setDiscoverTargets</code> was set to <code>true</code>.
func (c *Target) SetRemoteLocationsWithParams(v *TargetSetRemoteLocationsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.setRemoteLocations", Params: v})
}

// SetRemoteLocations - Enables target discovery for the specified locations, when <code>setDiscoverTargets</code> was set to <code>true</code>.
// locations - List of remote locations.
func (c *Target) SetRemoteLocations(locations []*TargetRemoteLocation) (*gcdmessage.ChromeResponse, error) {
	var v TargetSetRemoteLocationsParams
	v.Locations = locations
	return c.SetRemoteLocationsWithParams(&v)
}

type TargetSendMessageToTargetParams struct {
	//
	TargetId string `json:"targetId"`
	//
	Message string `json:"message"`
}

// SendMessageToTargetWithParams - Sends protocol message to the target with given id.
func (c *Target) SendMessageToTargetWithParams(v *TargetSendMessageToTargetParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.sendMessageToTarget", Params: v})
}

// SendMessageToTarget - Sends protocol message to the target with given id.
// targetId -
// message -
func (c *Target) SendMessageToTarget(targetId string, message string) (*gcdmessage.ChromeResponse, error) {
	var v TargetSendMessageToTargetParams
	v.TargetId = targetId
	v.Message = message
	return c.SendMessageToTargetWithParams(&v)
}

type TargetGetTargetInfoParams struct {
	//
	TargetId string `json:"targetId"`
}

// GetTargetInfoWithParams - Returns information about a target.
// Returns -  targetInfo -
func (c *Target) GetTargetInfoWithParams(v *TargetGetTargetInfoParams) (*TargetTargetInfo, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.getTargetInfo", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TargetInfo *TargetTargetInfo
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

	return chromeData.Result.TargetInfo, nil
}

// GetTargetInfo - Returns information about a target.
// targetId -
// Returns -  targetInfo -
func (c *Target) GetTargetInfo(targetId string) (*TargetTargetInfo, error) {
	var v TargetGetTargetInfoParams
	v.TargetId = targetId
	return c.GetTargetInfoWithParams(&v)
}

type TargetActivateTargetParams struct {
	//
	TargetId string `json:"targetId"`
}

// ActivateTargetWithParams - Activates (focuses) the target.
func (c *Target) ActivateTargetWithParams(v *TargetActivateTargetParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.activateTarget", Params: v})
}

// ActivateTarget - Activates (focuses) the target.
// targetId -
func (c *Target) ActivateTarget(targetId string) (*gcdmessage.ChromeResponse, error) {
	var v TargetActivateTargetParams
	v.TargetId = targetId
	return c.ActivateTargetWithParams(&v)
}

type TargetCloseTargetParams struct {
	//
	TargetId string `json:"targetId"`
}

// CloseTargetWithParams - Closes the target. If the target is a page that gets closed too.
// Returns -  success -
func (c *Target) CloseTargetWithParams(v *TargetCloseTargetParams) (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.closeTarget", Params: v})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Success, nil
}

// CloseTarget - Closes the target. If the target is a page that gets closed too.
// targetId -
// Returns -  success -
func (c *Target) CloseTarget(targetId string) (bool, error) {
	var v TargetCloseTargetParams
	v.TargetId = targetId
	return c.CloseTargetWithParams(&v)
}

type TargetAttachToTargetParams struct {
	//
	TargetId string `json:"targetId"`
}

// AttachToTargetWithParams - Attaches to the target with given id.
// Returns -  success - Whether attach succeeded.
func (c *Target) AttachToTargetWithParams(v *TargetAttachToTargetParams) (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.attachToTarget", Params: v})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Success, nil
}

// AttachToTarget - Attaches to the target with given id.
// targetId -
// Returns -  success - Whether attach succeeded.
func (c *Target) AttachToTarget(targetId string) (bool, error) {
	var v TargetAttachToTargetParams
	v.TargetId = targetId
	return c.AttachToTargetWithParams(&v)
}

type TargetDetachFromTargetParams struct {
	//
	TargetId string `json:"targetId"`
}

// DetachFromTargetWithParams - Detaches from the target with given id.
func (c *Target) DetachFromTargetWithParams(v *TargetDetachFromTargetParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.detachFromTarget", Params: v})
}

// DetachFromTarget - Detaches from the target with given id.
// targetId -
func (c *Target) DetachFromTarget(targetId string) (*gcdmessage.ChromeResponse, error) {
	var v TargetDetachFromTargetParams
	v.TargetId = targetId
	return c.DetachFromTargetWithParams(&v)
}

// CreateBrowserContext - Creates a new empty BrowserContext. Similar to an incognito profile but you can have more than one.
// Returns -  browserContextId - The id of the context created.
func (c *Target) CreateBrowserContext() (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.createBrowserContext"})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			BrowserContextId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.BrowserContextId, nil
}

type TargetDisposeBrowserContextParams struct {
	//
	BrowserContextId string `json:"browserContextId"`
}

// DisposeBrowserContextWithParams - Deletes a BrowserContext, will fail of any open page uses it.
// Returns -  success -
func (c *Target) DisposeBrowserContextWithParams(v *TargetDisposeBrowserContextParams) (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.disposeBrowserContext", Params: v})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Success bool
		}
	}

	if resp == nil {
		return false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, err
	}

	return chromeData.Result.Success, nil
}

// DisposeBrowserContext - Deletes a BrowserContext, will fail of any open page uses it.
// browserContextId -
// Returns -  success -
func (c *Target) DisposeBrowserContext(browserContextId string) (bool, error) {
	var v TargetDisposeBrowserContextParams
	v.BrowserContextId = browserContextId
	return c.DisposeBrowserContextWithParams(&v)
}

type TargetCreateTargetParams struct {
	// The initial URL the page will be navigated to.
	Url string `json:"url"`
	// Frame width in DIP (headless chrome only).
	Width int `json:"width,omitempty"`
	// Frame height in DIP (headless chrome only).
	Height int `json:"height,omitempty"`
	// The browser context to create the page in (headless chrome only).
	BrowserContextId string `json:"browserContextId,omitempty"`
}

// CreateTargetWithParams - Creates a new page.
// Returns -  targetId - The id of the page opened.
func (c *Target) CreateTargetWithParams(v *TargetCreateTargetParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.createTarget", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			TargetId string
		}
	}

	if resp == nil {
		return "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", err
	}

	return chromeData.Result.TargetId, nil
}

// CreateTarget - Creates a new page.
// url - The initial URL the page will be navigated to.
// width - Frame width in DIP (headless chrome only).
// height - Frame height in DIP (headless chrome only).
// browserContextId - The browser context to create the page in (headless chrome only).
// Returns -  targetId - The id of the page opened.
func (c *Target) CreateTarget(url string, width int, height int, browserContextId string) (string, error) {
	var v TargetCreateTargetParams
	v.Url = url
	v.Width = width
	v.Height = height
	v.BrowserContextId = browserContextId
	return c.CreateTargetWithParams(&v)
}

// GetTargets - Retrieves a list of available targets.
// Returns -  targetInfos - The list of targets.
func (c *Target) GetTargets() ([]*TargetTargetInfo, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Target.getTargets"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TargetInfos []*TargetTargetInfo
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

	return chromeData.Result.TargetInfos, nil
}
