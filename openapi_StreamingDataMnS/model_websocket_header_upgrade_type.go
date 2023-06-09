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

// WebsocketHeaderUpgradeType Header value for the upgrade to WebSocket request and response.
type WebsocketHeaderUpgradeType string

// List of websocketHeaderUpgrade-Type
const (
	WEBSOCKET WebsocketHeaderUpgradeType = "websocket"
)

// All allowed values of WebsocketHeaderUpgradeType enum
var AllowedWebsocketHeaderUpgradeTypeEnumValues = []WebsocketHeaderUpgradeType{
	"websocket",
}

func (v *WebsocketHeaderUpgradeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WebsocketHeaderUpgradeType(value)
	for _, existing := range AllowedWebsocketHeaderUpgradeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WebsocketHeaderUpgradeType", value)
}

// NewWebsocketHeaderUpgradeTypeFromValue returns a pointer to a valid WebsocketHeaderUpgradeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWebsocketHeaderUpgradeTypeFromValue(v string) (*WebsocketHeaderUpgradeType, error) {
	ev := WebsocketHeaderUpgradeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WebsocketHeaderUpgradeType: valid values are %v", v, AllowedWebsocketHeaderUpgradeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WebsocketHeaderUpgradeType) IsValid() bool {
	for _, existing := range AllowedWebsocketHeaderUpgradeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to websocketHeaderUpgrade-Type value
func (v WebsocketHeaderUpgradeType) Ptr() *WebsocketHeaderUpgradeType {
	return &v
}

type NullableWebsocketHeaderUpgradeType struct {
	value *WebsocketHeaderUpgradeType
	isSet bool
}

func (v NullableWebsocketHeaderUpgradeType) Get() *WebsocketHeaderUpgradeType {
	return v.value
}

func (v *NullableWebsocketHeaderUpgradeType) Set(val *WebsocketHeaderUpgradeType) {
	v.value = val
	v.isSet = true
}

func (v NullableWebsocketHeaderUpgradeType) IsSet() bool {
	return v.isSet
}

func (v *NullableWebsocketHeaderUpgradeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebsocketHeaderUpgradeType(val *WebsocketHeaderUpgradeType) *NullableWebsocketHeaderUpgradeType {
	return &NullableWebsocketHeaderUpgradeType{value: val, isSet: true}
}

func (v NullableWebsocketHeaderUpgradeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebsocketHeaderUpgradeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
