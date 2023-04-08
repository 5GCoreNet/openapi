/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
	"fmt"
)

// AcrMgntEventAnyOf the model 'AcrMgntEventAnyOf'
type AcrMgntEventAnyOf string

// List of AcrMgntEvent_anyOf
const (
	UP_PATH_CHG AcrMgntEventAnyOf = "UP_PATH_CHG"
	ACR_MONITORING AcrMgntEventAnyOf = "ACR_MONITORING"
	ACR_FACILITATION AcrMgntEventAnyOf = "ACR_FACILITATION"
	ACT_START_STOP AcrMgntEventAnyOf = "ACT_START_STOP"
)

// All allowed values of AcrMgntEventAnyOf enum
var AllowedAcrMgntEventAnyOfEnumValues = []AcrMgntEventAnyOf{
	"UP_PATH_CHG",
	"ACR_MONITORING",
	"ACR_FACILITATION",
	"ACT_START_STOP",
}

func (v *AcrMgntEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AcrMgntEventAnyOf(value)
	for _, existing := range AllowedAcrMgntEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AcrMgntEventAnyOf", value)
}

// NewAcrMgntEventAnyOfFromValue returns a pointer to a valid AcrMgntEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAcrMgntEventAnyOfFromValue(v string) (*AcrMgntEventAnyOf, error) {
	ev := AcrMgntEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AcrMgntEventAnyOf: valid values are %v", v, AllowedAcrMgntEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AcrMgntEventAnyOf) IsValid() bool {
	for _, existing := range AllowedAcrMgntEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AcrMgntEvent_anyOf value
func (v AcrMgntEventAnyOf) Ptr() *AcrMgntEventAnyOf {
	return &v
}

type NullableAcrMgntEventAnyOf struct {
	value *AcrMgntEventAnyOf
	isSet bool
}

func (v NullableAcrMgntEventAnyOf) Get() *AcrMgntEventAnyOf {
	return v.value
}

func (v *NullableAcrMgntEventAnyOf) Set(val *AcrMgntEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAcrMgntEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAcrMgntEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcrMgntEventAnyOf(val *AcrMgntEventAnyOf) *NullableAcrMgntEventAnyOf {
	return &NullableAcrMgntEventAnyOf{value: val, isSet: true}
}

func (v NullableAcrMgntEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcrMgntEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

