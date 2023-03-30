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

// NrfInfoServedBsfInfoValue struct for NrfInfoServedBsfInfoValue
type NrfInfoServedBsfInfoValue struct {
	BsfInfo *BsfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedBsfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BsfInfo
	err = json.Unmarshal(data, &dst.BsfInfo);
	if err == nil {
		jsonBsfInfo, _ := json.Marshal(dst.BsfInfo)
		if string(jsonBsfInfo) == "{}" { // empty struct
			dst.BsfInfo = nil
		} else {
			return nil // data stored in dst.BsfInfo, return on the first match
		}
	} else {
		dst.BsfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedBsfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedBsfInfoValue) MarshalJSON() ([]byte, error) {
	if src.BsfInfo != nil {
		return json.Marshal(&src.BsfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedBsfInfoValue struct {
	value *NrfInfoServedBsfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedBsfInfoValue) Get() *NrfInfoServedBsfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedBsfInfoValue) Set(val *NrfInfoServedBsfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedBsfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedBsfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedBsfInfoValue(val *NrfInfoServedBsfInfoValue) *NullableNrfInfoServedBsfInfoValue {
	return &NullableNrfInfoServedBsfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedBsfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedBsfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


