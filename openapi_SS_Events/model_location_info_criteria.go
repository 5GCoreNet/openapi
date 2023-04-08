/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
	"fmt"
)

// LocationInfoCriteria - Geographic location and reference UE details, where the UEs moving in and out to be monitored. 
type LocationInfoCriteria struct {
	Interface *interface{}
}

// interface{}AsLocationInfoCriteria is a convenience function that returns interface{} wrapped in LocationInfoCriteria
func InterfaceAsLocationInfoCriteria(v *interface{}) LocationInfoCriteria {
	return LocationInfoCriteria{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *LocationInfoCriteria) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(LocationInfoCriteria)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(LocationInfoCriteria)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src LocationInfoCriteria) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *LocationInfoCriteria) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableLocationInfoCriteria struct {
	value *LocationInfoCriteria
	isSet bool
}

func (v NullableLocationInfoCriteria) Get() *LocationInfoCriteria {
	return v.value
}

func (v *NullableLocationInfoCriteria) Set(val *LocationInfoCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationInfoCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationInfoCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationInfoCriteria(val *LocationInfoCriteria) *NullableLocationInfoCriteria {
	return &NullableLocationInfoCriteria{value: val, isSet: true}
}

func (v NullableLocationInfoCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationInfoCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


