// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains DOMStorage functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// DOM Storage identifier.
type DOMStorageStorageId struct {
	SecurityOrigin string `json:"securityOrigin"` // Security origin for the storage.
	IsLocalStorage bool   `json:"isLocalStorage"` // Whether the storage is local storage (not session storage).
}

//
type DOMStorageDomStorageItemsClearedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
	} `json:"Params,omitempty"`
}

//
type DOMStorageDomStorageItemRemovedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
		Key       string               `json:"key"`       //
	} `json:"Params,omitempty"`
}

//
type DOMStorageDomStorageItemAddedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
		Key       string               `json:"key"`       //
		NewValue  string               `json:"newValue"`  //
	} `json:"Params,omitempty"`
}

//
type DOMStorageDomStorageItemUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		StorageId *DOMStorageStorageId `json:"storageId"` //
		Key       string               `json:"key"`       //
		OldValue  string               `json:"oldValue"`  //
		NewValue  string               `json:"newValue"`  //
	} `json:"Params,omitempty"`
}

type DOMStorage struct {
	target gcdmessage.ChromeTargeter
}

func NewDOMStorage(target gcdmessage.ChromeTargeter) *DOMStorage {
	c := &DOMStorage{target: target}
	return c
}

// Enables storage tracking, storage events will now be delivered to the client.
func (c *DOMStorage) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.enable"})
}

// Disables storage tracking, prevents storage events from being sent to the client.
func (c *DOMStorage) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.disable"})
}

type DOMStorageClearParams struct {
	//
	StorageId *DOMStorageStorageId `json:"storageId"`
}

// ClearWithParams -
func (c *DOMStorage) ClearWithParams(v *DOMStorageClearParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.clear", Params: v})
}

// Clear -
// storageId -
func (c *DOMStorage) Clear(storageId *DOMStorageStorageId) (*gcdmessage.ChromeResponse, error) {
	var v DOMStorageClearParams
	v.StorageId = storageId
	return c.ClearWithParams(&v)
}

type DOMStorageGetDOMStorageItemsParams struct {
	//
	StorageId *DOMStorageStorageId `json:"storageId"`
}

// GetDOMStorageItemsWithParams -
// Returns -  entries -
func (c *DOMStorage) GetDOMStorageItemsWithParams(v *DOMStorageGetDOMStorageItemsParams) ([]string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.getDOMStorageItems", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Entries []string
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

	return chromeData.Result.Entries, nil
}

// GetDOMStorageItems -
// storageId -
// Returns -  entries -
func (c *DOMStorage) GetDOMStorageItems(storageId *DOMStorageStorageId) ([]string, error) {
	var v DOMStorageGetDOMStorageItemsParams
	v.StorageId = storageId
	return c.GetDOMStorageItemsWithParams(&v)
}

type DOMStorageSetDOMStorageItemParams struct {
	//
	StorageId *DOMStorageStorageId `json:"storageId"`
	//
	Key string `json:"key"`
	//
	Value string `json:"value"`
}

// SetDOMStorageItemWithParams -
func (c *DOMStorage) SetDOMStorageItemWithParams(v *DOMStorageSetDOMStorageItemParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.setDOMStorageItem", Params: v})
}

// SetDOMStorageItem -
// storageId -
// key -
// value -
func (c *DOMStorage) SetDOMStorageItem(storageId *DOMStorageStorageId, key string, value string) (*gcdmessage.ChromeResponse, error) {
	var v DOMStorageSetDOMStorageItemParams
	v.StorageId = storageId
	v.Key = key
	v.Value = value
	return c.SetDOMStorageItemWithParams(&v)
}

type DOMStorageRemoveDOMStorageItemParams struct {
	//
	StorageId *DOMStorageStorageId `json:"storageId"`
	//
	Key string `json:"key"`
}

// RemoveDOMStorageItemWithParams -
func (c *DOMStorage) RemoveDOMStorageItemWithParams(v *DOMStorageRemoveDOMStorageItemParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "DOMStorage.removeDOMStorageItem", Params: v})
}

// RemoveDOMStorageItem -
// storageId -
// key -
func (c *DOMStorage) RemoveDOMStorageItem(storageId *DOMStorageStorageId, key string) (*gcdmessage.ChromeResponse, error) {
	var v DOMStorageRemoveDOMStorageItemParams
	v.StorageId = storageId
	v.Key = key
	return c.RemoveDOMStorageItemWithParams(&v)
}
