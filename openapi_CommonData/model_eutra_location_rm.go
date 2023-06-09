/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// EutraLocationRm This data type is defined in the same way as the 'EutraLocation' data type, but with the OpenAPI 'nullable: true' property.
type EutraLocationRm struct {
	EutraLocation *EutraLocation
	NullValue     *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EutraLocationRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EutraLocation
	err = json.Unmarshal(data, &dst.EutraLocation)
	if err == nil {
		jsonEutraLocation, _ := json.Marshal(dst.EutraLocation)
		if string(jsonEutraLocation) == "{}" { // empty struct
			dst.EutraLocation = nil
		} else {
			return nil // data stored in dst.EutraLocation, return on the first match
		}
	} else {
		dst.EutraLocation = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue)
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EutraLocationRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EutraLocationRm) MarshalJSON() ([]byte, error) {
	if src.EutraLocation != nil {
		return json.Marshal(&src.EutraLocation)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEutraLocationRm struct {
	value *EutraLocationRm
	isSet bool
}

func (v NullableEutraLocationRm) Get() *EutraLocationRm {
	return v.value
}

func (v *NullableEutraLocationRm) Set(val *EutraLocationRm) {
	v.value = val
	v.isSet = true
}

func (v NullableEutraLocationRm) IsSet() bool {
	return v.isSet
}

func (v *NullableEutraLocationRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEutraLocationRm(val *EutraLocationRm) *NullableEutraLocationRm {
	return &NullableEutraLocationRm{value: val, isSet: true}
}

func (v NullableEutraLocationRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEutraLocationRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
