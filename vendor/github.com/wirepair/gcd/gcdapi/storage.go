// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Storage functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Usage for a storage type.
type StorageUsageForType struct {
	StorageType string  `json:"storageType"` // Name of storage type. enum values: appcache, cookies, file_systems, indexeddb, local_storage, shader_cache, websql, service_workers, cache_storage, all, other
	Usage       float64 `json:"usage"`       // Storage usage (bytes).
}

// A cache's contents have been modified.
type StorageCacheStorageContentUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin    string `json:"origin"`    // Origin to update.
		CacheName string `json:"cacheName"` // Name of cache in origin.
	} `json:"Params,omitempty"`
}

// A cache has been added/deleted.
type StorageCacheStorageListUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin string `json:"origin"` // Origin to update.
	} `json:"Params,omitempty"`
}

// The origin's IndexedDB object store has been modified.
type StorageIndexedDBContentUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin          string `json:"origin"`          // Origin to update.
		DatabaseName    string `json:"databaseName"`    // Database to update.
		ObjectStoreName string `json:"objectStoreName"` // ObjectStore to update.
	} `json:"Params,omitempty"`
}

// The origin's IndexedDB database list has been modified.
type StorageIndexedDBListUpdatedEvent struct {
	Method string `json:"method"`
	Params struct {
		Origin string `json:"origin"` // Origin to update.
	} `json:"Params,omitempty"`
}

type Storage struct {
	target gcdmessage.ChromeTargeter
}

func NewStorage(target gcdmessage.ChromeTargeter) *Storage {
	c := &Storage{target: target}
	return c
}

type StorageClearDataForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
	// Comma separated origin names.
	StorageTypes string `json:"storageTypes"`
}

// ClearDataForOriginWithParams - Clears storage for origin.
func (c *Storage) ClearDataForOriginWithParams(v *StorageClearDataForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.clearDataForOrigin", Params: v})
}

// ClearDataForOrigin - Clears storage for origin.
// origin - Security origin.
// storageTypes - Comma separated origin names.
func (c *Storage) ClearDataForOrigin(origin string, storageTypes string) (*gcdmessage.ChromeResponse, error) {
	var v StorageClearDataForOriginParams
	v.Origin = origin
	v.StorageTypes = storageTypes
	return c.ClearDataForOriginWithParams(&v)
}

type StorageGetUsageAndQuotaParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// GetUsageAndQuotaWithParams - Returns usage and quota in bytes.
// Returns -  usage - Storage usage (bytes). quota - Storage quota (bytes). usageBreakdown - Storage usage per type (bytes).
func (c *Storage) GetUsageAndQuotaWithParams(v *StorageGetUsageAndQuotaParams) (float64, float64, []*StorageUsageForType, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.getUsageAndQuota", Params: v})
	if err != nil {
		return 0, 0, nil, err
	}

	var chromeData struct {
		Result struct {
			Usage          float64
			Quota          float64
			UsageBreakdown []*StorageUsageForType
		}
	}

	if resp == nil {
		return 0, 0, nil, &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return 0, 0, nil, &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return 0, 0, nil, err
	}

	return chromeData.Result.Usage, chromeData.Result.Quota, chromeData.Result.UsageBreakdown, nil
}

// GetUsageAndQuota - Returns usage and quota in bytes.
// origin - Security origin.
// Returns -  usage - Storage usage (bytes). quota - Storage quota (bytes). usageBreakdown - Storage usage per type (bytes).
func (c *Storage) GetUsageAndQuota(origin string) (float64, float64, []*StorageUsageForType, error) {
	var v StorageGetUsageAndQuotaParams
	v.Origin = origin
	return c.GetUsageAndQuotaWithParams(&v)
}

type StorageTrackCacheStorageForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// TrackCacheStorageForOriginWithParams - Registers origin to be notified when an update occurs to its cache storage list.
func (c *Storage) TrackCacheStorageForOriginWithParams(v *StorageTrackCacheStorageForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.trackCacheStorageForOrigin", Params: v})
}

// TrackCacheStorageForOrigin - Registers origin to be notified when an update occurs to its cache storage list.
// origin - Security origin.
func (c *Storage) TrackCacheStorageForOrigin(origin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageTrackCacheStorageForOriginParams
	v.Origin = origin
	return c.TrackCacheStorageForOriginWithParams(&v)
}

type StorageTrackIndexedDBForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// TrackIndexedDBForOriginWithParams - Registers origin to be notified when an update occurs to its IndexedDB.
func (c *Storage) TrackIndexedDBForOriginWithParams(v *StorageTrackIndexedDBForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.trackIndexedDBForOrigin", Params: v})
}

// TrackIndexedDBForOrigin - Registers origin to be notified when an update occurs to its IndexedDB.
// origin - Security origin.
func (c *Storage) TrackIndexedDBForOrigin(origin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageTrackIndexedDBForOriginParams
	v.Origin = origin
	return c.TrackIndexedDBForOriginWithParams(&v)
}

type StorageUntrackCacheStorageForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// UntrackCacheStorageForOriginWithParams - Unregisters origin from receiving notifications for cache storage.
func (c *Storage) UntrackCacheStorageForOriginWithParams(v *StorageUntrackCacheStorageForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.untrackCacheStorageForOrigin", Params: v})
}

// UntrackCacheStorageForOrigin - Unregisters origin from receiving notifications for cache storage.
// origin - Security origin.
func (c *Storage) UntrackCacheStorageForOrigin(origin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageUntrackCacheStorageForOriginParams
	v.Origin = origin
	return c.UntrackCacheStorageForOriginWithParams(&v)
}

type StorageUntrackIndexedDBForOriginParams struct {
	// Security origin.
	Origin string `json:"origin"`
}

// UntrackIndexedDBForOriginWithParams - Unregisters origin from receiving notifications for IndexedDB.
func (c *Storage) UntrackIndexedDBForOriginWithParams(v *StorageUntrackIndexedDBForOriginParams) (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Storage.untrackIndexedDBForOrigin", Params: v})
}

// UntrackIndexedDBForOrigin - Unregisters origin from receiving notifications for IndexedDB.
// origin - Security origin.
func (c *Storage) UntrackIndexedDBForOrigin(origin string) (*gcdmessage.ChromeResponse, error) {
	var v StorageUntrackIndexedDBForOriginParams
	v.Origin = origin
	return c.UntrackIndexedDBForOriginWithParams(&v)
}
