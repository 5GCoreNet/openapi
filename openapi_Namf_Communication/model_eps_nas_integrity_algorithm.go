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

// EpsNasIntegrityAlgorithm Indicates the supported EPS NAS Integrity Algorithm
type EpsNasIntegrityAlgorithm struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EpsNasIntegrityAlgorithm) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(EpsNasIntegrityAlgorithm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EpsNasIntegrityAlgorithm) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEpsNasIntegrityAlgorithm struct {
	value *EpsNasIntegrityAlgorithm
	isSet bool
}

func (v NullableEpsNasIntegrityAlgorithm) Get() *EpsNasIntegrityAlgorithm {
	return v.value
}

func (v *NullableEpsNasIntegrityAlgorithm) Set(val *EpsNasIntegrityAlgorithm) {
	v.value = val
	v.isSet = true
}

func (v NullableEpsNasIntegrityAlgorithm) IsSet() bool {
	return v.isSet
}

func (v *NullableEpsNasIntegrityAlgorithm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpsNasIntegrityAlgorithm(val *EpsNasIntegrityAlgorithm) *NullableEpsNasIntegrityAlgorithm {
	return &NullableEpsNasIntegrityAlgorithm{value: val, isSet: true}
}

func (v NullableEpsNasIntegrityAlgorithm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpsNasIntegrityAlgorithm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
