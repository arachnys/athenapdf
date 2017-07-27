// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains HeapProfiler functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Sampling Heap Profile node. Holds callsite information, allocation statistics and child nodes.
type HeapProfilerSamplingHeapProfileNode struct {
	CallFrame *RuntimeCallFrame                      `json:"callFrame"` // Function location.
	SelfSize  float64                                `json:"selfSize"`  // Allocations size in bytes for the node excluding children.
	Children  []*HeapProfilerSamplingHeapProfileNode `json:"children"`  // Child nodes.
}

// Profile.
type HeapProfilerSamplingHeapProfile struct {
	Head *HeapProfilerSamplingHeapProfileNode `json:"head"` //
}

//
type HeapProfilerAddHeapSnapshotChunkEvent struct {
	Method string `json:"method"`
	Params struct {
		Chunk string `json:"chunk"` //
	} `json:"Params,omitempty"`
}

//
type HeapProfilerReportHeapSnapshotProgressEvent struct {
	Method string `json:"method"`
	Params struct {
		Done     int  `json:"done"`               //
		Total    int  `json:"total"`              //
		Finished bool `json:"finished,omitempty"` //
	} `json:"Params,omitempty"`
}

// If heap objects tracking has been started then backend regularly sends a current value for last seen object id and corresponding timestamp. If the were changes in the heap since last event then one or more heapStatsUpdate events will be sent before a new lastSeenObjectId event.
type HeapProfilerLastSeenObjectIdEvent struct {
	Method string `json:"method"`
	Params struct {
		LastSeenObjectId int     `json:"lastSeenObjectId"` //
		Timestamp        float64 `json:"timestamp"`        //
	} `json:"Params,omitempty"`
}

// If heap objects tracking has been started then backend may send update for one or more fragments
type HeapProfilerHeapStatsUpdateEvent struct {
	Method string `json:"method"`
	Params struct {
		StatsUpdate []int `json:"statsUpdate"` // An array of triplets. Each triplet describes a fragment. The first integer is the fragment index, the second integer is a total count of objects for the fragment, the third integer is a total size of the objects for the fragment.
	} `json:"Params,omitempty"`
}

type HeapProfiler struct {
	target gcdmessage.ChromeTargeter
}

func NewHeapProfiler(target gcdmessage.ChromeTargeter) *HeapProfiler {
	c := &HeapProfiler{target: target}
	return c
}

//
func (c *HeapProfiler) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.enable"})
}

//
func (c *HeapProfiler) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.disable"})
}

type HeapProfilerStartTrackingHeapObjectsParams struct {
	//
	TrackAllocations bool `json:"trackAllocations,omitempty"`
}

// StartTrackingHeapObjectsWithParams -
func (c *HeapProfiler) StartTrackingHeapObjectsWithParams(v *HeapProfilerStartTrackingHeapObjectsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.startTrackingHeapObjects", Params: v})
}

// StartTrackingHeapObjects -
// trackAllocations -
func (c *HeapProfiler) StartTrackingHeapObjects(trackAllocations bool) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerStartTrackingHeapObjectsParams
	v.TrackAllocations = trackAllocations
	return c.StartTrackingHeapObjectsWithParams(&v)
}

type HeapProfilerStopTrackingHeapObjectsParams struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
	ReportProgress bool `json:"reportProgress,omitempty"`
}

// StopTrackingHeapObjectsWithParams -
func (c *HeapProfiler) StopTrackingHeapObjectsWithParams(v *HeapProfilerStopTrackingHeapObjectsParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.stopTrackingHeapObjects", Params: v})
}

// StopTrackingHeapObjects -
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken when the tracking is stopped.
func (c *HeapProfiler) StopTrackingHeapObjects(reportProgress bool) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerStopTrackingHeapObjectsParams
	v.ReportProgress = reportProgress
	return c.StopTrackingHeapObjectsWithParams(&v)
}

type HeapProfilerTakeHeapSnapshotParams struct {
	// If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
	ReportProgress bool `json:"reportProgress,omitempty"`
}

// TakeHeapSnapshotWithParams -
func (c *HeapProfiler) TakeHeapSnapshotWithParams(v *HeapProfilerTakeHeapSnapshotParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.takeHeapSnapshot", Params: v})
}

