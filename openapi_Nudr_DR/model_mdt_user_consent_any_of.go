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

// MdtUserConsentAnyOf the model 'MdtUserConsentAnyOf'
type MdtUserConsentAnyOf string

// List of MdtUserConsent_anyOf
const (
	NOT_GIVEN MdtUserConsentAnyOf = "CONSENT_NOT_GIVEN"
	GIVEN MdtUserConsentAnyOf = "CONSENT_GIVEN"
)

// All allowed values of MdtUserConsentAnyOf enum
var AllowedMdtUserConsentAnyOfEnumValues = []MdtUserConsentAnyOf{
	"CONSENT_NOT_GIVEN",
	"CONSENT_GIVEN",
}

func (v *MdtUserConsentAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MdtUserConsentAnyOf(value)
	for _, existing := range AllowedMdtUserConsentAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MdtUserConsentAnyOf", value)
}

// NewMdtUserConsentAnyOfFromValue returns a pointer to a valid MdtUserConsentAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMdtUserConsentAnyOfFromValue(v string) (*MdtUserConsentAnyOf, error) {
	ev := MdtUserConsentAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MdtUserConsentAnyOf: valid values are %v", v, AllowedMdtUserConsentAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MdtUserConsentAnyOf) IsValid() bool {
	for _, existing := range AllowedMdtUserConsentAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MdtUserConsent_anyOf value
func (v MdtUserConsentAnyOf) Ptr() *MdtUserConsentAnyOf {
	return &v
}

type NullableMdtUserConsentAnyOf struct {
	value *MdtUserConsentAnyOf
	isSet bool
}

func (v NullableMdtUserConsentAnyOf) Get() *MdtUserConsentAnyOf {
	return v.value
}

func (v *NullableMdtUserConsentAnyOf) Set(val *MdtUserConsentAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMdtUserConsentAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMdtUserConsentAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMdtUserConsentAnyOf(val *MdtUserConsentAnyOf) *NullableMdtUserConsentAnyOf {
	return &NullableMdtUserConsentAnyOf{value: val, isSet: true}
}

func (v NullableMdtUserConsentAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMdtUserConsentAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

