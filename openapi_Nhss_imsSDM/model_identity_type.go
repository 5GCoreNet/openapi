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

// IdentityType Represents the type of IMS Public Identity
type IdentityType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IdentityType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IdentityType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IdentityType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIdentityType struct {
	value *IdentityType
	isSet bool
}

func (v NullableIdentityType) Get() *IdentityType {
	return v.value
}

func (v *NullableIdentityType) Set(val *IdentityType) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityType) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityType(val *IdentityType) *NullableIdentityType {
	return &NullableIdentityType{value: val, isSet: true}
}

func (v NullableIdentityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

