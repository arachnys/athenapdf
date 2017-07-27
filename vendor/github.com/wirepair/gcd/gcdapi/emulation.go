// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Emulation functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Screen orientation.
type EmulationScreenOrientation struct {
	Type  string `json:"type"`  // Orientation type.
	Angle int    `json:"angle"` // Orientation angle.
}

type Emulation struct {
	target gcdmessage.ChromeTargeter
}

func NewEmulation(target gcdmessage.ChromeTargeter) *Emulation {
	c := &Emulation{target: target}
	return c
}

type EmulationSetDeviceMetricsOverrideParams struct {
	// Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Width int `json:"width"`
	// Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
	Height int `json:"height"`
	// Overriding device scale factor value. 0 disables the override.
	DeviceScaleFactor float64 `json:"deviceScaleFactor"`
	// Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
	Mobile bool `json:"mobile"`
	// Whether a view that exceeds the available browser window area should be scaled down to fit.
	FitWindow bool `json:"fitWindow"`
	// Scale to apply to resulting view image. Ignored in |fitWindow| mode.
	Scale float64 `json:"scale,omitempty"`
	// Not used.
	OffsetX float64 `json:"offsetX,omitempty"`
	// Not used.
	OffsetY float64 `json:"offsetY,omitempty"`
	// Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	ScreenWidth int `json:"screenWidth,omitempty"`
	// Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	ScreenHeight int `json:"screenHeight,omitempty"`
	// Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	PositionX int `json:"positionX,omitempty"`
	// Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
	PositionY int `json:"positionY,omitempty"`
	// Screen orientation override.
	ScreenOrientation *EmulationScreenOrientation `json:"screenOrientation,omitempty"`
}

// SetDeviceMetricsOverrideWithParams - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
func (c *Emulation) SetDeviceMetricsOverrideWithParams(v *EmulationSetDeviceMetricsOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setDeviceMetricsOverride", Params: v})
}

// SetDeviceMetricsOverride - Overrides the values of device screen dimensions (window.screen.width, window.screen.height, window.innerWidth, window.innerHeight, and "device-width"/"device-height"-related CSS media query results).
// width - Overriding width value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// height - Overriding height value in pixels (minimum 0, maximum 10000000). 0 disables the override.
// deviceScaleFactor - Overriding device scale factor value. 0 disables the override.
// mobile - Whether to emulate mobile device. This includes viewport meta tag, overlay scrollbars, text autosizing and more.
// fitWindow - Whether a view that exceeds the available browser window area should be scaled down to fit.
// scale - Scale to apply to resulting view image. Ignored in |fitWindow| mode.
// offsetX - Not used.
// offsetY - Not used.
// screenWidth - Overriding screen width value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// screenHeight - Overriding screen height value in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// positionX - Overriding view X position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// positionY - Overriding view Y position on screen in pixels (minimum 0, maximum 10000000). Only used for |mobile==true|.
// screenOrientation - Screen orientation override.
func (c *Emulation) SetDeviceMetricsOverride(width int, height int, deviceScaleFactor float64, mobile bool, fitWindow bool, scale float64, offsetX float64, offsetY float64, screenWidth int, screenHeight int, positionX int, positionY int, screenOrientation *EmulationScreenOrientation) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetDeviceMetricsOverrideParams
	v.Width = width
	v.Height = height
	v.DeviceScaleFactor = deviceScaleFactor
	v.Mobile = mobile
	v.FitWindow = fitWindow
	v.Scale = scale
	v.OffsetX = offsetX
	v.OffsetY = offsetY
	v.ScreenWidth = screenWidth
	v.ScreenHeight = screenHeight
	v.PositionX = positionX
	v.PositionY = positionY
	v.ScreenOrientation = screenOrientation
	return c.SetDeviceMetricsOverrideWithParams(&v)
}

// Clears the overriden device metrics.
func (c *Emulation) ClearDeviceMetricsOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.clearDeviceMetricsOverride"})
}

type EmulationForceViewportParams struct {
	// X coordinate of top-left corner of the area (CSS pixels).
	X float64 `json:"x"`
	// Y coordinate of top-left corner of the area (CSS pixels).
	Y float64 `json:"y"`
	// Scale to apply to the area (relative to a page scale of 1.0).
	Scale float64 `json:"scale"`
}

// ForceViewportWithParams - Overrides the visible area of the page. The change is hidden from the page, i.e. the observable scroll position and page scale does not change. In effect, the command moves the specified area of the page into the top-left corner of the frame.
func (c *Emulation) ForceViewportWithParams(v *EmulationForceViewportParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.forceViewport", Params: v})
}

