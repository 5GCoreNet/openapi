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

// NrfInfoServedUdsfInfoValue struct for NrfInfoServedUdsfInfoValue
type NrfInfoServedUdsfInfoValue struct {
	UdsfInfo       *UdsfInfo
	MapOfInterface *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedUdsfInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UdsfInfo
	err = json.Unmarshal(data, &dst.UdsfInfo)
	if err == nil {
		jsonUdsfInfo, _ := json.Marshal(dst.UdsfInfo)
		if string(jsonUdsfInfo) == "{}" { // empty struct
			dst.UdsfInfo = nil
		} else {
			return nil // data stored in dst.UdsfInfo, return on the first match
		}
	} else {
		dst.UdsfInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedUdsfInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedUdsfInfoValue) MarshalJSON() ([]byte, error) {
	if src.UdsfInfo != nil {
		return json.Marshal(&src.UdsfInfo)
	}

	if src.MapOfInterface != nil {
		return json.Marshal(&src.MapOfInterface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedUdsfInfoValue struct {
	value *NrfInfoServedUdsfInfoValue
	isSet bool
}

func (v NullableNrfInfoServedUdsfInfoValue) Get() *NrfInfoServedUdsfInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedUdsfInfoValue) Set(val *NrfInfoServedUdsfInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedUdsfInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedUdsfInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedUdsfInfoValue(val *NrfInfoServedUdsfInfoValue) *NullableNrfInfoServedUdsfInfoValue {
	return &NullableNrfInfoServedUdsfInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedUdsfInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedUdsfInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
