/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
	"fmt"
)

// GbaTypeAnyOf the model 'GbaTypeAnyOf'
type GbaTypeAnyOf string

// List of GbaType_anyOf
const (
	_3_G_GBA GbaTypeAnyOf = "3G_GBA"
	_2_G_GBA GbaTypeAnyOf = "2G_GBA"
	GBA_DIGEST GbaTypeAnyOf = "GBA_DIGEST"
)

// All allowed values of GbaTypeAnyOf enum
var AllowedGbaTypeAnyOfEnumValues = []GbaTypeAnyOf{
	"3G_GBA",
	"2G_GBA",
	"GBA_DIGEST",
}

func (v *GbaTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := GbaTypeAnyOf(value)
	for _, existing := range AllowedGbaTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid GbaTypeAnyOf", value)
}

// NewGbaTypeAnyOfFromValue returns a pointer to a valid GbaTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewGbaTypeAnyOfFromValue(v string) (*GbaTypeAnyOf, error) {
	ev := GbaTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for GbaTypeAnyOf: valid values are %v", v, AllowedGbaTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v GbaTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedGbaTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to GbaType_anyOf value
func (v GbaTypeAnyOf) Ptr() *GbaTypeAnyOf {
	return &v
}

type NullableGbaTypeAnyOf struct {
	value *GbaTypeAnyOf
	isSet bool
}

func (v NullableGbaTypeAnyOf) Get() *GbaTypeAnyOf {
	return v.value
}

func (v *NullableGbaTypeAnyOf) Set(val *GbaTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableGbaTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableGbaTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGbaTypeAnyOf(val *GbaTypeAnyOf) *NullableGbaTypeAnyOf {
	return &NullableGbaTypeAnyOf{value: val, isSet: true}
}

func (v NullableGbaTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGbaTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
