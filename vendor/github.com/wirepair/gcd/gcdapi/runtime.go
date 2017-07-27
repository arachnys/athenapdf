// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Runtime functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Mirror object referencing original JavaScript object.
type RuntimeRemoteObject struct {
	Type                string                `json:"type"`                          // Object type.
	Subtype             string                `json:"subtype,omitempty"`             // Object subtype hint. Specified for <code>object</code> type values only.
	ClassName           string                `json:"className,omitempty"`           // Object class (constructor) name. Specified for <code>object</code> type values only.
	Value               interface{}           `json:"value,omitempty"`               // Remote object value in case of primitive values or JSON values (if it was requested).
	UnserializableValue string                `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified does not have <code>value</code>, but gets this property. enum values: Infinity, NaN, -Infinity, -0
	Description         string                `json:"description,omitempty"`         // String representation of the object.
	ObjectId            string                `json:"objectId,omitempty"`            // Unique object identifier (for non-primitive values).
	Preview             *RuntimeObjectPreview `json:"preview,omitempty"`             // Preview containing abbreviated property values. Specified for <code>object</code> type values only.
	CustomPreview       *RuntimeCustomPreview `json:"customPreview,omitempty"`       //
}

// No Description.
type RuntimeCustomPreview struct {
	Header                     string `json:"header"`                     //
	HasBody                    bool   `json:"hasBody"`                    //
	FormatterObjectId          string `json:"formatterObjectId"`          //
	BindRemoteObjectFunctionId string `json:"bindRemoteObjectFunctionId"` //
	ConfigObjectId             string `json:"configObjectId,omitempty"`   //
}

// Object containing abbreviated remote object value.
type RuntimeObjectPreview struct {
	Type        string                    `json:"type"`                  // Object type.
	Subtype     string                    `json:"subtype,omitempty"`     // Object subtype hint. Specified for <code>object</code> type values only.
	Description string                    `json:"description,omitempty"` // String representation of the object.
	Overflow    bool                      `json:"overflow"`              // True iff some of the properties or entries of the original object did not fit.
	Properties  []*RuntimePropertyPreview `json:"properties"`            // List of the properties.
	Entries     []*RuntimeEntryPreview    `json:"entries,omitempty"`     // List of the entries. Specified for <code>map</code> and <code>set</code> subtype values only.
}

// No Description.
type RuntimePropertyPreview struct {
	Name         string                `json:"name"`                   // Property name.
	Type         string                `json:"type"`                   // Object type. Accessor means that the property itself is an accessor property.
	Value        string                `json:"value,omitempty"`        // User-friendly property value string.
	ValuePreview *RuntimeObjectPreview `json:"valuePreview,omitempty"` // Nested value preview.
	Subtype      string                `json:"subtype,omitempty"`      // Object subtype hint. Specified for <code>object</code> type values only.
}

// No Description.
type RuntimeEntryPreview struct {
	Key   *RuntimeObjectPreview `json:"key,omitempty"` // Preview of the key. Specified for map-like collection entries.
	Value *RuntimeObjectPreview `json:"value"`         // Preview of the value.
}

// Object property descriptor.
type RuntimePropertyDescriptor struct {
	Name         string               `json:"name"`                // Property name or symbol description.
	Value        *RuntimeRemoteObject `json:"value,omitempty"`     // The value associated with the property.
	Writable     bool                 `json:"writable,omitempty"`  // True if the value associated with the property may be changed (data descriptors only).
	Get          *RuntimeRemoteObject `json:"get,omitempty"`       // A function which serves as a getter for the property, or <code>undefined</code> if there is no getter (accessor descriptors only).
	Set          *RuntimeRemoteObject `json:"set,omitempty"`       // A function which serves as a setter for the property, or <code>undefined</code> if there is no setter (accessor descriptors only).
	Configurable bool                 `json:"configurable"`        // True if the type of this property descriptor may be changed and if the property may be deleted from the corresponding object.
	Enumerable   bool                 `json:"enumerable"`          // True if this property shows up during enumeration of the properties on the corresponding object.
	WasThrown    bool                 `json:"wasThrown,omitempty"` // True if the result was thrown during the evaluation.
	IsOwn        bool                 `json:"isOwn,omitempty"`     // True if the property is owned for the object.
	Symbol       *RuntimeRemoteObject `json:"symbol,omitempty"`    // Property symbol object, if the property is of the <code>symbol</code> type.
}

// Object internal property descriptor. This property isn't normally visible in JavaScript code.
type RuntimeInternalPropertyDescriptor struct {
	Name  string               `json:"name"`            // Conventional property name.
	Value *RuntimeRemoteObject `json:"value,omitempty"` // The value associated with the property.
}

// Represents function call argument. Either remote object id <code>objectId</code>, primitive <code>value</code>, unserializable primitive value or neither of (for undefined) them should be specified.
type RuntimeCallArgument struct {
	Value               interface{} `json:"value,omitempty"`               // Primitive value.
	UnserializableValue string      `json:"unserializableValue,omitempty"` // Primitive value which can not be JSON-stringified. enum values: Infinity, NaN, -Infinity, -0
	ObjectId            string      `json:"objectId,omitempty"`            // Remote object handle.
}

// Description of an isolated world.
type RuntimeExecutionContextDescription struct {
	Id      int                    `json:"id"`                // Unique id of the execution context. It can be used to specify in which execution context script evaluation should be performed.
	Origin  string                 `json:"origin"`            // Execution context origin.
	Name    string                 `json:"name"`              // Human readable name describing given context.
	AuxData map[string]interface{} `json:"auxData,omitempty"` // Embedder-specific auxiliary data.
}

// Detailed information about exception (or error) that was thrown during script compilation or execution.
type RuntimeExceptionDetails struct {
	ExceptionId        int                  `json:"exceptionId"`                  // Exception id.
	Text               string               `json:"text"`                         // Exception text, which should be used together with exception object when available.
	LineNumber         int                  `json:"lineNumber"`                   // Line number of the exception location (0-based).
	ColumnNumber       int                  `json:"columnNumber"`                 // Column number of the exception location (0-based).
	ScriptId           string               `json:"scriptId,omitempty"`           // Script ID of the exception location.
	Url                string               `json:"url,omitempty"`                // URL of the exception location, to be used when the script was not reported.
	StackTrace         *RuntimeStackTrace   `json:"stackTrace,omitempty"`         // JavaScript stack trace if available.
	Exception          *RuntimeRemoteObject `json:"exception,omitempty"`          // Exception object if available.
	ExecutionContextId int                  `json:"executionContextId,omitempty"` // Identifier of the context where exception happened.
}

// Stack entry for runtime errors and assertions.
type RuntimeCallFrame struct {
	FunctionName string `json:"functionName"` // JavaScript function name.
	ScriptId     string `json:"scriptId"`     // JavaScript script id.
	Url          string `json:"url"`          // JavaScript script name or url.
	LineNumber   int    `json:"lineNumber"`   // JavaScript script line number (0-based).
	ColumnNumber int    `json:"columnNumber"` // JavaScript script column number (0-based).
}

// Call frames for assertions or error messages.
type RuntimeStackTrace struct {
	Description          string              `json:"description,omitempty"`          // String label of this stack trace. For async traces this may be a name of the function that initiated the async call.
	CallFrames           []*RuntimeCallFrame `json:"callFrames"`                     // JavaScript function name.
	Parent               *RuntimeStackTrace  `json:"parent,omitempty"`               // Asynchronous JavaScript stack trace that preceded this stack, if available.
	PromiseCreationFrame *RuntimeCallFrame   `json:"promiseCreationFrame,omitempty"` // Creation frame of the Promise which produced the next synchronous trace when resolved, if available.
}

// Issued when new execution context is created.
type RuntimeExecutionContextCreatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Context *RuntimeExecutionContextDescription `json:"context"` // A newly created execution context.
	} `json:"Params,omitempty"`
}

// Issued when execution context is destroyed.
type RuntimeExecutionContextDestroyedEvent struct {
	Method string `json:"method"`
	Params struct {
		ExecutionContextId int `json:"executionContextId"` // Id of the destroyed context
	} `json:"Params,omitempty"`
}

// Issued when exception was thrown and unhandled.
type RuntimeExceptionThrownEvent struct {
	Method string `json:"method"`
	Params struct {
		Timestamp        float64                  `json:"timestamp"`        // Timestamp of the exception.
		ExceptionDetails *RuntimeExceptionDetails `json:"exceptionDetails"` //
	} `json:"Params,omitempty"`
}

// Issued when unhandled exception was revoked.
type RuntimeExceptionRevokedEvent struct {
	Method string `json:"method"`
	Params struct {
		Reason      string `json:"reason"`      // Reason describing why exception was revoked.
		ExceptionId int    `json:"exceptionId"` // The id of revoked exception, as reported in <code>exceptionUnhandled</code>.
	} `json:"Params,omitempty"`
}

// Issued when console API was called.
type RuntimeConsoleAPICalledEvent struct {
	Method string `json:"method"`
	Params struct {
		Type               string                 `json:"type"`                 // Type of the call.
		Args               []*RuntimeRemoteObject `json:"args"`                 // Call arguments.
		ExecutionContextId int                    `json:"executionContextId"`   // Identifier of the context where the call was made.
		Timestamp          float64                `json:"timestamp"`            // Call timestamp.
		StackTrace         *RuntimeStackTrace     `json:"stackTrace,omitempty"` // Stack trace captured when the call was made.
		Context            string                 `json:"context,omitempty"`    // Console context descriptor for calls on non-default console context (not console.*): 'anonymous#unique-logger-id' for call on unnamed context, 'name#unique-logger-id' for call on named context.
	} `json:"Params,omitempty"`
}

// Issued when object should be inspected (for example, as a result of inspect() command line API call).
type RuntimeInspectRequestedEvent struct {
	Method string `json:"method"`
	Params struct {
		Object *RuntimeRemoteObject   `json:"object"` //
		Hints  map[string]interface{} `json:"hints"`  //
	} `json:"Params,omitempty"`
}

type Runtime struct {
	target gcdmessage.ChromeTargeter
}

func NewRuntime(target gcdmessage.ChromeTargeter) *Runtime {
	c := &Runtime{target: target}
	return c
}

type RuntimeEvaluateParams struct {
	// Expression to evaluate.
	Expression string `json:"expression"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	Silent bool `json:"silent,omitempty"`
	// Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ContextId int `json:"contextId,omitempty"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`
	// Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

// EvaluateWithParams - Evaluates expression on global object.
// Returns -  result - Evaluation result. exceptionDetails - Exception details.
func (c *Runtime) EvaluateWithParams(v *RuntimeEvaluateParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.evaluate", Params: v})
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

// Evaluate - Evaluates expression on global object.
// expression - Expression to evaluate.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// includeCommandLineAPI - Determines whether Command Line API should be available during the evaluation.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
// contextId - Specifies in which execution context to perform evaluation. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// userGesture - Whether execution should be treated as initiated by user in the UI.
// awaitPromise - Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
// Returns -  result - Evaluation result. exceptionDetails - Exception details.
func (c *Runtime) Evaluate(expression string, objectGroup string, includeCommandLineAPI bool, silent bool, contextId int, returnByValue bool, generatePreview bool, userGesture bool, awaitPromise bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v RuntimeEvaluateParams
	v.Expression = expression
	v.ObjectGroup = objectGroup
	v.IncludeCommandLineAPI = includeCommandLineAPI
	v.Silent = silent
	v.ContextId = contextId
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	v.UserGesture = userGesture
	v.AwaitPromise = awaitPromise
	return c.EvaluateWithParams(&v)
}

type RuntimeAwaitPromiseParams struct {
	// Identifier of the promise.
	PromiseObjectId string `json:"promiseObjectId"`
	// Whether the result is expected to be a JSON object that should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

// AwaitPromiseWithParams - Add handler to promise with given promise object id.
// Returns -  result - Promise result. Will contain rejected value if promise was rejected. exceptionDetails - Exception details if stack strace is available.
func (c *Runtime) AwaitPromiseWithParams(v *RuntimeAwaitPromiseParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.awaitPromise", Params: v})
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

// AwaitPromise - Add handler to promise with given promise object id.
// promiseObjectId - Identifier of the promise.
// returnByValue - Whether the result is expected to be a JSON object that should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// Returns -  result - Promise result. Will contain rejected value if promise was rejected. exceptionDetails - Exception details if stack strace is available.
func (c *Runtime) AwaitPromise(promiseObjectId string, returnByValue bool, generatePreview bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v RuntimeAwaitPromiseParams
	v.PromiseObjectId = promiseObjectId
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	return c.AwaitPromiseWithParams(&v)
}

type RuntimeCallFunctionOnParams struct {
	// Identifier of the object to call function on.
	ObjectId string `json:"objectId"`
	// Declaration of the function to call.
	FunctionDeclaration string `json:"functionDeclaration"`
	// Call arguments. All call arguments must belong to the same JavaScript world as the target object.
	Arguments []*RuntimeCallArgument `json:"arguments,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	Silent bool `json:"silent,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should be treated as initiated by user in the UI.
	UserGesture bool `json:"userGesture,omitempty"`
	// Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

// CallFunctionOnWithParams - Calls function with given declaration on the given object. Object group of the result is inherited from the target object.
// Returns -  result - Call result. exceptionDetails - Exception details.
func (c *Runtime) CallFunctionOnWithParams(v *RuntimeCallFunctionOnParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.callFunctionOn", Params: v})
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

// CallFunctionOn - Calls function with given declaration on the given object. Object group of the result is inherited from the target object.
// objectId - Identifier of the object to call function on.
// functionDeclaration - Declaration of the function to call.
// arguments - Call arguments. All call arguments must belong to the same JavaScript world as the target object.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
// returnByValue - Whether the result is expected to be a JSON object which should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// userGesture - Whether execution should be treated as initiated by user in the UI.
// awaitPromise - Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
// Returns -  result - Call result. exceptionDetails - Exception details.
func (c *Runtime) CallFunctionOn(objectId string, functionDeclaration string, arguments []*RuntimeCallArgument, silent bool, returnByValue bool, generatePreview bool, userGesture bool, awaitPromise bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v RuntimeCallFunctionOnParams
	v.ObjectId = objectId
	v.FunctionDeclaration = functionDeclaration
	v.Arguments = arguments
	v.Silent = silent
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	v.UserGesture = userGesture
	v.AwaitPromise = awaitPromise
	return c.CallFunctionOnWithParams(&v)
}

type RuntimeGetPropertiesParams struct {
	// Identifier of the object to return properties for.
	ObjectId string `json:"objectId"`
	// If true, returns properties belonging only to the element itself, not to its prototype chain.
	OwnProperties bool `json:"ownProperties,omitempty"`
	// If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
	AccessorPropertiesOnly bool `json:"accessorPropertiesOnly,omitempty"`
	// Whether preview should be generated for the results.
	GeneratePreview bool `json:"generatePreview,omitempty"`
}

// GetPropertiesWithParams - Returns properties of a given object. Object group of the result is inherited from the target object.
// Returns -  result - Object properties. internalProperties - Internal object properties (only of the element itself). exceptionDetails - Exception details.
func (c *Runtime) GetPropertiesWithParams(v *RuntimeGetPropertiesParams) ([]*RuntimePropertyDescriptor, []*RuntimeInternalPropertyDescriptor, *RuntimeExceptionDetails, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.getProperties", Params: v})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			Result             []*RuntimePropertyDescriptor
			InternalProperties []*RuntimeInternalPropertyDescriptor
			ExceptionDetails   *RuntimeExceptionDetails
		}
	}

	if resp == nil {
		return nil, nil, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, nil, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, nil, nil, err
	}

	return chromeData.Result.Result, chromeData.Result.InternalProperties, chromeData.Result.ExceptionDetails, nil
}

// GetProperties - Returns properties of a given object. Object group of the result is inherited from the target object.
// objectId - Identifier of the object to return properties for.
// ownProperties - If true, returns properties belonging only to the element itself, not to its prototype chain.
// accessorPropertiesOnly - If true, returns accessor properties (with getter/setter) only; internal properties are not returned either.
// generatePreview - Whether preview should be generated for the results.
// Returns -  result - Object properties. internalProperties - Internal object properties (only of the element itself). exceptionDetails - Exception details.
func (c *Runtime) GetProperties(objectId string, ownProperties bool, accessorPropertiesOnly bool, generatePreview bool) ([]*RuntimePropertyDescriptor, []*RuntimeInternalPropertyDescriptor, *RuntimeExceptionDetails, error) {
	var v RuntimeGetPropertiesParams
	v.ObjectId = objectId
	v.OwnProperties = ownProperties
	v.AccessorPropertiesOnly = accessorPropertiesOnly
	v.GeneratePreview = generatePreview
	return c.GetPropertiesWithParams(&v)
}

type RuntimeReleaseObjectParams struct {
	// Identifier of the object to release.
	ObjectId string `json:"objectId"`
}

// ReleaseObjectWithParams - Releases remote object with given id.
func (c *Runtime) ReleaseObjectWithParams(v *RuntimeReleaseObjectParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.releaseObject", Params: v})
}

// ReleaseObject - Releases remote object with given id.
// objectId - Identifier of the object to release.
func (c *Runtime) ReleaseObject(objectId string) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeReleaseObjectParams
	v.ObjectId = objectId
	return c.ReleaseObjectWithParams(&v)
}

type RuntimeReleaseObjectGroupParams struct {
	// Symbolic object group name.
	ObjectGroup string `json:"objectGroup"`
}

// ReleaseObjectGroupWithParams - Releases all remote objects that belong to a given group.
func (c *Runtime) ReleaseObjectGroupWithParams(v *RuntimeReleaseObjectGroupParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.releaseObjectGroup", Params: v})
}

// ReleaseObjectGroup - Releases all remote objects that belong to a given group.
// objectGroup - Symbolic object group name.
func (c *Runtime) ReleaseObjectGroup(objectGroup string) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeReleaseObjectGroupParams
	v.ObjectGroup = objectGroup
	return c.ReleaseObjectGroupWithParams(&v)
}

// Tells inspected instance to run if it was waiting for debugger to attach.
func (c *Runtime) RunIfWaitingForDebugger() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.runIfWaitingForDebugger"})
}

// Enables reporting of execution contexts creation by means of <code>executionContextCreated</code> event. When the reporting gets enabled the event will be sent immediately for each existing execution context.
func (c *Runtime) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.enable"})
}

