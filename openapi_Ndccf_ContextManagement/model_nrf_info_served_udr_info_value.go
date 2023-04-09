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

// NrfInfoServedUdrInfoValue struct for NrfInfoServedUdrInfoValue
type NrfInfoServedUdrInfoValue struct {
	UdrInfo        *UdrInfo
	MapOfInterface *map[string]interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrfInfoServedUdrInfoValue) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UdrInfo
	err = json.Unmarshal(data, &dst.UdrInfo)
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrfInfoServedUdrInfoValue)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrfInfoServedUdrInfoValue) MarshalJSON() ([]byte, error) {
	if src.UdrInfo != nil {
		return json.Marshal(&src.UdrInfo)
	}

	if src.MapOfInterface != nil {
		return json.Marshal(&src.MapOfInterface)
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
