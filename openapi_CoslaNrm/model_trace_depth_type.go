/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
	"fmt"
)

// TraceDepthType Specifies how detailed information should be recorded in the Network Element. The Trace Depth is a paremeter for Trace Session level, i.e., the Trace Depth is the same for all of the NEs to be traced in the same Trace Session. See 3GPP TS 32.422 clause 5.3 for additional details.
type TraceDepthType string

// List of traceDepth-Type
const (
	MINIMUM TraceDepthType = "MINIMUM"
	MEDIUM TraceDepthType = "MEDIUM"
	MAXIMUM TraceDepthType = "MAXIMUM"
	VENDORMINIMUM TraceDepthType = "VENDORMINIMUM"
	VENDORMEDIUM TraceDepthType = "VENDORMEDIUM"
	VENDORMAXIMUM TraceDepthType = "VENDORMAXIMUM"
)

// All allowed values of TraceDepthType enum
var AllowedTraceDepthTypeEnumValues = []TraceDepthType{
	"MINIMUM",
	"MEDIUM",
	"MAXIMUM",
	"VENDORMINIMUM",
	"VENDORMEDIUM",
	"VENDORMAXIMUM",
}

func (v *TraceDepthType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TraceDepthType(value)
	for _, existing := range AllowedTraceDepthTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TraceDepthType", value)
}

// NewTraceDepthTypeFromValue returns a pointer to a valid TraceDepthType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTraceDepthTypeFromValue(v string) (*TraceDepthType, error) {
	ev := TraceDepthType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TraceDepthType: valid values are %v", v, AllowedTraceDepthTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TraceDepthType) IsValid() bool {
	for _, existing := range AllowedTraceDepthTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to traceDepth-Type value
func (v TraceDepthType) Ptr() *TraceDepthType {
	return &v
}

type NullableTraceDepthType struct {
	value *TraceDepthType
	isSet bool
}

func (v NullableTraceDepthType) Get() *TraceDepthType {
	return v.value
}

func (v *NullableTraceDepthType) Set(val *TraceDepthType) {
	v.value = val
	v.isSet = true
}

func (v NullableTraceDepthType) IsSet() bool {
	return v.isSet
}

func (v *NullableTraceDepthType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTraceDepthType(val *TraceDepthType) *NullableTraceDepthType {
	return &NullableTraceDepthType{value: val, isSet: true}
}

func (v NullableTraceDepthType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTraceDepthType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

