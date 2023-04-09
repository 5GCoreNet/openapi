/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// RmState Describes the registration management state of a UE
type RmState struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RmState) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RmState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RmState) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
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
