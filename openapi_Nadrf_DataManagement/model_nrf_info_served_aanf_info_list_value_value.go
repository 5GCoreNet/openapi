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

// NrfInfoServedAanfInfoListValueValue struct for NrfInfoServedAanfInfoListValueValue
type NrfInfoServedAanfInfoListValueValue struct {
	AanfInfo       *AanfInfo
	MapOfInterface *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedAanfInfoListValueValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AanfInfo
	err = json.Unmarshal(data, &dst.AanfInfo)
	if err == nil {
		jsonAanfInfo, _ := json.Marshal(dst.AanfInfo)
		if string(jsonAanfInfo) == "{}" { // empty struct
			dst.AanfInfo = nil
		} else {
			return nil // data stored in dst.AanfInfo, return on the first match
		}
	} else {
		dst.AanfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedAanfInfoListValueValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedAanfInfoListValueValue) MarshalJSON() ([]byte, error) {
	if src.AanfInfo != nil {
		return json.Marshal(&src.AanfInfo)
	}

	if src.MapOfInterface != nil {
		return json.Marshal(&src.MapOfInterface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedAanfInfoListValueValue struct {
	value *NrfInfoServedAanfInfoListValueValue
	isSet bool
}

func (v NullableNrfInfoServedAanfInfoListValueValue) Get() *NrfInfoServedAanfInfoListValueValue {
	return v.value
}

func (v *NullableNrfInfoServedAanfInfoListValueValue) Set(val *NrfInfoServedAanfInfoListValueValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedAanfInfoListValueValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedAanfInfoListValueValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedAanfInfoListValueValue(val *NrfInfoServedAanfInfoListValueValue) *NullableNrfInfoServedAanfInfoListValueValue {
	return &NullableNrfInfoServedAanfInfoListValueValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedAanfInfoListValueValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedAanfInfoListValueValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
