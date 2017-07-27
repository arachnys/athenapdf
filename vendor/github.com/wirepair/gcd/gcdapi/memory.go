// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Memory functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

type Memory struct {
	target gcdmessage.ChromeTargeter
}

func NewMemory(target gcdmessage.ChromeTargeter) *Memory {
	c := &Memory{target: target}
	return c
}

// GetDOMCounters -
// Returns -  documents -  nodes -  jsEventListeners -
func (c *Memory) GetDOMCounters() (int, int, int, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.getDOMCounters"})
	if err != nil {
		return 0, 0, 0, err
	}

	var chromeData struct {
		Result struct {
			Documents        int
			Nodes            int
			JsEventListeners int
		}
	}

	if resp == nil {
		return 0, 0, 0, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, 0, 0, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, 0, 0, err
	}

	return chromeData.Result.Documents, chromeData.Result.Nodes, chromeData.Result.JsEventListeners, nil
}

type MemorySetPressureNotificationsSuppressedParams struct {
	// If true, memory pressure notifications will be suppressed.
	Suppressed bool `json:"suppressed"`
}

// SetPressureNotificationsSuppressedWithParams - Enable/disable suppressing memory pressure notifications in all processes.
func (c *Memory) SetPressureNotificationsSuppressedWithParams(v *MemorySetPressureNotificationsSuppressedParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.setPressureNotificationsSuppressed", Params: v})
}

// SetPressureNotificationsSuppressed - Enable/disable suppressing memory pressure notifications in all processes.
// suppressed - If true, memory pressure notifications will be suppressed.
func (c *Memory) SetPressureNotificationsSuppressed(suppressed bool) (*gcdmessage.ChromeResponse, error) {
	var v MemorySetPressureNotificationsSuppressedParams
	v.Suppressed = suppressed
	return c.SetPressureNotificationsSuppressedWithParams(&v)
}

type MemorySimulatePressureNotificationParams struct {
	// Memory pressure level of the notification. enum values: moderate, critical
	Level string `json:"level"`
}

// SimulatePressureNotificationWithParams - Simulate a memory pressure notification in all processes.
func (c *Memory) SimulatePressureNotificationWithParams(v *MemorySimulatePressureNotificationParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.simulatePressureNotification", Params: v})
}

// SimulatePressureNotification - Simulate a memory pressure notification in all processes.
// level - Memory pressure level of the notification. enum values: moderate, critical
func (c *Memory) SimulatePressureNotification(level string) (*gcdmessage.ChromeResponse, error) {
	var v MemorySimulatePressureNotificationParams
	v.Level = level
	return c.SimulatePressureNotificationWithParams(&v)
}
