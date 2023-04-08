/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_PP

import (
	"encoding/json"
	"fmt"
)

// TrafficProfileAnyOf the model 'TrafficProfileAnyOf'
type TrafficProfileAnyOf string

// List of TrafficProfile_anyOf
const (
	SINGLE_TRANS_UL TrafficProfileAnyOf = "SINGLE_TRANS_UL"
	SINGLE_TRANS_DL TrafficProfileAnyOf = "SINGLE_TRANS_DL"
	DUAL_TRANS_UL_FIRST TrafficProfileAnyOf = "DUAL_TRANS_UL_FIRST"
	DUAL_TRANS_DL_FIRST TrafficProfileAnyOf = "DUAL_TRANS_DL_FIRST"
	MULTI_TRANS TrafficProfileAnyOf = "MULTI_TRANS"
)

// All allowed values of TrafficProfileAnyOf enum
var AllowedTrafficProfileAnyOfEnumValues = []TrafficProfileAnyOf{
	"SINGLE_TRANS_UL",
	"SINGLE_TRANS_DL",
	"DUAL_TRANS_UL_FIRST",
	"DUAL_TRANS_DL_FIRST",
	"MULTI_TRANS",
}

func (v *TrafficProfileAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrafficProfileAnyOf(value)
	for _, existing := range AllowedTrafficProfileAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrafficProfileAnyOf", value)
}

// NewTrafficProfileAnyOfFromValue returns a pointer to a valid TrafficProfileAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrafficProfileAnyOfFromValue(v string) (*TrafficProfileAnyOf, error) {
	ev := TrafficProfileAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrafficProfileAnyOf: valid values are %v", v, AllowedTrafficProfileAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrafficProfileAnyOf) IsValid() bool {
	for _, existing := range AllowedTrafficProfileAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrafficProfile_anyOf value
func (v TrafficProfileAnyOf) Ptr() *TrafficProfileAnyOf {
	return &v
}

type NullableTrafficProfileAnyOf struct {
	value *TrafficProfileAnyOf
	isSet bool
}

func (v NullableTrafficProfileAnyOf) Get() *TrafficProfileAnyOf {
	return v.value
}

func (v *NullableTrafficProfileAnyOf) Set(val *TrafficProfileAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficProfileAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficProfileAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficProfileAnyOf(val *TrafficProfileAnyOf) *NullableTrafficProfileAnyOf {
	return &NullableTrafficProfileAnyOf{value: val, isSet: true}
}

func (v NullableTrafficProfileAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficProfileAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

