/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// AdministrativeState the model 'AdministrativeState'
type AdministrativeState string

// List of AdministrativeState
const (
	LOCKED   AdministrativeState = "LOCKED"
	UNLOCKED AdministrativeState = "UNLOCKED"
)

// All allowed values of AdministrativeState enum
var AllowedAdministrativeStateEnumValues = []AdministrativeState{
	"LOCKED",
	"UNLOCKED",
}

func (v *AdministrativeState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AdministrativeState(value)
	for _, existing := range AllowedAdministrativeStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AdministrativeState", value)
}

// NewAdministrativeStateFromValue returns a pointer to a valid AdministrativeState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAdministrativeStateFromValue(v string) (*AdministrativeState, error) {
	ev := AdministrativeState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AdministrativeState: valid values are %v", v, AllowedAdministrativeStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AdministrativeState) IsValid() bool {
	for _, existing := range AllowedAdministrativeStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AdministrativeState value
func (v AdministrativeState) Ptr() *AdministrativeState {
	return &v
}

type NullableAdministrativeState struct {
	value *AdministrativeState
	isSet bool
}

func (v NullableAdministrativeState) Get() *AdministrativeState {
	return v.value
}

func (v *NullableAdministrativeState) Set(val *AdministrativeState) {
	v.value = val
	v.isSet = true
}

func (v NullableAdministrativeState) IsSet() bool {
	return v.isSet
}

func (v *NullableAdministrativeState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdministrativeState(val *AdministrativeState) *NullableAdministrativeState {
	return &NullableAdministrativeState{value: val, isSet: true}
}

func (v NullableAdministrativeState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdministrativeState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
