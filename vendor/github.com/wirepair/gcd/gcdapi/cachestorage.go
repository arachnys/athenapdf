// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains CacheStorage functionality.
// API Version: 1.2

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Data entry.
type CacheStorageDataEntry struct {
	Request      string  `json:"request"`      // Request url spec.
	Response     string  `json:"response"`     // Response status text.
	ResponseTime float64 `json:"responseTime"` // Number of seconds since epoch.
}

// Cache identifier.
type CacheStorageCache struct {
	CacheId        string `json:"cacheId"`        // An opaque unique id of the cache.
	SecurityOrigin string `json:"securityOrigin"` // Security origin of the cache.
	CacheName      string `json:"cacheName"`      // The name of the cache.
}

type CacheStorage struct {
	target gcdmessage.ChromeTargeter
}

func NewCacheStorage(target gcdmessage.ChromeTargeter) *CacheStorage {
	c := &CacheStorage{target: target}
	return c
}

type CacheStorageRequestCacheNamesParams struct {
	// Security origin.
	SecurityOrigin string `json:"securityOrigin"`
}

// RequestCacheNamesWithParams - Requests cache names.
// Returns -  caches - Caches for the security origin.
func (c *CacheStorage) RequestCacheNamesWithParams(v *CacheStorageRequestCacheNamesParams) ([]*CacheStorageCache, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.requestCacheNames", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Caches []*CacheStorageCache
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

	return chromeData.Result.Caches, nil
}

// RequestCacheNames - Requests cache names.
// securityOrigin - Security origin.
// Returns -  caches - Caches for the security origin.
func (c *CacheStorage) RequestCacheNames(securityOrigin string) ([]*CacheStorageCache, error) {
	var v CacheStorageRequestCacheNamesParams
	v.SecurityOrigin = securityOrigin
	return c.RequestCacheNamesWithParams(&v)
}

type CacheStorageRequestEntriesParams struct {
	// ID of cache to get entries from.
	CacheId string `json:"cacheId"`
	// Number of records to skip.
	SkipCount int `json:"skipCount"`
	// Number of records to fetch.
	PageSize int `json:"pageSize"`
}

// RequestEntriesWithParams - Requests data from cache.
// Returns -  cacheDataEntries - Array of object store data entries. hasMore - If true, there are more entries to fetch in the given range.
func (c *CacheStorage) RequestEntriesWithParams(v *CacheStorageRequestEntriesParams) ([]*CacheStorageDataEntry, bool, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.requestEntries", Params: v})
	if err != nil {
		return nil, false, err
	}

	var chromeData struct {
		Result struct {
			CacheDataEntries []*CacheStorageDataEntry
			HasMore          bool
		}
	}

	if resp == nil {
		return nil, false, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, false, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, false, err
	}

	return chromeData.Result.CacheDataEntries, chromeData.Result.HasMore, nil
}

// RequestEntries - Requests data from cache.
// cacheId - ID of cache to get entries from.
// skipCount - Number of records to skip.
// pageSize - Number of records to fetch.
// Returns -  cacheDataEntries - Array of object store data entries. hasMore - If true, there are more entries to fetch in the given range.
func (c *CacheStorage) RequestEntries(cacheId string, skipCount int, pageSize int) ([]*CacheStorageDataEntry, bool, error) {
	var v CacheStorageRequestEntriesParams
	v.CacheId = cacheId
	v.SkipCount = skipCount
	v.PageSize = pageSize
	return c.RequestEntriesWithParams(&v)
}

type CacheStorageDeleteCacheParams struct {
	// Id of cache for deletion.
	CacheId string `json:"cacheId"`
}

// DeleteCacheWithParams - Deletes a cache.
func (c *CacheStorage) DeleteCacheWithParams(v *CacheStorageDeleteCacheParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.deleteCache", Params: v})
}

// DeleteCache - Deletes a cache.
// cacheId - Id of cache for deletion.
func (c *CacheStorage) DeleteCache(cacheId string) (*gcdmessage.ChromeResponse, error) {
	var v CacheStorageDeleteCacheParams
	v.CacheId = cacheId
	return c.DeleteCacheWithParams(&v)
}

type CacheStorageDeleteEntryParams struct {
	// Id of cache where the entry will be deleted.
	CacheId string `json:"cacheId"`
	// URL spec of the request.
	Request string `json:"request"`
}

// DeleteEntryWithParams - Deletes a cache entry.
func (c *CacheStorage) DeleteEntryWithParams(v *CacheStorageDeleteEntryParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "CacheStorage.deleteEntry", Params: v})
}

// DeleteEntry - Deletes a cache entry.
// cacheId - Id of cache where the entry will be deleted.
// request - URL spec of the request.
func (c *CacheStorage) DeleteEntry(cacheId string, request string) (*gcdmessage.ChromeResponse, error) {
	var v CacheStorageDeleteEntryParams
	v.CacheId = cacheId
	v.Request = request
	return c.DeleteEntryWithParams(&v)
}
