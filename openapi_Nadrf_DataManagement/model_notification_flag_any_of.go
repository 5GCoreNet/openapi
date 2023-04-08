/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// NotificationFlagAnyOf the model 'NotificationFlagAnyOf'
type NotificationFlagAnyOf string

// List of NotificationFlag_anyOf
const (
	ACTIVATE NotificationFlagAnyOf = "ACTIVATE"
	DEACTIVATE NotificationFlagAnyOf = "DEACTIVATE"
	RETRIEVAL NotificationFlagAnyOf = "RETRIEVAL"
)

// All allowed values of NotificationFlagAnyOf enum
var AllowedNotificationFlagAnyOfEnumValues = []NotificationFlagAnyOf{
	"ACTIVATE",
	"DEACTIVATE",
	"RETRIEVAL",
}

func (v *NotificationFlagAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationFlagAnyOf(value)
	for _, existing := range AllowedNotificationFlagAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationFlagAnyOf", value)
}

// NewNotificationFlagAnyOfFromValue returns a pointer to a valid NotificationFlagAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationFlagAnyOfFromValue(v string) (*NotificationFlagAnyOf, error) {
	ev := NotificationFlagAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationFlagAnyOf: valid values are %v", v, AllowedNotificationFlagAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationFlagAnyOf) IsValid() bool {
	for _, existing := range AllowedNotificationFlagAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationFlag_anyOf value
func (v NotificationFlagAnyOf) Ptr() *NotificationFlagAnyOf {
	return &v
}

type NullableNotificationFlagAnyOf struct {
	value *NotificationFlagAnyOf
	isSet bool
}

func (v NullableNotificationFlagAnyOf) Get() *NotificationFlagAnyOf {
	return v.value
}

func (v *NullableNotificationFlagAnyOf) Set(val *NotificationFlagAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationFlagAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationFlagAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationFlagAnyOf(val *NotificationFlagAnyOf) *NullableNotificationFlagAnyOf {
	return &NullableNotificationFlagAnyOf{value: val, isSet: true}
}

func (v NullableNotificationFlagAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationFlagAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