// ForceViewport - Overrides the visible area of the page. The change is hidden from the page, i.e. the observable scroll position and page scale does not change. In effect, the command moves the specified area of the page into the top-left corner of the frame.
// x - X coordinate of top-left corner of the area (CSS pixels).
// y - Y coordinate of top-left corner of the area (CSS pixels).
// scale - Scale to apply to the area (relative to a page scale of 1.0).
func (c *Emulation) ForceViewport(x float64, y float64, scale float64) (*gcdmessage.ChromeResponse, error) {
	var v EmulationForceViewportParams
	v.X = x
	v.Y = y
	v.Scale = scale
	return c.ForceViewportWithParams(&v)
}

// Resets the visible area of the page to the original viewport, undoing any effects of the <code>forceViewport</code> command.
func (c *Emulation) ResetViewport() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.resetViewport"})
}

// Requests that page scale factor is reset to initial values.
func (c *Emulation) ResetPageScaleFactor() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.resetPageScaleFactor"})
}

type EmulationSetPageScaleFactorParams struct {
	// Page scale factor.
	PageScaleFactor float64 `json:"pageScaleFactor"`
}

// SetPageScaleFactorWithParams - Sets a specified page scale factor.
func (c *Emulation) SetPageScaleFactorWithParams(v *EmulationSetPageScaleFactorParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setPageScaleFactor", Params: v})
}

// SetPageScaleFactor - Sets a specified page scale factor.
// pageScaleFactor - Page scale factor.
func (c *Emulation) SetPageScaleFactor(pageScaleFactor float64) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetPageScaleFactorParams
	v.PageScaleFactor = pageScaleFactor
	return c.SetPageScaleFactorWithParams(&v)
}

type EmulationSetVisibleSizeParams struct {
	// Frame width (DIP).
	Width int `json:"width"`
	// Frame height (DIP).
	Height int `json:"height"`
}

// SetVisibleSizeWithParams - Resizes the frame/viewport of the page. Note that this does not affect the frame's container (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported on Android.
func (c *Emulation) SetVisibleSizeWithParams(v *EmulationSetVisibleSizeParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setVisibleSize", Params: v})
}

// SetVisibleSize - Resizes the frame/viewport of the page. Note that this does not affect the frame's container (e.g. browser window). Can be used to produce screenshots of the specified size. Not supported on Android.
// width - Frame width (DIP).
// height - Frame height (DIP).
func (c *Emulation) SetVisibleSize(width int, height int) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetVisibleSizeParams
	v.Width = width
	v.Height = height
	return c.SetVisibleSizeWithParams(&v)
}

type EmulationSetScriptExecutionDisabledParams struct {
	// Whether script execution should be disabled in the page.
	Value bool `json:"value"`
}

// SetScriptExecutionDisabledWithParams - Switches script execution in the page.
func (c *Emulation) SetScriptExecutionDisabledWithParams(v *EmulationSetScriptExecutionDisabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setScriptExecutionDisabled", Params: v})
}

// SetScriptExecutionDisabled - Switches script execution in the page.
// value - Whether script execution should be disabled in the page.
func (c *Emulation) SetScriptExecutionDisabled(value bool) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetScriptExecutionDisabledParams
	v.Value = value
	return c.SetScriptExecutionDisabledWithParams(&v)
}

type EmulationSetGeolocationOverrideParams struct {
	// Mock latitude
	Latitude float64 `json:"latitude,omitempty"`
	// Mock longitude
	Longitude float64 `json:"longitude,omitempty"`
	// Mock accuracy
	Accuracy float64 `json:"accuracy,omitempty"`
}

// SetGeolocationOverrideWithParams - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
func (c *Emulation) SetGeolocationOverrideWithParams(v *EmulationSetGeolocationOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setGeolocationOverride", Params: v})
}

// SetGeolocationOverride - Overrides the Geolocation Position or Error. Omitting any of the parameters emulates position unavailable.
// latitude - Mock latitude
// longitude - Mock longitude
// accuracy - Mock accuracy
func (c *Emulation) SetGeolocationOverride(latitude float64, longitude float64, accuracy float64) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetGeolocationOverrideParams
	v.Latitude = latitude
	v.Longitude = longitude
	v.Accuracy = accuracy
	return c.SetGeolocationOverrideWithParams(&v)
}

// Clears the overriden Geolocation Position and Error.
func (c *Emulation) ClearGeolocationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.clearGeolocationOverride"})
}

type EmulationSetTouchEmulationEnabledParams struct {
	// Whether the touch event emulation should be enabled.
	Enabled bool `json:"enabled"`
	// Touch/gesture events configuration. Default: current platform.
	Configuration string `json:"configuration,omitempty"`
}

