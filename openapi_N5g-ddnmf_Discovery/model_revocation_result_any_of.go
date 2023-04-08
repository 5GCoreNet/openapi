/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-ddnmf_Discovery

import (
	"encoding/json"
	"fmt"
)

// RevocationResultAnyOf the model 'RevocationResultAnyOf'
type RevocationResultAnyOf string

// List of RevocationResult_anyOf
const (
	SUCCESSFUL RevocationResultAnyOf = "SUCCESSFUL"
	FAILED RevocationResultAnyOf = "FAILED"
)

// All allowed values of RevocationResultAnyOf enum
var AllowedRevocationResultAnyOfEnumValues = []RevocationResultAnyOf{
	"SUCCESSFUL",
	"FAILED",
}

func (v *RevocationResultAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RevocationResultAnyOf(value)
	for _, existing := range AllowedRevocationResultAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RevocationResultAnyOf", value)
}

// NewRevocationResultAnyOfFromValue returns a pointer to a valid RevocationResultAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRevocationResultAnyOfFromValue(v string) (*RevocationResultAnyOf, error) {
	ev := RevocationResultAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RevocationResultAnyOf: valid values are %v", v, AllowedRevocationResultAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RevocationResultAnyOf) IsValid() bool {
	for _, existing := range AllowedRevocationResultAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RevocationResult_anyOf value
func (v RevocationResultAnyOf) Ptr() *RevocationResultAnyOf {
	return &v
}

type NullableRevocationResultAnyOf struct {
	value *RevocationResultAnyOf
	isSet bool
}

func (v NullableRevocationResultAnyOf) Get() *RevocationResultAnyOf {
	return v.value
}

func (v *NullableRevocationResultAnyOf) Set(val *RevocationResultAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRevocationResultAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRevocationResultAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevocationResultAnyOf(val *RevocationResultAnyOf) *NullableRevocationResultAnyOf {
	return &NullableRevocationResultAnyOf{value: val, isSet: true}
}

func (v NullableRevocationResultAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevocationResultAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

