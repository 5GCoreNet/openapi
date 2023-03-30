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

// OperationMode struct for OperationMode
type OperationMode struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OperationMode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(OperationMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *OperationMode) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOperationMode struct {
	value *OperationMode
	isSet bool
}

func (v NullableOperationMode) Get() *OperationMode {
	return v.value
}

func (v *NullableOperationMode) Set(val *OperationMode) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationMode) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationMode(val *OperationMode) *NullableOperationMode {
	return &NullableOperationMode{value: val, isSet: true}
}

func (v NullableOperationMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


