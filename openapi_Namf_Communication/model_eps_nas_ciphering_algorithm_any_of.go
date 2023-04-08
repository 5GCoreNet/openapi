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

// EpsNasCipheringAlgorithmAnyOf the model 'EpsNasCipheringAlgorithmAnyOf'
type EpsNasCipheringAlgorithmAnyOf string

// List of EpsNasCipheringAlgorithm_anyOf
const (
	EEA0 EpsNasCipheringAlgorithmAnyOf = "EEA0"
	EEA1 EpsNasCipheringAlgorithmAnyOf = "EEA1"
	EEA2 EpsNasCipheringAlgorithmAnyOf = "EEA2"
	EEA3 EpsNasCipheringAlgorithmAnyOf = "EEA3"
)

// All allowed values of EpsNasCipheringAlgorithmAnyOf enum
var AllowedEpsNasCipheringAlgorithmAnyOfEnumValues = []EpsNasCipheringAlgorithmAnyOf{
	"EEA0",
	"EEA1",
	"EEA2",
	"EEA3",
}

func (v *EpsNasCipheringAlgorithmAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EpsNasCipheringAlgorithmAnyOf(value)
	for _, existing := range AllowedEpsNasCipheringAlgorithmAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EpsNasCipheringAlgorithmAnyOf", value)
}

// NewEpsNasCipheringAlgorithmAnyOfFromValue returns a pointer to a valid EpsNasCipheringAlgorithmAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEpsNasCipheringAlgorithmAnyOfFromValue(v string) (*EpsNasCipheringAlgorithmAnyOf, error) {
	ev := EpsNasCipheringAlgorithmAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EpsNasCipheringAlgorithmAnyOf: valid values are %v", v, AllowedEpsNasCipheringAlgorithmAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EpsNasCipheringAlgorithmAnyOf) IsValid() bool {
	for _, existing := range AllowedEpsNasCipheringAlgorithmAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EpsNasCipheringAlgorithm_anyOf value
func (v EpsNasCipheringAlgorithmAnyOf) Ptr() *EpsNasCipheringAlgorithmAnyOf {
	return &v
}

type NullableEpsNasCipheringAlgorithmAnyOf struct {
	value *EpsNasCipheringAlgorithmAnyOf
	isSet bool
}

func (v NullableEpsNasCipheringAlgorithmAnyOf) Get() *EpsNasCipheringAlgorithmAnyOf {
	return v.value
}

func (v *NullableEpsNasCipheringAlgorithmAnyOf) Set(val *EpsNasCipheringAlgorithmAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEpsNasCipheringAlgorithmAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEpsNasCipheringAlgorithmAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEpsNasCipheringAlgorithmAnyOf(val *EpsNasCipheringAlgorithmAnyOf) *NullableEpsNasCipheringAlgorithmAnyOf {
	return &NullableEpsNasCipheringAlgorithmAnyOf{value: val, isSet: true}
}

func (v NullableEpsNasCipheringAlgorithmAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEpsNasCipheringAlgorithmAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

