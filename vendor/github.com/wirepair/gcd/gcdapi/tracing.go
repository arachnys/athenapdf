// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Tracing functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// No Description.
type TracingTraceConfig struct {
	RecordMode           string                 `json:"recordMode,omitempty"`           // Controls how the trace buffer stores data.
	EnableSampling       bool                   `json:"enableSampling,omitempty"`       // Turns on JavaScript stack sampling.
	EnableSystrace       bool                   `json:"enableSystrace,omitempty"`       // Turns on system tracing.
	EnableArgumentFilter bool                   `json:"enableArgumentFilter,omitempty"` // Turns on argument filter.
	IncludedCategories   []string               `json:"includedCategories,omitempty"`   // Included category filters.
	ExcludedCategories   []string               `json:"excludedCategories,omitempty"`   // Excluded category filters.
	SyntheticDelays      []string               `json:"syntheticDelays,omitempty"`      // Configuration to synthesize the delays in tracing.
	MemoryDumpConfig     map[string]interface{} `json:"memoryDumpConfig,omitempty"`     // Configuration for memory dump triggers. Used only when "memory-infra" category is enabled.
}

//
type TracingBufferUsageEvent struct {
	Method string `json:"method"`
	Params struct {
		PercentFull float64 `json:"percentFull,omitempty"` // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
		EventCount  float64 `json:"eventCount,omitempty"`  // An approximate number of events in the trace log.
		Value       float64 `json:"value,omitempty"`       // A number in range [0..1] that indicates the used size of event buffer as a fraction of its total size.
	} `json:"Params,omitempty"`
}

// Signals that tracing is stopped and there is no trace buffers pending flush, all data were delivered via dataCollected events.
type TracingTracingCompleteEvent struct {
	Method string `json:"method"`
	Params struct {
		Stream            string `json:"stream,omitempty"`            // A handle of the stream that holds resulting trace data.
		StreamCompression string `json:"streamCompression,omitempty"` // Compression format of returned stream. enum values: none, gzip
	} `json:"Params,omitempty"`
}

type Tracing struct {
	target gcdmessage.ChromeTargeter
}

func NewTracing(target gcdmessage.ChromeTargeter) *Tracing {
	c := &Tracing{target: target}
	return c
}

// Stop trace events collection.
func (c *Tracing) End() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.end"})
}

// GetCategories - Gets supported tracing categories.
// Returns -  categories - A list of supported tracing categories.
func (c *Tracing) GetCategories() ([]string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.getCategories"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Categories []string
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

	return chromeData.Result.Categories, nil
}

type TracingRecordClockSyncMarkerParams struct {
	// The ID of this clock sync marker
	SyncId string `json:"syncId"`
}

// RecordClockSyncMarkerWithParams - Record a clock sync marker in the trace.
func (c *Tracing) RecordClockSyncMarkerWithParams(v *TracingRecordClockSyncMarkerParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.recordClockSyncMarker", Params: v})
}

// RecordClockSyncMarker - Record a clock sync marker in the trace.
// syncId - The ID of this clock sync marker
func (c *Tracing) RecordClockSyncMarker(syncId string) (*gcdmessage.ChromeResponse, error) {
	var v TracingRecordClockSyncMarkerParams
	v.SyncId = syncId
	return c.RecordClockSyncMarkerWithParams(&v)
}

// RequestMemoryDump - Request a global memory dump.
// Returns -  dumpGuid - GUID of the resulting global memory dump. success - True iff the global memory dump succeeded.
func (c *Tracing) RequestMemoryDump() (string, bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.requestMemoryDump"})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		Result struct {
			DumpGuid string
			Success  bool
		}
	}

	if resp == nil {
		return "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", false, err
	}

	return chromeData.Result.DumpGuid, chromeData.Result.Success, nil
}

type TracingStartParams struct {
	// Category/tag filter
	Categories string `json:"categories,omitempty"`
	// Tracing options
	Options string `json:"options,omitempty"`
	// If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
	BufferUsageReportingInterval float64 `json:"bufferUsageReportingInterval,omitempty"`
	// Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to `ReportEvents`).
	TransferMode string `json:"transferMode,omitempty"`
	// Compression format to use. This only applies when using `ReturnAsStream` transfer mode (defaults to `none`) enum values: none, gzip
	StreamCompression string `json:"streamCompression,omitempty"`
	//
	TraceConfig *TracingTraceConfig `json:"traceConfig,omitempty"`
}

// StartWithParams - Start trace events collection.
func (c *Tracing) StartWithParams(v *TracingStartParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Tracing.start", Params: v})
}

// Start - Start trace events collection.
// categories - Category/tag filter
// options - Tracing options
// bufferUsageReportingInterval - If set, the agent will issue bufferUsage events at this interval, specified in milliseconds
// transferMode - Whether to report trace events as series of dataCollected events or to save trace to a stream (defaults to `ReportEvents`).
// streamCompression - Compression format to use. This only applies when using `ReturnAsStream` transfer mode (defaults to `none`) enum values: none, gzip
// traceConfig -
func (c *Tracing) Start(categories string, options string, bufferUsageReportingInterval float64, transferMode string, streamCompression string, traceConfig *TracingTraceConfig) (*gcdmessage.ChromeResponse, error) {
	var v TracingStartParams
	v.Categories = categories
	v.Options = options
	v.BufferUsageReportingInterval = bufferUsageReportingInterval
	v.TransferMode = transferMode
	v.StreamCompression = streamCompression
	v.TraceConfig = traceConfig
	return c.StartWithParams(&v)
}
