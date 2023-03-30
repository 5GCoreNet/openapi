/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedUdrInfoValue struct for NrfInfoServedUdrInfoValue
type NrfInfoServedUdrInfoValue struct {
	UdrInfo *UdrInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedUdrInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UdrInfo
	err = json.Unmarshal(data, &dst.UdrInfo);
	if err == nil {
		jsonUdrInfo, _ := json.Marshal(dst.UdrInfo)
		if string(jsonUdrInfo) == "{}" { // empty struct
			dst.UdrInfo = nil
		} else {
			return nil // data stored in dst.UdrInfo, return on the first match
		}
	} else {
		dst.UdrInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedUdrInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedUdrInfoValue) MarshalJSON() ([]byte, error) {
	if src.UdrInfo != nil {
		return json.Marshal(&src.UdrInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedUdrInfoValue struct {
	value *NrfInfoServedUdrInfoValue
	isSet bool
}

func (v NullableNrfInfoServedUdrInfoValue) Get() *NrfInfoServedUdrInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedUdrInfoValue) Set(val *NrfInfoServedUdrInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedUdrInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedUdrInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedUdrInfoValue(val *NrfInfoServedUdrInfoValue) *NullableNrfInfoServedUdrInfoValue {
	return &NullableNrfInfoServedUdrInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedUdrInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedUdrInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


