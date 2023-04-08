/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// SubscribedEventAnyOf the model 'SubscribedEventAnyOf'
type SubscribedEventAnyOf string

// List of SubscribedEvent_anyOf
const (
	UP_PATH_CHANGE SubscribedEventAnyOf = "UP_PATH_CHANGE"
)

// All allowed values of SubscribedEventAnyOf enum
var AllowedSubscribedEventAnyOfEnumValues = []SubscribedEventAnyOf{
	"UP_PATH_CHANGE",
}

func (v *SubscribedEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubscribedEventAnyOf(value)
	for _, existing := range AllowedSubscribedEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubscribedEventAnyOf", value)
}

// NewSubscribedEventAnyOfFromValue returns a pointer to a valid SubscribedEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubscribedEventAnyOfFromValue(v string) (*SubscribedEventAnyOf, error) {
	ev := SubscribedEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubscribedEventAnyOf: valid values are %v", v, AllowedSubscribedEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubscribedEventAnyOf) IsValid() bool {
	for _, existing := range AllowedSubscribedEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubscribedEvent_anyOf value
func (v SubscribedEventAnyOf) Ptr() *SubscribedEventAnyOf {
	return &v
}

type NullableSubscribedEventAnyOf struct {
	value *SubscribedEventAnyOf
	isSet bool
}

func (v NullableSubscribedEventAnyOf) Get() *SubscribedEventAnyOf {
	return v.value
}

func (v *NullableSubscribedEventAnyOf) Set(val *SubscribedEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribedEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribedEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribedEventAnyOf(val *SubscribedEventAnyOf) *NullableSubscribedEventAnyOf {
	return &NullableSubscribedEventAnyOf{value: val, isSet: true}
}

func (v NullableSubscribedEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribedEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

