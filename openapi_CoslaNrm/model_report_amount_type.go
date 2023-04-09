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

// ReportAmountType See details in 3GPP TS 32.422 clause 5.10.6.
type ReportAmountType string

// List of reportAmount-Type
const (
	_1       ReportAmountType = "1"
	_2       ReportAmountType = "2"
	_4       ReportAmountType = "4"
	_8       ReportAmountType = "8"
	_16      ReportAmountType = "16"
	_32      ReportAmountType = "32"
	_64      ReportAmountType = "64"
	INFINITY ReportAmountType = "INFINITY"
)

// All allowed values of ReportAmountType enum
var AllowedReportAmountTypeEnumValues = []ReportAmountType{
	"1",
	"2",
	"4",
	"8",
	"16",
	"32",
	"64",
	"INFINITY",
}

func (v *ReportAmountType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ReportAmountType(value)
	for _, existing := range AllowedReportAmountTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ReportAmountType", value)
}

// NewReportAmountTypeFromValue returns a pointer to a valid ReportAmountType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewReportAmountTypeFromValue(v string) (*ReportAmountType, error) {
	ev := ReportAmountType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ReportAmountType: valid values are %v", v, AllowedReportAmountTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ReportAmountType) IsValid() bool {
	for _, existing := range AllowedReportAmountTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to reportAmount-Type value
func (v ReportAmountType) Ptr() *ReportAmountType {
	return &v
}

type NullableReportAmountType struct {
	value *ReportAmountType
	isSet bool
}

func (v NullableReportAmountType) Get() *ReportAmountType {
	return v.value
}

func (v *NullableReportAmountType) Set(val *ReportAmountType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportAmountType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportAmountType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportAmountType(val *ReportAmountType) *NullableReportAmountType {
	return &NullableReportAmountType{value: val, isSet: true}
}

func (v NullableReportAmountType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportAmountType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
