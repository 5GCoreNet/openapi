/*
NhssUEAU

HSS UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_UEAU

import (
	"encoding/json"
	"fmt"
)

// AvTypeAnyOf the model 'AvTypeAnyOf'
type AvTypeAnyOf string

// List of AvType_anyOf
const (
	_5_G_HE_AKA AvTypeAnyOf = "5G_HE_AKA"
	EAP_AKA_PRIME AvTypeAnyOf = "EAP_AKA_PRIME"
)

// All allowed values of AvTypeAnyOf enum
var AllowedAvTypeAnyOfEnumValues = []AvTypeAnyOf{
	"5G_HE_AKA",
	"EAP_AKA_PRIME",
}

func (v *AvTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AvTypeAnyOf(value)
	for _, existing := range AllowedAvTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AvTypeAnyOf", value)
}

// NewAvTypeAnyOfFromValue returns a pointer to a valid AvTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAvTypeAnyOfFromValue(v string) (*AvTypeAnyOf, error) {
	ev := AvTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AvTypeAnyOf: valid values are %v", v, AllowedAvTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AvTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAvTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AvType_anyOf value
func (v AvTypeAnyOf) Ptr() *AvTypeAnyOf {
	return &v
}

type NullableAvTypeAnyOf struct {
	value *AvTypeAnyOf
	isSet bool
}

func (v NullableAvTypeAnyOf) Get() *AvTypeAnyOf {
	return v.value
}

func (v *NullableAvTypeAnyOf) Set(val *AvTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAvTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAvTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvTypeAnyOf(val *AvTypeAnyOf) *NullableAvTypeAnyOf {
	return &NullableAvTypeAnyOf{value: val, isSet: true}
}

func (v NullableAvTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

