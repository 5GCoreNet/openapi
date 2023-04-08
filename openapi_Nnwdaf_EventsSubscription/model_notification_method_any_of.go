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

// NotificationMethodAnyOf the model 'NotificationMethodAnyOf'
type NotificationMethodAnyOf string

// List of NotificationMethod_anyOf
const (
	PERIODIC NotificationMethodAnyOf = "PERIODIC"
	THRESHOLD NotificationMethodAnyOf = "THRESHOLD"
)

// All allowed values of NotificationMethodAnyOf enum
var AllowedNotificationMethodAnyOfEnumValues = []NotificationMethodAnyOf{
	"PERIODIC",
	"THRESHOLD",
}

func (v *NotificationMethodAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NotificationMethodAnyOf(value)
	for _, existing := range AllowedNotificationMethodAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NotificationMethodAnyOf", value)
}

// NewNotificationMethodAnyOfFromValue returns a pointer to a valid NotificationMethodAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNotificationMethodAnyOfFromValue(v string) (*NotificationMethodAnyOf, error) {
	ev := NotificationMethodAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NotificationMethodAnyOf: valid values are %v", v, AllowedNotificationMethodAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NotificationMethodAnyOf) IsValid() bool {
	for _, existing := range AllowedNotificationMethodAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NotificationMethod_anyOf value
func (v NotificationMethodAnyOf) Ptr() *NotificationMethodAnyOf {
	return &v
}

type NullableNotificationMethodAnyOf struct {
	value *NotificationMethodAnyOf
	isSet bool
}

func (v NullableNotificationMethodAnyOf) Get() *NotificationMethodAnyOf {
	return v.value
}

func (v *NullableNotificationMethodAnyOf) Set(val *NotificationMethodAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationMethodAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationMethodAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationMethodAnyOf(val *NotificationMethodAnyOf) *NullableNotificationMethodAnyOf {
	return &NullableNotificationMethodAnyOf{value: val, isSet: true}
}

func (v NullableNotificationMethodAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationMethodAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

