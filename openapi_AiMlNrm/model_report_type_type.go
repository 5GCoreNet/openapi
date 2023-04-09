/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
	"fmt"
)

// ReportTypeType Report type for logged NR MDT. See details in 3GPP TS 32.422 clause 5.10.27.
type ReportTypeType string

// List of reportType-Type
const (
	PERIODICAL      ReportTypeType = "PERIODICAL"
	EVENT_TRIGGERED ReportTypeType = "EVENT_TRIGGERED"
)

// All allowed values of ReportTypeType enum
var AllowedReportTypeTypeEnumValues = []ReportTypeType{
	"PERIODICAL",
	"EVENT_TRIGGERED",
}

func (v *ReportTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportTypeType(value)
	for _, existing := range AllowedReportTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportTypeType", value)
}

// NewReportTypeTypeFromValue returns a pointer to a valid ReportTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportTypeTypeFromValue(v string) (*ReportTypeType, error) {
	ev := ReportTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportTypeType: valid values are %v", v, AllowedReportTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportTypeType) IsValid() bool {
	for _, existing := range AllowedReportTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reportType-Type value
func (v ReportTypeType) Ptr() *ReportTypeType {
	return &v
}

type NullableReportTypeType struct {
	value *ReportTypeType
	isSet bool
}

func (v NullableReportTypeType) Get() *ReportTypeType {
	return v.value
}

func (v *NullableReportTypeType) Set(val *ReportTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportTypeType(val *ReportTypeType) *NullableReportTypeType {
	return &NullableReportTypeType{value: val, isSet: true}
}

func (v NullableReportTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
