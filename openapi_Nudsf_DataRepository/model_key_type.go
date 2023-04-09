/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
	"fmt"
)

// KeyType Represents the type of a key.
type KeyType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *KeyType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(KeyType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *KeyType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableKeyType struct {
	value *KeyType
	isSet bool
}

func (v NullableKeyType) Get() *KeyType {
	return v.value
}

func (v *NullableKeyType) Set(val *KeyType) {
	v.value = val
	v.isSet = true
}

func (v NullableKeyType) IsSet() bool {
	return v.isSet
}

func (v *NullableKeyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKeyType(val *KeyType) *NullableKeyType {
	return &NullableKeyType{value: val, isSet: true}
}

func (v NullableKeyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKeyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
