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

// EpsNasCipheringAlgorithm Indicates the supported EPS NAS Ciphering Algorithm
type EpsNasCipheringAlgorithm struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EpsNasCipheringAlgorithm) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(EpsNasCipheringAlgorithm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EpsNasCipheringAlgorithm) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEpsNasCipheringAlgorithm struct {
	value *EpsNasCipheringAlgorithm
	isSet bool
}

func (v NullableEpsNasCipheringAlgorithm) Get() *EpsNasCipheringAlgorithm {
	return v.value
}

func (v *NullableEpsNasCipheringAlgorithm) Set(val *EpsNasCipheringAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableEpsNasCipheringAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableEpsNasCipheringAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpsNasCipheringAlgorithm(val *EpsNasCipheringAlgorithm) *NullableEpsNasCipheringAlgorithm {
	return &NullableEpsNasCipheringAlgorithm{value: val, isSet: true}
}

func (v NullableEpsNasCipheringAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpsNasCipheringAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
