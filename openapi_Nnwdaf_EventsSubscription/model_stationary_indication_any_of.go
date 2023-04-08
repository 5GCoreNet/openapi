/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// StationaryIndicationAnyOf the model 'StationaryIndicationAnyOf'
type StationaryIndicationAnyOf string

// List of StationaryIndication_anyOf
const (
	STATIONARY StationaryIndicationAnyOf = "STATIONARY"
	MOBILE StationaryIndicationAnyOf = "MOBILE"
)

// All allowed values of StationaryIndicationAnyOf enum
var AllowedStationaryIndicationAnyOfEnumValues = []StationaryIndicationAnyOf{
	"STATIONARY",
	"MOBILE",
}

func (v *StationaryIndicationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StationaryIndicationAnyOf(value)
	for _, existing := range AllowedStationaryIndicationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StationaryIndicationAnyOf", value)
}

// NewStationaryIndicationAnyOfFromValue returns a pointer to a valid StationaryIndicationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStationaryIndicationAnyOfFromValue(v string) (*StationaryIndicationAnyOf, error) {
	ev := StationaryIndicationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StationaryIndicationAnyOf: valid values are %v", v, AllowedStationaryIndicationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StationaryIndicationAnyOf) IsValid() bool {
	for _, existing := range AllowedStationaryIndicationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StationaryIndication_anyOf value
func (v StationaryIndicationAnyOf) Ptr() *StationaryIndicationAnyOf {
	return &v
}

type NullableStationaryIndicationAnyOf struct {
	value *StationaryIndicationAnyOf
	isSet bool
}

func (v NullableStationaryIndicationAnyOf) Get() *StationaryIndicationAnyOf {
	return v.value
}

func (v *NullableStationaryIndicationAnyOf) Set(val *StationaryIndicationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableStationaryIndicationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableStationaryIndicationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStationaryIndicationAnyOf(val *StationaryIndicationAnyOf) *NullableStationaryIndicationAnyOf {
	return &NullableStationaryIndicationAnyOf{value: val, isSet: true}
}

func (v NullableStationaryIndicationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStationaryIndicationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

