// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Performance functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Run-time execution metric.
type PerformanceMetric struct {
	Name  string  `json:"name"`  // Metric name.
	Value float64 `json:"value"` // Metric value.
}

// Current values of the metrics.
type PerformanceMetricsEvent struct {
	Method string `json:"method"`
	Params struct {
		Metrics []*PerformanceMetric `json:"metrics"` // Current values of the metrics.
		Title   string               `json:"title"`   // Timestamp title.
	} `json:"Params,omitempty"`
}

type Performance struct {
	target gcdmessage.ChromeTargeter
}

func NewPerformance(target gcdmessage.ChromeTargeter) *Performance {
	c := &Performance{target: target}
	return c
}

// Disable collecting and reporting metrics.
func (c *Performance) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Performance.disable"})
}

// Enable collecting and reporting metrics.
func (c *Performance) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Performance.enable"})
}

// GetMetrics - Retrieve current values of run-time metrics.
// Returns -  metrics - Current values for run-time metrics.
func (c *Performance) GetMetrics() ([]*PerformanceMetric, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Performance.getMetrics"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Metrics []*PerformanceMetric
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

	return chromeData.Result.Metrics, nil
}
