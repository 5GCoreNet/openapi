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

// DnaiChangeTypeRm It can take the values  as specified for DnaiChangeType but with the OpenAPI 'nullable: true' property.
type DnaiChangeTypeRm struct {
	DnaiChangeType *DnaiChangeType
	NullValue      *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnaiChangeTypeRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DnaiChangeType
	err = json.Unmarshal(data, &dst.DnaiChangeType)
	if err == nil {
		jsonDnaiChangeType, _ := json.Marshal(dst.DnaiChangeType)
		if string(jsonDnaiChangeType) == "{}" { // empty struct
			dst.DnaiChangeType = nil
		} else {
			return nil // data stored in dst.DnaiChangeType, return on the first match
		}
	} else {
		dst.DnaiChangeType = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnaiChangeTypeRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnaiChangeTypeRm) MarshalJSON() ([]byte, error) {
	if src.DnaiChangeType != nil {
		return json.Marshal(&src.DnaiChangeType)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnaiChangeTypeRm struct {
	value *DnaiChangeTypeRm
	isSet bool
}

func (v NullableDnaiChangeTypeRm) Get() *DnaiChangeTypeRm {
	return v.value
}

func (v *NullableDnaiChangeTypeRm) Set(val *DnaiChangeTypeRm) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaiChangeTypeRm) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaiChangeTypeRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaiChangeTypeRm(val *DnaiChangeTypeRm) *NullableDnaiChangeTypeRm {
	return &NullableDnaiChangeTypeRm{value: val, isSet: true}
}

func (v NullableDnaiChangeTypeRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaiChangeTypeRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
