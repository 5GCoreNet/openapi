/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// Area - Provides area information.
type Area struct {
	Interface *interface{}
}

// interface{}AsArea is a convenience function that returns interface{} wrapped in Area
func InterfaceAsArea(v *interface{}) Area {
	return Area{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *Area) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(Area)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(Area)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src Area) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *Area) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableArea struct {
	value *Area
	isSet bool
}

func (v NullableArea) Get() *Area {
	return v.value
}

func (v *NullableArea) Set(val *Area) {
	v.value = val
	v.isSet = true
}

func (v NullableArea) IsSet() bool {
	return v.isSet
}

func (v *NullableArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArea(val *Area) *NullableArea {
	return &NullableArea{value: val, isSet: true}
}

func (v NullableArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
