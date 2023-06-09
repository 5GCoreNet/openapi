/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// AmfEventType Describes the supported event types of Namf_EventExposure Service
type AmfEventType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmfEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AmfEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmfEventType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmfEventType struct {
	value *AmfEventType
	isSet bool
}

func (v NullableAmfEventType) Get() *AmfEventType {
	return v.value
}

func (v *NullableAmfEventType) Set(val *AmfEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventType(val *AmfEventType) *NullableAmfEventType {
	return &NullableAmfEventType{value: val, isSet: true}
}

func (v NullableAmfEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
