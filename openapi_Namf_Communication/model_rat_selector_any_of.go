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

// RatSelectorAnyOf the model 'RatSelectorAnyOf'
type RatSelectorAnyOf string

// List of RatSelector_anyOf
const (
	E_UTRA RatSelectorAnyOf = "E-UTRA"
	NR RatSelectorAnyOf = "NR"
)

// All allowed values of RatSelectorAnyOf enum
var AllowedRatSelectorAnyOfEnumValues = []RatSelectorAnyOf{
	"E-UTRA",
	"NR",
}

func (v *RatSelectorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RatSelectorAnyOf(value)
	for _, existing := range AllowedRatSelectorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RatSelectorAnyOf", value)
}

// NewRatSelectorAnyOfFromValue returns a pointer to a valid RatSelectorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRatSelectorAnyOfFromValue(v string) (*RatSelectorAnyOf, error) {
	ev := RatSelectorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RatSelectorAnyOf: valid values are %v", v, AllowedRatSelectorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RatSelectorAnyOf) IsValid() bool {
	for _, existing := range AllowedRatSelectorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RatSelector_anyOf value
func (v RatSelectorAnyOf) Ptr() *RatSelectorAnyOf {
	return &v
}

type NullableRatSelectorAnyOf struct {
	value *RatSelectorAnyOf
	isSet bool
}

func (v NullableRatSelectorAnyOf) Get() *RatSelectorAnyOf {
	return v.value
}

func (v *NullableRatSelectorAnyOf) Set(val *RatSelectorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRatSelectorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRatSelectorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatSelectorAnyOf(val *RatSelectorAnyOf) *NullableRatSelectorAnyOf {
	return &NullableRatSelectorAnyOf{value: val, isSet: true}
}

func (v NullableRatSelectorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatSelectorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
