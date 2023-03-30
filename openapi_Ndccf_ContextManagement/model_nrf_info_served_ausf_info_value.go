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

// NrfInfoServedAusfInfoValue struct for NrfInfoServedAusfInfoValue
type NrfInfoServedAusfInfoValue struct {
	AusfInfo *AusfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedAusfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AusfInfo
	err = json.Unmarshal(data, &dst.AusfInfo);
	if err == nil {
		jsonAusfInfo, _ := json.Marshal(dst.AusfInfo)
		if string(jsonAusfInfo) == "{}" { // empty struct
			dst.AusfInfo = nil
		} else {
			return nil // data stored in dst.AusfInfo, return on the first match
		}
	} else {
		dst.AusfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedAusfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedAusfInfoValue) MarshalJSON() ([]byte, error) {
	if src.AusfInfo != nil {
		return json.Marshal(&src.AusfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedAusfInfoValue struct {
	value *NrfInfoServedAusfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedAusfInfoValue) Get() *NrfInfoServedAusfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedAusfInfoValue) Set(val *NrfInfoServedAusfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedAusfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedAusfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedAusfInfoValue(val *NrfInfoServedAusfInfoValue) *NullableNrfInfoServedAusfInfoValue {
	return &NullableNrfInfoServedAusfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedAusfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedAusfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


