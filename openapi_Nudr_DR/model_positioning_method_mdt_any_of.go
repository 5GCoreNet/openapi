/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// PositioningMethodMdtAnyOf the model 'PositioningMethodMdtAnyOf'
type PositioningMethodMdtAnyOf string

// List of PositioningMethodMdt_anyOf
const (
	GNSS PositioningMethodMdtAnyOf = "GNSS"
	E_CELL_ID PositioningMethodMdtAnyOf = "E_CELL_ID"
)

// All allowed values of PositioningMethodMdtAnyOf enum
var AllowedPositioningMethodMdtAnyOfEnumValues = []PositioningMethodMdtAnyOf{
	"GNSS",
	"E_CELL_ID",
}

func (v *PositioningMethodMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PositioningMethodMdtAnyOf(value)
	for _, existing := range AllowedPositioningMethodMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PositioningMethodMdtAnyOf", value)
}

// NewPositioningMethodMdtAnyOfFromValue returns a pointer to a valid PositioningMethodMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPositioningMethodMdtAnyOfFromValue(v string) (*PositioningMethodMdtAnyOf, error) {
	ev := PositioningMethodMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PositioningMethodMdtAnyOf: valid values are %v", v, AllowedPositioningMethodMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PositioningMethodMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedPositioningMethodMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PositioningMethodMdt_anyOf value
func (v PositioningMethodMdtAnyOf) Ptr() *PositioningMethodMdtAnyOf {
	return &v
}

type NullablePositioningMethodMdtAnyOf struct {
	value *PositioningMethodMdtAnyOf
	isSet bool
}

func (v NullablePositioningMethodMdtAnyOf) Get() *PositioningMethodMdtAnyOf {
	return v.value
}

func (v *NullablePositioningMethodMdtAnyOf) Set(val *PositioningMethodMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningMethodMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningMethodMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningMethodMdtAnyOf(val *PositioningMethodMdtAnyOf) *NullablePositioningMethodMdtAnyOf {
	return &NullablePositioningMethodMdtAnyOf{value: val, isSet: true}
}

func (v NullablePositioningMethodMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningMethodMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

