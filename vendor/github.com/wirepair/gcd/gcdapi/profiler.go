// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Profiler functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Profile node. Holds callsite information, execution statistics and child nodes.
type ProfilerProfileNode struct {
	Id            int                         `json:"id"`                      // Unique id of the node.
	CallFrame     *RuntimeCallFrame           `json:"callFrame"`               // Function location.
	HitCount      int                         `json:"hitCount,omitempty"`      // Number of samples where this node was on top of the call stack.
	Children      []int                       `json:"children,omitempty"`      // Child node ids.
	DeoptReason   string                      `json:"deoptReason,omitempty"`   // The reason of being not optimized. The function may be deoptimized or marked as don't optimize.
	PositionTicks []*ProfilerPositionTickInfo `json:"positionTicks,omitempty"` // An array of source position ticks.
}

// Profile.
type ProfilerProfile struct {
	Nodes      []*ProfilerProfileNode `json:"nodes"`                // The list of profile nodes. First item is the root node.
	StartTime  float64                `json:"startTime"`            // Profiling start timestamp in microseconds.
	EndTime    float64                `json:"endTime"`              // Profiling end timestamp in microseconds.
	Samples    []int                  `json:"samples,omitempty"`    // Ids of samples top nodes.
	TimeDeltas []int                  `json:"timeDeltas,omitempty"` // Time intervals between adjacent samples in microseconds. The first delta is relative to the profile startTime.
}

// Specifies a number of samples attributed to a certain source position.
type ProfilerPositionTickInfo struct {
	Line  int `json:"line"`  // Source line number (1-based).
	Ticks int `json:"ticks"` // Number of samples attributed to the source line.
}

// Coverage data for a source range.
type ProfilerCoverageRange struct {
	StartOffset int `json:"startOffset"` // JavaScript script source offset for the range start.
	EndOffset   int `json:"endOffset"`   // JavaScript script source offset for the range end.
	Count       int `json:"count"`       // Collected execution count of the source range.
}

// Coverage data for a JavaScript function.
type ProfilerFunctionCoverage struct {
	FunctionName    string                   `json:"functionName"`    // JavaScript function name.
	Ranges          []*ProfilerCoverageRange `json:"ranges"`          // Source ranges inside the function with coverage data.
	IsBlockCoverage bool                     `json:"isBlockCoverage"` // Whether coverage data for this function has block granularity.
}

// Coverage data for a JavaScript script.
type ProfilerScriptCoverage struct {
	ScriptId  string                      `json:"scriptId"`  // JavaScript script id.
	Url       string                      `json:"url"`       // JavaScript script name or url.
	Functions []*ProfilerFunctionCoverage `json:"functions"` // Functions contained in the script that has coverage data.
}

// Sent when new profile recording is started using console.profile() call.
type ProfilerConsoleProfileStartedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id       string            `json:"id"`              //
		Location *DebuggerLocation `json:"location"`        // Location of console.profile().
		Title    string            `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
	} `json:"Params,omitempty"`
}

//
type ProfilerConsoleProfileFinishedEvent struct {
	Method string `json:"method"`
	Params struct {
		Id       string            `json:"id"`              //
		Location *DebuggerLocation `json:"location"`        // Location of console.profileEnd().
		Profile  *ProfilerProfile  `json:"profile"`         //
		Title    string            `json:"title,omitempty"` // Profile title passed as an argument to console.profile().
	} `json:"Params,omitempty"`
}

type Profiler struct {
	target gcdmessage.ChromeTargeter
}

func NewProfiler(target gcdmessage.ChromeTargeter) *Profiler {
	c := &Profiler{target: target}
	return c
}

//
func (c *Profiler) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.enable"})
}

//
func (c *Profiler) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.disable"})
}

type ProfilerSetSamplingIntervalParams struct {
	// New sampling interval in microseconds.
	Interval int `json:"interval"`
}

// SetSamplingIntervalWithParams - Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
func (c *Profiler) SetSamplingIntervalWithParams(v *ProfilerSetSamplingIntervalParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.setSamplingInterval", Params: v})
}

// SetSamplingInterval - Changes CPU profiler sampling interval. Must be called before CPU profiles recording started.
// interval - New sampling interval in microseconds.
func (c *Profiler) SetSamplingInterval(interval int) (*gcdmessage.ChromeResponse, error) {
	var v ProfilerSetSamplingIntervalParams
	v.Interval = interval
	return c.SetSamplingIntervalWithParams(&v)
}

//
func (c *Profiler) Start() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.start"})
}

// Stop -
// Returns -  profile - Recorded profile.
func (c *Profiler) Stop() (*ProfilerProfile, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.stop"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Profile *ProfilerProfile
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

type ProfilerStartPreciseCoverageParams struct {
	// Collect accurate call counts beyond simple 'covered' or 'not covered'.
	CallCount bool `json:"callCount,omitempty"`
}

// StartPreciseCoverageWithParams - Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code coverage may be incomplete. Enabling prevents running optimized code and resets execution counters.
func (c *Profiler) StartPreciseCoverageWithParams(v *ProfilerStartPreciseCoverageParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.startPreciseCoverage", Params: v})
}

// StartPreciseCoverage - Enable precise code coverage. Coverage data for JavaScript executed before enabling precise code coverage may be incomplete. Enabling prevents running optimized code and resets execution counters.
// callCount - Collect accurate call counts beyond simple 'covered' or 'not covered'.
func (c *Profiler) StartPreciseCoverage(callCount bool) (*gcdmessage.ChromeResponse, error) {
	var v ProfilerStartPreciseCoverageParams
	v.CallCount = callCount
	return c.StartPreciseCoverageWithParams(&v)
}

// Disable precise code coverage. Disabling releases unnecessary execution count records and allows executing optimized code.
func (c *Profiler) StopPreciseCoverage() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.stopPreciseCoverage"})
}

// TakePreciseCoverage - Collect coverage data for the current isolate, and resets execution counters. Precise code coverage needs to have started.
// Returns -  result - Coverage data for the current isolate.
func (c *Profiler) TakePreciseCoverage() ([]*ProfilerScriptCoverage, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.takePreciseCoverage"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result []*ProfilerScriptCoverage
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

	return chromeData.Result.Result, nil
}

// GetBestEffortCoverage - Collect coverage data for the current isolate. The coverage data may be incomplete due to garbage collection.
// Returns -  result - Coverage data for the current isolate.
func (c *Profiler) GetBestEffortCoverage() ([]*ProfilerScriptCoverage, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Profiler.getBestEffortCoverage"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result []*ProfilerScriptCoverage
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

	return chromeData.Result.Result, nil
}
