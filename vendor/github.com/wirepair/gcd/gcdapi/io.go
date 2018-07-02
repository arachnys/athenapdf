// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains IO functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

type IO struct {
	target gcdmessage.ChromeTargeter
}

func NewIO(target gcdmessage.ChromeTargeter) *IO {
	c := &IO{target: target}
	return c
}

type IOCloseParams struct {
	// Handle of the stream to close.
	Handle string `json:"handle"`
}

// CloseWithParams - Close the stream, discard any temporary backing storage.
func (c *IO) CloseWithParams(v *IOCloseParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IO.close", Params: v})
}

// Close - Close the stream, discard any temporary backing storage.
// handle - Handle of the stream to close.
func (c *IO) Close(handle string) (*gcdmessage.ChromeResponse, error) {
	var v IOCloseParams
	v.Handle = handle
	return c.CloseWithParams(&v)
}

type IOReadParams struct {
	// Handle of the stream to read.
	Handle string `json:"handle"`
	// Seek to the specified offset before reading (if not specificed, proceed with offset following the last read). Some types of streams may only support sequential reads.
	Offset int `json:"offset,omitempty"`
	// Maximum number of bytes to read (left upon the agent discretion if not specified).
	Size int `json:"size,omitempty"`
}

// ReadWithParams - Read a chunk of the stream
// Returns -  base64Encoded - Set if the data is base64-encoded data - Data that were read. eof - Set if the end-of-file condition occured while reading.
func (c *IO) ReadWithParams(v *IOReadParams) (bool, string, bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IO.read", Params: v})
	if err != nil {
		return false, "", false, err
	}

	var chromeData struct {
		Result struct {
			Base64Encoded bool
			Data          string
			Eof           bool
		}
	}

	if resp == nil {
		return false, "", false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return false, "", false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return false, "", false, err
	}

	return chromeData.Result.Base64Encoded, chromeData.Result.Data, chromeData.Result.Eof, nil
}

// Read - Read a chunk of the stream
// handle - Handle of the stream to read.
// offset - Seek to the specified offset before reading (if not specificed, proceed with offset following the last read). Some types of streams may only support sequential reads.
// size - Maximum number of bytes to read (left upon the agent discretion if not specified).
// Returns -  base64Encoded - Set if the data is base64-encoded data - Data that were read. eof - Set if the end-of-file condition occured while reading.
func (c *IO) Read(handle string, offset int, size int) (bool, string, bool, error) {
	var v IOReadParams
	v.Handle = handle
	v.Offset = offset
	v.Size = size
	return c.ReadWithParams(&v)
}

type IOResolveBlobParams struct {
	// Object id of a Blob object wrapper.
	ObjectId string `json:"objectId"`
}

// ResolveBlobWithParams - Return UUID of Blob object specified by a remote object id.
// Returns -  uuid - UUID of the specified Blob.
func (c *IO) ResolveBlobWithParams(v *IOResolveBlobParams) (string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IO.resolveBlob", Params: v})
	if err != nil {
		return "", err
	}

	var chromeData struct {
		Result struct {
			Uuid string
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

	return chromeData.Result.Uuid, nil
}

// ResolveBlob - Return UUID of Blob object specified by a remote object id.
// objectId - Object id of a Blob object wrapper.
// Returns -  uuid - UUID of the specified Blob.
func (c *IO) ResolveBlob(objectId string) (string, error) {
	var v IOResolveBlobParams
	v.ObjectId = objectId
	return c.ResolveBlobWithParams(&v)
}
