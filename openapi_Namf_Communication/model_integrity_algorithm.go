/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// IntegrityAlgorithm Indicates the supported Integrity Algorithm
type IntegrityAlgorithm struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IntegrityAlgorithm) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(IntegrityAlgorithm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IntegrityAlgorithm) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIntegrityAlgorithm struct {
	value *IntegrityAlgorithm
	isSet bool
}

func (v NullableIntegrityAlgorithm) Get() *IntegrityAlgorithm {
	return v.value
}

func (v *NullableIntegrityAlgorithm) Set(val *IntegrityAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegrityAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegrityAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegrityAlgorithm(val *IntegrityAlgorithm) *NullableIntegrityAlgorithm {
	return &NullableIntegrityAlgorithm{value: val, isSet: true}
}

func (v NullableIntegrityAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegrityAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
