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

// AssociationTypeAnyOf the model 'AssociationTypeAnyOf'
type AssociationTypeAnyOf string

// List of AssociationType_anyOf
const (
	IMEI_CHANGE AssociationTypeAnyOf = "IMEI_CHANGE"
	IMEISV_CHANGE AssociationTypeAnyOf = "IMEISV_CHANGE"
)

// All allowed values of AssociationTypeAnyOf enum
var AllowedAssociationTypeAnyOfEnumValues = []AssociationTypeAnyOf{
	"IMEI_CHANGE",
	"IMEISV_CHANGE",
}

func (v *AssociationTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssociationTypeAnyOf(value)
	for _, existing := range AllowedAssociationTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssociationTypeAnyOf", value)
}

// NewAssociationTypeAnyOfFromValue returns a pointer to a valid AssociationTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssociationTypeAnyOfFromValue(v string) (*AssociationTypeAnyOf, error) {
	ev := AssociationTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssociationTypeAnyOf: valid values are %v", v, AllowedAssociationTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssociationTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAssociationTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssociationType_anyOf value
func (v AssociationTypeAnyOf) Ptr() *AssociationTypeAnyOf {
	return &v
}

type NullableAssociationTypeAnyOf struct {
	value *AssociationTypeAnyOf
	isSet bool
}

func (v NullableAssociationTypeAnyOf) Get() *AssociationTypeAnyOf {
	return v.value
}

func (v *NullableAssociationTypeAnyOf) Set(val *AssociationTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationTypeAnyOf(val *AssociationTypeAnyOf) *NullableAssociationTypeAnyOf {
	return &NullableAssociationTypeAnyOf{value: val, isSet: true}
}

func (v NullableAssociationTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

