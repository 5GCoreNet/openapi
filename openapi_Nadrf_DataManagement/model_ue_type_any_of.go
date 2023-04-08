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

// UeTypeAnyOf the model 'UeTypeAnyOf'
type UeTypeAnyOf string

// List of UeType_anyOf
const (
	AERIAL_UE UeTypeAnyOf = "AERIAL_UE"
)

// All allowed values of UeTypeAnyOf enum
var AllowedUeTypeAnyOfEnumValues = []UeTypeAnyOf{
	"AERIAL_UE",
}

func (v *UeTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UeTypeAnyOf(value)
	for _, existing := range AllowedUeTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UeTypeAnyOf", value)
}

// NewUeTypeAnyOfFromValue returns a pointer to a valid UeTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUeTypeAnyOfFromValue(v string) (*UeTypeAnyOf, error) {
	ev := UeTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UeTypeAnyOf: valid values are %v", v, AllowedUeTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UeTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedUeTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UeType_anyOf value
func (v UeTypeAnyOf) Ptr() *UeTypeAnyOf {
	return &v
}

type NullableUeTypeAnyOf struct {
	value *UeTypeAnyOf
	isSet bool
}

func (v NullableUeTypeAnyOf) Get() *UeTypeAnyOf {
	return v.value
}

func (v *NullableUeTypeAnyOf) Set(val *UeTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUeTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUeTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeTypeAnyOf(val *UeTypeAnyOf) *NullableUeTypeAnyOf {
	return &NullableUeTypeAnyOf{value: val, isSet: true}
}

func (v NullableUeTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

