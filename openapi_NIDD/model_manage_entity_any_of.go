/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NIDD

import (
	"encoding/json"
	"fmt"
)

// ManageEntityAnyOf the model 'ManageEntityAnyOf'
type ManageEntityAnyOf string

// List of ManageEntity_anyOf
const (
	UE ManageEntityAnyOf = "UE"
	AS ManageEntityAnyOf = "AS"
)

// All allowed values of ManageEntityAnyOf enum
var AllowedManageEntityAnyOfEnumValues = []ManageEntityAnyOf{
	"UE",
	"AS",
}

func (v *ManageEntityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ManageEntityAnyOf(value)
	for _, existing := range AllowedManageEntityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ManageEntityAnyOf", value)
}

// NewManageEntityAnyOfFromValue returns a pointer to a valid ManageEntityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewManageEntityAnyOfFromValue(v string) (*ManageEntityAnyOf, error) {
	ev := ManageEntityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ManageEntityAnyOf: valid values are %v", v, AllowedManageEntityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ManageEntityAnyOf) IsValid() bool {
	for _, existing := range AllowedManageEntityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManageEntity_anyOf value
func (v ManageEntityAnyOf) Ptr() *ManageEntityAnyOf {
	return &v
}

type NullableManageEntityAnyOf struct {
	value *ManageEntityAnyOf
	isSet bool
}

func (v NullableManageEntityAnyOf) Get() *ManageEntityAnyOf {
	return v.value
}

func (v *NullableManageEntityAnyOf) Set(val *ManageEntityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableManageEntityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableManageEntityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManageEntityAnyOf(val *ManageEntityAnyOf) *NullableManageEntityAnyOf {
	return &NullableManageEntityAnyOf{value: val, isSet: true}
}

func (v NullableManageEntityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManageEntityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

