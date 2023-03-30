/*
Nnsacf_NSAC

Nnsacf_NSAC Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnsacf_NSAC

import (
	"encoding/json"
	"fmt"
)

// EACMode EAC mode. Possible values are - ACTIVE - DEACTIVE 
type EACMode struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EACMode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(EACMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EACMode) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEACMode struct {
	value *EACMode
	isSet bool
}

func (v NullableEACMode) Get() *EACMode {
	return v.value
}

func (v *NullableEACMode) Set(val *EACMode) {
	v.value = val
	v.isSet = true
}

func (v NullableEACMode) IsSet() bool {
	return v.isSet
}

func (v *NullableEACMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEACMode(val *EACMode) *NullableEACMode {
	return &NullableEACMode{value: val, isSet: true}
}

func (v NullableEACMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEACMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


