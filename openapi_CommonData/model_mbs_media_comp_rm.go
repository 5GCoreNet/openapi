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

// MbsMediaCompRm This data type is defined in the same way as the MbsMediaComp data type, but with the OpenAPI nullable property set to true.
type MbsMediaCompRm struct {
	MbsMediaComp *MbsMediaComp
	NullValue    *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsMediaCompRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MbsMediaComp
	err = json.Unmarshal(data, &dst.MbsMediaComp)
	if err == nil {
		jsonMbsMediaComp, _ := json.Marshal(dst.MbsMediaComp)
		if string(jsonMbsMediaComp) == "{}" { // empty struct
			dst.MbsMediaComp = nil
		} else {
			return nil // data stored in dst.MbsMediaComp, return on the first match
		}
	} else {
		dst.MbsMediaComp = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(MbsMediaCompRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsMediaCompRm) MarshalJSON() ([]byte, error) {
	if src.MbsMediaComp != nil {
		return json.Marshal(&src.MbsMediaComp)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsMediaCompRm struct {
	value *MbsMediaCompRm
	isSet bool
}

func (v NullableMbsMediaCompRm) Get() *MbsMediaCompRm {
	return v.value
}

func (v *NullableMbsMediaCompRm) Set(val *MbsMediaCompRm) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsMediaCompRm) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsMediaCompRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsMediaCompRm(val *MbsMediaCompRm) *NullableMbsMediaCompRm {
	return &NullableMbsMediaCompRm{value: val, isSet: true}
}

func (v NullableMbsMediaCompRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsMediaCompRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
