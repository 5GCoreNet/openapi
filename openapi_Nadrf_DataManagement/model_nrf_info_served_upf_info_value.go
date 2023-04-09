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

// NrfInfoServedUpfInfoValue struct for NrfInfoServedUpfInfoValue
type NrfInfoServedUpfInfoValue struct {
	UpfInfo        *UpfInfo
	MapOfInterface *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedUpfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UpfInfo
	err = json.Unmarshal(data, &dst.UpfInfo)
	if err == nil {
		jsonUpfInfo, _ := json.Marshal(dst.UpfInfo)
		if string(jsonUpfInfo) == "{}" { // empty struct
			dst.UpfInfo = nil
		} else {
			return nil // data stored in dst.UpfInfo, return on the first match
		}
	} else {
		dst.UpfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedUpfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedUpfInfoValue) MarshalJSON() ([]byte, error) {
	if src.UpfInfo != nil {
		return json.Marshal(&src.UpfInfo)
	}

	if src.MapOfInterface != nil {
		return json.Marshal(&src.MapOfInterface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedUpfInfoValue struct {
	value *NrfInfoServedUpfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedUpfInfoValue) Get() *NrfInfoServedUpfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedUpfInfoValue) Set(val *NrfInfoServedUpfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedUpfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedUpfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedUpfInfoValue(val *NrfInfoServedUpfInfoValue) *NullableNrfInfoServedUpfInfoValue {
	return &NullableNrfInfoServedUpfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedUpfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedUpfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
