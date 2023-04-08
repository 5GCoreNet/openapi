/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// LocationRequestTypeAnyOf the model 'LocationRequestTypeAnyOf'
type LocationRequestTypeAnyOf string

// List of LocationRequestType_anyOf
const (
	NI_LR LocationRequestTypeAnyOf = "NI_LR"
	MT_LR LocationRequestTypeAnyOf = "MT_LR"
	MO_LR LocationRequestTypeAnyOf = "MO_LR"
)

// All allowed values of LocationRequestTypeAnyOf enum
var AllowedLocationRequestTypeAnyOfEnumValues = []LocationRequestTypeAnyOf{
	"NI_LR",
	"MT_LR",
	"MO_LR",
}

func (v *LocationRequestTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocationRequestTypeAnyOf(value)
	for _, existing := range AllowedLocationRequestTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocationRequestTypeAnyOf", value)
}

// NewLocationRequestTypeAnyOfFromValue returns a pointer to a valid LocationRequestTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocationRequestTypeAnyOfFromValue(v string) (*LocationRequestTypeAnyOf, error) {
	ev := LocationRequestTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocationRequestTypeAnyOf: valid values are %v", v, AllowedLocationRequestTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocationRequestTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedLocationRequestTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocationRequestType_anyOf value
func (v LocationRequestTypeAnyOf) Ptr() *LocationRequestTypeAnyOf {
	return &v
}

type NullableLocationRequestTypeAnyOf struct {
	value *LocationRequestTypeAnyOf
	isSet bool
}

func (v NullableLocationRequestTypeAnyOf) Get() *LocationRequestTypeAnyOf {
	return v.value
}

func (v *NullableLocationRequestTypeAnyOf) Set(val *LocationRequestTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationRequestTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationRequestTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationRequestTypeAnyOf(val *LocationRequestTypeAnyOf) *NullableLocationRequestTypeAnyOf {
	return &NullableLocationRequestTypeAnyOf{value: val, isSet: true}
}

func (v NullableLocationRequestTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationRequestTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

