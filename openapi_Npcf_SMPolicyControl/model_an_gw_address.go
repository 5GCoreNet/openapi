/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// AnGwAddress Describes the address of the access network gateway control node.
type AnGwAddress struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AnGwAddress) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			return nil // data stored in dst.Interface, return on the first match
		}
	} else {
		dst.Interface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AnGwAddress)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AnGwAddress) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAnGwAddress struct {
	value *AnGwAddress
	isSet bool
}

func (v NullableAnGwAddress) Get() *AnGwAddress {
	return v.value
}

func (v *NullableAnGwAddress) Set(val *AnGwAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAnGwAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAnGwAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnGwAddress(val *AnGwAddress) *NullableAnGwAddress {
	return &NullableAnGwAddress{value: val, isSet: true}
}

func (v NullableAnGwAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnGwAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
