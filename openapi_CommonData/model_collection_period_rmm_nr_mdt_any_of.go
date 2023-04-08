/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRmmNrMdtAnyOf the model 'CollectionPeriodRmmNrMdtAnyOf'
type CollectionPeriodRmmNrMdtAnyOf string

// List of CollectionPeriodRmmNrMdt_anyOf
const (
	_1024 CollectionPeriodRmmNrMdtAnyOf = "1024"
	_2048 CollectionPeriodRmmNrMdtAnyOf = "2048"
	_5120 CollectionPeriodRmmNrMdtAnyOf = "5120"
	_10240 CollectionPeriodRmmNrMdtAnyOf = "10240"
	_60000 CollectionPeriodRmmNrMdtAnyOf = "60000"
)

// All allowed values of CollectionPeriodRmmNrMdtAnyOf enum
var AllowedCollectionPeriodRmmNrMdtAnyOfEnumValues = []CollectionPeriodRmmNrMdtAnyOf{
	"1024",
	"2048",
	"5120",
	"10240",
	"60000",
}

func (v *CollectionPeriodRmmNrMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodRmmNrMdtAnyOf(value)
	for _, existing := range AllowedCollectionPeriodRmmNrMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodRmmNrMdtAnyOf", value)
}

// NewCollectionPeriodRmmNrMdtAnyOfFromValue returns a pointer to a valid CollectionPeriodRmmNrMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodRmmNrMdtAnyOfFromValue(v string) (*CollectionPeriodRmmNrMdtAnyOf, error) {
	ev := CollectionPeriodRmmNrMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodRmmNrMdtAnyOf: valid values are %v", v, AllowedCollectionPeriodRmmNrMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodRmmNrMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodRmmNrMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CollectionPeriodRmmNrMdt_anyOf value
func (v CollectionPeriodRmmNrMdtAnyOf) Ptr() *CollectionPeriodRmmNrMdtAnyOf {
	return &v
}

type NullableCollectionPeriodRmmNrMdtAnyOf struct {
	value *CollectionPeriodRmmNrMdtAnyOf
	isSet bool
}

func (v NullableCollectionPeriodRmmNrMdtAnyOf) Get() *CollectionPeriodRmmNrMdtAnyOf {
	return v.value
}

func (v *NullableCollectionPeriodRmmNrMdtAnyOf) Set(val *CollectionPeriodRmmNrMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRmmNrMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRmmNrMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRmmNrMdtAnyOf(val *CollectionPeriodRmmNrMdtAnyOf) *NullableCollectionPeriodRmmNrMdtAnyOf {
	return &NullableCollectionPeriodRmmNrMdtAnyOf{value: val, isSet: true}
}

func (v NullableCollectionPeriodRmmNrMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRmmNrMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

