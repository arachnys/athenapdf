// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Debugger functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Location in the source code.
type DebuggerLocation struct {
	ScriptId     string `json:"scriptId"`               // Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	LineNumber   int    `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int    `json:"columnNumber,omitempty"` // Column number in the script (0-based).
}

// Location in the source code.
type DebuggerScriptPosition struct {
	LineNumber   int `json:"lineNumber"`   //
	ColumnNumber int `json:"columnNumber"` //
}

// JavaScript call frame. Array of call frames form the call stack.
type DebuggerCallFrame struct {
	CallFrameId      string               `json:"callFrameId"`                // Call frame identifier. This identifier is only valid while the virtual machine is paused.
	FunctionName     string               `json:"functionName"`               // Name of the JavaScript function called on this call frame.
	FunctionLocation *DebuggerLocation    `json:"functionLocation,omitempty"` // Location in the source code.
	Location         *DebuggerLocation    `json:"location"`                   // Location in the source code.
	ScopeChain       []*DebuggerScope     `json:"scopeChain"`                 // Scope chain for this call frame.
	This             *RuntimeRemoteObject `json:"this"`                       // <code>this</code> object for this call frame.
	ReturnValue      *RuntimeRemoteObject `json:"returnValue,omitempty"`      // The value being returned, if the function is at return point.
}

// Scope description.
type DebuggerScope struct {
	Type          string               `json:"type"`                    // Scope type.
	Object        *RuntimeRemoteObject `json:"object"`                  // Object representing the scope. For <code>global</code> and <code>with</code> scopes it represents the actual object; for the rest of the scopes, it is artificial transient object enumerating scope variables as its properties.
	Name          string               `json:"name,omitempty"`          //
	StartLocation *DebuggerLocation    `json:"startLocation,omitempty"` // Location in the source code where scope starts
	EndLocation   *DebuggerLocation    `json:"endLocation,omitempty"`   // Location in the source code where scope ends
}

// Search match for resource.
type DebuggerSearchMatch struct {
	LineNumber  float64 `json:"lineNumber"`  // Line number in resource content.
	LineContent string  `json:"lineContent"` // Line with match content.
}

// No Description.
type DebuggerBreakLocation struct {
	ScriptId     string `json:"scriptId"`               // Script identifier as reported in the <code>Debugger.scriptParsed</code>.
	LineNumber   int    `json:"lineNumber"`             // Line number in the script (0-based).
	ColumnNumber int    `json:"columnNumber,omitempty"` // Column number in the script (0-based).
	Type         string `json:"type,omitempty"`         //
}

