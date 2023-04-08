/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// PartialRecordMethodAnyOf the model 'PartialRecordMethodAnyOf'
type PartialRecordMethodAnyOf string

// List of PartialRecordMethod_anyOf
const (
	DEFAULT PartialRecordMethodAnyOf = "DEFAULT"
	INDIVIDUAL PartialRecordMethodAnyOf = "INDIVIDUAL"
)

// All allowed values of PartialRecordMethodAnyOf enum
var AllowedPartialRecordMethodAnyOfEnumValues = []PartialRecordMethodAnyOf{
	"DEFAULT",
	"INDIVIDUAL",
}

func (v *PartialRecordMethodAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PartialRecordMethodAnyOf(value)
	for _, existing := range AllowedPartialRecordMethodAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PartialRecordMethodAnyOf", value)
}

// NewPartialRecordMethodAnyOfFromValue returns a pointer to a valid PartialRecordMethodAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPartialRecordMethodAnyOfFromValue(v string) (*PartialRecordMethodAnyOf, error) {
	ev := PartialRecordMethodAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PartialRecordMethodAnyOf: valid values are %v", v, AllowedPartialRecordMethodAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PartialRecordMethodAnyOf) IsValid() bool {
	for _, existing := range AllowedPartialRecordMethodAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PartialRecordMethod_anyOf value
func (v PartialRecordMethodAnyOf) Ptr() *PartialRecordMethodAnyOf {
	return &v
}

type NullablePartialRecordMethodAnyOf struct {
	value *PartialRecordMethodAnyOf
	isSet bool
}

func (v NullablePartialRecordMethodAnyOf) Get() *PartialRecordMethodAnyOf {
	return v.value
}

func (v *NullablePartialRecordMethodAnyOf) Set(val *PartialRecordMethodAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePartialRecordMethodAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePartialRecordMethodAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartialRecordMethodAnyOf(val *PartialRecordMethodAnyOf) *NullablePartialRecordMethodAnyOf {
	return &NullablePartialRecordMethodAnyOf{value: val, isSet: true}
}

func (v NullablePartialRecordMethodAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartialRecordMethodAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

