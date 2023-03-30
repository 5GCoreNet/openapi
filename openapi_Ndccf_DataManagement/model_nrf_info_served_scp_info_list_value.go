/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedScpInfoListValue struct for NrfInfoServedScpInfoListValue
type NrfInfoServedScpInfoListValue struct {
	ScpInfo *ScpInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedScpInfoListValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ScpInfo
	err = json.Unmarshal(data, &dst.ScpInfo);
	if err == nil {
		jsonScpInfo, _ := json.Marshal(dst.ScpInfo)
		if string(jsonScpInfo) == "{}" { // empty struct
			dst.ScpInfo = nil
		} else {
			return nil // data stored in dst.ScpInfo, return on the first match
		}
	} else {
		dst.ScpInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedScpInfoListValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedScpInfoListValue) MarshalJSON() ([]byte, error) {
	if src.ScpInfo != nil {
		return json.Marshal(&src.ScpInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedScpInfoListValue struct {
	value *NrfInfoServedScpInfoListValue
	isSet bool
}

func (v NullableNrfInfoServedScpInfoListValue) Get() *NrfInfoServedScpInfoListValue {
	return v.value
}

func (v *NullableNrfInfoServedScpInfoListValue) Set(val *NrfInfoServedScpInfoListValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedScpInfoListValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedScpInfoListValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedScpInfoListValue(val *NrfInfoServedScpInfoListValue) *NullableNrfInfoServedScpInfoListValue {
	return &NullableNrfInfoServedScpInfoListValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedScpInfoListValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedScpInfoListValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


