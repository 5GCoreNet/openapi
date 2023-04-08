/*
Nmfaf_3daDataManagement

MFAF 3GPP DCCF Adaptor (3DA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3daDataManagement

import (
	"encoding/json"
	"fmt"
)

// ReportingOptions - Represents reporting options for processed notifications.
type ReportingOptions struct {
	Interface *interface{}
}

// interface{}AsReportingOptions is a convenience function that returns interface{} wrapped in ReportingOptions
func InterfaceAsReportingOptions(v *interface{}) ReportingOptions {
	return ReportingOptions{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ReportingOptions) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface
	err = newStrictDecoder(data).Decode(&dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			match++
		}
	} else {
		dst.Interface = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ReportingOptions)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ReportingOptions)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ReportingOptions) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ReportingOptions) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableReportingOptions struct {
	value *ReportingOptions
	isSet bool
}

func (v NullableReportingOptions) Get() *ReportingOptions {
	return v.value
}

func (v *NullableReportingOptions) Set(val *ReportingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingOptions(val *ReportingOptions) *NullableReportingOptions {
	return &NullableReportingOptions{value: val, isSet: true}
}

func (v NullableReportingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


