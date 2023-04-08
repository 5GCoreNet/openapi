/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// WlanOrderingCriterionAnyOf the model 'WlanOrderingCriterionAnyOf'
type WlanOrderingCriterionAnyOf string

// List of WlanOrderingCriterion_anyOf
const (
	TIME_SLOT_START WlanOrderingCriterionAnyOf = "TIME_SLOT_START"
	NUMBER_OF_UES WlanOrderingCriterionAnyOf = "NUMBER_OF_UES"
	RSSI WlanOrderingCriterionAnyOf = "RSSI"
	RTT WlanOrderingCriterionAnyOf = "RTT"
	TRAFFIC_INFO WlanOrderingCriterionAnyOf = "TRAFFIC_INFO"
)

// All allowed values of WlanOrderingCriterionAnyOf enum
var AllowedWlanOrderingCriterionAnyOfEnumValues = []WlanOrderingCriterionAnyOf{
	"TIME_SLOT_START",
	"NUMBER_OF_UES",
	"RSSI",
	"RTT",
	"TRAFFIC_INFO",
}

func (v *WlanOrderingCriterionAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := WlanOrderingCriterionAnyOf(value)
	for _, existing := range AllowedWlanOrderingCriterionAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid WlanOrderingCriterionAnyOf", value)
}

// NewWlanOrderingCriterionAnyOfFromValue returns a pointer to a valid WlanOrderingCriterionAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewWlanOrderingCriterionAnyOfFromValue(v string) (*WlanOrderingCriterionAnyOf, error) {
	ev := WlanOrderingCriterionAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for WlanOrderingCriterionAnyOf: valid values are %v", v, AllowedWlanOrderingCriterionAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v WlanOrderingCriterionAnyOf) IsValid() bool {
	for _, existing := range AllowedWlanOrderingCriterionAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to WlanOrderingCriterion_anyOf value
func (v WlanOrderingCriterionAnyOf) Ptr() *WlanOrderingCriterionAnyOf {
	return &v
}

type NullableWlanOrderingCriterionAnyOf struct {
	value *WlanOrderingCriterionAnyOf
	isSet bool
}

func (v NullableWlanOrderingCriterionAnyOf) Get() *WlanOrderingCriterionAnyOf {
	return v.value
}

func (v *NullableWlanOrderingCriterionAnyOf) Set(val *WlanOrderingCriterionAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanOrderingCriterionAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanOrderingCriterionAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanOrderingCriterionAnyOf(val *WlanOrderingCriterionAnyOf) *NullableWlanOrderingCriterionAnyOf {
	return &NullableWlanOrderingCriterionAnyOf{value: val, isSet: true}
}

func (v NullableWlanOrderingCriterionAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanOrderingCriterionAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

