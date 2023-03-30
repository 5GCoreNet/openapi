/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// RmState Describes the registration management state of a UE
type RmState struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RmState) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RmState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RmState) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRmState struct {
	value *RmState
	isSet bool
}

func (v NullableRmState) Get() *RmState {
	return v.value
}

func (v *NullableRmState) Set(val *RmState) {
	v.value = val
	v.isSet = true
}

func (v NullableRmState) IsSet() bool {
	return v.isSet
}

func (v *NullableRmState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRmState(val *RmState) *NullableRmState {
	return &NullableRmState{value: val, isSet: true}
}

func (v NullableRmState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRmState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


