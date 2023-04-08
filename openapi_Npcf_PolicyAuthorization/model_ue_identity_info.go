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

// UeIdentityInfo Represents 5GS-Level UE identities.
type UeIdentityInfo struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeIdentityInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface);
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeIdentityInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeIdentityInfo) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeIdentityInfo struct {
	value *UeIdentityInfo
	isSet bool
}

func (v NullableUeIdentityInfo) Get() *UeIdentityInfo {
	return v.value
}

func (v *NullableUeIdentityInfo) Set(val *UeIdentityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeIdentityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeIdentityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeIdentityInfo(val *UeIdentityInfo) *NullableUeIdentityInfo {
	return &NullableUeIdentityInfo{value: val, isSet: true}
}

func (v NullableUeIdentityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeIdentityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


