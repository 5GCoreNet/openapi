/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedAmfInfoValue struct for NrfInfoServedAmfInfoValue
type NrfInfoServedAmfInfoValue struct {
	AmfInfo *AmfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedAmfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AmfInfo
	err = json.Unmarshal(data, &dst.AmfInfo);
	if err == nil {
		jsonAmfInfo, _ := json.Marshal(dst.AmfInfo)
		if string(jsonAmfInfo) == "{}" { // empty struct
			dst.AmfInfo = nil
		} else {
			return nil // data stored in dst.AmfInfo, return on the first match
		}
	} else {
		dst.AmfInfo = nil
	}

	// try to unmarshal JSON data into map[string]interface{}
	err = json.Unmarshal(data, &dst.map[string]interface{});
	if err == nil {
		jsonmap[string]interface{}, _ := json.Marshal(dst.map[string]interface{})
		if string(jsonmap[string]interface{}) == "{}" { // empty struct
			dst.map[string]interface{} = nil
		} else {
			return nil // data stored in dst.map[string]interface{}, return on the first match
		}
	} else {
		dst.map[string]interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedAmfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedAmfInfoValue) MarshalJSON() ([]byte, error) {
	if src.AmfInfo != nil {
		return json.Marshal(&src.AmfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedAmfInfoValue struct {
	value *NrfInfoServedAmfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedAmfInfoValue) Get() *NrfInfoServedAmfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedAmfInfoValue) Set(val *NrfInfoServedAmfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedAmfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedAmfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedAmfInfoValue(val *NrfInfoServedAmfInfoValue) *NullableNrfInfoServedAmfInfoValue {
	return &NullableNrfInfoServedAmfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedAmfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedAmfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

