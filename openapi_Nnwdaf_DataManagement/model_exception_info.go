/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// ExceptionInfo - Represents the exceptions information provided by the AF.
type ExceptionInfo struct {
	Interface *interface{}
}

// interface{}AsExceptionInfo is a convenience function that returns interface{} wrapped in ExceptionInfo
func InterfaceAsExceptionInfo(v *interface{}) ExceptionInfo {
	return ExceptionInfo{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExceptionInfo) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ExceptionInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ExceptionInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExceptionInfo) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ExceptionInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableExceptionInfo struct {
	value *ExceptionInfo
	isSet bool
}

func (v NullableExceptionInfo) Get() *ExceptionInfo {
	return v.value
}

func (v *NullableExceptionInfo) Set(val *ExceptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionInfo(val *ExceptionInfo) *NullableExceptionInfo {
	return &NullableExceptionInfo{value: val, isSet: true}
}

func (v NullableExceptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


