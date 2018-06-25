// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains SystemInfo functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Describes a single graphics processor (GPU).
type SystemInfoGPUDevice struct {
	VendorId     float64 `json:"vendorId"`     // PCI ID of the GPU vendor, if available; 0 otherwise.
	DeviceId     float64 `json:"deviceId"`     // PCI ID of the GPU device, if available; 0 otherwise.
	VendorString string  `json:"vendorString"` // String description of the GPU vendor, if the PCI ID is not available.
	DeviceString string  `json:"deviceString"` // String description of the GPU device, if the PCI ID is not available.
}

// Provides information about the GPU(s) on the system.
type SystemInfoGPUInfo struct {
	Devices              []*SystemInfoGPUDevice `json:"devices"`                 // The graphics devices on the system. Element 0 is the primary GPU.
	AuxAttributes        map[string]interface{} `json:"auxAttributes,omitempty"` // An optional dictionary of additional GPU related attributes.
	FeatureStatus        map[string]interface{} `json:"featureStatus,omitempty"` // An optional dictionary of graphics features and their status.
	DriverBugWorkarounds []string               `json:"driverBugWorkarounds"`    // An optional array of GPU driver bug workarounds.
}

type SystemInfo struct {
	target gcdmessage.ChromeTargeter
}

func NewSystemInfo(target gcdmessage.ChromeTargeter) *SystemInfo {
	c := &SystemInfo{target: target}
	return c
}

// GetInfo - Returns information about the system.
// Returns -  gpu - Information about the GPUs on the system. modelName - A platform-dependent description of the model of the machine. On Mac OS, this is, for example, 'MacBookPro'. Will be the empty string if not supported. modelVersion - A platform-dependent description of the version of the machine. On Mac OS, this is, for example, '10.1'. Will be the empty string if not supported. commandLine - The command line string used to launch the browser. Will be the empty string if not supported.
func (c *SystemInfo) GetInfo() (*SystemInfoGPUInfo, string, string, string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "SystemInfo.getInfo"})
	if err != nil {
		return nil, "", "", "", err
	}

	var chromeData struct {
		Result struct {
			Gpu          *SystemInfoGPUInfo
			ModelName    string
			ModelVersion string
			CommandLine  string
		}
	}

	if resp == nil {
		return nil, "", "", "", &gcdmessage.ChromeEmptyResponseErr{}
	}

	// test if error first
	cerr := &gcdmessage.ChromeErrorResponse{}
	json.Unmarshal(resp.Data, cerr)
	if cerr != nil && cerr.Error != nil {
		return nil, "", "", "", &gcdmessage.ChromeRequestErr{Resp: cerr}
	}

	if err := json.Unmarshal(resp.Data, &chromeData); err != nil {
		return nil, "", "", "", err
	}

	return chromeData.Result.Gpu, chromeData.Result.ModelName, chromeData.Result.ModelVersion, chromeData.Result.CommandLine, nil
}
