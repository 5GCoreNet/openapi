/*
Npcf_EventExposure

PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// FailureOneOf the model 'FailureOneOf'
type FailureOneOf string

// List of Failure_oneOf
const (
	UNSPECIFIED FailureOneOf = "UNSPECIFIED"
	UE_NOT_REACHABLE FailureOneOf = "UE_NOT_REACHABLE"
	UNKNOWN FailureOneOf = "UNKNOWN"
	UE_TEMP_UNREACHABLE FailureOneOf = "UE_TEMP_UNREACHABLE"
)

// All allowed values of FailureOneOf enum
var AllowedFailureOneOfEnumValues = []FailureOneOf{
	"UNSPECIFIED",
	"UE_NOT_REACHABLE",
	"UNKNOWN",
	"UE_TEMP_UNREACHABLE",
}

func (v *FailureOneOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := FailureOneOf(value)
	for _, existing := range AllowedFailureOneOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid FailureOneOf", value)
}

// NewFailureOneOfFromValue returns a pointer to a valid FailureOneOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewFailureOneOfFromValue(v string) (*FailureOneOf, error) {
	ev := FailureOneOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for FailureOneOf: valid values are %v", v, AllowedFailureOneOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v FailureOneOf) IsValid() bool {
	for _, existing := range AllowedFailureOneOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Failure_oneOf value
func (v FailureOneOf) Ptr() *FailureOneOf {
	return &v
}

type NullableFailureOneOf struct {
	value *FailureOneOf
	isSet bool
}

func (v NullableFailureOneOf) Get() *FailureOneOf {
	return v.value
}

func (v *NullableFailureOneOf) Set(val *FailureOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureOneOf(val *FailureOneOf) *NullableFailureOneOf {
	return &NullableFailureOneOf{value: val, isSet: true}
}

func (v NullableFailureOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

