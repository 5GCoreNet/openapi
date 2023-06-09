/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// WlanPerTsPerformanceInfo WLAN performance information per Time Slot during the analytics target period.
type WlanPerTsPerformanceInfo struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *WlanPerTsPerformanceInfo) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(WlanPerTsPerformanceInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *WlanPerTsPerformanceInfo) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableWlanPerTsPerformanceInfo struct {
	value *WlanPerTsPerformanceInfo
	isSet bool
}

func (v NullableWlanPerTsPerformanceInfo) Get() *WlanPerTsPerformanceInfo {
	return v.value
}

func (v *NullableWlanPerTsPerformanceInfo) Set(val *WlanPerTsPerformanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanPerTsPerformanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanPerTsPerformanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanPerTsPerformanceInfo(val *WlanPerTsPerformanceInfo) *NullableWlanPerTsPerformanceInfo {
	return &NullableWlanPerTsPerformanceInfo{value: val, isSet: true}
}

func (v NullableWlanPerTsPerformanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanPerTsPerformanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
