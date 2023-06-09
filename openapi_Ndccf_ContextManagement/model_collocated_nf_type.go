/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// CollocatedNfType NF types for a collocated NF
type CollocatedNfType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CollocatedNfType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CollocatedNfType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CollocatedNfType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCollocatedNfType struct {
	value *CollocatedNfType
	isSet bool
}

func (v NullableCollocatedNfType) Get() *CollocatedNfType {
	return v.value
}

func (v *NullableCollocatedNfType) Set(val *CollocatedNfType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollocatedNfType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollocatedNfType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollocatedNfType(val *CollocatedNfType) *NullableCollocatedNfType {
	return &NullableCollocatedNfType{value: val, isSet: true}
}

func (v NullableCollocatedNfType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollocatedNfType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
