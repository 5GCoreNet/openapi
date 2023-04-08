/*
Nhss_gbaSDM

Nhss Subscriber Data Management Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_gbaSDM

import (
	"encoding/json"
	"fmt"
)

// SecFeatureAnyOf the model 'SecFeatureAnyOf'
type SecFeatureAnyOf string

// List of SecFeature_anyOf
const (
	GPL_U SecFeatureAnyOf = "GPL_U"
)

// All allowed values of SecFeatureAnyOf enum
var AllowedSecFeatureAnyOfEnumValues = []SecFeatureAnyOf{
	"GPL_U",
}

func (v *SecFeatureAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecFeatureAnyOf(value)
	for _, existing := range AllowedSecFeatureAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecFeatureAnyOf", value)
}

// NewSecFeatureAnyOfFromValue returns a pointer to a valid SecFeatureAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecFeatureAnyOfFromValue(v string) (*SecFeatureAnyOf, error) {
	ev := SecFeatureAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecFeatureAnyOf: valid values are %v", v, AllowedSecFeatureAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecFeatureAnyOf) IsValid() bool {
	for _, existing := range AllowedSecFeatureAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecFeature_anyOf value
func (v SecFeatureAnyOf) Ptr() *SecFeatureAnyOf {
	return &v
}

type NullableSecFeatureAnyOf struct {
	value *SecFeatureAnyOf
	isSet bool
}

func (v NullableSecFeatureAnyOf) Get() *SecFeatureAnyOf {
	return v.value
}

func (v *NullableSecFeatureAnyOf) Set(val *SecFeatureAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecFeatureAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecFeatureAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecFeatureAnyOf(val *SecFeatureAnyOf) *NullableSecFeatureAnyOf {
	return &NullableSecFeatureAnyOf{value: val, isSet: true}
}

func (v NullableSecFeatureAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecFeatureAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

