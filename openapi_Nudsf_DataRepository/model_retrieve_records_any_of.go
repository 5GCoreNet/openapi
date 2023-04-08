/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
	"fmt"
)

// RetrieveRecordsAnyOf the model 'RetrieveRecordsAnyOf'
type RetrieveRecordsAnyOf string

// List of RetrieveRecords_anyOf
const (
	ONLY_META RetrieveRecordsAnyOf = "ONLY_META"
	META_AND_BLOCKS RetrieveRecordsAnyOf = "META_AND_BLOCKS"
)

// All allowed values of RetrieveRecordsAnyOf enum
var AllowedRetrieveRecordsAnyOfEnumValues = []RetrieveRecordsAnyOf{
	"ONLY_META",
	"META_AND_BLOCKS",
}

func (v *RetrieveRecordsAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RetrieveRecordsAnyOf(value)
	for _, existing := range AllowedRetrieveRecordsAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RetrieveRecordsAnyOf", value)
}

// NewRetrieveRecordsAnyOfFromValue returns a pointer to a valid RetrieveRecordsAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRetrieveRecordsAnyOfFromValue(v string) (*RetrieveRecordsAnyOf, error) {
	ev := RetrieveRecordsAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RetrieveRecordsAnyOf: valid values are %v", v, AllowedRetrieveRecordsAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RetrieveRecordsAnyOf) IsValid() bool {
	for _, existing := range AllowedRetrieveRecordsAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RetrieveRecords_anyOf value
func (v RetrieveRecordsAnyOf) Ptr() *RetrieveRecordsAnyOf {
	return &v
}

type NullableRetrieveRecordsAnyOf struct {
	value *RetrieveRecordsAnyOf
	isSet bool
}

func (v NullableRetrieveRecordsAnyOf) Get() *RetrieveRecordsAnyOf {
	return v.value
}

func (v *NullableRetrieveRecordsAnyOf) Set(val *RetrieveRecordsAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveRecordsAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveRecordsAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveRecordsAnyOf(val *RetrieveRecordsAnyOf) *NullableRetrieveRecordsAnyOf {
	return &NullableRetrieveRecordsAnyOf{value: val, isSet: true}
}

func (v NullableRetrieveRecordsAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveRecordsAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
