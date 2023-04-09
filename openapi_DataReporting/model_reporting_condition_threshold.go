/*
3gpp-data-reporting

API for 3GPP Data Reporting.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_DataReporting

import (
	"encoding/json"
	"fmt"
)

// ReportingConditionThreshold struct for ReportingConditionThreshold
type ReportingConditionThreshold struct {
	Float32 *float32
	Float64 *float64
	Int32   *int32
	Int64   *int64
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportingConditionThreshold) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into float32
	err = json.Unmarshal(data, &dst.Float32)
	if err == nil {
		jsonFloat32, _ := json.Marshal(dst.Float32)
		if string(jsonFloat32) == "{}" { // empty struct
			dst.Float32 = nil
		} else {
			return nil // data stored in dst.Float32, return on the first match
		}
	} else {
		dst.Float32 = nil
	}

	// try to unmarshal JSON data into float64
	err = json.Unmarshal(data, &dst.Float64)
	if err == nil {
		jsonFloat64, _ := json.Marshal(dst.Float64)
		if string(jsonFloat64) == "{}" { // empty struct
			dst.Float64 = nil
		} else {
			return nil // data stored in dst.Float64, return on the first match
		}
	} else {
		dst.Float64 = nil
	}

	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			return nil // data stored in dst.Int32, return on the first match
		}
	} else {
		dst.Int32 = nil
	}

	// try to unmarshal JSON data into int64
	err = json.Unmarshal(data, &dst.Int64)
	if err == nil {
		jsonInt64, _ := json.Marshal(dst.Int64)
		if string(jsonInt64) == "{}" { // empty struct
			dst.Int64 = nil
		} else {
			return nil // data stored in dst.Int64, return on the first match
		}
	} else {
		dst.Int64 = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ReportingConditionThreshold)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportingConditionThreshold) MarshalJSON() ([]byte, error) {
	if src.Float32 != nil {
		return json.Marshal(&src.Float32)
	}

	if src.Float64 != nil {
		return json.Marshal(&src.Float64)
	}

	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	if src.Int64 != nil {
		return json.Marshal(&src.Int64)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportingConditionThreshold struct {
	value *ReportingConditionThreshold
	isSet bool
}

func (v NullableReportingConditionThreshold) Get() *ReportingConditionThreshold {
	return v.value
}

func (v *NullableReportingConditionThreshold) Set(val *ReportingConditionThreshold) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingConditionThreshold) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingConditionThreshold) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingConditionThreshold(val *ReportingConditionThreshold) *NullableReportingConditionThreshold {
	return &NullableReportingConditionThreshold{value: val, isSet: true}
}

func (v NullableReportingConditionThreshold) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingConditionThreshold) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
