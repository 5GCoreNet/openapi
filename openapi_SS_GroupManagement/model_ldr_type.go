/*
SS_GroupManagement

API for SEAL Group management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_GroupManagement

import (
	"encoding/json"
	"fmt"
)

// LdrType Indicates LDR types.
type LdrType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LdrType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(LdrType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LdrType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLdrType struct {
	value *LdrType
	isSet bool
}

func (v NullableLdrType) Get() *LdrType {
	return v.value
}

func (v *NullableLdrType) Set(val *LdrType) {
	v.value = val
	v.isSet = true
}

func (v NullableLdrType) IsSet() bool {
	return v.isSet
}

func (v *NullableLdrType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdrType(val *LdrType) *NullableLdrType {
	return &NullableLdrType{value: val, isSet: true}
}

func (v NullableLdrType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdrType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


