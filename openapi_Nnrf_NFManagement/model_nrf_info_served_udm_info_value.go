/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// NrfInfoServedUdmInfoValue struct for NrfInfoServedUdmInfoValue
type NrfInfoServedUdmInfoValue struct {
	UdmInfo *UdmInfo
	map[string]interface{} *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedUdmInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UdmInfo
	err = json.Unmarshal(data, &dst.UdmInfo);
	if err == nil {
		jsonUdmInfo, _ := json.Marshal(dst.UdmInfo)
		if string(jsonUdmInfo) == "{}" { // empty struct
			dst.UdmInfo = nil
		} else {
			return nil // data stored in dst.UdmInfo, return on the first match
		}
	} else {
		dst.UdmInfo = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedUdmInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedUdmInfoValue) MarshalJSON() ([]byte, error) {
	if src.UdmInfo != nil {
		return json.Marshal(&src.UdmInfo)
	}

	if src.map[string]interface{} != nil {
		return json.Marshal(&src.map[string]interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrfInfoServedUdmInfoValue struct {
	value *NrfInfoServedUdmInfoValue
	isSet bool
}

func (v NullableNrfInfoServedUdmInfoValue) Get() *NrfInfoServedUdmInfoValue {
	return v.value
}

func (v *NullableNrfInfoServedUdmInfoValue) Set(val *NrfInfoServedUdmInfoValue) {
	v.value = val
	v.isSet = true
}

func (v NullableNrfInfoServedUdmInfoValue) IsSet() bool {
	return v.isSet
}

func (v *NullableNrfInfoServedUdmInfoValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrfInfoServedUdmInfoValue(val *NrfInfoServedUdmInfoValue) *NullableNrfInfoServedUdmInfoValue {
	return &NullableNrfInfoServedUdmInfoValue{value: val, isSet: true}
}

func (v NullableNrfInfoServedUdmInfoValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrfInfoServedUdmInfoValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


