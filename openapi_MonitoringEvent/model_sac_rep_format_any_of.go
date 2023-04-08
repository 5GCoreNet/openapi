/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// SACRepFormatAnyOf the model 'SACRepFormatAnyOf'
type SACRepFormatAnyOf string

// List of SACRepFormat_anyOf
const (
	NUMERICAL SACRepFormatAnyOf = "NUMERICAL"
	PERCENTAGE SACRepFormatAnyOf = "PERCENTAGE"
)

// All allowed values of SACRepFormatAnyOf enum
var AllowedSACRepFormatAnyOfEnumValues = []SACRepFormatAnyOf{
	"NUMERICAL",
	"PERCENTAGE",
}

func (v *SACRepFormatAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SACRepFormatAnyOf(value)
	for _, existing := range AllowedSACRepFormatAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SACRepFormatAnyOf", value)
}

// NewSACRepFormatAnyOfFromValue returns a pointer to a valid SACRepFormatAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSACRepFormatAnyOfFromValue(v string) (*SACRepFormatAnyOf, error) {
	ev := SACRepFormatAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SACRepFormatAnyOf: valid values are %v", v, AllowedSACRepFormatAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SACRepFormatAnyOf) IsValid() bool {
	for _, existing := range AllowedSACRepFormatAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SACRepFormat_anyOf value
func (v SACRepFormatAnyOf) Ptr() *SACRepFormatAnyOf {
	return &v
}

type NullableSACRepFormatAnyOf struct {
	value *SACRepFormatAnyOf
	isSet bool
}

func (v NullableSACRepFormatAnyOf) Get() *SACRepFormatAnyOf {
	return v.value
}

func (v *NullableSACRepFormatAnyOf) Set(val *SACRepFormatAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSACRepFormatAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSACRepFormatAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACRepFormatAnyOf(val *SACRepFormatAnyOf) *NullableSACRepFormatAnyOf {
	return &NullableSACRepFormatAnyOf{value: val, isSet: true}
}

func (v NullableSACRepFormatAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACRepFormatAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

