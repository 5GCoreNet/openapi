/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// MICOModeIndicationAnyOf the model 'MICOModeIndicationAnyOf'
type MICOModeIndicationAnyOf string

// List of MICOModeIndication_anyOf
const (
	MICO_MODE MICOModeIndicationAnyOf = "MICO_MODE"
	NO_MICO_MODE MICOModeIndicationAnyOf = "NO_MICO_MODE"
)

// All allowed values of MICOModeIndicationAnyOf enum
var AllowedMICOModeIndicationAnyOfEnumValues = []MICOModeIndicationAnyOf{
	"MICO_MODE",
	"NO_MICO_MODE",
}

func (v *MICOModeIndicationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MICOModeIndicationAnyOf(value)
	for _, existing := range AllowedMICOModeIndicationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MICOModeIndicationAnyOf", value)
}

// NewMICOModeIndicationAnyOfFromValue returns a pointer to a valid MICOModeIndicationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMICOModeIndicationAnyOfFromValue(v string) (*MICOModeIndicationAnyOf, error) {
	ev := MICOModeIndicationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MICOModeIndicationAnyOf: valid values are %v", v, AllowedMICOModeIndicationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MICOModeIndicationAnyOf) IsValid() bool {
	for _, existing := range AllowedMICOModeIndicationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MICOModeIndication_anyOf value
func (v MICOModeIndicationAnyOf) Ptr() *MICOModeIndicationAnyOf {
	return &v
}

type NullableMICOModeIndicationAnyOf struct {
	value *MICOModeIndicationAnyOf
	isSet bool
}

func (v NullableMICOModeIndicationAnyOf) Get() *MICOModeIndicationAnyOf {
	return v.value
}

func (v *NullableMICOModeIndicationAnyOf) Set(val *MICOModeIndicationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMICOModeIndicationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMICOModeIndicationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMICOModeIndicationAnyOf(val *MICOModeIndicationAnyOf) *NullableMICOModeIndicationAnyOf {
	return &NullableMICOModeIndicationAnyOf{value: val, isSet: true}
}

func (v NullableMICOModeIndicationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMICOModeIndicationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

