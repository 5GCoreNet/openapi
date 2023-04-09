/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// TraceReportingFormatType Specifies whether file-based or streaming reporting shall be used for this Trace Session. See 3GPP TS 32.422 clause 5.11 for additional details.
type TraceReportingFormatType string

// List of traceReportingFormat-Type
const (
	FILE_BASED TraceReportingFormatType = "FILE-BASED"
	STREAMING  TraceReportingFormatType = "STREAMING"
)

// All allowed values of TraceReportingFormatType enum
var AllowedTraceReportingFormatTypeEnumValues = []TraceReportingFormatType{
	"FILE-BASED",
	"STREAMING",
}

func (v *TraceReportingFormatType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TraceReportingFormatType(value)
	for _, existing := range AllowedTraceReportingFormatTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TraceReportingFormatType", value)
}

// NewTraceReportingFormatTypeFromValue returns a pointer to a valid TraceReportingFormatType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTraceReportingFormatTypeFromValue(v string) (*TraceReportingFormatType, error) {
	ev := TraceReportingFormatType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TraceReportingFormatType: valid values are %v", v, AllowedTraceReportingFormatTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TraceReportingFormatType) IsValid() bool {
	for _, existing := range AllowedTraceReportingFormatTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to traceReportingFormat-Type value
func (v TraceReportingFormatType) Ptr() *TraceReportingFormatType {
	return &v
}

type NullableTraceReportingFormatType struct {
	value *TraceReportingFormatType
	isSet bool
}

func (v NullableTraceReportingFormatType) Get() *TraceReportingFormatType {
	return v.value
}

func (v *NullableTraceReportingFormatType) Set(val *TraceReportingFormatType) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceReportingFormatType) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceReportingFormatType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceReportingFormatType(val *TraceReportingFormatType) *NullableTraceReportingFormatType {
	return &NullableTraceReportingFormatType{value: val, isSet: true}
}

func (v NullableTraceReportingFormatType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceReportingFormatType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
