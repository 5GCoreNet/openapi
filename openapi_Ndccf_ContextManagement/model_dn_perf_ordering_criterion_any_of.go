/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// DnPerfOrderingCriterionAnyOf the model 'DnPerfOrderingCriterionAnyOf'
type DnPerfOrderingCriterionAnyOf string

// List of DnPerfOrderingCriterion_anyOf
const (
	AVERAGE_TRAFFIC_RATE DnPerfOrderingCriterionAnyOf = "AVERAGE_TRAFFIC_RATE"
	MAXIMUM_TRAFFIC_RATE DnPerfOrderingCriterionAnyOf = "MAXIMUM_TRAFFIC_RATE"
	AVERAGE_PACKET_DELAY DnPerfOrderingCriterionAnyOf = "AVERAGE_PACKET_DELAY"
	MAXIMUM_PACKET_DELAY DnPerfOrderingCriterionAnyOf = "MAXIMUM_PACKET_DELAY"
	AVERAGE_PACKET_LOSS_RATE DnPerfOrderingCriterionAnyOf = "AVERAGE_PACKET_LOSS_RATE"
)

// All allowed values of DnPerfOrderingCriterionAnyOf enum
var AllowedDnPerfOrderingCriterionAnyOfEnumValues = []DnPerfOrderingCriterionAnyOf{
	"AVERAGE_TRAFFIC_RATE",
	"MAXIMUM_TRAFFIC_RATE",
	"AVERAGE_PACKET_DELAY",
	"MAXIMUM_PACKET_DELAY",
	"AVERAGE_PACKET_LOSS_RATE",
}

func (v *DnPerfOrderingCriterionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DnPerfOrderingCriterionAnyOf(value)
	for _, existing := range AllowedDnPerfOrderingCriterionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DnPerfOrderingCriterionAnyOf", value)
}

// NewDnPerfOrderingCriterionAnyOfFromValue returns a pointer to a valid DnPerfOrderingCriterionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDnPerfOrderingCriterionAnyOfFromValue(v string) (*DnPerfOrderingCriterionAnyOf, error) {
	ev := DnPerfOrderingCriterionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DnPerfOrderingCriterionAnyOf: valid values are %v", v, AllowedDnPerfOrderingCriterionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DnPerfOrderingCriterionAnyOf) IsValid() bool {
	for _, existing := range AllowedDnPerfOrderingCriterionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DnPerfOrderingCriterion_anyOf value
func (v DnPerfOrderingCriterionAnyOf) Ptr() *DnPerfOrderingCriterionAnyOf {
	return &v
}

type NullableDnPerfOrderingCriterionAnyOf struct {
	value *DnPerfOrderingCriterionAnyOf
	isSet bool
}

func (v NullableDnPerfOrderingCriterionAnyOf) Get() *DnPerfOrderingCriterionAnyOf {
	return v.value
}

func (v *NullableDnPerfOrderingCriterionAnyOf) Set(val *DnPerfOrderingCriterionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDnPerfOrderingCriterionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDnPerfOrderingCriterionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnPerfOrderingCriterionAnyOf(val *DnPerfOrderingCriterionAnyOf) *NullableDnPerfOrderingCriterionAnyOf {
	return &NullableDnPerfOrderingCriterionAnyOf{value: val, isSet: true}
}

func (v NullableDnPerfOrderingCriterionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnPerfOrderingCriterionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

