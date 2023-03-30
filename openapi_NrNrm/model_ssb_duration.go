/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
	"fmt"
)

// SsbDuration the model 'SsbDuration'
type SsbDuration int32

// List of SsbDuration
const (
	_1 SsbDuration = 1
	_2 SsbDuration = 2
	_3 SsbDuration = 3
	_4 SsbDuration = 4
	_5 SsbDuration = 5
)

// All allowed values of SsbDuration enum
var AllowedSsbDurationEnumValues = []SsbDuration{
	1,
	2,
	3,
	4,
	5,
}

func (v *SsbDuration) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SsbDuration(value)
	for _, existing := range AllowedSsbDurationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SsbDuration", value)
}

// NewSsbDurationFromValue returns a pointer to a valid SsbDuration
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSsbDurationFromValue(v int32) (*SsbDuration, error) {
	ev := SsbDuration(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SsbDuration: valid values are %v", v, AllowedSsbDurationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SsbDuration) IsValid() bool {
	for _, existing := range AllowedSsbDurationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SsbDuration value
func (v SsbDuration) Ptr() *SsbDuration {
	return &v
}

type NullableSsbDuration struct {
	value *SsbDuration
	isSet bool
}

func (v NullableSsbDuration) Get() *SsbDuration {
	return v.value
}

func (v *NullableSsbDuration) Set(val *SsbDuration) {
	v.value = val
	v.isSet = true
}

func (v NullableSsbDuration) IsSet() bool {
	return v.isSet
}

func (v *NullableSsbDuration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsbDuration(val *SsbDuration) *NullableSsbDuration {
	return &NullableSsbDuration{value: val, isSet: true}
}

func (v NullableSsbDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsbDuration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

