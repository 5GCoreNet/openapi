/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// NotificationControlAnyOf the model 'NotificationControlAnyOf'
type NotificationControlAnyOf string

// List of NotificationControl_anyOf
const (
	REQUESTED NotificationControlAnyOf = "REQUESTED"
	NOT_REQUESTED NotificationControlAnyOf = "NOT_REQUESTED"
)

// All allowed values of NotificationControlAnyOf enum
var AllowedNotificationControlAnyOfEnumValues = []NotificationControlAnyOf{
	"REQUESTED",
	"NOT_REQUESTED",
}

func (v *NotificationControlAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationControlAnyOf(value)
	for _, existing := range AllowedNotificationControlAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationControlAnyOf", value)
}

// NewNotificationControlAnyOfFromValue returns a pointer to a valid NotificationControlAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationControlAnyOfFromValue(v string) (*NotificationControlAnyOf, error) {
	ev := NotificationControlAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationControlAnyOf: valid values are %v", v, AllowedNotificationControlAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationControlAnyOf) IsValid() bool {
	for _, existing := range AllowedNotificationControlAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationControl_anyOf value
func (v NotificationControlAnyOf) Ptr() *NotificationControlAnyOf {
	return &v
}

type NullableNotificationControlAnyOf struct {
	value *NotificationControlAnyOf
	isSet bool
}

func (v NullableNotificationControlAnyOf) Get() *NotificationControlAnyOf {
	return v.value
}

func (v *NullableNotificationControlAnyOf) Set(val *NotificationControlAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationControlAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationControlAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationControlAnyOf(val *NotificationControlAnyOf) *NullableNotificationControlAnyOf {
	return &NullableNotificationControlAnyOf{value: val, isSet: true}
}

func (v NullableNotificationControlAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationControlAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
