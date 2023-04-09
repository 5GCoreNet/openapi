/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_StreamingDataMnS

import (
	"encoding/json"
	"fmt"
)

// WebsocketHeaderConnectionType Header value for the upgrade request and response.
type WebsocketHeaderConnectionType string

// List of websocketHeaderConnection-Type
const (
	UPGRADE WebsocketHeaderConnectionType = "Upgrade"
)

// All allowed values of WebsocketHeaderConnectionType enum
var AllowedWebsocketHeaderConnectionTypeEnumValues = []WebsocketHeaderConnectionType{
	"Upgrade",
}

func (v *WebsocketHeaderConnectionType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebsocketHeaderConnectionType(value)
	for _, existing := range AllowedWebsocketHeaderConnectionTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebsocketHeaderConnectionType", value)
}

// NewWebsocketHeaderConnectionTypeFromValue returns a pointer to a valid WebsocketHeaderConnectionType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebsocketHeaderConnectionTypeFromValue(v string) (*WebsocketHeaderConnectionType, error) {
	ev := WebsocketHeaderConnectionType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebsocketHeaderConnectionType: valid values are %v", v, AllowedWebsocketHeaderConnectionTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebsocketHeaderConnectionType) IsValid() bool {
	for _, existing := range AllowedWebsocketHeaderConnectionTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to websocketHeaderConnection-Type value
func (v WebsocketHeaderConnectionType) Ptr() *WebsocketHeaderConnectionType {
	return &v
}

type NullableWebsocketHeaderConnectionType struct {
	value *WebsocketHeaderConnectionType
	isSet bool
}

func (v NullableWebsocketHeaderConnectionType) Get() *WebsocketHeaderConnectionType {
	return v.value
}

func (v *NullableWebsocketHeaderConnectionType) Set(val *WebsocketHeaderConnectionType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsocketHeaderConnectionType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsocketHeaderConnectionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsocketHeaderConnectionType(val *WebsocketHeaderConnectionType) *NullableWebsocketHeaderConnectionType {
	return &NullableWebsocketHeaderConnectionType{value: val, isSet: true}
}

func (v NullableWebsocketHeaderConnectionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsocketHeaderConnectionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