// SetTouchEmulationEnabledWithParams - Toggles mouse event-based touch event emulation.
func (c *Emulation) SetTouchEmulationEnabledWithParams(v *EmulationSetTouchEmulationEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setTouchEmulationEnabled", Params: v})
}

// SetTouchEmulationEnabled - Toggles mouse event-based touch event emulation.
// enabled - Whether the touch event emulation should be enabled.
// configuration - Touch/gesture events configuration. Default: current platform.
func (c *Emulation) SetTouchEmulationEnabled(enabled bool, configuration string) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetTouchEmulationEnabledParams
	v.Enabled = enabled
	v.Configuration = configuration
	return c.SetTouchEmulationEnabledWithParams(&v)
}

type EmulationSetEmulatedMediaParams struct {
	// Media type to emulate. Empty string disables the override.
	Media string `json:"media"`
}

// SetEmulatedMediaWithParams - Emulates the given media for CSS media queries.
func (c *Emulation) SetEmulatedMediaWithParams(v *EmulationSetEmulatedMediaParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setEmulatedMedia", Params: v})
}

// SetEmulatedMedia - Emulates the given media for CSS media queries.
// media - Media type to emulate. Empty string disables the override.
func (c *Emulation) SetEmulatedMedia(media string) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetEmulatedMediaParams
	v.Media = media
	return c.SetEmulatedMediaWithParams(&v)
}

type EmulationSetCPUThrottlingRateParams struct {
	// Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
	Rate float64 `json:"rate"`
}

// SetCPUThrottlingRateWithParams - Enables CPU throttling to emulate slow CPUs.
func (c *Emulation) SetCPUThrottlingRateWithParams(v *EmulationSetCPUThrottlingRateParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setCPUThrottlingRate", Params: v})
}

// SetCPUThrottlingRate - Enables CPU throttling to emulate slow CPUs.
// rate - Throttling rate as a slowdown factor (1 is no throttle, 2 is 2x slowdown, etc).
func (c *Emulation) SetCPUThrottlingRate(rate float64) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetCPUThrottlingRateParams
	v.Rate = rate
	return c.SetCPUThrottlingRateWithParams(&v)
}

// CanEmulate - Tells whether emulation is supported.
// Returns -  result - True if emulation is supported.
func (c *Emulation) CanEmulate() (bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.canEmulate"})
	if err != nil {
		return false, err
	}

	var chromeData struct {
		Result struct {
			Result bool
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

	return chromeData.Result.Result, nil
}

type EmulationSetVirtualTimePolicyParams struct {
	//  enum values: advance, pause, pauseIfNetworkFetchesPending
	Policy string `json:"policy"`
	// If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
	Budget int `json:"budget,omitempty"`
}

// SetVirtualTimePolicyWithParams - Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget.
func (c *Emulation) SetVirtualTimePolicyWithParams(v *EmulationSetVirtualTimePolicyParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setVirtualTimePolicy", Params: v})
}

// SetVirtualTimePolicy - Turns on virtual time for all frames (replacing real-time with a synthetic time source) and sets the current virtual time policy.  Note this supersedes any previous time budget.
// policy -  enum values: advance, pause, pauseIfNetworkFetchesPending
// budget - If set, after this many virtual milliseconds have elapsed virtual time will be paused and a virtualTimeBudgetExpired event is sent.
func (c *Emulation) SetVirtualTimePolicy(policy string, budget int) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetVirtualTimePolicyParams
	v.Policy = policy
	v.Budget = budget
	return c.SetVirtualTimePolicyWithParams(&v)
}

type EmulationSetDefaultBackgroundColorOverrideParams struct {
	// RGBA of the default background color. If not specified, any existing override will be cleared.
	Color *DOMRGBA `json:"color,omitempty"`
}

// SetDefaultBackgroundColorOverrideWithParams - Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one.
func (c *Emulation) SetDefaultBackgroundColorOverrideWithParams(v *EmulationSetDefaultBackgroundColorOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Emulation.setDefaultBackgroundColorOverride", Params: v})
}

// SetDefaultBackgroundColorOverride - Sets or clears an override of the default background color of the frame. This override is used if the content does not specify one.
// color - RGBA of the default background color. If not specified, any existing override will be cleared.
func (c *Emulation) SetDefaultBackgroundColorOverride(color *DOMRGBA) (*gcdmessage.ChromeResponse, error) {
	var v EmulationSetDefaultBackgroundColorOverrideParams
	v.Color = color
	return c.SetDefaultBackgroundColorOverrideWithParams(&v)
}
