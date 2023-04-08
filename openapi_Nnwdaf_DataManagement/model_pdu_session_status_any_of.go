/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// PduSessionStatusAnyOf the model 'PduSessionStatusAnyOf'
type PduSessionStatusAnyOf string

// List of PduSessionStatus_anyOf
const (
	ACTIVATED PduSessionStatusAnyOf = "ACTIVATED"
	DEACTIVATED PduSessionStatusAnyOf = "DEACTIVATED"
)

// All allowed values of PduSessionStatusAnyOf enum
var AllowedPduSessionStatusAnyOfEnumValues = []PduSessionStatusAnyOf{
	"ACTIVATED",
	"DEACTIVATED",
}

func (v *PduSessionStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PduSessionStatusAnyOf(value)
	for _, existing := range AllowedPduSessionStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PduSessionStatusAnyOf", value)
}

// NewPduSessionStatusAnyOfFromValue returns a pointer to a valid PduSessionStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPduSessionStatusAnyOfFromValue(v string) (*PduSessionStatusAnyOf, error) {
	ev := PduSessionStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PduSessionStatusAnyOf: valid values are %v", v, AllowedPduSessionStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PduSessionStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedPduSessionStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PduSessionStatus_anyOf value
func (v PduSessionStatusAnyOf) Ptr() *PduSessionStatusAnyOf {
	return &v
}

type NullablePduSessionStatusAnyOf struct {
	value *PduSessionStatusAnyOf
	isSet bool
}

func (v NullablePduSessionStatusAnyOf) Get() *PduSessionStatusAnyOf {
	return v.value
}

func (v *NullablePduSessionStatusAnyOf) Set(val *PduSessionStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionStatusAnyOf(val *PduSessionStatusAnyOf) *NullablePduSessionStatusAnyOf {
	return &NullablePduSessionStatusAnyOf{value: val, isSet: true}
}

func (v NullablePduSessionStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

