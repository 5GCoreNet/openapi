/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedSeppInfoListValue struct for NrfInfoServedSeppInfoListValue
type NrfInfoServedSeppInfoListValue struct {
	SeppInfo       *SeppInfo
	MapOfInterface *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedSeppInfoListValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SeppInfo
	err = json.Unmarshal(data, &dst.SeppInfo)
	if err == nil {
		jsonSeppInfo, _ := json.Marshal(dst.SeppInfo)
		if string(jsonSeppInfo) == "{}" { // empty struct
			dst.SeppInfo = nil
		} else {
			return nil // data stored in dst.SeppInfo, return on the first match
		}
	} else {
		dst.SeppInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedSeppInfoListValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedSeppInfoListValue) MarshalJSON() ([]byte, error) {
	if src.SeppInfo != nil {
		return json.Marshal(&src.SeppInfo)
	}

	if src.MapOfInterface != nil {
		return json.Marshal(&src.MapOfInterface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedSeppInfoListValue struct {
	value *NrfInfoServedSeppInfoListValue
	isSet bool
}

func (v NullableNrfInfoServedSeppInfoListValue) Get() *NrfInfoServedSeppInfoListValue {
	return v.value
}

func (v *NullableNrfInfoServedSeppInfoListValue) Set(val *NrfInfoServedSeppInfoListValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedSeppInfoListValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedSeppInfoListValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedSeppInfoListValue(val *NrfInfoServedSeppInfoListValue) *NullableNrfInfoServedSeppInfoListValue {
	return &NullableNrfInfoServedSeppInfoListValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedSeppInfoListValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedSeppInfoListValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
