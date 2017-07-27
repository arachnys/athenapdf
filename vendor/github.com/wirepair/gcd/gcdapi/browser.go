// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Browser functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Browser window bounds information
type BrowserBounds struct {
	Left        int    `json:"left,omitempty"`        // The offset from the left edge of the screen to the window in pixels.
	Top         int    `json:"top,omitempty"`         // The offset from the top edge of the screen to the window in pixels.
	Width       int    `json:"width,omitempty"`       // The window width in pixels.
	Height      int    `json:"height,omitempty"`      // The window height in pixels.
	WindowState string `json:"windowState,omitempty"` // The window state. Default to normal. enum values: normal, minimized, maximized, fullscreen
}

type Browser struct {
	target gcdmessage.ChromeTargeter
}

func NewBrowser(target gcdmessage.ChromeTargeter) *Browser {
	c := &Browser{target: target}
	return c
}

type BrowserGetWindowForTargetParams struct {
	// Devtools agent host id.
	TargetId string `json:"targetId"`
}

// GetWindowForTargetWithParams - Get the browser window that contains the devtools target.
// Returns -  windowId - Browser window id. bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (c *Browser) GetWindowForTargetWithParams(v *BrowserGetWindowForTargetParams) (int, *BrowserBounds, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getWindowForTarget", Params: v})
	if err != nil {
		return 0, nil, err
	}

	var chromeData struct {
		Result struct {
			WindowId int
			Bounds   *BrowserBounds
		}
	}

	if resp == nil {
		return 0, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, nil, err
	}

	return chromeData.Result.WindowId, chromeData.Result.Bounds, nil
}

// GetWindowForTarget - Get the browser window that contains the devtools target.
// targetId - Devtools agent host id.
// Returns -  windowId - Browser window id. bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (c *Browser) GetWindowForTarget(targetId string) (int, *BrowserBounds, error) {
	var v BrowserGetWindowForTargetParams
	v.TargetId = targetId
	return c.GetWindowForTargetWithParams(&v)
}

type BrowserSetWindowBoundsParams struct {
	// Browser window id.
	WindowId int `json:"windowId"`
	// New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
	Bounds *BrowserBounds `json:"bounds"`
}

// SetWindowBoundsWithParams - Set position and/or size of the browser window.
func (c *Browser) SetWindowBoundsWithParams(v *BrowserSetWindowBoundsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.setWindowBounds", Params: v})
}

// SetWindowBounds - Set position and/or size of the browser window.
// windowId - Browser window id.
// bounds - New window bounds. The 'minimized', 'maximized' and 'fullscreen' states cannot be combined with 'left', 'top', 'width' or 'height'. Leaves unspecified fields unchanged.
func (c *Browser) SetWindowBounds(windowId int, bounds *BrowserBounds) (*gcdmessage.ChromeResponse, error) {
	var v BrowserSetWindowBoundsParams
	v.WindowId = windowId
	v.Bounds = bounds
	return c.SetWindowBoundsWithParams(&v)
}

type BrowserGetWindowBoundsParams struct {
	// Browser window id.
	WindowId int `json:"windowId"`
}

// GetWindowBoundsWithParams - Get position and size of the browser window.
// Returns -  bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (c *Browser) GetWindowBoundsWithParams(v *BrowserGetWindowBoundsParams) (*BrowserBounds, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Browser.getWindowBounds", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Bounds *BrowserBounds
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

	return chromeData.Result.Bounds, nil
}

// GetWindowBounds - Get position and size of the browser window.
// windowId - Browser window id.
// Returns -  bounds - Bounds information of the window. When window state is 'minimized', the restored window position and size are returned.
func (c *Browser) GetWindowBounds(windowId int) (*BrowserBounds, error) {
	var v BrowserGetWindowBoundsParams
	v.WindowId = windowId
	return c.GetWindowBoundsWithParams(&v)
}
