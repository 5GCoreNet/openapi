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

// UuaaMmStatusAnyOf the model 'UuaaMmStatusAnyOf'
type UuaaMmStatusAnyOf string

// List of UuaaMmStatus_anyOf
const (
	SUCCESS UuaaMmStatusAnyOf = "SUCCESS"
	PENDING UuaaMmStatusAnyOf = "PENDING"
	FAILED UuaaMmStatusAnyOf = "FAILED"
)

// All allowed values of UuaaMmStatusAnyOf enum
var AllowedUuaaMmStatusAnyOfEnumValues = []UuaaMmStatusAnyOf{
	"SUCCESS",
	"PENDING",
	"FAILED",
}

func (v *UuaaMmStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UuaaMmStatusAnyOf(value)
	for _, existing := range AllowedUuaaMmStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UuaaMmStatusAnyOf", value)
}

// NewUuaaMmStatusAnyOfFromValue returns a pointer to a valid UuaaMmStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUuaaMmStatusAnyOfFromValue(v string) (*UuaaMmStatusAnyOf, error) {
	ev := UuaaMmStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UuaaMmStatusAnyOf: valid values are %v", v, AllowedUuaaMmStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UuaaMmStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedUuaaMmStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UuaaMmStatus_anyOf value
func (v UuaaMmStatusAnyOf) Ptr() *UuaaMmStatusAnyOf {
	return &v
}

type NullableUuaaMmStatusAnyOf struct {
	value *UuaaMmStatusAnyOf
	isSet bool
}

func (v NullableUuaaMmStatusAnyOf) Get() *UuaaMmStatusAnyOf {
	return v.value
}

func (v *NullableUuaaMmStatusAnyOf) Set(val *UuaaMmStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUuaaMmStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUuaaMmStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUuaaMmStatusAnyOf(val *UuaaMmStatusAnyOf) *NullableUuaaMmStatusAnyOf {
	return &NullableUuaaMmStatusAnyOf{value: val, isSet: true}
}

func (v NullableUuaaMmStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUuaaMmStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

