/*
Nhss_UECM

HSS UE Context Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_UECM

import (
	"encoding/json"
	"fmt"
)

// ImeiUpdateInfo - It represents the request body of the IMEI update request sent by UDM to HSS, and contains the IMSI of the UE and the new IMEI(SV)
type ImeiUpdateInfo struct {
	Interface *interface{}
}

// interface{}AsImeiUpdateInfo is a convenience function that returns interface{} wrapped in ImeiUpdateInfo
func InterfaceAsImeiUpdateInfo(v *interface{}) ImeiUpdateInfo {
	return ImeiUpdateInfo{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ImeiUpdateInfo) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface
	err = newStrictDecoder(data).Decode(&dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			match++
		}
	} else {
		dst.Interface = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ImeiUpdateInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ImeiUpdateInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ImeiUpdateInfo) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ImeiUpdateInfo) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableImeiUpdateInfo struct {
	value *ImeiUpdateInfo
	isSet bool
}

func (v NullableImeiUpdateInfo) Get() *ImeiUpdateInfo {
	return v.value
}

func (v *NullableImeiUpdateInfo) Set(val *ImeiUpdateInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableImeiUpdateInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableImeiUpdateInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImeiUpdateInfo(val *ImeiUpdateInfo) *NullableImeiUpdateInfo {
	return &NullableImeiUpdateInfo{value: val, isSet: true}
}

func (v NullableImeiUpdateInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImeiUpdateInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
