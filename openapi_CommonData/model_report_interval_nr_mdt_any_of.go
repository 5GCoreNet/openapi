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

// ReportIntervalNrMdtAnyOf the model 'ReportIntervalNrMdtAnyOf'
type ReportIntervalNrMdtAnyOf string

// List of ReportIntervalNrMdt_anyOf
const (
	_120 ReportIntervalNrMdtAnyOf = "120"
	_240 ReportIntervalNrMdtAnyOf = "240"
	_480 ReportIntervalNrMdtAnyOf = "480"
	_640 ReportIntervalNrMdtAnyOf = "640"
	_1024 ReportIntervalNrMdtAnyOf = "1024"
	_2048 ReportIntervalNrMdtAnyOf = "2048"
	_5120 ReportIntervalNrMdtAnyOf = "5120"
	_10240 ReportIntervalNrMdtAnyOf = "10240"
	_20480 ReportIntervalNrMdtAnyOf = "20480"
	_40960 ReportIntervalNrMdtAnyOf = "40960"
	_60000 ReportIntervalNrMdtAnyOf = "60000"
	_360000 ReportIntervalNrMdtAnyOf = "360000"
	_720000 ReportIntervalNrMdtAnyOf = "720000"
	_1800000 ReportIntervalNrMdtAnyOf = "1800000"
	_3600000 ReportIntervalNrMdtAnyOf = "3600000"
)

// All allowed values of ReportIntervalNrMdtAnyOf enum
var AllowedReportIntervalNrMdtAnyOfEnumValues = []ReportIntervalNrMdtAnyOf{
	"120",
	"240",
	"480",
	"640",
	"1024",
	"2048",
	"5120",
	"10240",
	"20480",
	"40960",
	"60000",
	"360000",
	"720000",
	"1800000",
	"3600000",
}

func (v *ReportIntervalNrMdtAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportIntervalNrMdtAnyOf(value)
	for _, existing := range AllowedReportIntervalNrMdtAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportIntervalNrMdtAnyOf", value)
}

// NewReportIntervalNrMdtAnyOfFromValue returns a pointer to a valid ReportIntervalNrMdtAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportIntervalNrMdtAnyOfFromValue(v string) (*ReportIntervalNrMdtAnyOf, error) {
	ev := ReportIntervalNrMdtAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportIntervalNrMdtAnyOf: valid values are %v", v, AllowedReportIntervalNrMdtAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportIntervalNrMdtAnyOf) IsValid() bool {
	for _, existing := range AllowedReportIntervalNrMdtAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportIntervalNrMdt_anyOf value
func (v ReportIntervalNrMdtAnyOf) Ptr() *ReportIntervalNrMdtAnyOf {
	return &v
}

type NullableReportIntervalNrMdtAnyOf struct {
	value *ReportIntervalNrMdtAnyOf
	isSet bool
}

func (v NullableReportIntervalNrMdtAnyOf) Get() *ReportIntervalNrMdtAnyOf {
	return v.value
}

func (v *NullableReportIntervalNrMdtAnyOf) Set(val *ReportIntervalNrMdtAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportIntervalNrMdtAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportIntervalNrMdtAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportIntervalNrMdtAnyOf(val *ReportIntervalNrMdtAnyOf) *NullableReportIntervalNrMdtAnyOf {
	return &NullableReportIntervalNrMdtAnyOf{value: val, isSet: true}
}

func (v NullableReportIntervalNrMdtAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportIntervalNrMdtAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

