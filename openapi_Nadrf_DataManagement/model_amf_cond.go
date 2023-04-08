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

// AmfCond Subscription to a set of AMFs, based on AMF Set Id and/or AMF Region Id
type AmfCond struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmfCond) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface);
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

	return fmt.Errorf("data failed to match schemas in anyOf(AmfCond)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmfCond) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmfCond struct {
	value *AmfCond
	isSet bool
}

func (v NullableAmfCond) Get() *AmfCond {
	return v.value
}

func (v *NullableAmfCond) Set(val *AmfCond) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfCond) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfCond(val *AmfCond) *NullableAmfCond {
	return &NullableAmfCond{value: val, isSet: true}
}

func (v NullableAmfCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


