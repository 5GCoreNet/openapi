/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbstf_DistSession

import (
	"encoding/json"
	"fmt"
)

// DistSessionStateAnyOf the model 'DistSessionStateAnyOf'
type DistSessionStateAnyOf string

// List of DistSessionState_anyOf
const (
	INACTIVE DistSessionStateAnyOf = "INACTIVE"
	ESTABLISHED DistSessionStateAnyOf = "ESTABLISHED"
	ACTIVE DistSessionStateAnyOf = "ACTIVE"
	DEACTIVATING DistSessionStateAnyOf = "DEACTIVATING"
)

// All allowed values of DistSessionStateAnyOf enum
var AllowedDistSessionStateAnyOfEnumValues = []DistSessionStateAnyOf{
	"INACTIVE",
	"ESTABLISHED",
	"ACTIVE",
	"DEACTIVATING",
}

func (v *DistSessionStateAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DistSessionStateAnyOf(value)
	for _, existing := range AllowedDistSessionStateAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DistSessionStateAnyOf", value)
}

// NewDistSessionStateAnyOfFromValue returns a pointer to a valid DistSessionStateAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDistSessionStateAnyOfFromValue(v string) (*DistSessionStateAnyOf, error) {
	ev := DistSessionStateAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DistSessionStateAnyOf: valid values are %v", v, AllowedDistSessionStateAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DistSessionStateAnyOf) IsValid() bool {
	for _, existing := range AllowedDistSessionStateAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DistSessionState_anyOf value
func (v DistSessionStateAnyOf) Ptr() *DistSessionStateAnyOf {
	return &v
}

type NullableDistSessionStateAnyOf struct {
	value *DistSessionStateAnyOf
	isSet bool
}

func (v NullableDistSessionStateAnyOf) Get() *DistSessionStateAnyOf {
	return v.value
}

func (v *NullableDistSessionStateAnyOf) Set(val *DistSessionStateAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDistSessionStateAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDistSessionStateAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDistSessionStateAnyOf(val *DistSessionStateAnyOf) *NullableDistSessionStateAnyOf {
	return &NullableDistSessionStateAnyOf{value: val, isSet: true}
}

func (v NullableDistSessionStateAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDistSessionStateAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

