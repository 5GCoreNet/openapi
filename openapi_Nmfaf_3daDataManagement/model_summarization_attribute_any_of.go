/*
Nmfaf_3daDataManagement

MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3daDataManagement

import (
	"encoding/json"
	"fmt"
)

// SummarizationAttributeAnyOf the model 'SummarizationAttributeAnyOf'
type SummarizationAttributeAnyOf string

// List of SummarizationAttribute_anyOf
const (
	SPACING SummarizationAttributeAnyOf = "SPACING"
	DURATION SummarizationAttributeAnyOf = "DURATION"
	OCCURRENCES SummarizationAttributeAnyOf = "OCCURRENCES"
	AVG_VAR SummarizationAttributeAnyOf = "AVG_VAR"
	FREQ_VAL SummarizationAttributeAnyOf = "FREQ_VAL"
	MIN_MAX SummarizationAttributeAnyOf = "MIN_MAX"
)

// All allowed values of SummarizationAttributeAnyOf enum
var AllowedSummarizationAttributeAnyOfEnumValues = []SummarizationAttributeAnyOf{
	"SPACING",
	"DURATION",
	"OCCURRENCES",
	"AVG_VAR",
	"FREQ_VAL",
	"MIN_MAX",
}

func (v *SummarizationAttributeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SummarizationAttributeAnyOf(value)
	for _, existing := range AllowedSummarizationAttributeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SummarizationAttributeAnyOf", value)
}

// NewSummarizationAttributeAnyOfFromValue returns a pointer to a valid SummarizationAttributeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSummarizationAttributeAnyOfFromValue(v string) (*SummarizationAttributeAnyOf, error) {
	ev := SummarizationAttributeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SummarizationAttributeAnyOf: valid values are %v", v, AllowedSummarizationAttributeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SummarizationAttributeAnyOf) IsValid() bool {
	for _, existing := range AllowedSummarizationAttributeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SummarizationAttribute_anyOf value
func (v SummarizationAttributeAnyOf) Ptr() *SummarizationAttributeAnyOf {
	return &v
}

type NullableSummarizationAttributeAnyOf struct {
	value *SummarizationAttributeAnyOf
	isSet bool
}

func (v NullableSummarizationAttributeAnyOf) Get() *SummarizationAttributeAnyOf {
	return v.value
}

func (v *NullableSummarizationAttributeAnyOf) Set(val *SummarizationAttributeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSummarizationAttributeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSummarizationAttributeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSummarizationAttributeAnyOf(val *SummarizationAttributeAnyOf) *NullableSummarizationAttributeAnyOf {
	return &NullableSummarizationAttributeAnyOf{value: val, isSet: true}
}

func (v NullableSummarizationAttributeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSummarizationAttributeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

