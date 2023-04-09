/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// TwifInfo Addressing information (IP addresses, FQDN) of the TWIF
type TwifInfo struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TwifInfo) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TwifInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TwifInfo) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTwifInfo struct {
	value *TwifInfo
	isSet bool
}

func (v NullableTwifInfo) Get() *TwifInfo {
	return v.value
}

func (v *NullableTwifInfo) Set(val *TwifInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTwifInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTwifInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTwifInfo(val *TwifInfo) *NullableTwifInfo {
	return &NullableTwifInfo{value: val, isSet: true}
}

func (v NullableTwifInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTwifInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
