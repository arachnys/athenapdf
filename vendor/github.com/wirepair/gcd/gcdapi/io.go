// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains IO functionality.
// API Version: 1.2

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

type IOReadParams struct {
	// Handle of the stream to read.
	Handle string `json:"handle"`
	// Seek to the specified offset before reading (if not specificed, proceed with offset following the last read).
	Offset int `json:"offset,omitempty"`
	// Maximum number of bytes to read (left upon the agent discretion if not specified).
	Size int `json:"size,omitempty"`
}

// ReadWithParams - Read a chunk of the stream
// Returns -  data - Data that were read. eof - Set if the end-of-file condition occured while reading.
func (c *IO) ReadWithParams(v *IOReadParams) (string, bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "IO.read", Params: v})
	if err != nil {
		return "", false, err
	}

	var chromeData struct {
		Result struct {
			Data string
			Eof  bool
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

	return chromeData.Result.Data, chromeData.Result.Eof, nil
}

// Read - Read a chunk of the stream
// handle - Handle of the stream to read.
// offset - Seek to the specified offset before reading (if not specificed, proceed with offset following the last read).
// size - Maximum number of bytes to read (left upon the agent discretion if not specified).
// Returns -  data - Data that were read. eof - Set if the end-of-file condition occured while reading.
func (c *IO) Read(handle string, offset int, size int) (string, bool, error) {
	var v IOReadParams
	v.Handle = handle
	v.Offset = offset
	v.Size = size
	return c.ReadWithParams(&v)
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
