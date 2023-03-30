/*
EES EEC Context Relocation API

API for EEC Context Relocation.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EECContextRelocation

import (
	"encoding/json"
	"fmt"
)

// EndPoint - The end point information to reach EAS.
type EndPoint struct {
	Interface{} *interface{}
}

// interface{}AsEndPoint is a convenience function that returns interface{} wrapped in EndPoint
func Interface{}AsEndPoint(v *interface{}) EndPoint {
	return EndPoint{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *EndPoint) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EndPoint)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EndPoint)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EndPoint) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EndPoint) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableEndPoint struct {
	value *EndPoint
	isSet bool
}

func (v NullableEndPoint) Get() *EndPoint {
	return v.value
}

func (v *NullableEndPoint) Set(val *EndPoint) {
	v.value = val
	v.isSet = true
}

func (v NullableEndPoint) IsSet() bool {
	return v.isSet
}

func (v *NullableEndPoint) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndPoint(val *EndPoint) *NullableEndPoint {
	return &NullableEndPoint{value: val, isSet: true}
}

func (v NullableEndPoint) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndPoint) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


