/*
EES ACR Status Update Service

EES ACR Status Update Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRStatusUpdate

import (
	"encoding/json"
	"fmt"
)

// ACRUpdateData Represents the parameters to update the information related to ACR (e.g. indicate the status of ACT, update the notification target address).
type ACRUpdateData struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ACRUpdateData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			return nil // data stored in dst.Interface, return on the first match
		}
	} else {
		dst.Interface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ACRUpdateData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ACRUpdateData) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableACRUpdateData struct {
	value *ACRUpdateData
	isSet bool
}

func (v NullableACRUpdateData) Get() *ACRUpdateData {
	return v.value
}

func (v *NullableACRUpdateData) Set(val *ACRUpdateData) {
	v.value = val
	v.isSet = true
}

func (v NullableACRUpdateData) IsSet() bool {
	return v.isSet
}

func (v *NullableACRUpdateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACRUpdateData(val *ACRUpdateData) *NullableACRUpdateData {
	return &NullableACRUpdateData{value: val, isSet: true}
}

func (v NullableACRUpdateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACRUpdateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
