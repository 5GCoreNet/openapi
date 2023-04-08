/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
	"fmt"
)

// ReportingLevelAnyOf the model 'ReportingLevelAnyOf'
type ReportingLevelAnyOf string

// List of ReportingLevel_anyOf
const (
	SER_ID_LEVEL ReportingLevelAnyOf = "SER_ID_LEVEL"
	RAT_GR_LEVEL ReportingLevelAnyOf = "RAT_GR_LEVEL"
	SPON_CON_LEVEL ReportingLevelAnyOf = "SPON_CON_LEVEL"
)

// All allowed values of ReportingLevelAnyOf enum
var AllowedReportingLevelAnyOfEnumValues = []ReportingLevelAnyOf{
	"SER_ID_LEVEL",
	"RAT_GR_LEVEL",
	"SPON_CON_LEVEL",
}

func (v *ReportingLevelAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportingLevelAnyOf(value)
	for _, existing := range AllowedReportingLevelAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportingLevelAnyOf", value)
}

// NewReportingLevelAnyOfFromValue returns a pointer to a valid ReportingLevelAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportingLevelAnyOfFromValue(v string) (*ReportingLevelAnyOf, error) {
	ev := ReportingLevelAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportingLevelAnyOf: valid values are %v", v, AllowedReportingLevelAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportingLevelAnyOf) IsValid() bool {
	for _, existing := range AllowedReportingLevelAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ReportingLevel_anyOf value
func (v ReportingLevelAnyOf) Ptr() *ReportingLevelAnyOf {
	return &v
}

type NullableReportingLevelAnyOf struct {
	value *ReportingLevelAnyOf
	isSet bool
}

func (v NullableReportingLevelAnyOf) Get() *ReportingLevelAnyOf {
	return v.value
}

func (v *NullableReportingLevelAnyOf) Set(val *ReportingLevelAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingLevelAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingLevelAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingLevelAnyOf(val *ReportingLevelAnyOf) *NullableReportingLevelAnyOf {
	return &NullableReportingLevelAnyOf{value: val, isSet: true}
}

func (v NullableReportingLevelAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingLevelAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

