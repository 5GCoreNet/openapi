/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// AccessType Represents the type of access (3GPP or non-3GPP)
type AccessType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccessType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessType struct {
	value *AccessType
	isSet bool
}

func (v NullableAccessType) Get() *AccessType {
	return v.value
}

func (v *NullableAccessType) Set(val *AccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessType(val *AccessType) *NullableAccessType {
	return &NullableAccessType{value: val, isSet: true}
}

func (v NullableAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