// Disables reporting of execution contexts creation.
func (c *Runtime) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.disable"})
}

// Discards collected exceptions and console API calls.
func (c *Runtime) DiscardConsoleEntries() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.discardConsoleEntries"})
}

type RuntimeSetCustomObjectFormatterEnabledParams struct {
	//
	Enabled bool `json:"enabled"`
}

// SetCustomObjectFormatterEnabledWithParams -
func (c *Runtime) SetCustomObjectFormatterEnabledWithParams(v *RuntimeSetCustomObjectFormatterEnabledParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.setCustomObjectFormatterEnabled", Params: v})
}

// SetCustomObjectFormatterEnabled -
// enabled -
func (c *Runtime) SetCustomObjectFormatterEnabled(enabled bool) (*gcdmessage.ChromeResponse, error) {
	var v RuntimeSetCustomObjectFormatterEnabledParams
	v.Enabled = enabled
	return c.SetCustomObjectFormatterEnabledWithParams(&v)
}

type RuntimeCompileScriptParams struct {
	// Expression to compile.
	Expression string `json:"expression"`
	// Source url to be set for the script.
	SourceURL string `json:"sourceURL"`
	// Specifies whether the compiled script should be persisted.
	PersistScript bool `json:"persistScript"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextId int `json:"executionContextId,omitempty"`
}

// CompileScriptWithParams - Compiles expression.
// Returns -  scriptId - Id of the script. exceptionDetails - Exception details.
func (c *Runtime) CompileScriptWithParams(v *RuntimeCompileScriptParams) (string, *RuntimeExceptionDetails, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.compileScript", Params: v})
	if err != nil {
		return "", nil, err
	}

	var chromeData struct {
		Result struct {
			ScriptId         string
			ExceptionDetails *RuntimeExceptionDetails
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

	return chromeData.Result.ScriptId, chromeData.Result.ExceptionDetails, nil
}

// CompileScript - Compiles expression.
// expression - Expression to compile.
// sourceURL - Source url to be set for the script.
// persistScript - Specifies whether the compiled script should be persisted.
// executionContextId - Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
// Returns -  scriptId - Id of the script. exceptionDetails - Exception details.
func (c *Runtime) CompileScript(expression string, sourceURL string, persistScript bool, executionContextId int) (string, *RuntimeExceptionDetails, error) {
	var v RuntimeCompileScriptParams
	v.Expression = expression
	v.SourceURL = sourceURL
	v.PersistScript = persistScript
	v.ExecutionContextId = executionContextId
	return c.CompileScriptWithParams(&v)
}

type RuntimeRunScriptParams struct {
	// Id of the script to run.
	ScriptId string `json:"scriptId"`
	// Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
	ExecutionContextId int `json:"executionContextId,omitempty"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
	// In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
	Silent bool `json:"silent,omitempty"`
	// Determines whether Command Line API should be available during the evaluation.
	IncludeCommandLineAPI bool `json:"includeCommandLineAPI,omitempty"`
	// Whether the result is expected to be a JSON object which should be sent by value.
	ReturnByValue bool `json:"returnByValue,omitempty"`
	// Whether preview should be generated for the result.
	GeneratePreview bool `json:"generatePreview,omitempty"`
	// Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
	AwaitPromise bool `json:"awaitPromise,omitempty"`
}

