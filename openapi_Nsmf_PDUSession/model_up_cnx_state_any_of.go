/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// UpCnxStateAnyOf the model 'UpCnxStateAnyOf'
type UpCnxStateAnyOf string

// List of UpCnxState_anyOf
const (
	ACTIVATED UpCnxStateAnyOf = "ACTIVATED"
	DEACTIVATED UpCnxStateAnyOf = "DEACTIVATED"
	ACTIVATING UpCnxStateAnyOf = "ACTIVATING"
	SUSPENDED UpCnxStateAnyOf = "SUSPENDED"
)

// All allowed values of UpCnxStateAnyOf enum
var AllowedUpCnxStateAnyOfEnumValues = []UpCnxStateAnyOf{
	"ACTIVATED",
	"DEACTIVATED",
	"ACTIVATING",
	"SUSPENDED",
}

func (v *UpCnxStateAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := UpCnxStateAnyOf(value)
	for _, existing := range AllowedUpCnxStateAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid UpCnxStateAnyOf", value)
}

// NewUpCnxStateAnyOfFromValue returns a pointer to a valid UpCnxStateAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewUpCnxStateAnyOfFromValue(v string) (*UpCnxStateAnyOf, error) {
	ev := UpCnxStateAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for UpCnxStateAnyOf: valid values are %v", v, AllowedUpCnxStateAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v UpCnxStateAnyOf) IsValid() bool {
	for _, existing := range AllowedUpCnxStateAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to UpCnxState_anyOf value
func (v UpCnxStateAnyOf) Ptr() *UpCnxStateAnyOf {
	return &v
}

type NullableUpCnxStateAnyOf struct {
	value *UpCnxStateAnyOf
	isSet bool
}

func (v NullableUpCnxStateAnyOf) Get() *UpCnxStateAnyOf {
	return v.value
}

func (v *NullableUpCnxStateAnyOf) Set(val *UpCnxStateAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableUpCnxStateAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUpCnxStateAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpCnxStateAnyOf(val *UpCnxStateAnyOf) *NullableUpCnxStateAnyOf {
	return &NullableUpCnxStateAnyOf{value: val, isSet: true}
}

func (v NullableUpCnxStateAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpCnxStateAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

