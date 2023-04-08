/*
Nucmf_UECapabilityManagement

Nucmf_UECapabilityManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nucmf_UERCM

import (
	"encoding/json"
	"fmt"
)

// RacFormatAnyOf the model 'RacFormatAnyOf'
type RacFormatAnyOf string

// List of RacFormat_anyOf
const (
	_5_GS RacFormatAnyOf = "5GS"
	EPS RacFormatAnyOf = "EPS"
)

// All allowed values of RacFormatAnyOf enum
var AllowedRacFormatAnyOfEnumValues = []RacFormatAnyOf{
	"5GS",
	"EPS",
}

func (v *RacFormatAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RacFormatAnyOf(value)
	for _, existing := range AllowedRacFormatAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RacFormatAnyOf", value)
}

// NewRacFormatAnyOfFromValue returns a pointer to a valid RacFormatAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRacFormatAnyOfFromValue(v string) (*RacFormatAnyOf, error) {
	ev := RacFormatAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RacFormatAnyOf: valid values are %v", v, AllowedRacFormatAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RacFormatAnyOf) IsValid() bool {
	for _, existing := range AllowedRacFormatAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RacFormat_anyOf value
func (v RacFormatAnyOf) Ptr() *RacFormatAnyOf {
	return &v
}

type NullableRacFormatAnyOf struct {
	value *RacFormatAnyOf
	isSet bool
}

func (v NullableRacFormatAnyOf) Get() *RacFormatAnyOf {
	return v.value
}

func (v *NullableRacFormatAnyOf) Set(val *RacFormatAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRacFormatAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRacFormatAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRacFormatAnyOf(val *RacFormatAnyOf) *NullableRacFormatAnyOf {
	return &NullableRacFormatAnyOf{value: val, isSet: true}
}

func (v NullableRacFormatAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRacFormatAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