// Fired when virtual machine parses script. This event is also fired for all known and uncollected scripts upon enabling debugger.
type DebuggerScriptParsedEvent struct {
	Method string `json:"method"`
	Params struct {
		ScriptId                string                 `json:"scriptId"`                          // Identifier of the script parsed.
		Url                     string                 `json:"url"`                               // URL or name of the script parsed (if any).
		StartLine               int                    `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
		StartColumn             int                    `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
		EndLine                 int                    `json:"endLine"`                           // Last line of the script.
		EndColumn               int                    `json:"endColumn"`                         // Length of the last line of the script.
		ExecutionContextId      int                    `json:"executionContextId"`                // Specifies script creation context.
		Hash                    string                 `json:"hash"`                              // Content hash of the script.
		ExecutionContextAuxData map[string]interface{} `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data.
		IsLiveEdit              bool                   `json:"isLiveEdit,omitempty"`              // True, if this script is generated as a result of the live edit operation.
		SourceMapURL            string                 `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
		HasSourceURL            bool                   `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
		IsModule                bool                   `json:"isModule,omitempty"`                // True, if this script is ES6 module.
		Length                  int                    `json:"length,omitempty"`                  // This script length.
		StackTrace              *RuntimeStackTrace     `json:"stackTrace,omitempty"`              // JavaScript top stack frame of where the script parsed event was triggered if available.
	} `json:"Params,omitempty"`
}

// Fired when virtual machine fails to parse the script.
type DebuggerScriptFailedToParseEvent struct {
	Method string `json:"method"`
	Params struct {
		ScriptId                string                 `json:"scriptId"`                          // Identifier of the script parsed.
		Url                     string                 `json:"url"`                               // URL or name of the script parsed (if any).
		StartLine               int                    `json:"startLine"`                         // Line offset of the script within the resource with given URL (for script tags).
		StartColumn             int                    `json:"startColumn"`                       // Column offset of the script within the resource with given URL.
		EndLine                 int                    `json:"endLine"`                           // Last line of the script.
		EndColumn               int                    `json:"endColumn"`                         // Length of the last line of the script.
		ExecutionContextId      int                    `json:"executionContextId"`                // Specifies script creation context.
		Hash                    string                 `json:"hash"`                              // Content hash of the script.
		ExecutionContextAuxData map[string]interface{} `json:"executionContextAuxData,omitempty"` // Embedder-specific auxiliary data.
		SourceMapURL            string                 `json:"sourceMapURL,omitempty"`            // URL of source map associated with script (if any).
		HasSourceURL            bool                   `json:"hasSourceURL,omitempty"`            // True, if this script has sourceURL.
		IsModule                bool                   `json:"isModule,omitempty"`                // True, if this script is ES6 module.
		Length                  int                    `json:"length,omitempty"`                  // This script length.
		StackTrace              *RuntimeStackTrace     `json:"stackTrace,omitempty"`              // JavaScript top stack frame of where the script parsed event was triggered if available.
	} `json:"Params,omitempty"`
}

// Fired when breakpoint is resolved to an actual script and location.
type DebuggerBreakpointResolvedEvent struct {
	Method string `json:"method"`
	Params struct {
		BreakpointId string            `json:"breakpointId"` // Breakpoint unique identifier.
		Location     *DebuggerLocation `json:"location"`     // Actual breakpoint location.
	} `json:"Params,omitempty"`
}

// Fired when the virtual machine stopped on breakpoint or exception or any other stop criteria.
type DebuggerPausedEvent struct {
	Method string `json:"method"`
	Params struct {
		CallFrames      []*DebuggerCallFrame   `json:"callFrames"`                // Call stack the virtual machine stopped on.
		Reason          string                 `json:"reason"`                    // Pause reason.
		Data            map[string]interface{} `json:"data,omitempty"`            // Object containing break-specific auxiliary properties.
		HitBreakpoints  []string               `json:"hitBreakpoints,omitempty"`  // Hit breakpoints IDs
		AsyncStackTrace *RuntimeStackTrace     `json:"asyncStackTrace,omitempty"` // Async stack trace, if any.
	} `json:"Params,omitempty"`
}

type Debugger struct {
	target gcdmessage.ChromeTargeter
}

func NewDebugger(target gcdmessage.ChromeTargeter) *Debugger {
	c := &Debugger{target: target}
	return c
}

// Enables debugger for the given page. Clients should not assume that the debugging has been enabled until the result for this command is received.
func (c *Debugger) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.enable"})
}

// Disables debugger for given page.
func (c *Debugger) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.disable"})
}

type DebuggerSetBreakpointsActiveParams struct {
	// New value for breakpoints active state.
	Active bool `json:"active"`
}

// SetBreakpointsActiveWithParams - Activates / deactivates all breakpoints on the page.
func (c *Debugger) SetBreakpointsActiveWithParams(v *DebuggerSetBreakpointsActiveParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointsActive", Params: v})
}

// SetBreakpointsActive - Activates / deactivates all breakpoints on the page.
// active - New value for breakpoints active state.
func (c *Debugger) SetBreakpointsActive(active bool) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetBreakpointsActiveParams
	v.Active = active
	return c.SetBreakpointsActiveWithParams(&v)
}

type DebuggerSetSkipAllPausesParams struct {
	// New value for skip pauses state.
	Skip bool `json:"skip"`
}

// SetSkipAllPausesWithParams - Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
func (c *Debugger) SetSkipAllPausesWithParams(v *DebuggerSetSkipAllPausesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setSkipAllPauses", Params: v})
}

// SetSkipAllPauses - Makes page not interrupt on any pauses (breakpoint, exception, dom exception etc).
// skip - New value for skip pauses state.
func (c *Debugger) SetSkipAllPauses(skip bool) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetSkipAllPausesParams
	v.Skip = skip
	return c.SetSkipAllPausesWithParams(&v)
}

type DebuggerSetBreakpointByUrlParams struct {
	// Line number to set breakpoint at.
	LineNumber int `json:"lineNumber"`
	// URL of the resources to set breakpoint on.
	Url string `json:"url,omitempty"`
	// Regex pattern for the URLs of the resources to set breakpoints on. Either <code>url</code> or <code>urlRegex</code> must be specified.
	UrlRegex string `json:"urlRegex,omitempty"`
	// Offset in the line to set breakpoint at.
	ColumnNumber int `json:"columnNumber,omitempty"`
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

// SetBreakpointByUrlWithParams - Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in <code>locations</code> property. Further matching script parsing will result in subsequent <code>breakpointResolved</code> events issued. This logical breakpoint will survive page reloads.
// Returns -  breakpointId - Id of the created breakpoint for further reference. locations - List of the locations this breakpoint resolved into upon addition.
func (c *Debugger) SetBreakpointByUrlWithParams(v *DebuggerSetBreakpointByUrlParams) (string, []*DebuggerLocation, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpointByUrl", Params: v})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		Result struct {
			BreakpointId string
			Locations    []*DebuggerLocation
		}
	}

	if resp == nil {
		return "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, err
	}

	return chromeData.Result.BreakpointId, chromeData.Result.Locations, nil
}

// SetBreakpointByUrl - Sets JavaScript breakpoint at given location specified either by URL or URL regex. Once this command is issued, all existing parsed scripts will have breakpoints resolved and returned in <code>locations</code> property. Further matching script parsing will result in subsequent <code>breakpointResolved</code> events issued. This logical breakpoint will survive page reloads.
// lineNumber - Line number to set breakpoint at.
// url - URL of the resources to set breakpoint on.
// urlRegex - Regex pattern for the URLs of the resources to set breakpoints on. Either <code>url</code> or <code>urlRegex</code> must be specified.
// columnNumber - Offset in the line to set breakpoint at.
// condition - Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
// Returns -  breakpointId - Id of the created breakpoint for further reference. locations - List of the locations this breakpoint resolved into upon addition.
func (c *Debugger) SetBreakpointByUrl(lineNumber int, url string, urlRegex string, columnNumber int, condition string) (string, []*DebuggerLocation, error) {
	var v DebuggerSetBreakpointByUrlParams
	v.LineNumber = lineNumber
	v.Url = url
	v.UrlRegex = urlRegex
	v.ColumnNumber = columnNumber
	v.Condition = condition
	return c.SetBreakpointByUrlWithParams(&v)
}

type DebuggerSetBreakpointParams struct {
	// Location to set breakpoint in.
	Location *DebuggerLocation `json:"location"`
	// Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
	Condition string `json:"condition,omitempty"`
}

// SetBreakpointWithParams - Sets JavaScript breakpoint at a given location.
// Returns -  breakpointId - Id of the created breakpoint for further reference. actualLocation - Location this breakpoint resolved into.
func (c *Debugger) SetBreakpointWithParams(v *DebuggerSetBreakpointParams) (string, *DebuggerLocation, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBreakpoint", Params: v})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		Result struct {
			BreakpointId   string
			ActualLocation *DebuggerLocation
		}
	}

	if resp == nil {
		return "", nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return "", nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return "", nil, err
	}

	return chromeData.Result.BreakpointId, chromeData.Result.ActualLocation, nil
}

// SetBreakpoint - Sets JavaScript breakpoint at a given location.
// location - Location to set breakpoint in.
// condition - Expression to use as a breakpoint condition. When specified, debugger will only stop on the breakpoint if this expression evaluates to true.
// Returns -  breakpointId - Id of the created breakpoint for further reference. actualLocation - Location this breakpoint resolved into.
func (c *Debugger) SetBreakpoint(location *DebuggerLocation, condition string) (string, *DebuggerLocation, error) {
	var v DebuggerSetBreakpointParams
	v.Location = location
	v.Condition = condition
	return c.SetBreakpointWithParams(&v)
}

type DebuggerRemoveBreakpointParams struct {
	//
	BreakpointId string `json:"breakpointId"`
}

// RemoveBreakpointWithParams - Removes JavaScript breakpoint.
func (c *Debugger) RemoveBreakpointWithParams(v *DebuggerRemoveBreakpointParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.removeBreakpoint", Params: v})
}

// RemoveBreakpoint - Removes JavaScript breakpoint.
// breakpointId -
func (c *Debugger) RemoveBreakpoint(breakpointId string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerRemoveBreakpointParams
	v.BreakpointId = breakpointId
	return c.RemoveBreakpointWithParams(&v)
}

type DebuggerGetPossibleBreakpointsParams struct {
	// Start of range to search possible breakpoint locations in.
	Start *DebuggerLocation `json:"start"`
	// End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
	End *DebuggerLocation `json:"end,omitempty"`
	// Only consider locations which are in the same (non-nested) function as start.
	RestrictToFunction bool `json:"restrictToFunction,omitempty"`
}

// GetPossibleBreakpointsWithParams - Returns possible locations for breakpoint. scriptId in start and end range locations should be the same.
// Returns -  locations - List of the possible breakpoint locations.
func (c *Debugger) GetPossibleBreakpointsWithParams(v *DebuggerGetPossibleBreakpointsParams) ([]*DebuggerBreakLocation, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getPossibleBreakpoints", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Locations []*DebuggerBreakLocation
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

	return chromeData.Result.Locations, nil
}

// GetPossibleBreakpoints - Returns possible locations for breakpoint. scriptId in start and end range locations should be the same.
// start - Start of range to search possible breakpoint locations in.
// end - End of range to search possible breakpoint locations in (excluding). When not specified, end of scripts is used as end of range.
// restrictToFunction - Only consider locations which are in the same (non-nested) function as start.
// Returns -  locations - List of the possible breakpoint locations.
func (c *Debugger) GetPossibleBreakpoints(start *DebuggerLocation, end *DebuggerLocation, restrictToFunction bool) ([]*DebuggerBreakLocation, error) {
	var v DebuggerGetPossibleBreakpointsParams
	v.Start = start
	v.End = end
	v.RestrictToFunction = restrictToFunction
	return c.GetPossibleBreakpointsWithParams(&v)
}

type DebuggerContinueToLocationParams struct {
	// Location to continue to.
	Location *DebuggerLocation `json:"location"`
	//
	TargetCallFrames string `json:"targetCallFrames,omitempty"`
}

// ContinueToLocationWithParams - Continues execution until specific location is reached.
func (c *Debugger) ContinueToLocationWithParams(v *DebuggerContinueToLocationParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.continueToLocation", Params: v})
}

// ContinueToLocation - Continues execution until specific location is reached.
// location - Location to continue to.
// targetCallFrames -
func (c *Debugger) ContinueToLocation(location *DebuggerLocation, targetCallFrames string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerContinueToLocationParams
	v.Location = location
	v.TargetCallFrames = targetCallFrames
	return c.ContinueToLocationWithParams(&v)
}

// Steps over the statement.
func (c *Debugger) StepOver() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepOver"})
}

// Steps into the function call.
func (c *Debugger) StepInto() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepInto"})
}

// Steps out of the function call.
func (c *Debugger) StepOut() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.stepOut"})
}

// Stops on the next JavaScript statement.
func (c *Debugger) Pause() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.pause"})
}

// Steps into next scheduled async task if any is scheduled before next pause. Returns success when async task is actually scheduled, returns error if no task were scheduled or another scheduleStepIntoAsync was called.
func (c *Debugger) ScheduleStepIntoAsync() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.scheduleStepIntoAsync"})
}

// Resumes JavaScript execution.
func (c *Debugger) Resume() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.resume"})
}

type DebuggerSearchInContentParams struct {
	// Id of the script to search in.
	ScriptId string `json:"scriptId"`
	// String to search for.
	Query string `json:"query"`
	// If true, search is case sensitive.
	CaseSensitive bool `json:"caseSensitive,omitempty"`
	// If true, treats string parameter as regex.
	IsRegex bool `json:"isRegex,omitempty"`
}

// SearchInContentWithParams - Searches for given string in script content.
// Returns -  result - List of search matches.
func (c *Debugger) SearchInContentWithParams(v *DebuggerSearchInContentParams) ([]*DebuggerSearchMatch, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.searchInContent", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result []*DebuggerSearchMatch
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

// SearchInContent - Searches for given string in script content.
// scriptId - Id of the script to search in.
// query - String to search for.
// caseSensitive - If true, search is case sensitive.
// isRegex - If true, treats string parameter as regex.
// Returns -  result - List of search matches.
func (c *Debugger) SearchInContent(scriptId string, query string, caseSensitive bool, isRegex bool) ([]*DebuggerSearchMatch, error) {
	var v DebuggerSearchInContentParams
	v.ScriptId = scriptId
	v.Query = query
	v.CaseSensitive = caseSensitive
	v.IsRegex = isRegex
	return c.SearchInContentWithParams(&v)
}

type DebuggerSetScriptSourceParams struct {
	// Id of the script to edit.
	ScriptId string `json:"scriptId"`
	// New content of the script.
	ScriptSource string `json:"scriptSource"`
	//  If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
	DryRun bool `json:"dryRun,omitempty"`
}

// SetScriptSourceWithParams - Edits JavaScript source live.
// Returns -  callFrames - New stack trace in case editing has happened while VM was stopped. stackChanged - Whether current call stack  was modified after applying the changes. asyncStackTrace - Async stack trace, if any. exceptionDetails - Exception details if any.
func (c *Debugger) SetScriptSourceWithParams(v *DebuggerSetScriptSourceParams) ([]*DebuggerCallFrame, bool, *RuntimeStackTrace, *RuntimeExceptionDetails, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setScriptSource", Params: v})
	if err != nil {
		return nil, false, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			CallFrames       []*DebuggerCallFrame
			StackChanged     bool
			AsyncStackTrace  *RuntimeStackTrace
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, false, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, false, nil, nil, err
	}

	return chromeData.Result.CallFrames, chromeData.Result.StackChanged, chromeData.Result.AsyncStackTrace, chromeData.Result.ExceptionDetails, nil
}

// SetScriptSource - Edits JavaScript source live.
// scriptId - Id of the script to edit.
// scriptSource - New content of the script.
// dryRun -  If true the change will not actually be applied. Dry run may be used to get result description without actually modifying the code.
// Returns -  callFrames - New stack trace in case editing has happened while VM was stopped. stackChanged - Whether current call stack  was modified after applying the changes. asyncStackTrace - Async stack trace, if any. exceptionDetails - Exception details if any.
func (c *Debugger) SetScriptSource(scriptId string, scriptSource string, dryRun bool) ([]*DebuggerCallFrame, bool, *RuntimeStackTrace, *RuntimeExceptionDetails, error) {
	var v DebuggerSetScriptSourceParams
	v.ScriptId = scriptId
	v.ScriptSource = scriptSource
	v.DryRun = dryRun
	return c.SetScriptSourceWithParams(&v)
}

type DebuggerRestartFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameId string `json:"callFrameId"`
}

// RestartFrameWithParams - Restarts particular call frame from the beginning.
// Returns -  callFrames - New stack trace. asyncStackTrace - Async stack trace, if any.
func (c *Debugger) RestartFrameWithParams(v *DebuggerRestartFrameParams) ([]*DebuggerCallFrame, *RuntimeStackTrace, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.restartFrame", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			CallFrames      []*DebuggerCallFrame
			AsyncStackTrace *RuntimeStackTrace
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.CallFrames, chromeData.Result.AsyncStackTrace, nil
}

// RestartFrame - Restarts particular call frame from the beginning.
// callFrameId - Call frame identifier to evaluate on.
// Returns -  callFrames - New stack trace. asyncStackTrace - Async stack trace, if any.
func (c *Debugger) RestartFrame(callFrameId string) ([]*DebuggerCallFrame, *RuntimeStackTrace, error) {
	var v DebuggerRestartFrameParams
	v.CallFrameId = callFrameId
	return c.RestartFrameWithParams(&v)
}

type DebuggerGetScriptSourceParams struct {
	// Id of the script to get source for.
	ScriptId string `json:"scriptId"`
}

// GetScriptSourceWithParams - Returns source for the script with given id.
// Returns -  scriptSource - Script source.
func (c *Debugger) GetScriptSourceWithParams(v *DebuggerGetScriptSourceParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.getScriptSource", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			ScriptSource string
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

	return chromeData.Result.ScriptSource, nil
}

// GetScriptSource - Returns source for the script with given id.
// scriptId - Id of the script to get source for.
// Returns -  scriptSource - Script source.
func (c *Debugger) GetScriptSource(scriptId string) (string, error) {
	var v DebuggerGetScriptSourceParams
	v.ScriptId = scriptId
	return c.GetScriptSourceWithParams(&v)
}

type DebuggerSetPauseOnExceptionsParams struct {
	// Pause on exceptions mode.
	State string `json:"state"`
}

// SetPauseOnExceptionsWithParams - Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
func (c *Debugger) SetPauseOnExceptionsWithParams(v *DebuggerSetPauseOnExceptionsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setPauseOnExceptions", Params: v})
}

// SetPauseOnExceptions - Defines pause on exceptions state. Can be set to stop on all exceptions, uncaught exceptions or no exceptions. Initial pause on exceptions state is <code>none</code>.
// state - Pause on exceptions mode.
func (c *Debugger) SetPauseOnExceptions(state string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetPauseOnExceptionsParams
	v.State = state
	return c.SetPauseOnExceptionsWithParams(&v)
}

type DebuggerEvaluateOnCallFrameParams struct {
	// Call frame identifier to evaluate on.
	CallFrameId string `json:"callFrameId"`
	// Expression to evaluate.
	Expression string `json:"expression"`
	// String object group name to put result into (allows rapid releasing resulting object handles using <code>releaseObjectGroup</code>).
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Specifies whether command line API should be available to the evaluated expression, defaults to false.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	Silent bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether to throw an exception if side effect cannot be ruled out during evaluation.
	ThrowOnSideEffect bool `json:"throwOnSideEffect,omitempty"`
}

// EvaluateOnCallFrameWithParams - Evaluates expression on a given call frame.
// Returns -  result - Object wrapper for the evaluation result. exceptionDetails - Exception details.
func (c *Debugger) EvaluateOnCallFrameWithParams(v *DebuggerEvaluateOnCallFrameParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.evaluateOnCallFrame", Params: v})
	if err != nil {
		return nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result           *RuntimeRemoteObject
			ExceptionDetails *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.ExceptionDetails, nil
}

// EvaluateOnCallFrame - Evaluates expression on a given call frame.
// callFrameId - Call frame identifier to evaluate on.
// expression - Expression to evaluate.
// objectGroup - String object group name to put result into (allows rapid releasing resulting object handles using <code>releaseObjectGroup</code>).
// includeCommandLineAPI - Specifies whether command line API should be available to the evaluated expression, defaults to false.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// throwOnSideEffect - Whether to throw an exception if side effect cannot be ruled out during evaluation.
// Returns -  result - Object wrapper for the evaluation result. exceptionDetails - Exception details.
func (c *Debugger) EvaluateOnCallFrame(callFrameId string, expression string, objectGroup string, includeCommandLineAPI bool, silent bool, returnByValue bool, generatePreview bool, throwOnSideEffect bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v DebuggerEvaluateOnCallFrameParams
	v.CallFrameId = callFrameId
	v.Expression = expression
	v.ObjectGroup = objectGroup
	v.IncludeCommandLineAPI = includeCommandLineAPI
	v.Silent = silent
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	v.ThrowOnSideEffect = throwOnSideEffect
	return c.EvaluateOnCallFrameWithParams(&v)
}

type DebuggerSetVariableValueParams struct {
	// 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
	ScopeNumber int `json:"scopeNumber"`
	// Variable name.
	VariableName string `json:"variableName"`
	// New variable value.
	NewValue *RuntimeCallArgument `json:"newValue"`
	// Id of callframe that holds variable.
	CallFrameId string `json:"callFrameId"`
}

// SetVariableValueWithParams - Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
func (c *Debugger) SetVariableValueWithParams(v *DebuggerSetVariableValueParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setVariableValue", Params: v})
}

// SetVariableValue - Changes value of variable in a callframe. Object-based scopes are not supported and must be mutated manually.
// scopeNumber - 0-based number of scope as was listed in scope chain. Only 'local', 'closure' and 'catch' scope types are allowed. Other scopes could be manipulated manually.
// variableName - Variable name.
// newValue - New variable value.
// callFrameId - Id of callframe that holds variable.
func (c *Debugger) SetVariableValue(scopeNumber int, variableName string, newValue *RuntimeCallArgument, callFrameId string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetVariableValueParams
	v.ScopeNumber = scopeNumber
	v.VariableName = variableName
	v.NewValue = newValue
	v.CallFrameId = callFrameId
	return c.SetVariableValueWithParams(&v)
}

type DebuggerSetAsyncCallStackDepthParams struct {
	// Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
	MaxDepth int `json:"maxDepth"`
}

// SetAsyncCallStackDepthWithParams - Enables or disables async call stacks tracking.
func (c *Debugger) SetAsyncCallStackDepthWithParams(v *DebuggerSetAsyncCallStackDepthParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setAsyncCallStackDepth", Params: v})
}

// SetAsyncCallStackDepth - Enables or disables async call stacks tracking.
// maxDepth - Maximum depth of async call stacks. Setting to <code>0</code> will effectively disable collecting async call stacks (default).
func (c *Debugger) SetAsyncCallStackDepth(maxDepth int) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetAsyncCallStackDepthParams
	v.MaxDepth = maxDepth
	return c.SetAsyncCallStackDepthWithParams(&v)
}

type DebuggerSetBlackboxPatternsParams struct {
	// Array of regexps that will be used to check script url for blackbox state.
	Patterns []string `json:"patterns"`
}

// SetBlackboxPatternsWithParams - Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
func (c *Debugger) SetBlackboxPatternsWithParams(v *DebuggerSetBlackboxPatternsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBlackboxPatterns", Params: v})
}

// SetBlackboxPatterns - Replace previous blackbox patterns with passed ones. Forces backend to skip stepping/pausing in scripts with url matching one of the patterns. VM will try to leave blackboxed script by performing 'step in' several times, finally resorting to 'step out' if unsuccessful.
// patterns - Array of regexps that will be used to check script url for blackbox state.
func (c *Debugger) SetBlackboxPatterns(patterns []string) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetBlackboxPatternsParams
	v.Patterns = patterns
	return c.SetBlackboxPatternsWithParams(&v)
}

type DebuggerSetBlackboxedRangesParams struct {
	// Id of the script.
	ScriptId string `json:"scriptId"`
	//
	Positions []*DebuggerScriptPosition `json:"positions"`
}

// SetBlackboxedRangesWithParams - Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
func (c *Debugger) SetBlackboxedRangesWithParams(v *DebuggerSetBlackboxedRangesParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Debugger.setBlackboxedRanges", Params: v})
}

// SetBlackboxedRanges - Makes backend skip steps in the script in blackboxed ranges. VM will try leave blacklisted scripts by performing 'step in' several times, finally resorting to 'step out' if unsuccessful. Positions array contains positions where blackbox state is changed. First interval isn't blackboxed. Array should be sorted.
// scriptId - Id of the script.
// positions -
func (c *Debugger) SetBlackboxedRanges(scriptId string, positions []*DebuggerScriptPosition) (*gcdmessage.ChromeResponse, error) {
	var v DebuggerSetBlackboxedRangesParams
	v.ScriptId = scriptId
	v.Positions = positions
	return c.SetBlackboxedRangesWithParams(&v)
}