// TakeHeapSnapshot -
// reportProgress - If true 'reportHeapSnapshotProgress' events will be generated while snapshot is being taken.
func (c *HeapProfiler) TakeHeapSnapshot(reportProgress bool) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerTakeHeapSnapshotParams
	v.ReportProgress = reportProgress
	return c.TakeHeapSnapshotWithParams(&v)
}

//
func (c *HeapProfiler) CollectGarbage() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.collectGarbage"})
}

type HeapProfilerGetObjectByHeapObjectIdParams struct {
	//
	ObjectId string `json:"objectId"`
	// Symbolic group name that can be used to release multiple objects.
	ObjectGroup string `json:"objectGroup,omitempty"`
}

// GetObjectByHeapObjectIdWithParams -
// Returns -  result - Evaluation result.
func (c *HeapProfiler) GetObjectByHeapObjectIdWithParams(v *HeapProfilerGetObjectByHeapObjectIdParams) (*RuntimeRemoteObject, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getObjectByHeapObjectId", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Result *RuntimeRemoteObject
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

// GetObjectByHeapObjectId -
// objectId -
// objectGroup - Symbolic group name that can be used to release multiple objects.
// Returns -  result - Evaluation result.
func (c *HeapProfiler) GetObjectByHeapObjectId(objectId string, objectGroup string) (*RuntimeRemoteObject, error) {
	var v HeapProfilerGetObjectByHeapObjectIdParams
	v.ObjectId = objectId
	v.ObjectGroup = objectGroup
	return c.GetObjectByHeapObjectIdWithParams(&v)
}

type HeapProfilerAddInspectedHeapObjectParams struct {
	// Heap snapshot object id to be accessible by means of $x command line API.
	HeapObjectId string `json:"heapObjectId"`
}

// AddInspectedHeapObjectWithParams - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
func (c *HeapProfiler) AddInspectedHeapObjectWithParams(v *HeapProfilerAddInspectedHeapObjectParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.addInspectedHeapObject", Params: v})
}

// AddInspectedHeapObject - Enables console to refer to the node with given id via $x (see Command Line API for more details $x functions).
// heapObjectId - Heap snapshot object id to be accessible by means of $x command line API.
func (c *HeapProfiler) AddInspectedHeapObject(heapObjectId string) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerAddInspectedHeapObjectParams
	v.HeapObjectId = heapObjectId
	return c.AddInspectedHeapObjectWithParams(&v)
}

type HeapProfilerGetHeapObjectIdParams struct {
	// Identifier of the object to get heap object id for.
	ObjectId string `json:"objectId"`
}

// GetHeapObjectIdWithParams -
// Returns -  heapSnapshotObjectId - Id of the heap snapshot object corresponding to the passed remote object id.
func (c *HeapProfiler) GetHeapObjectIdWithParams(v *HeapProfilerGetHeapObjectIdParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.getHeapObjectId", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			HeapSnapshotObjectId string
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

	return chromeData.Result.HeapSnapshotObjectId, nil
}

// GetHeapObjectId -
// objectId - Identifier of the object to get heap object id for.
// Returns -  heapSnapshotObjectId - Id of the heap snapshot object corresponding to the passed remote object id.
func (c *HeapProfiler) GetHeapObjectId(objectId string) (string, error) {
	var v HeapProfilerGetHeapObjectIdParams
	v.ObjectId = objectId
	return c.GetHeapObjectIdWithParams(&v)
}

type HeapProfilerStartSamplingParams struct {
	// Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
	SamplingInterval float64 `json:"samplingInterval,omitempty"`
}

// StartSamplingWithParams -
func (c *HeapProfiler) StartSamplingWithParams(v *HeapProfilerStartSamplingParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.startSampling", Params: v})
}

// StartSampling -
// samplingInterval - Average sample interval in bytes. Poisson distribution is used for the intervals. The default value is 32768 bytes.
func (c *HeapProfiler) StartSampling(samplingInterval float64) (*gcdmessage.ChromeResponse, error) {
	var v HeapProfilerStartSamplingParams
	v.SamplingInterval = samplingInterval
	return c.StartSamplingWithParams(&v)
}

// StopSampling -
// Returns -  profile - Recorded sampling heap profile.
func (c *HeapProfiler) StopSampling() (*HeapProfilerSamplingHeapProfile, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "HeapProfiler.stopSampling"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Profile *HeapProfilerSamplingHeapProfile
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
