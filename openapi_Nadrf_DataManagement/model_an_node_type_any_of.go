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

// AnNodeTypeAnyOf the model 'AnNodeTypeAnyOf'
type AnNodeTypeAnyOf string

// List of AnNodeType_anyOf
const (
	GNB AnNodeTypeAnyOf = "GNB"
	NG_ENB AnNodeTypeAnyOf = "NG_ENB"
)

// All allowed values of AnNodeTypeAnyOf enum
var AllowedAnNodeTypeAnyOfEnumValues = []AnNodeTypeAnyOf{
	"GNB",
	"NG_ENB",
}

func (v *AnNodeTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnNodeTypeAnyOf(value)
	for _, existing := range AllowedAnNodeTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnNodeTypeAnyOf", value)
}

// NewAnNodeTypeAnyOfFromValue returns a pointer to a valid AnNodeTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnNodeTypeAnyOfFromValue(v string) (*AnNodeTypeAnyOf, error) {
	ev := AnNodeTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnNodeTypeAnyOf: valid values are %v", v, AllowedAnNodeTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnNodeTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAnNodeTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnNodeType_anyOf value
func (v AnNodeTypeAnyOf) Ptr() *AnNodeTypeAnyOf {
	return &v
}

type NullableAnNodeTypeAnyOf struct {
	value *AnNodeTypeAnyOf
	isSet bool
}

func (v NullableAnNodeTypeAnyOf) Get() *AnNodeTypeAnyOf {
	return v.value
}

func (v *NullableAnNodeTypeAnyOf) Set(val *AnNodeTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnNodeTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnNodeTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnNodeTypeAnyOf(val *AnNodeTypeAnyOf) *NullableAnNodeTypeAnyOf {
	return &NullableAnNodeTypeAnyOf{value: val, isSet: true}
}

func (v NullableAnNodeTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnNodeTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

