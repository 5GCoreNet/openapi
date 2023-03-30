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

// NrfInfoServedMbSmfInfoListValueValue struct for NrfInfoServedMbSmfInfoListValueValue
type NrfInfoServedMbSmfInfoListValueValue struct {
	MbSmfInfo *MbSmfInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedMbSmfInfoListValueValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into MbSmfInfo
	err = json.Unmarshal(data, &dst.MbSmfInfo);
	if err == nil {
		jsonMbSmfInfo, _ := json.Marshal(dst.MbSmfInfo)
		if string(jsonMbSmfInfo) == "{}" { // empty struct
			dst.MbSmfInfo = nil
		} else {
			return nil // data stored in dst.MbSmfInfo, return on the first match
		}
	} else {
		dst.MbSmfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedMbSmfInfoListValueValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedMbSmfInfoListValueValue) MarshalJSON() ([]byte, error) {
	if src.MbSmfInfo != nil {
		return json.Marshal(&src.MbSmfInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedMbSmfInfoListValueValue struct {
	value *NrfInfoServedMbSmfInfoListValueValue
	isSet bool
}

func (v NullableNrfInfoServedMbSmfInfoListValueValue) Get() *NrfInfoServedMbSmfInfoListValueValue {
	return v.value
}

func (v *NullableNrfInfoServedMbSmfInfoListValueValue) Set(val *NrfInfoServedMbSmfInfoListValueValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedMbSmfInfoListValueValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedMbSmfInfoListValueValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedMbSmfInfoListValueValue(val *NrfInfoServedMbSmfInfoListValueValue) *NullableNrfInfoServedMbSmfInfoListValueValue {
	return &NullableNrfInfoServedMbSmfInfoListValueValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedMbSmfInfoListValueValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedMbSmfInfoListValueValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


