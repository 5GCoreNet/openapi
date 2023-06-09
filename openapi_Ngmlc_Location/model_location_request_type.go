/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// LocationRequestType NI-LR, MT-LR or MO-LR
type LocationRequestType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LocationRequestType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(LocationRequestType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LocationRequestType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocationRequestType struct {
	value *LocationRequestType
	isSet bool
}

func (v NullableLocationRequestType) Get() *LocationRequestType {
	return v.value
}

func (v *NullableLocationRequestType) Set(val *LocationRequestType) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRequestType) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRequestType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRequestType(val *LocationRequestType) *NullableLocationRequestType {
	return &NullableLocationRequestType{value: val, isSet: true}
}

func (v NullableLocationRequestType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRequestType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
