/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_StreamingDataMnS

import (
	"encoding/json"
	"fmt"
)

// TimeToTriggerL1Type See details in 3GPP TS 32.422 clause 5.10.Z.
type TimeToTriggerL1Type string

// List of timeToTriggerL1-Type
const (
	_0MS TimeToTriggerL1Type = "0ms"
	_40MS TimeToTriggerL1Type = "40ms"
	_64MS TimeToTriggerL1Type = "64ms"
	_80MS TimeToTriggerL1Type = "80ms"
	_100MS TimeToTriggerL1Type = "100ms"
	_128MS TimeToTriggerL1Type = "128ms"
	_160MS TimeToTriggerL1Type = "160ms"
	_256MS TimeToTriggerL1Type = "256ms"
	_320MS TimeToTriggerL1Type = "320ms"
	_480MS TimeToTriggerL1Type = "480ms"
	_512MS TimeToTriggerL1Type = "512ms"
	_640MS TimeToTriggerL1Type = "640ms"
	_1024MS TimeToTriggerL1Type = "1024ms"
	_1280MS TimeToTriggerL1Type = "1280ms"
	_2560MS TimeToTriggerL1Type = "2560ms"
	_5120MS TimeToTriggerL1Type = "5120ms"
)

// All allowed values of TimeToTriggerL1Type enum
var AllowedTimeToTriggerL1TypeEnumValues = []TimeToTriggerL1Type{
	"0ms",
	"40ms",
	"64ms",
	"80ms",
	"100ms",
	"128ms",
	"160ms",
	"256ms",
	"320ms",
	"480ms",
	"512ms",
	"640ms",
	"1024ms",
	"1280ms",
	"2560ms",
	"5120ms",
}

func (v *TimeToTriggerL1Type) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TimeToTriggerL1Type(value)
	for _, existing := range AllowedTimeToTriggerL1TypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TimeToTriggerL1Type", value)
}

// NewTimeToTriggerL1TypeFromValue returns a pointer to a valid TimeToTriggerL1Type
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTimeToTriggerL1TypeFromValue(v string) (*TimeToTriggerL1Type, error) {
	ev := TimeToTriggerL1Type(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TimeToTriggerL1Type: valid values are %v", v, AllowedTimeToTriggerL1TypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TimeToTriggerL1Type) IsValid() bool {
	for _, existing := range AllowedTimeToTriggerL1TypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to timeToTriggerL1-Type value
func (v TimeToTriggerL1Type) Ptr() *TimeToTriggerL1Type {
	return &v
}

type NullableTimeToTriggerL1Type struct {
	value *TimeToTriggerL1Type
	isSet bool
}

func (v NullableTimeToTriggerL1Type) Get() *TimeToTriggerL1Type {
	return v.value
}

func (v *NullableTimeToTriggerL1Type) Set(val *TimeToTriggerL1Type) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeToTriggerL1Type) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeToTriggerL1Type) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeToTriggerL1Type(val *TimeToTriggerL1Type) *NullableTimeToTriggerL1Type {
	return &NullableTimeToTriggerL1Type{value: val, isSet: true}
}

func (v NullableTimeToTriggerL1Type) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeToTriggerL1Type) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
