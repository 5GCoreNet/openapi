/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NefEvent Represents Network Exposure Events.
type NefEvent struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NefEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NefEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NefEvent) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNefEvent struct {
	value *NefEvent
	isSet bool
}

func (v NullableNefEvent) Get() *NefEvent {
	return v.value
}

func (v *NullableNefEvent) Set(val *NefEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNefEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNefEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNefEvent(val *NefEvent) *NullableNefEvent {
	return &NullableNefEvent{value: val, isSet: true}
}

func (v NullableNefEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNefEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
