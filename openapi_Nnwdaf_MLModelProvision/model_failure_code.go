/*
Nnwdaf_MLModelProvision

Nnwdaf_MLModelProvision API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_MLModelProvision

import (
	"encoding/json"
	"fmt"
)

// FailureCode Possible values are - UNAVAILABLE_ML_MODEL: Indicates the requested ML model for the event is unavailable.
type FailureCode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FailureCode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(FailureCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FailureCode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFailureCode struct {
	value *FailureCode
	isSet bool
}

func (v NullableFailureCode) Get() *FailureCode {
	return v.value
}

func (v *NullableFailureCode) Set(val *FailureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureCode(val *FailureCode) *NullableFailureCode {
	return &NullableFailureCode{value: val, isSet: true}
}

func (v NullableFailureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
