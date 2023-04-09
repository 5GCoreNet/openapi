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

// TnapIdRm This data type is defined in the same way as the 'TnapId' data type, but with the OpenAPI 'nullable: true' property.
type TnapIdRm struct {
	NullValue *NullValue
	TnapId    *TnapId
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TnapIdRm) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into TnapId
	err = json.Unmarshal(data, &dst.TnapId)
	if err == nil {
		jsonTnapId, _ := json.Marshal(dst.TnapId)
		if string(jsonTnapId) == "{}" { // empty struct
			dst.TnapId = nil
		} else {
			return nil // data stored in dst.TnapId, return on the first match
		}
	} else {
		dst.TnapId = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TnapIdRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TnapIdRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.TnapId != nil {
		return json.Marshal(&src.TnapId)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTnapIdRm struct {
	value *TnapIdRm
	isSet bool
}

func (v NullableTnapIdRm) Get() *TnapIdRm {
	return v.value
}

func (v *NullableTnapIdRm) Set(val *TnapIdRm) {
	v.value = val
	v.isSet = true
}

func (v NullableTnapIdRm) IsSet() bool {
	return v.isSet
}

func (v *NullableTnapIdRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTnapIdRm(val *TnapIdRm) *NullableTnapIdRm {
	return &NullableTnapIdRm{value: val, isSet: true}
}

func (v NullableTnapIdRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTnapIdRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
