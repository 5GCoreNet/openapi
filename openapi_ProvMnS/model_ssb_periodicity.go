/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// SsbPeriodicity the model 'SsbPeriodicity'
type SsbPeriodicity int32

// List of SsbPeriodicity
const (
	_5 SsbPeriodicity = 5
	_10 SsbPeriodicity = 10
	_20 SsbPeriodicity = 20
	_40 SsbPeriodicity = 40
	_80 SsbPeriodicity = 80
	_160 SsbPeriodicity = 160
)

// All allowed values of SsbPeriodicity enum
var AllowedSsbPeriodicityEnumValues = []SsbPeriodicity{
	5,
	10,
	20,
	40,
	80,
	160,
}

func (v *SsbPeriodicity) UnmarshalJSON(src []byte) error {
	var value int32
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SsbPeriodicity(value)
	for _, existing := range AllowedSsbPeriodicityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SsbPeriodicity", value)
}

// NewSsbPeriodicityFromValue returns a pointer to a valid SsbPeriodicity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSsbPeriodicityFromValue(v int32) (*SsbPeriodicity, error) {
	ev := SsbPeriodicity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SsbPeriodicity: valid values are %v", v, AllowedSsbPeriodicityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SsbPeriodicity) IsValid() bool {
	for _, existing := range AllowedSsbPeriodicityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SsbPeriodicity value
func (v SsbPeriodicity) Ptr() *SsbPeriodicity {
	return &v
}

type NullableSsbPeriodicity struct {
	value *SsbPeriodicity
	isSet bool
}

func (v NullableSsbPeriodicity) Get() *SsbPeriodicity {
	return v.value
}

func (v *NullableSsbPeriodicity) Set(val *SsbPeriodicity) {
	v.value = val
	v.isSet = true
}

func (v NullableSsbPeriodicity) IsSet() bool {
	return v.isSet
}

func (v *NullableSsbPeriodicity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSsbPeriodicity(val *SsbPeriodicity) *NullableSsbPeriodicity {
	return &NullableSsbPeriodicity{value: val, isSet: true}
}

func (v NullableSsbPeriodicity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSsbPeriodicity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
