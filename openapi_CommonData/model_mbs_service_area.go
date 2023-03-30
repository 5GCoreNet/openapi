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

// MbsServiceArea MBS Service Area
type MbsServiceArea struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsServiceArea) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MbsServiceArea)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsServiceArea) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsServiceArea struct {
	value *MbsServiceArea
	isSet bool
}

func (v NullableMbsServiceArea) Get() *MbsServiceArea {
	return v.value
}

func (v *NullableMbsServiceArea) Set(val *MbsServiceArea) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsServiceArea) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsServiceArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsServiceArea(val *MbsServiceArea) *NullableMbsServiceArea {
	return &NullableMbsServiceArea{value: val, isSet: true}
}

func (v NullableMbsServiceArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsServiceArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


