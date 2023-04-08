/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
	"fmt"
)

// LossOfConnectivityReasonAnyOf the model 'LossOfConnectivityReasonAnyOf'
type LossOfConnectivityReasonAnyOf string

// List of LossOfConnectivityReason_anyOf
const (
	DEREGISTERED LossOfConnectivityReasonAnyOf = "DEREGISTERED"
	MAX_DETECTION_TIME_EXPIRED LossOfConnectivityReasonAnyOf = "MAX_DETECTION_TIME_EXPIRED"
	PURGED LossOfConnectivityReasonAnyOf = "PURGED"
)

// All allowed values of LossOfConnectivityReasonAnyOf enum
var AllowedLossOfConnectivityReasonAnyOfEnumValues = []LossOfConnectivityReasonAnyOf{
	"DEREGISTERED",
	"MAX_DETECTION_TIME_EXPIRED",
	"PURGED",
}

func (v *LossOfConnectivityReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LossOfConnectivityReasonAnyOf(value)
	for _, existing := range AllowedLossOfConnectivityReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LossOfConnectivityReasonAnyOf", value)
}

// NewLossOfConnectivityReasonAnyOfFromValue returns a pointer to a valid LossOfConnectivityReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLossOfConnectivityReasonAnyOfFromValue(v string) (*LossOfConnectivityReasonAnyOf, error) {
	ev := LossOfConnectivityReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LossOfConnectivityReasonAnyOf: valid values are %v", v, AllowedLossOfConnectivityReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LossOfConnectivityReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedLossOfConnectivityReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LossOfConnectivityReason_anyOf value
func (v LossOfConnectivityReasonAnyOf) Ptr() *LossOfConnectivityReasonAnyOf {
	return &v
}

type NullableLossOfConnectivityReasonAnyOf struct {
	value *LossOfConnectivityReasonAnyOf
	isSet bool
}

func (v NullableLossOfConnectivityReasonAnyOf) Get() *LossOfConnectivityReasonAnyOf {
	return v.value
}

func (v *NullableLossOfConnectivityReasonAnyOf) Set(val *LossOfConnectivityReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLossOfConnectivityReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLossOfConnectivityReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLossOfConnectivityReasonAnyOf(val *LossOfConnectivityReasonAnyOf) *NullableLossOfConnectivityReasonAnyOf {
	return &NullableLossOfConnectivityReasonAnyOf{value: val, isSet: true}
}

func (v NullableLossOfConnectivityReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLossOfConnectivityReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

