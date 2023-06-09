/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// ServAuthInfo Indicates the result of the Policy Authorization service request from the AF.
type ServAuthInfo struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServAuthInfo) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServAuthInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServAuthInfo) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServAuthInfo struct {
	value *ServAuthInfo
	isSet bool
}

func (v NullableServAuthInfo) Get() *ServAuthInfo {
	return v.value
}

func (v *NullableServAuthInfo) Set(val *ServAuthInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServAuthInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServAuthInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServAuthInfo(val *ServAuthInfo) *NullableServAuthInfo {
	return &NullableServAuthInfo{value: val, isSet: true}
}

func (v NullableServAuthInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServAuthInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
