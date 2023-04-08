/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
	"fmt"
)

// SubTerminationReasonAnyOf the model 'SubTerminationReasonAnyOf'
type SubTerminationReasonAnyOf string

// List of SubTerminationReason_anyOf
const (
	INVALID_SUBSCRIPTION SubTerminationReasonAnyOf = "INVALID_SUBSCRIPTION"
)

// All allowed values of SubTerminationReasonAnyOf enum
var AllowedSubTerminationReasonAnyOfEnumValues = []SubTerminationReasonAnyOf{
	"INVALID_SUBSCRIPTION",
}

func (v *SubTerminationReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubTerminationReasonAnyOf(value)
	for _, existing := range AllowedSubTerminationReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubTerminationReasonAnyOf", value)
}

// NewSubTerminationReasonAnyOfFromValue returns a pointer to a valid SubTerminationReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubTerminationReasonAnyOfFromValue(v string) (*SubTerminationReasonAnyOf, error) {
	ev := SubTerminationReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubTerminationReasonAnyOf: valid values are %v", v, AllowedSubTerminationReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubTerminationReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedSubTerminationReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubTerminationReason_anyOf value
func (v SubTerminationReasonAnyOf) Ptr() *SubTerminationReasonAnyOf {
	return &v
}

type NullableSubTerminationReasonAnyOf struct {
	value *SubTerminationReasonAnyOf
	isSet bool
}

func (v NullableSubTerminationReasonAnyOf) Get() *SubTerminationReasonAnyOf {
	return v.value
}

func (v *NullableSubTerminationReasonAnyOf) Set(val *SubTerminationReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubTerminationReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubTerminationReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubTerminationReasonAnyOf(val *SubTerminationReasonAnyOf) *NullableSubTerminationReasonAnyOf {
	return &NullableSubTerminationReasonAnyOf{value: val, isSet: true}
}

func (v NullableSubTerminationReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubTerminationReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

