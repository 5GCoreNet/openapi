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

// SliceMbrRm SliceMbr with nullable: true
type SliceMbrRm struct {
	NullValue *NullValue
	SliceMbr  *SliceMbr
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SliceMbrRm) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into SliceMbr
	err = json.Unmarshal(data, &dst.SliceMbr)
	if err == nil {
		jsonSliceMbr, _ := json.Marshal(dst.SliceMbr)
		if string(jsonSliceMbr) == "{}" { // empty struct
			dst.SliceMbr = nil
		} else {
			return nil // data stored in dst.SliceMbr, return on the first match
		}
	} else {
		dst.SliceMbr = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SliceMbrRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SliceMbrRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.SliceMbr != nil {
		return json.Marshal(&src.SliceMbr)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSliceMbrRm struct {
	value *SliceMbrRm
	isSet bool
}

func (v NullableSliceMbrRm) Get() *SliceMbrRm {
	return v.value
}

func (v *NullableSliceMbrRm) Set(val *SliceMbrRm) {
	v.value = val
	v.isSet = true
}

func (v NullableSliceMbrRm) IsSet() bool {
	return v.isSet
}

func (v *NullableSliceMbrRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSliceMbrRm(val *SliceMbrRm) *NullableSliceMbrRm {
	return &NullableSliceMbrRm{value: val, isSet: true}
}

func (v NullableSliceMbrRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSliceMbrRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
