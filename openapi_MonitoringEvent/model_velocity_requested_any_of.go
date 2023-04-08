/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// VelocityRequestedAnyOf the model 'VelocityRequestedAnyOf'
type VelocityRequestedAnyOf string

// List of VelocityRequested_anyOf
const (
	NOT_REQUESTED VelocityRequestedAnyOf = "VELOCITY_IS_NOT_REQUESTED"
	REQUESTED VelocityRequestedAnyOf = "VELOCITY_IS_REQUESTED"
)

// All allowed values of VelocityRequestedAnyOf enum
var AllowedVelocityRequestedAnyOfEnumValues = []VelocityRequestedAnyOf{
	"VELOCITY_IS_NOT_REQUESTED",
	"VELOCITY_IS_REQUESTED",
}

func (v *VelocityRequestedAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := VelocityRequestedAnyOf(value)
	for _, existing := range AllowedVelocityRequestedAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid VelocityRequestedAnyOf", value)
}

// NewVelocityRequestedAnyOfFromValue returns a pointer to a valid VelocityRequestedAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewVelocityRequestedAnyOfFromValue(v string) (*VelocityRequestedAnyOf, error) {
	ev := VelocityRequestedAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for VelocityRequestedAnyOf: valid values are %v", v, AllowedVelocityRequestedAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v VelocityRequestedAnyOf) IsValid() bool {
	for _, existing := range AllowedVelocityRequestedAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to VelocityRequested_anyOf value
func (v VelocityRequestedAnyOf) Ptr() *VelocityRequestedAnyOf {
	return &v
}

type NullableVelocityRequestedAnyOf struct {
	value *VelocityRequestedAnyOf
	isSet bool
}

func (v NullableVelocityRequestedAnyOf) Get() *VelocityRequestedAnyOf {
	return v.value
}

func (v *NullableVelocityRequestedAnyOf) Set(val *VelocityRequestedAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityRequestedAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityRequestedAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityRequestedAnyOf(val *VelocityRequestedAnyOf) *NullableVelocityRequestedAnyOf {
	return &NullableVelocityRequestedAnyOf{value: val, isSet: true}
}

func (v NullableVelocityRequestedAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityRequestedAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
