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

// NrfInfoServedPcfInfoValue struct for NrfInfoServedPcfInfoValue
type NrfInfoServedPcfInfoValue struct {
	PcfInfo        *PcfInfo
	MapOfInterface *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedPcfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PcfInfo
	err = json.Unmarshal(data, &dst.PcfInfo)
	if err == nil {
		jsonPcfInfo, _ := json.Marshal(dst.PcfInfo)
		if string(jsonPcfInfo) == "{}" { // empty struct
			dst.PcfInfo = nil
		} else {
			return nil // data stored in dst.PcfInfo, return on the first match
		}
	} else {
		dst.PcfInfo = nil
	}

	// try to unmarshal JSON data into map[string]interface{}
	err = json.Unmarshal(data, &dst.MapOfInterface)
	if err == nil {
		jsonMapOfInterface, _ := json.Marshal(dst.MapOfInterface)
		if string(jsonMapOfInterface) == "{}" { // empty struct
			dst.MapOfInterface = nil
		} else {
			return nil // data stored in dst.MapOfInterface, return on the first match
		}
	} else {
		dst.MapOfInterface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedPcfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedPcfInfoValue) MarshalJSON() ([]byte, error) {
	if src.PcfInfo != nil {
		return json.Marshal(&src.PcfInfo)
	}

	if src.MapOfInterface != nil {
		return json.Marshal(&src.MapOfInterface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedPcfInfoValue struct {
	value *NrfInfoServedPcfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedPcfInfoValue) Get() *NrfInfoServedPcfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedPcfInfoValue) Set(val *NrfInfoServedPcfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedPcfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedPcfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedPcfInfoValue(val *NrfInfoServedPcfInfoValue) *NullableNrfInfoServedPcfInfoValue {
	return &NullableNrfInfoServedPcfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedPcfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedPcfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
