/*
VAE_SessionOrientedService

API for VAE_SessionOrientedService   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_SessionOrientedService

import (
	"encoding/json"
	"fmt"
)

// ActionAnyOf the model 'ActionAnyOf'
type ActionAnyOf string

// List of Action_anyOf
const (
	ESTABLISHMENT ActionAnyOf = "ESTABLISHMENT"
	UPDATE ActionAnyOf = "UPDATE"
)

// All allowed values of ActionAnyOf enum
var AllowedActionAnyOfEnumValues = []ActionAnyOf{
	"ESTABLISHMENT",
	"UPDATE",
}

func (v *ActionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ActionAnyOf(value)
	for _, existing := range AllowedActionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ActionAnyOf", value)
}

// NewActionAnyOfFromValue returns a pointer to a valid ActionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewActionAnyOfFromValue(v string) (*ActionAnyOf, error) {
	ev := ActionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ActionAnyOf: valid values are %v", v, AllowedActionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ActionAnyOf) IsValid() bool {
	for _, existing := range AllowedActionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Action_anyOf value
func (v ActionAnyOf) Ptr() *ActionAnyOf {
	return &v
}

type NullableActionAnyOf struct {
	value *ActionAnyOf
	isSet bool
}

func (v NullableActionAnyOf) Get() *ActionAnyOf {
	return v.value
}

func (v *NullableActionAnyOf) Set(val *ActionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableActionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableActionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActionAnyOf(val *ActionAnyOf) *NullableActionAnyOf {
	return &NullableActionAnyOf{value: val, isSet: true}
}

func (v NullableActionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

