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

// ReportingTrigger The enumeration ReportingTrigger defines Reporting Triggers for MDT in the trace. See 3GPP TS 32.42] for further  description of the values. It shall comply with the provisions defined in table 5.6.3.8-1.
type ReportingTrigger struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportingTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ReportingTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportingTrigger) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportingTrigger struct {
	value *ReportingTrigger
	isSet bool
}

func (v NullableReportingTrigger) Get() *ReportingTrigger {
	return v.value
}

func (v *NullableReportingTrigger) Set(val *ReportingTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingTrigger(val *ReportingTrigger) *NullableReportingTrigger {
	return &NullableReportingTrigger{value: val, isSet: true}
}

func (v NullableReportingTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
