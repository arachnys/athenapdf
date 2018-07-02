// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DeviceOrientation functionality.
// API Version: 1.3

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

type DeviceOrientation struct {
	target gcdmessage.ChromeTargeter
}

func NewDeviceOrientation(target gcdmessage.ChromeTargeter) *DeviceOrientation {
	c := &DeviceOrientation{target: target}
	return c
}

// Clears the overridden Device Orientation.
func (c *DeviceOrientation) ClearDeviceOrientationOverride() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DeviceOrientation.clearDeviceOrientationOverride"})
}

type DeviceOrientationSetDeviceOrientationOverrideParams struct {
	// Mock alpha
	Alpha float64 `json:"alpha"`
	// Mock beta
	Beta float64 `json:"beta"`
	// Mock gamma
	Gamma float64 `json:"gamma"`
}

// SetDeviceOrientationOverrideWithParams - Overrides the Device Orientation.
func (c *DeviceOrientation) SetDeviceOrientationOverrideWithParams(v *DeviceOrientationSetDeviceOrientationOverrideParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DeviceOrientation.setDeviceOrientationOverride", Params: v})
}

// SetDeviceOrientationOverride - Overrides the Device Orientation.
// alpha - Mock alpha
// beta - Mock beta
// gamma - Mock gamma
func (c *DeviceOrientation) SetDeviceOrientationOverride(alpha float64, beta float64, gamma float64) (*gcdmessage.ChromeResponse, error) {
	var v DeviceOrientationSetDeviceOrientationOverrideParams
	v.Alpha = alpha
	v.Beta = beta
	v.Gamma = gamma
	return c.SetDeviceOrientationOverrideWithParams(&v)
}
