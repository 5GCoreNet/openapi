/*
Eees_EASDiscovery

API for EAS Discovery. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EASDiscovery

import (
	"encoding/json"
	"fmt"
)

// RequestorId - Represents identifier of the requestor.
type RequestorId struct {
	Interface *interface{}
}

// interface{}AsRequestorId is a convenience function that returns interface{} wrapped in RequestorId
func InterfaceAsRequestorId(v *interface{}) RequestorId {
	return RequestorId{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *RequestorId) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(RequestorId)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(RequestorId)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src RequestorId) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *RequestorId) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableRequestorId struct {
	value *RequestorId
	isSet bool
}

func (v NullableRequestorId) Get() *RequestorId {
	return v.value
}

func (v *NullableRequestorId) Set(val *RequestorId) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestorId) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestorId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestorId(val *RequestorId) *NullableRequestorId {
	return &NullableRequestorId{value: val, isSet: true}
}

func (v NullableRequestorId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestorId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


