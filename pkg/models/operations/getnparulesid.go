// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"encoding/json"
	"fmt"
	"github.com/speakeasy-sdks/ns_npa_go_sdk/pkg/models/shared"
	"net/http"
)

type GetNpaRulesIDRequest struct {
	// Return values only from specified fields
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// npa policy id
	ID int `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetNpaRulesIDRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetNpaRulesIDRequest) GetID() int {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetNpaRulesID200ApplicationJSONStatus string

const (
	GetNpaRulesID200ApplicationJSONStatusSuccess GetNpaRulesID200ApplicationJSONStatus = "success"
	GetNpaRulesID200ApplicationJSONStatusError   GetNpaRulesID200ApplicationJSONStatus = "error"
)

func (e GetNpaRulesID200ApplicationJSONStatus) ToPointer() *GetNpaRulesID200ApplicationJSONStatus {
	return &e
}

func (e *GetNpaRulesID200ApplicationJSONStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "success":
		fallthrough
	case "error":
		*e = GetNpaRulesID200ApplicationJSONStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GetNpaRulesID200ApplicationJSONStatus: %v", v)
	}
}

// GetNpaRulesID200ApplicationJSON - successful operation
type GetNpaRulesID200ApplicationJSON struct {
	Data   *shared.NpaPolicyResponseItem          `json:"data,omitempty"`
	Status *GetNpaRulesID200ApplicationJSONStatus `json:"status,omitempty"`
}

func (o *GetNpaRulesID200ApplicationJSON) GetData() *shared.NpaPolicyResponseItem {
	if o == nil {
		return nil
	}
	return o.Data
}

func (o *GetNpaRulesID200ApplicationJSON) GetStatus() *GetNpaRulesID200ApplicationJSONStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type GetNpaRulesIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// successful operation
	GetNpaRulesID200ApplicationJSONObject *GetNpaRulesID200ApplicationJSON
}

func (o *GetNpaRulesIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetNpaRulesIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetNpaRulesIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetNpaRulesIDResponse) GetGetNpaRulesID200ApplicationJSONObject() *GetNpaRulesID200ApplicationJSON {
	if o == nil {
		return nil
	}
	return o.GetNpaRulesID200ApplicationJSONObject
}
