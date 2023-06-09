/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// PsUserState - User state in PS domain
type PsUserState struct {
	Interface *interface{}
}

// interface{}AsPsUserState is a convenience function that returns interface{} wrapped in PsUserState
func InterfaceAsPsUserState(v *interface{}) PsUserState {
	return PsUserState{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PsUserState) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface
	err = newStrictDecoder(data).Decode(&dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			match++
		}
	} else {
		dst.Interface = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PsUserState)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PsUserState)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PsUserState) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PsUserState) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullablePsUserState struct {
	value *PsUserState
	isSet bool
}

func (v NullablePsUserState) Get() *PsUserState {
	return v.value
}

func (v *NullablePsUserState) Set(val *PsUserState) {
	v.value = val
	v.isSet = true
}

func (v NullablePsUserState) IsSet() bool {
	return v.isSet
}

func (v *NullablePsUserState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePsUserState(val *PsUserState) *NullablePsUserState {
	return &NullablePsUserState{value: val, isSet: true}
}

func (v NullablePsUserState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePsUserState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
