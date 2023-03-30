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

// AppDetectionNotifType Indicates the notification type for Application Detection Control.
type AppDetectionNotifType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AppDetectionNotifType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AppDetectionNotifType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AppDetectionNotifType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAppDetectionNotifType struct {
	value *AppDetectionNotifType
	isSet bool
}

func (v NullableAppDetectionNotifType) Get() *AppDetectionNotifType {
	return v.value
}

func (v *NullableAppDetectionNotifType) Set(val *AppDetectionNotifType) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDetectionNotifType) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDetectionNotifType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDetectionNotifType(val *AppDetectionNotifType) *NullableAppDetectionNotifType {
	return &NullableAppDetectionNotifType{value: val, isSet: true}
}

func (v NullableAppDetectionNotifType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDetectionNotifType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


