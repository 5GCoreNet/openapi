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

// AerialUeIndicationAnyOf the model 'AerialUeIndicationAnyOf'
type AerialUeIndicationAnyOf string

// List of AerialUeIndication_anyOf
const (
	ALLOWED AerialUeIndicationAnyOf = "AERIAL_UE_ALLOWED"
	NOT_ALLOWED AerialUeIndicationAnyOf = "AERIAL_UE_NOT_ALLOWED"
)

// All allowed values of AerialUeIndicationAnyOf enum
var AllowedAerialUeIndicationAnyOfEnumValues = []AerialUeIndicationAnyOf{
	"AERIAL_UE_ALLOWED",
	"AERIAL_UE_NOT_ALLOWED",
}

func (v *AerialUeIndicationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AerialUeIndicationAnyOf(value)
	for _, existing := range AllowedAerialUeIndicationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AerialUeIndicationAnyOf", value)
}

// NewAerialUeIndicationAnyOfFromValue returns a pointer to a valid AerialUeIndicationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAerialUeIndicationAnyOfFromValue(v string) (*AerialUeIndicationAnyOf, error) {
	ev := AerialUeIndicationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AerialUeIndicationAnyOf: valid values are %v", v, AllowedAerialUeIndicationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AerialUeIndicationAnyOf) IsValid() bool {
	for _, existing := range AllowedAerialUeIndicationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AerialUeIndication_anyOf value
func (v AerialUeIndicationAnyOf) Ptr() *AerialUeIndicationAnyOf {
	return &v
}

type NullableAerialUeIndicationAnyOf struct {
	value *AerialUeIndicationAnyOf
	isSet bool
}

func (v NullableAerialUeIndicationAnyOf) Get() *AerialUeIndicationAnyOf {
	return v.value
}

func (v *NullableAerialUeIndicationAnyOf) Set(val *AerialUeIndicationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAerialUeIndicationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAerialUeIndicationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAerialUeIndicationAnyOf(val *AerialUeIndicationAnyOf) *NullableAerialUeIndicationAnyOf {
	return &NullableAerialUeIndicationAnyOf{value: val, isSet: true}
}

func (v NullableAerialUeIndicationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAerialUeIndicationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

