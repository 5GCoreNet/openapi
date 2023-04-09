/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// ReportIntervalNrMdt The enumeration ReportIntervalNrMdt defines Report Interval in NR for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.17-1.
type ReportIntervalNrMdt struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportIntervalNrMdt) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportIntervalNrMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportIntervalNrMdt) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportIntervalNrMdt struct {
	value *ReportIntervalNrMdt
	isSet bool
}

func (v NullableReportIntervalNrMdt) Get() *ReportIntervalNrMdt {
	return v.value
}

func (v *NullableReportIntervalNrMdt) Set(val *ReportIntervalNrMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableReportIntervalNrMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableReportIntervalNrMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportIntervalNrMdt(val *ReportIntervalNrMdt) *NullableReportIntervalNrMdt {
	return &NullableReportIntervalNrMdt{value: val, isSet: true}
}

func (v NullableReportIntervalNrMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportIntervalNrMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
