/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// CodeWordIndAnyOf the model 'CodeWordIndAnyOf'
type CodeWordIndAnyOf string

// List of CodeWordInd_anyOf
const (
	UE CodeWordIndAnyOf = "CODEWORD_CHECK_IN_UE"
	GMLC CodeWordIndAnyOf = "CODEWORD_CHECK_IN_GMLC"
)

// All allowed values of CodeWordIndAnyOf enum
var AllowedCodeWordIndAnyOfEnumValues = []CodeWordIndAnyOf{
	"CODEWORD_CHECK_IN_UE",
	"CODEWORD_CHECK_IN_GMLC",
}

func (v *CodeWordIndAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CodeWordIndAnyOf(value)
	for _, existing := range AllowedCodeWordIndAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CodeWordIndAnyOf", value)
}

// NewCodeWordIndAnyOfFromValue returns a pointer to a valid CodeWordIndAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCodeWordIndAnyOfFromValue(v string) (*CodeWordIndAnyOf, error) {
	ev := CodeWordIndAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CodeWordIndAnyOf: valid values are %v", v, AllowedCodeWordIndAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CodeWordIndAnyOf) IsValid() bool {
	for _, existing := range AllowedCodeWordIndAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CodeWordInd_anyOf value
func (v CodeWordIndAnyOf) Ptr() *CodeWordIndAnyOf {
	return &v
}

type NullableCodeWordIndAnyOf struct {
	value *CodeWordIndAnyOf
	isSet bool
}

func (v NullableCodeWordIndAnyOf) Get() *CodeWordIndAnyOf {
	return v.value
}

func (v *NullableCodeWordIndAnyOf) Set(val *CodeWordIndAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCodeWordIndAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCodeWordIndAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCodeWordIndAnyOf(val *CodeWordIndAnyOf) *NullableCodeWordIndAnyOf {
	return &NullableCodeWordIndAnyOf{value: val, isSet: true}
}

func (v NullableCodeWordIndAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCodeWordIndAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
