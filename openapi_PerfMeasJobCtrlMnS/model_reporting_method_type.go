/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMeasJobCtrlMnS

import (
	"encoding/json"
	"fmt"
)

// ReportingMethodType the model 'ReportingMethodType'
type ReportingMethodType string

// List of reportingMethod-Type
const (
	FILE      ReportingMethodType = "file"
	STREAMING ReportingMethodType = "streaming"
)

// All allowed values of ReportingMethodType enum
var AllowedReportingMethodTypeEnumValues = []ReportingMethodType{
	"file",
	"streaming",
}

func (v *ReportingMethodType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportingMethodType(value)
	for _, existing := range AllowedReportingMethodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportingMethodType", value)
}

// NewReportingMethodTypeFromValue returns a pointer to a valid ReportingMethodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportingMethodTypeFromValue(v string) (*ReportingMethodType, error) {
	ev := ReportingMethodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportingMethodType: valid values are %v", v, AllowedReportingMethodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportingMethodType) IsValid() bool {
	for _, existing := range AllowedReportingMethodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reportingMethod-Type value
func (v ReportingMethodType) Ptr() *ReportingMethodType {
	return &v
}

type NullableReportingMethodType struct {
	value *ReportingMethodType
	isSet bool
}

func (v NullableReportingMethodType) Get() *ReportingMethodType {
	return v.value
}

func (v *NullableReportingMethodType) Set(val *ReportingMethodType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingMethodType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingMethodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingMethodType(val *ReportingMethodType) *NullableReportingMethodType {
	return &NullableReportingMethodType{value: val, isSet: true}
}

func (v NullableReportingMethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingMethodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
