/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// PositioningMethodAnyOf the model 'PositioningMethodAnyOf'
type PositioningMethodAnyOf string

// List of PositioningMethod_anyOf
const (
	CELLID PositioningMethodAnyOf = "CELLID"
	ECID PositioningMethodAnyOf = "ECID"
	OTDOA PositioningMethodAnyOf = "OTDOA"
	BAROMETRIC_PRESSURE PositioningMethodAnyOf = "BAROMETRIC_PRESSURE"
	WLAN PositioningMethodAnyOf = "WLAN"
	BLUETOOTH PositioningMethodAnyOf = "BLUETOOTH"
	MBS PositioningMethodAnyOf = "MBS"
	MOTION_SENSOR PositioningMethodAnyOf = "MOTION_SENSOR"
	DL_TDOA PositioningMethodAnyOf = "DL_TDOA"
	DL_AOD PositioningMethodAnyOf = "DL_AOD"
	MULTI_RTT PositioningMethodAnyOf = "MULTI-RTT"
	NR_ECID PositioningMethodAnyOf = "NR_ECID"
	UL_TDOA PositioningMethodAnyOf = "UL_TDOA"
	UL_AOA PositioningMethodAnyOf = "UL_AOA"
	NETWORK_SPECIFIC PositioningMethodAnyOf = "NETWORK_SPECIFIC"
)

// All allowed values of PositioningMethodAnyOf enum
var AllowedPositioningMethodAnyOfEnumValues = []PositioningMethodAnyOf{
	"CELLID",
	"ECID",
	"OTDOA",
	"BAROMETRIC_PRESSURE",
	"WLAN",
	"BLUETOOTH",
	"MBS",
	"MOTION_SENSOR",
	"DL_TDOA",
	"DL_AOD",
	"MULTI-RTT",
	"NR_ECID",
	"UL_TDOA",
	"UL_AOA",
	"NETWORK_SPECIFIC",
}

func (v *PositioningMethodAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PositioningMethodAnyOf(value)
	for _, existing := range AllowedPositioningMethodAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PositioningMethodAnyOf", value)
}

// NewPositioningMethodAnyOfFromValue returns a pointer to a valid PositioningMethodAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPositioningMethodAnyOfFromValue(v string) (*PositioningMethodAnyOf, error) {
	ev := PositioningMethodAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PositioningMethodAnyOf: valid values are %v", v, AllowedPositioningMethodAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PositioningMethodAnyOf) IsValid() bool {
	for _, existing := range AllowedPositioningMethodAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PositioningMethod_anyOf value
func (v PositioningMethodAnyOf) Ptr() *PositioningMethodAnyOf {
	return &v
}

type NullablePositioningMethodAnyOf struct {
	value *PositioningMethodAnyOf
	isSet bool
}

func (v NullablePositioningMethodAnyOf) Get() *PositioningMethodAnyOf {
	return v.value
}

func (v *NullablePositioningMethodAnyOf) Set(val *PositioningMethodAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningMethodAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningMethodAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningMethodAnyOf(val *PositioningMethodAnyOf) *NullablePositioningMethodAnyOf {
	return &NullablePositioningMethodAnyOf{value: val, isSet: true}
}

func (v NullablePositioningMethodAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningMethodAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

