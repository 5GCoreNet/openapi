/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UECM

import (
	"encoding/json"
	"fmt"
)

// UeReachableIndAnyOf the model 'UeReachableIndAnyOf'
type UeReachableIndAnyOf string

// List of UeReachableInd_anyOf
const (
	REACHABLE UeReachableIndAnyOf = "REACHABLE"
	NOT_REACHABLE UeReachableIndAnyOf = "NOT_REACHABLE"
	UNKNOWN UeReachableIndAnyOf = "UNKNOWN"
)

// All allowed values of UeReachableIndAnyOf enum
var AllowedUeReachableIndAnyOfEnumValues = []UeReachableIndAnyOf{
	"REACHABLE",
	"NOT_REACHABLE",
	"UNKNOWN",
}

func (v *UeReachableIndAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UeReachableIndAnyOf(value)
	for _, existing := range AllowedUeReachableIndAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UeReachableIndAnyOf", value)
}

// NewUeReachableIndAnyOfFromValue returns a pointer to a valid UeReachableIndAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUeReachableIndAnyOfFromValue(v string) (*UeReachableIndAnyOf, error) {
	ev := UeReachableIndAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UeReachableIndAnyOf: valid values are %v", v, AllowedUeReachableIndAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UeReachableIndAnyOf) IsValid() bool {
	for _, existing := range AllowedUeReachableIndAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UeReachableInd_anyOf value
func (v UeReachableIndAnyOf) Ptr() *UeReachableIndAnyOf {
	return &v
}

type NullableUeReachableIndAnyOf struct {
	value *UeReachableIndAnyOf
	isSet bool
}

func (v NullableUeReachableIndAnyOf) Get() *UeReachableIndAnyOf {
	return v.value
}

func (v *NullableUeReachableIndAnyOf) Set(val *UeReachableIndAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUeReachableIndAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUeReachableIndAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeReachableIndAnyOf(val *UeReachableIndAnyOf) *NullableUeReachableIndAnyOf {
	return &NullableUeReachableIndAnyOf{value: val, isSet: true}
}

func (v NullableUeReachableIndAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeReachableIndAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

