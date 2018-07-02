// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Memory functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Heap profile sample.
type MemorySamplingProfileNode struct {
	Size  float64  `json:"size"`  // Size of the sampled allocation.
	Total float64  `json:"total"` // Total bytes attributed to this sample.
	Stack []string `json:"stack"` // Execution stack at the point of allocation.
}

// Array of heap profile samples.
type MemorySamplingProfile struct {
	Samples []*MemorySamplingProfileNode `json:"samples"` //
}

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

//
func (c *Memory) PrepareForLeakDetection() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.prepareForLeakDetection"})
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

type MemoryStartSamplingParams struct {
	// Average number of bytes between samples.
	SamplingInterval int `json:"samplingInterval,omitempty"`
	// Do not randomize intervals between samples.
	SuppressRandomness bool `json:"suppressRandomness,omitempty"`
}

// StartSamplingWithParams - Start collecting native memory profile.
func (c *Memory) StartSamplingWithParams(v *MemoryStartSamplingParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.startSampling", Params: v})
}

// StartSampling - Start collecting native memory profile.
// samplingInterval - Average number of bytes between samples.
// suppressRandomness - Do not randomize intervals between samples.
func (c *Memory) StartSampling(samplingInterval int, suppressRandomness bool) (*gcdmessage.ChromeResponse, error) {
	var v MemoryStartSamplingParams
	v.SamplingInterval = samplingInterval
	v.SuppressRandomness = suppressRandomness
	return c.StartSamplingWithParams(&v)
}

// Stop collecting native memory profile.
func (c *Memory) StopSampling() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.stopSampling"})
}

// GetAllTimeSamplingProfile - Retrieve native memory allocations profile collected since renderer process startup.
// Returns -  profile -
func (c *Memory) GetAllTimeSamplingProfile() (*MemorySamplingProfile, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.getAllTimeSamplingProfile"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Profile *MemorySamplingProfile
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

	return chromeData.Result.Profile, nil
}

// GetBrowserSamplingProfile - Retrieve native memory allocations profile collected since browser process startup.
// Returns -  profile -
func (c *Memory) GetBrowserSamplingProfile() (*MemorySamplingProfile, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.getBrowserSamplingProfile"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Profile *MemorySamplingProfile
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

	return chromeData.Result.Profile, nil
}

// GetSamplingProfile - Retrieve native memory allocations profile collected since last `startSampling` call.
// Returns -  profile -
func (c *Memory) GetSamplingProfile() (*MemorySamplingProfile, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Memory.getSamplingProfile"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Profile *MemorySamplingProfile
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

	return chromeData.Result.Profile, nil
}
