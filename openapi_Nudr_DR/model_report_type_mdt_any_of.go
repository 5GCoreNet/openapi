/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// ReportTypeMdtAnyOf the model 'ReportTypeMdtAnyOf'
type ReportTypeMdtAnyOf string

// List of ReportTypeMdt_anyOf
const (
	PERIODICAL ReportTypeMdtAnyOf = "PERIODICAL"
	EVENT_TRIGGED ReportTypeMdtAnyOf = "EVENT_TRIGGED"
)

// All allowed values of ReportTypeMdtAnyOf enum
var AllowedReportTypeMdtAnyOfEnumValues = []ReportTypeMdtAnyOf{
	"PERIODICAL",
	"EVENT_TRIGGED",
}

func (v *ReportTypeMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportTypeMdtAnyOf(value)
	for _, existing := range AllowedReportTypeMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportTypeMdtAnyOf", value)
}

// NewReportTypeMdtAnyOfFromValue returns a pointer to a valid ReportTypeMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportTypeMdtAnyOfFromValue(v string) (*ReportTypeMdtAnyOf, error) {
	ev := ReportTypeMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportTypeMdtAnyOf: valid values are %v", v, AllowedReportTypeMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportTypeMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedReportTypeMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportTypeMdt_anyOf value
func (v ReportTypeMdtAnyOf) Ptr() *ReportTypeMdtAnyOf {
	return &v
}

type NullableReportTypeMdtAnyOf struct {
	value *ReportTypeMdtAnyOf
	isSet bool
}

func (v NullableReportTypeMdtAnyOf) Get() *ReportTypeMdtAnyOf {
	return v.value
}

func (v *NullableReportTypeMdtAnyOf) Set(val *ReportTypeMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportTypeMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportTypeMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportTypeMdtAnyOf(val *ReportTypeMdtAnyOf) *NullableReportTypeMdtAnyOf {
	return &NullableReportTypeMdtAnyOf{value: val, isSet: true}
}

func (v NullableReportTypeMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportTypeMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

