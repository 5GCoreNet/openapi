/*
Nnef_EventExposure

NEF Event Exposure Service.   © 2022 , 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_EventExposure

import (
	"encoding/json"
	"fmt"
)

// NefEvent Represents Network Exposure Events.
type NefEvent struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NefEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NefEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NefEvent) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
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


