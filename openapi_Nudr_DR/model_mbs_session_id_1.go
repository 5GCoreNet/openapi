/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// MbsSessionId1 MBS Session Identifier
type MbsSessionId1 struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsSessionId1) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MbsSessionId1)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsSessionId1) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsSessionId1 struct {
	value *MbsSessionId1
	isSet bool
}

func (v NullableMbsSessionId1) Get() *MbsSessionId1 {
	return v.value
}

func (v *NullableMbsSessionId1) Set(val *MbsSessionId1) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionId1) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionId1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionId1(val *MbsSessionId1) *NullableMbsSessionId1 {
	return &NullableMbsSessionId1{value: val, isSet: true}
}

func (v NullableMbsSessionId1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionId1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

