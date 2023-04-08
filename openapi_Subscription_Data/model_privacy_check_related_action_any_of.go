/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// PrivacyCheckRelatedActionAnyOf the model 'PrivacyCheckRelatedActionAnyOf'
type PrivacyCheckRelatedActionAnyOf string

// List of PrivacyCheckRelatedAction_anyOf
const (
	NOT_ALLOWED PrivacyCheckRelatedActionAnyOf = "LOCATION_NOT_ALLOWED"
	ALLOWED_WITH_NOTIFICATION PrivacyCheckRelatedActionAnyOf = "LOCATION_ALLOWED_WITH_NOTIFICATION"
	ALLOWED_WITHOUT_NOTIFICATION PrivacyCheckRelatedActionAnyOf = "LOCATION_ALLOWED_WITHOUT_NOTIFICATION"
	ALLOWED_WITHOUT_RESPONSE PrivacyCheckRelatedActionAnyOf = "LOCATION_ALLOWED_WITHOUT_RESPONSE"
	RESTRICTED_WITHOUT_RESPONSE PrivacyCheckRelatedActionAnyOf = "LOCATION_RESTRICTED_WITHOUT_RESPONSE"
)

// All allowed values of PrivacyCheckRelatedActionAnyOf enum
var AllowedPrivacyCheckRelatedActionAnyOfEnumValues = []PrivacyCheckRelatedActionAnyOf{
	"LOCATION_NOT_ALLOWED",
	"LOCATION_ALLOWED_WITH_NOTIFICATION",
	"LOCATION_ALLOWED_WITHOUT_NOTIFICATION",
	"LOCATION_ALLOWED_WITHOUT_RESPONSE",
	"LOCATION_RESTRICTED_WITHOUT_RESPONSE",
}

func (v *PrivacyCheckRelatedActionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrivacyCheckRelatedActionAnyOf(value)
	for _, existing := range AllowedPrivacyCheckRelatedActionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrivacyCheckRelatedActionAnyOf", value)
}

// NewPrivacyCheckRelatedActionAnyOfFromValue returns a pointer to a valid PrivacyCheckRelatedActionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrivacyCheckRelatedActionAnyOfFromValue(v string) (*PrivacyCheckRelatedActionAnyOf, error) {
	ev := PrivacyCheckRelatedActionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrivacyCheckRelatedActionAnyOf: valid values are %v", v, AllowedPrivacyCheckRelatedActionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrivacyCheckRelatedActionAnyOf) IsValid() bool {
	for _, existing := range AllowedPrivacyCheckRelatedActionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrivacyCheckRelatedAction_anyOf value
func (v PrivacyCheckRelatedActionAnyOf) Ptr() *PrivacyCheckRelatedActionAnyOf {
	return &v
}

type NullablePrivacyCheckRelatedActionAnyOf struct {
	value *PrivacyCheckRelatedActionAnyOf
	isSet bool
}

func (v NullablePrivacyCheckRelatedActionAnyOf) Get() *PrivacyCheckRelatedActionAnyOf {
	return v.value
}

func (v *NullablePrivacyCheckRelatedActionAnyOf) Set(val *PrivacyCheckRelatedActionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePrivacyCheckRelatedActionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePrivacyCheckRelatedActionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrivacyCheckRelatedActionAnyOf(val *PrivacyCheckRelatedActionAnyOf) *NullablePrivacyCheckRelatedActionAnyOf {
	return &NullablePrivacyCheckRelatedActionAnyOf{value: val, isSet: true}
}

func (v NullablePrivacyCheckRelatedActionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrivacyCheckRelatedActionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

