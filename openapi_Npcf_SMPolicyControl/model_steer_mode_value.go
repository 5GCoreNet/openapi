/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// SteerModeValue Indicates the steering mode value determined by the PCF.
type SteerModeValue struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SteerModeValue) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(SteerModeValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SteerModeValue) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSteerModeValue struct {
	value *SteerModeValue
	isSet bool
}

func (v NullableSteerModeValue) Get() *SteerModeValue {
	return v.value
}

func (v *NullableSteerModeValue) Set(val *SteerModeValue) {
	v.value = val
	v.isSet = true
}

func (v NullableSteerModeValue) IsSet() bool {
	return v.isSet
}

func (v *NullableSteerModeValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSteerModeValue(val *SteerModeValue) *NullableSteerModeValue {
	return &NullableSteerModeValue{value: val, isSet: true}
}

func (v NullableSteerModeValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSteerModeValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


