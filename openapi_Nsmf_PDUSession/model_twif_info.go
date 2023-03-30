/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// TwifInfo Addressing information (IP addresses, FQDN) of the TWIF
type TwifInfo struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TwifInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TwifInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TwifInfo) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
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


