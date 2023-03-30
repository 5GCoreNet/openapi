/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// AccessType Indicates whether the access is  via 3GPP or via non-3GPP.
type AccessType string

// List of AccessType
const (
	_3_GPP_ACCESS AccessType = "3GPP_ACCESS"
	NON_3_GPP_ACCESS AccessType = "NON_3GPP_ACCESS"
)

// All allowed values of AccessType enum
var AllowedAccessTypeEnumValues = []AccessType{
	"3GPP_ACCESS",
	"NON_3GPP_ACCESS",
}

func (v *AccessType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessType(value)
	for _, existing := range AllowedAccessTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessType", value)
}

// NewAccessTypeFromValue returns a pointer to a valid AccessType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessTypeFromValue(v string) (*AccessType, error) {
	ev := AccessType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessType: valid values are %v", v, AllowedAccessTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessType) IsValid() bool {
	for _, existing := range AllowedAccessTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessType value
func (v AccessType) Ptr() *AccessType {
	return &v
}

type NullableAccessType struct {
	value *AccessType
	isSet bool
}

func (v NullableAccessType) Get() *AccessType {
	return v.value
}

func (v *NullableAccessType) Set(val *AccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessType(val *AccessType) *NullableAccessType {
	return &NullableAccessType{value: val, isSet: true}
}

func (v NullableAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

