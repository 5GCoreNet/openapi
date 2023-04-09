/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
	"fmt"
)

// JobTypeType Specifies whether the TraceJob represents only MDT, Logged MBSFN MDT, Trace or a combined Trace and MDT job. Applicable for Trace, MDT, RCEF and RLF reporting. See 3GPP TS 32.422 clause 5.9a for additional details.
type JobTypeType string

// List of jobType-Type
const (
	IMMEDIATE_MDT_ONLY      JobTypeType = "IMMEDIATE_MDT_ONLY"
	LOGGED_MDT_ONLY         JobTypeType = "LOGGED_MDT_ONLY"
	TRACE_ONLY              JobTypeType = "TRACE_ONLY"
	IMMEDIATE_MDT_AND_TRACE JobTypeType = "IMMEDIATE_MDT AND TRACE"
	RLF_REPORT_ONLY         JobTypeType = "RLF_REPORT_ONLY"
	RCEF_REPORT_ONLY        JobTypeType = "RCEF_REPORT_ONLY"
	LOGGED_MBSFN_MDT        JobTypeType = "LOGGED_MBSFN_MDT"
)

// All allowed values of JobTypeType enum
var AllowedJobTypeTypeEnumValues = []JobTypeType{
	"IMMEDIATE_MDT_ONLY",
	"LOGGED_MDT_ONLY",
	"TRACE_ONLY",
	"IMMEDIATE_MDT AND TRACE",
	"RLF_REPORT_ONLY",
	"RCEF_REPORT_ONLY",
	"LOGGED_MBSFN_MDT",
}

func (v *JobTypeType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := JobTypeType(value)
	for _, existing := range AllowedJobTypeTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid JobTypeType", value)
}

// NewJobTypeTypeFromValue returns a pointer to a valid JobTypeType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewJobTypeTypeFromValue(v string) (*JobTypeType, error) {
	ev := JobTypeType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for JobTypeType: valid values are %v", v, AllowedJobTypeTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v JobTypeType) IsValid() bool {
	for _, existing := range AllowedJobTypeTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to jobType-Type value
func (v JobTypeType) Ptr() *JobTypeType {
	return &v
}

type NullableJobTypeType struct {
	value *JobTypeType
	isSet bool
}

func (v NullableJobTypeType) Get() *JobTypeType {
	return v.value
}

func (v *NullableJobTypeType) Set(val *JobTypeType) {
	v.value = val
	v.isSet = true
}

func (v NullableJobTypeType) IsSet() bool {
	return v.isSet
}

func (v *NullableJobTypeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobTypeType(val *JobTypeType) *NullableJobTypeType {
	return &NullableJobTypeType{value: val, isSet: true}
}

func (v NullableJobTypeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobTypeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
