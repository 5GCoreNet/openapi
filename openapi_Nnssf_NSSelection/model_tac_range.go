/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSelection

import (
	"encoding/json"
	"fmt"
)

// TacRange - Range of TACs (Tracking Area Codes)
type TacRange struct {
	Interface *interface{}
}

// interface{}AsTacRange is a convenience function that returns interface{} wrapped in TacRange
func InterfaceAsTacRange(v *interface{}) TacRange {
	return TacRange{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *TacRange) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface
	err = newStrictDecoder(data).Decode(&dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			match++
		}
	} else {
		dst.Interface = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(TacRange)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(TacRange)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src TacRange) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *TacRange) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableTacRange struct {
	value *TacRange
	isSet bool
}

func (v NullableTacRange) Get() *TacRange {
	return v.value
}

func (v *NullableTacRange) Set(val *TacRange) {
	v.value = val
	v.isSet = true
}

func (v NullableTacRange) IsSet() bool {
	return v.isSet
}

func (v *NullableTacRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTacRange(val *TacRange) *NullableTacRange {
	return &NullableTacRange{value: val, isSet: true}
}

func (v NullableTacRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTacRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


