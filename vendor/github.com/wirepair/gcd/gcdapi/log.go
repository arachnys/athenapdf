// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Log functionality.
// API Version: 1.2

package gcdapi

import (
	"github.com/wirepair/gcd/gcdmessage"
)

// Log entry.
type LogLogEntry struct {
	Source           string             `json:"source"`                     // Log entry source.
	Level            string             `json:"level"`                      // Log entry severity.
	Text             string             `json:"text"`                       // Logged text.
	Timestamp        float64            `json:"timestamp"`                  // Timestamp when this entry was added.
	Url              string             `json:"url,omitempty"`              // URL of the resource if known.
	LineNumber       int                `json:"lineNumber,omitempty"`       // Line number in the resource.
	StackTrace       *RuntimeStackTrace `json:"stackTrace,omitempty"`       // JavaScript stack trace.
	NetworkRequestId string             `json:"networkRequestId,omitempty"` // Identifier of the network request associated with this entry.
	WorkerId         string             `json:"workerId,omitempty"`         // Identifier of the worker associated with this entry.
}

// Violation configuration setting.
type LogViolationSetting struct {
	Name      string  `json:"name"`      // Violation type.
	Threshold float64 `json:"threshold"` // Time threshold to trigger upon.
}

// Issued when new message was logged.
type LogEntryAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		Entry *LogLogEntry `json:"entry"` // The entry.
	} `json:"Params,omitempty"`
}

type Log struct {
	target gcdmessage.ChromeTargeter
}

func NewLog(target gcdmessage.ChromeTargeter) *Log {
	c := &Log{target: target}
	return c
}

// Enables log domain, sends the entries collected so far to the client by means of the <code>entryAdded</code> notification.
func (c *Log) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.enable"})
}

// Disables log domain, prevents further log entries from being reported to the client.
func (c *Log) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.disable"})
}

// Clears the log.
func (c *Log) Clear() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.clear"})
}

type LogStartViolationsReportParams struct {
	// Configuration for violations.
	Config []*LogViolationSetting `json:"config"`
}

// StartViolationsReportWithParams - start violation reporting.
func (c *Log) StartViolationsReportWithParams(v *LogStartViolationsReportParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.startViolationsReport", Params: v})
}

// StartViolationsReport - start violation reporting.
// config - Configuration for violations.
func (c *Log) StartViolationsReport(config []*LogViolationSetting) (*gcdmessage.ChromeResponse, error) {
	var v LogStartViolationsReportParams
	v.Config = config
	return c.StartViolationsReportWithParams(&v)
}

// Stop violation reporting.
func (c *Log) StopViolationsReport() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Log.stopViolationsReport"})
}