// RunScriptWithParams - Runs script with given id in a given context.
// Returns -  result - Run result. exceptionDetails - Exception details.
func (c *Runtime) RunScriptWithParams(v *RuntimeRunScriptParams) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Runtime.runScript", Params: v})
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

// RunScript - Runs script with given id in a given context.
// scriptId - Id of the script to run.
// executionContextId - Specifies in which execution context to perform script run. If the parameter is omitted the evaluation will be performed in the context of the inspected page.
// objectGroup - Symbolic group name that can be used to release multiple objects.
// silent - In silent mode exceptions thrown during evaluation are not reported and do not pause execution. Overrides <code>setPauseOnException</code> state.
// includeCommandLineAPI - Determines whether Command Line API should be available during the evaluation.
// returnByValue - Whether the result is expected to be a JSON object which should be sent by value.
// generatePreview - Whether preview should be generated for the result.
// awaitPromise - Whether execution should wait for promise to be resolved. If the result of evaluation is not a Promise, it's considered to be an error.
// Returns -  result - Run result. exceptionDetails - Exception details.
func (c *Runtime) RunScript(scriptId string, executionContextId int, objectGroup string, silent bool, includeCommandLineAPI bool, returnByValue bool, generatePreview bool, awaitPromise bool) (*RuntimeRemoteObject, *RuntimeExceptionDetails, error) {
	var v RuntimeRunScriptParams
	v.ScriptId = scriptId
	v.ExecutionContextId = executionContextId
	v.ObjectGroup = objectGroup
	v.Silent = silent
	v.IncludeCommandLineAPI = includeCommandLineAPI
	v.ReturnByValue = returnByValue
	v.GeneratePreview = generatePreview
	v.AwaitPromise = awaitPromise
	return c.RunScriptWithParams(&v)
}
