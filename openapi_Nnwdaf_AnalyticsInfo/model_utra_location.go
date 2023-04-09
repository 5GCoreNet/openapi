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

// UtraLocation - Exactly one of cgi, sai or lai shall be present.
type UtraLocation struct {
	Interface *interface{}
}

// interface{}AsUtraLocation is a convenience function that returns interface{} wrapped in UtraLocation
func InterfaceAsUtraLocation(v *interface{}) UtraLocation {
	return UtraLocation{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *UtraLocation) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(UtraLocation)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UtraLocation)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UtraLocation) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UtraLocation) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableUtraLocation struct {
	value *UtraLocation
	isSet bool
}

func (v NullableUtraLocation) Get() *UtraLocation {
	return v.value
}

func (v *NullableUtraLocation) Set(val *UtraLocation) {
	v.value = val
	v.isSet = true
}

func (v NullableUtraLocation) IsSet() bool {
	return v.isSet
}

func (v *NullableUtraLocation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUtraLocation(val *UtraLocation) *NullableUtraLocation {
	return &NullableUtraLocation{value: val, isSet: true}
}

func (v NullableUtraLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUtraLocation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
