/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
	"fmt"
)

// RevokedCauseAnyOf the model 'RevokedCauseAnyOf'
type RevokedCauseAnyOf string

// List of RevokedCause_anyOf
const (
	NOT_ALLOWED RevokedCauseAnyOf = "NOT_ALLOWED"
	EXCLUDED_FROM_GROUP RevokedCauseAnyOf = "EXCLUDED_FROM_GROUP"
	GPSI_REMOVED RevokedCauseAnyOf = "GPSI_REMOVED"
)

// All allowed values of RevokedCauseAnyOf enum
var AllowedRevokedCauseAnyOfEnumValues = []RevokedCauseAnyOf{
	"NOT_ALLOWED",
	"EXCLUDED_FROM_GROUP",
	"GPSI_REMOVED",
}

func (v *RevokedCauseAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RevokedCauseAnyOf(value)
	for _, existing := range AllowedRevokedCauseAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RevokedCauseAnyOf", value)
}

// NewRevokedCauseAnyOfFromValue returns a pointer to a valid RevokedCauseAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRevokedCauseAnyOfFromValue(v string) (*RevokedCauseAnyOf, error) {
	ev := RevokedCauseAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RevokedCauseAnyOf: valid values are %v", v, AllowedRevokedCauseAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RevokedCauseAnyOf) IsValid() bool {
	for _, existing := range AllowedRevokedCauseAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RevokedCause_anyOf value
func (v RevokedCauseAnyOf) Ptr() *RevokedCauseAnyOf {
	return &v
}

type NullableRevokedCauseAnyOf struct {
	value *RevokedCauseAnyOf
	isSet bool
}

func (v NullableRevokedCauseAnyOf) Get() *RevokedCauseAnyOf {
	return v.value
}

func (v *NullableRevokedCauseAnyOf) Set(val *RevokedCauseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokedCauseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokedCauseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokedCauseAnyOf(val *RevokedCauseAnyOf) *NullableRevokedCauseAnyOf {
	return &NullableRevokedCauseAnyOf{value: val, isSet: true}
}

func (v NullableRevokedCauseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokedCauseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

