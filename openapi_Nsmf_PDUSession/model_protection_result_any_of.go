/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// ProtectionResultAnyOf the model 'ProtectionResultAnyOf'
type ProtectionResultAnyOf string

// List of ProtectionResult_anyOf
const (
	PERFORMED ProtectionResultAnyOf = "PERFORMED"
	NOT_PERFORMED ProtectionResultAnyOf = "NOT_PERFORMED"
)

// All allowed values of ProtectionResultAnyOf enum
var AllowedProtectionResultAnyOfEnumValues = []ProtectionResultAnyOf{
	"PERFORMED",
	"NOT_PERFORMED",
}

func (v *ProtectionResultAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ProtectionResultAnyOf(value)
	for _, existing := range AllowedProtectionResultAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ProtectionResultAnyOf", value)
}

// NewProtectionResultAnyOfFromValue returns a pointer to a valid ProtectionResultAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewProtectionResultAnyOfFromValue(v string) (*ProtectionResultAnyOf, error) {
	ev := ProtectionResultAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ProtectionResultAnyOf: valid values are %v", v, AllowedProtectionResultAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ProtectionResultAnyOf) IsValid() bool {
	for _, existing := range AllowedProtectionResultAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ProtectionResult_anyOf value
func (v ProtectionResultAnyOf) Ptr() *ProtectionResultAnyOf {
	return &v
}

type NullableProtectionResultAnyOf struct {
	value *ProtectionResultAnyOf
	isSet bool
}

func (v NullableProtectionResultAnyOf) Get() *ProtectionResultAnyOf {
	return v.value
}

func (v *NullableProtectionResultAnyOf) Set(val *ProtectionResultAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProtectionResultAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProtectionResultAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProtectionResultAnyOf(val *ProtectionResultAnyOf) *NullableProtectionResultAnyOf {
	return &NullableProtectionResultAnyOf{value: val, isSet: true}
}

func (v NullableProtectionResultAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProtectionResultAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

