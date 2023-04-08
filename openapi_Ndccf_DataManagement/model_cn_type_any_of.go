/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// CnTypeAnyOf the model 'CnTypeAnyOf'
type CnTypeAnyOf string

// List of CnType_anyOf
const (
	SINGLE_4_G CnTypeAnyOf = "SINGLE_4G"
	SINGLE_5_G CnTypeAnyOf = "SINGLE_5G"
	DUAL_4_G5_G CnTypeAnyOf = "DUAL_4G5G"
)

// All allowed values of CnTypeAnyOf enum
var AllowedCnTypeAnyOfEnumValues = []CnTypeAnyOf{
	"SINGLE_4G",
	"SINGLE_5G",
	"DUAL_4G5G",
}

func (v *CnTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CnTypeAnyOf(value)
	for _, existing := range AllowedCnTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CnTypeAnyOf", value)
}

// NewCnTypeAnyOfFromValue returns a pointer to a valid CnTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCnTypeAnyOfFromValue(v string) (*CnTypeAnyOf, error) {
	ev := CnTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CnTypeAnyOf: valid values are %v", v, AllowedCnTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CnTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedCnTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CnType_anyOf value
func (v CnTypeAnyOf) Ptr() *CnTypeAnyOf {
	return &v
}

type NullableCnTypeAnyOf struct {
	value *CnTypeAnyOf
	isSet bool
}

func (v NullableCnTypeAnyOf) Get() *CnTypeAnyOf {
	return v.value
}

func (v *NullableCnTypeAnyOf) Set(val *CnTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCnTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCnTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnTypeAnyOf(val *CnTypeAnyOf) *NullableCnTypeAnyOf {
	return &NullableCnTypeAnyOf{value: val, isSet: true}
}

func (v NullableCnTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

