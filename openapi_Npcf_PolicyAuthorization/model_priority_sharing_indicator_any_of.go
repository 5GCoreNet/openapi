/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// PrioritySharingIndicatorAnyOf the model 'PrioritySharingIndicatorAnyOf'
type PrioritySharingIndicatorAnyOf string

// List of PrioritySharingIndicator_anyOf
const (
	ENABLED PrioritySharingIndicatorAnyOf = "ENABLED"
	DISABLED PrioritySharingIndicatorAnyOf = "DISABLED"
)

// All allowed values of PrioritySharingIndicatorAnyOf enum
var AllowedPrioritySharingIndicatorAnyOfEnumValues = []PrioritySharingIndicatorAnyOf{
	"ENABLED",
	"DISABLED",
}

func (v *PrioritySharingIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PrioritySharingIndicatorAnyOf(value)
	for _, existing := range AllowedPrioritySharingIndicatorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PrioritySharingIndicatorAnyOf", value)
}

// NewPrioritySharingIndicatorAnyOfFromValue returns a pointer to a valid PrioritySharingIndicatorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPrioritySharingIndicatorAnyOfFromValue(v string) (*PrioritySharingIndicatorAnyOf, error) {
	ev := PrioritySharingIndicatorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PrioritySharingIndicatorAnyOf: valid values are %v", v, AllowedPrioritySharingIndicatorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PrioritySharingIndicatorAnyOf) IsValid() bool {
	for _, existing := range AllowedPrioritySharingIndicatorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PrioritySharingIndicator_anyOf value
func (v PrioritySharingIndicatorAnyOf) Ptr() *PrioritySharingIndicatorAnyOf {
	return &v
}

type NullablePrioritySharingIndicatorAnyOf struct {
	value *PrioritySharingIndicatorAnyOf
	isSet bool
}

func (v NullablePrioritySharingIndicatorAnyOf) Get() *PrioritySharingIndicatorAnyOf {
	return v.value
}

func (v *NullablePrioritySharingIndicatorAnyOf) Set(val *PrioritySharingIndicatorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePrioritySharingIndicatorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePrioritySharingIndicatorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePrioritySharingIndicatorAnyOf(val *PrioritySharingIndicatorAnyOf) *NullablePrioritySharingIndicatorAnyOf {
	return &NullablePrioritySharingIndicatorAnyOf{value: val, isSet: true}
}

func (v NullablePrioritySharingIndicatorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePrioritySharingIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
