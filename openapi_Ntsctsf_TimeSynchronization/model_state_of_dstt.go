/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_TimeSynchronization

import (
	"encoding/json"
	"fmt"
)

// StateOfDstt - Contains the PTP port state of a DS-TT.
type StateOfDstt struct {
	Interface *interface{}
}

// interface{}AsStateOfDstt is a convenience function that returns interface{} wrapped in StateOfDstt
func InterfaceAsStateOfDstt(v *interface{}) StateOfDstt {
	return StateOfDstt{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *StateOfDstt) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(StateOfDstt)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(StateOfDstt)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src StateOfDstt) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *StateOfDstt) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableStateOfDstt struct {
	value *StateOfDstt
	isSet bool
}

func (v NullableStateOfDstt) Get() *StateOfDstt {
	return v.value
}

func (v *NullableStateOfDstt) Set(val *StateOfDstt) {
	v.value = val
	v.isSet = true
}

func (v NullableStateOfDstt) IsSet() bool {
	return v.isSet
}

func (v *NullableStateOfDstt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateOfDstt(val *StateOfDstt) *NullableStateOfDstt {
	return &NullableStateOfDstt{value: val, isSet: true}
}

func (v NullableStateOfDstt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateOfDstt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
