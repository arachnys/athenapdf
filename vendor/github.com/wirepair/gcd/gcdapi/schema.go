// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Schema functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Description of the protocol domain.
type SchemaDomain struct {
	Name    string `json:"name"`    // Domain name.
	Version string `json:"version"` // Domain version.
}

type Schema struct {
	target gcdmessage.ChromeTargeter
}

func NewSchema(target gcdmessage.ChromeTargeter) *Schema {
	c := &Schema{target: target}
	return c
}

// GetDomains - Returns supported domains.
// Returns -  domains - List of supported domains.
func (c *Schema) GetDomains() ([]*SchemaDomain, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Schema.getDomains"})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			Domains []*SchemaDomain
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

	return chromeData.Result.Domains, nil
}
