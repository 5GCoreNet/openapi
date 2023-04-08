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

// TermCauseAnyOf the model 'TermCauseAnyOf'
type TermCauseAnyOf string

// List of TermCause_anyOf
const (
	USER_CONSENT_REVOKED TermCauseAnyOf = "USER_CONSENT_REVOKED"
	NWDAF_OVERLOAD TermCauseAnyOf = "NWDAF_OVERLOAD"
	UE_LEFT_AREA TermCauseAnyOf = "UE_LEFT_AREA"
)

// All allowed values of TermCauseAnyOf enum
var AllowedTermCauseAnyOfEnumValues = []TermCauseAnyOf{
	"USER_CONSENT_REVOKED",
	"NWDAF_OVERLOAD",
	"UE_LEFT_AREA",
}

func (v *TermCauseAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TermCauseAnyOf(value)
	for _, existing := range AllowedTermCauseAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TermCauseAnyOf", value)
}

// NewTermCauseAnyOfFromValue returns a pointer to a valid TermCauseAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTermCauseAnyOfFromValue(v string) (*TermCauseAnyOf, error) {
	ev := TermCauseAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TermCauseAnyOf: valid values are %v", v, AllowedTermCauseAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TermCauseAnyOf) IsValid() bool {
	for _, existing := range AllowedTermCauseAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TermCause_anyOf value
func (v TermCauseAnyOf) Ptr() *TermCauseAnyOf {
	return &v
}

type NullableTermCauseAnyOf struct {
	value *TermCauseAnyOf
	isSet bool
}

func (v NullableTermCauseAnyOf) Get() *TermCauseAnyOf {
	return v.value
}

func (v *NullableTermCauseAnyOf) Set(val *TermCauseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableTermCauseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableTermCauseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermCauseAnyOf(val *TermCauseAnyOf) *NullableTermCauseAnyOf {
	return &NullableTermCauseAnyOf{value: val, isSet: true}
}

func (v NullableTermCauseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermCauseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

