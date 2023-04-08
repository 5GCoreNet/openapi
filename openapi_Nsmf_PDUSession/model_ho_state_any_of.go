/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// HoStateAnyOf the model 'HoStateAnyOf'
type HoStateAnyOf string

// List of HoState_anyOf
const (
	NONE HoStateAnyOf = "NONE"
	PREPARING HoStateAnyOf = "PREPARING"
	PREPARED HoStateAnyOf = "PREPARED"
	COMPLETED HoStateAnyOf = "COMPLETED"
	CANCELLED HoStateAnyOf = "CANCELLED"
)

// All allowed values of HoStateAnyOf enum
var AllowedHoStateAnyOfEnumValues = []HoStateAnyOf{
	"NONE",
	"PREPARING",
	"PREPARED",
	"COMPLETED",
	"CANCELLED",
}

func (v *HoStateAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HoStateAnyOf(value)
	for _, existing := range AllowedHoStateAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HoStateAnyOf", value)
}

// NewHoStateAnyOfFromValue returns a pointer to a valid HoStateAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHoStateAnyOfFromValue(v string) (*HoStateAnyOf, error) {
	ev := HoStateAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HoStateAnyOf: valid values are %v", v, AllowedHoStateAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HoStateAnyOf) IsValid() bool {
	for _, existing := range AllowedHoStateAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HoState_anyOf value
func (v HoStateAnyOf) Ptr() *HoStateAnyOf {
	return &v
}

type NullableHoStateAnyOf struct {
	value *HoStateAnyOf
	isSet bool
}

func (v NullableHoStateAnyOf) Get() *HoStateAnyOf {
	return v.value
}

func (v *NullableHoStateAnyOf) Set(val *HoStateAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHoStateAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHoStateAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHoStateAnyOf(val *HoStateAnyOf) *NullableHoStateAnyOf {
	return &NullableHoStateAnyOf{value: val, isSet: true}
}

func (v NullableHoStateAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHoStateAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

