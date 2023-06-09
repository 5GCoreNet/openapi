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

// TaiRm This data type is defined in the same way as the 'Tai' data type, but with the OpenAPI 'nullable: true' property.
type TaiRm struct {
	NullValue *NullValue
	Tai       *Tai
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TaiRm) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into Tai
	err = json.Unmarshal(data, &dst.Tai)
	if err == nil {
		jsonTai, _ := json.Marshal(dst.Tai)
		if string(jsonTai) == "{}" { // empty struct
			dst.Tai = nil
		} else {
			return nil // data stored in dst.Tai, return on the first match
		}
	} else {
		dst.Tai = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TaiRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TaiRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.Tai != nil {
		return json.Marshal(&src.Tai)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTaiRm struct {
	value *TaiRm
	isSet bool
}

func (v NullableTaiRm) Get() *TaiRm {
	return v.value
}

func (v *NullableTaiRm) Set(val *TaiRm) {
	v.value = val
	v.isSet = true
}

func (v NullableTaiRm) IsSet() bool {
	return v.isSet
}

func (v *NullableTaiRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTaiRm(val *TaiRm) *NullableTaiRm {
	return &NullableTaiRm{value: val, isSet: true}
}

func (v NullableTaiRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTaiRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
