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

// NrfInfoServedGmlcInfoValue struct for NrfInfoServedGmlcInfoValue
type NrfInfoServedGmlcInfoValue struct {
	GmlcInfo *GmlcInfo
	MapOfInterface *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedGmlcInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into GmlcInfo
	err = json.Unmarshal(data, &dst.GmlcInfo);
	if err == nil {
		jsonGmlcInfo, _ := json.Marshal(dst.GmlcInfo)
		if string(jsonGmlcInfo) == "{}" { // empty struct
			dst.GmlcInfo = nil
		} else {
			return nil // data stored in dst.GmlcInfo, return on the first match
		}
	} else {
		dst.GmlcInfo = nil
	}

	// try to unmarshal JSON data into map[string]interface{}
	err = json.Unmarshal(data, &dst.MapOfInterface);
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedGmlcInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedGmlcInfoValue) MarshalJSON() ([]byte, error) {
	if src.GmlcInfo != nil {
		return json.Marshal(&src.GmlcInfo)
	}

	if src.MapOfInterface != nil {
		return json.Marshal(&src.MapOfInterface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedGmlcInfoValue struct {
	value *NrfInfoServedGmlcInfoValue
	isSet bool
}

func (v NullableNrfInfoServedGmlcInfoValue) Get() *NrfInfoServedGmlcInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedGmlcInfoValue) Set(val *NrfInfoServedGmlcInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedGmlcInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedGmlcInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedGmlcInfoValue(val *NrfInfoServedGmlcInfoValue) *NullableNrfInfoServedGmlcInfoValue {
	return &NullableNrfInfoServedGmlcInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedGmlcInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedGmlcInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


