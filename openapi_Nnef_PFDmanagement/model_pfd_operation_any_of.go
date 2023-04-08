/*
Nnef_PFDmanagement Service API

Packet Flow Description Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_PFDmanagement

import (
	"encoding/json"
	"fmt"
)

// PfdOperationAnyOf the model 'PfdOperationAnyOf'
type PfdOperationAnyOf string

// List of PfdOperation_anyOf
const (
	RETRIEVE PfdOperationAnyOf = "RETRIEVE"
	FULLPULL PfdOperationAnyOf = "FULLPULL"
	PARTIALPULL PfdOperationAnyOf = "PARTIALPULL"
	REMOVE PfdOperationAnyOf = "REMOVE"
)

// All allowed values of PfdOperationAnyOf enum
var AllowedPfdOperationAnyOfEnumValues = []PfdOperationAnyOf{
	"RETRIEVE",
	"FULLPULL",
	"PARTIALPULL",
	"REMOVE",
}

func (v *PfdOperationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PfdOperationAnyOf(value)
	for _, existing := range AllowedPfdOperationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PfdOperationAnyOf", value)
}

// NewPfdOperationAnyOfFromValue returns a pointer to a valid PfdOperationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPfdOperationAnyOfFromValue(v string) (*PfdOperationAnyOf, error) {
	ev := PfdOperationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PfdOperationAnyOf: valid values are %v", v, AllowedPfdOperationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PfdOperationAnyOf) IsValid() bool {
	for _, existing := range AllowedPfdOperationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PfdOperation_anyOf value
func (v PfdOperationAnyOf) Ptr() *PfdOperationAnyOf {
	return &v
}

type NullablePfdOperationAnyOf struct {
	value *PfdOperationAnyOf
	isSet bool
}

func (v NullablePfdOperationAnyOf) Get() *PfdOperationAnyOf {
	return v.value
}

func (v *NullablePfdOperationAnyOf) Set(val *PfdOperationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePfdOperationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePfdOperationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePfdOperationAnyOf(val *PfdOperationAnyOf) *NullablePfdOperationAnyOf {
	return &NullablePfdOperationAnyOf{value: val, isSet: true}
}

func (v NullablePfdOperationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePfdOperationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

