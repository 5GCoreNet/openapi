/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// PlmnRange - Range of PLMN IDs
type PlmnRange struct {
	Interface *interface{}
}

// interface{}AsPlmnRange is a convenience function that returns interface{} wrapped in PlmnRange
func InterfaceAsPlmnRange(v *interface{}) PlmnRange {
	return PlmnRange{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PlmnRange) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(PlmnRange)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PlmnRange)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PlmnRange) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PlmnRange) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullablePlmnRange struct {
	value *PlmnRange
	isSet bool
}

func (v NullablePlmnRange) Get() *PlmnRange {
	return v.value
}

func (v *NullablePlmnRange) Set(val *PlmnRange) {
	v.value = val
	v.isSet = true
}

func (v NullablePlmnRange) IsSet() bool {
	return v.isSet
}

func (v *NullablePlmnRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlmnRange(val *PlmnRange) *NullablePlmnRange {
	return &NullablePlmnRange{value: val, isSet: true}
}

func (v NullablePlmnRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlmnRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
