/*
3gpp-service-parameter

API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ServiceParameter

import (
	"encoding/json"
	"fmt"
)

// EventAnyOf the model 'EventAnyOf'
type EventAnyOf string

// List of Event_anyOf
const (
	SUCCESS_UE_POL_DEL_SP EventAnyOf = "SUCCESS_UE_POL_DEL_SP"
	UNSUCCESS_UE_POL_DEL_SP EventAnyOf = "UNSUCCESS_UE_POL_DEL_SP"
)

// All allowed values of EventAnyOf enum
var AllowedEventAnyOfEnumValues = []EventAnyOf{
	"SUCCESS_UE_POL_DEL_SP",
	"UNSUCCESS_UE_POL_DEL_SP",
}

func (v *EventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventAnyOf(value)
	for _, existing := range AllowedEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventAnyOf", value)
}

// NewEventAnyOfFromValue returns a pointer to a valid EventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventAnyOfFromValue(v string) (*EventAnyOf, error) {
	ev := EventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventAnyOf: valid values are %v", v, AllowedEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventAnyOf) IsValid() bool {
	for _, existing := range AllowedEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Event_anyOf value
func (v EventAnyOf) Ptr() *EventAnyOf {
	return &v
}

type NullableEventAnyOf struct {
	value *EventAnyOf
	isSet bool
}

func (v NullableEventAnyOf) Get() *EventAnyOf {
	return v.value
}

func (v *NullableEventAnyOf) Set(val *EventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventAnyOf(val *EventAnyOf) *NullableEventAnyOf {
	return &NullableEventAnyOf{value: val, isSet: true}
}

func (v NullableEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

