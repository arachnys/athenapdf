// AUTO-GENERATED Chrome Remote Debugger Protocol API Client
// This file contains Database functionality.
// API Version: 1.3

package gcdapi

import (
	"encoding/json"
	"github.com/wirepair/gcd/gcdmessage"
)

// Database object.
type DatabaseDatabase struct {
	Id      string `json:"id"`      // Database ID.
	Domain  string `json:"domain"`  // Database domain.
	Name    string `json:"name"`    // Database name.
	Version string `json:"version"` // Database version.
}

// Database error.
type DatabaseError struct {
	Message string `json:"message"` // Error message.
	Code    int    `json:"code"`    // Error code.
}

//
type DatabaseAddDatabaseEvent struct {
	Method string `json:"method"`
	Params struct {
		Database *DatabaseDatabase `json:"database"` //
	} `json:"Params,omitempty"`
}

type Database struct {
	target gcdmessage.ChromeTargeter
}

func NewDatabase(target gcdmessage.ChromeTargeter) *Database {
	c := &Database{target: target}
	return c
}

// Disables database tracking, prevents database events from being sent to the client.
func (c *Database) Disable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.disable"})
}

// Enables database tracking, database events will now be delivered to the client.
func (c *Database) Enable() (*gcdmessage.ChromeResponse, error) {
	return gcdmessage.SendDefaultRequest(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.enable"})
}

type DatabaseExecuteSQLParams struct {
	//
	DatabaseId string `json:"databaseId"`
	//
	Query string `json:"query"`
}

// ExecuteSQLWithParams -
// Returns -  columnNames -  values -  sqlError -
func (c *Database) ExecuteSQLWithParams(v *DatabaseExecuteSQLParams) ([]string, []interface{}, *DatabaseError, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.executeSQL", Params: v})
	if err != nil {
		return nil, nil, nil, err
	}

	var chromeData struct {
		Result struct {
			ColumnNames []string
			Values      []interface{}
			SqlError    *DatabaseError
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

	return chromeData.Result.ColumnNames, chromeData.Result.Values, chromeData.Result.SqlError, nil
}

// ExecuteSQL -
// databaseId -
// query -
// Returns -  columnNames -  values -  sqlError -
func (c *Database) ExecuteSQL(databaseId string, query string) ([]string, []interface{}, *DatabaseError, error) {
	var v DatabaseExecuteSQLParams
	v.DatabaseId = databaseId
	v.Query = query
	return c.ExecuteSQLWithParams(&v)
}

type DatabaseGetDatabaseTableNamesParams struct {
	//
	DatabaseId string `json:"databaseId"`
}

// GetDatabaseTableNamesWithParams -
// Returns -  tableNames -
func (c *Database) GetDatabaseTableNamesWithParams(v *DatabaseGetDatabaseTableNamesParams) ([]string, error) {
	resp, err := gcdmessage.SendCustomReturn(c.target, c.target.GetSendCh(), &gcdmessage.ParamRequest{Id: c.target.GetId(), Method: "Database.getDatabaseTableNames", Params: v})
	if err != nil {
		return nil, err
	}

	var chromeData struct {
		Result struct {
			TableNames []string
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

	return chromeData.Result.TableNames, nil
}

// GetDatabaseTableNames -
// databaseId -
// Returns -  tableNames -
func (c *Database) GetDatabaseTableNames(databaseId string) ([]string, error) {
	var v DatabaseGetDatabaseTableNamesParams
	v.DatabaseId = databaseId
	return c.GetDatabaseTableNamesWithParams(&v)
}
