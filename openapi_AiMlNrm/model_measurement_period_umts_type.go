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

// MeasurementPeriodUmtsType See details in 3GPP TS 32.422 clause 5.10.22.
type MeasurementPeriodUmtsType string

// List of measurementPeriodUmts-Type
const (
	_1000MS  MeasurementPeriodUmtsType = "1000ms"
	_2000MS  MeasurementPeriodUmtsType = "2000ms"
	_3000MS  MeasurementPeriodUmtsType = "3000ms"
	_4000MS  MeasurementPeriodUmtsType = "4000ms"
	_6000MS  MeasurementPeriodUmtsType = "6000ms"
	_8000MS  MeasurementPeriodUmtsType = "8000ms"
	_12000MS MeasurementPeriodUmtsType = "12000ms"
	_16000MS MeasurementPeriodUmtsType = "16000ms"
	_20000MS MeasurementPeriodUmtsType = "20000ms"
	_24000MS MeasurementPeriodUmtsType = "24000ms"
	_28000MS MeasurementPeriodUmtsType = "28000ms"
	_32000MS MeasurementPeriodUmtsType = "32000ms"
	_64000MS MeasurementPeriodUmtsType = "64000ms"
)

// All allowed values of MeasurementPeriodUmtsType enum
var AllowedMeasurementPeriodUmtsTypeEnumValues = []MeasurementPeriodUmtsType{
	"1000ms",
	"2000ms",
	"3000ms",
	"4000ms",
	"6000ms",
	"8000ms",
	"12000ms",
	"16000ms",
	"20000ms",
	"24000ms",
	"28000ms",
	"32000ms",
	"64000ms",
}

func (v *MeasurementPeriodUmtsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MeasurementPeriodUmtsType(value)
	for _, existing := range AllowedMeasurementPeriodUmtsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MeasurementPeriodUmtsType", value)
}

// NewMeasurementPeriodUmtsTypeFromValue returns a pointer to a valid MeasurementPeriodUmtsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMeasurementPeriodUmtsTypeFromValue(v string) (*MeasurementPeriodUmtsType, error) {
	ev := MeasurementPeriodUmtsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MeasurementPeriodUmtsType: valid values are %v", v, AllowedMeasurementPeriodUmtsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MeasurementPeriodUmtsType) IsValid() bool {
	for _, existing := range AllowedMeasurementPeriodUmtsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to measurementPeriodUmts-Type value
func (v MeasurementPeriodUmtsType) Ptr() *MeasurementPeriodUmtsType {
	return &v
}

type NullableMeasurementPeriodUmtsType struct {
	value *MeasurementPeriodUmtsType
	isSet bool
}

func (v NullableMeasurementPeriodUmtsType) Get() *MeasurementPeriodUmtsType {
	return v.value
}

func (v *NullableMeasurementPeriodUmtsType) Set(val *MeasurementPeriodUmtsType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementPeriodUmtsType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementPeriodUmtsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementPeriodUmtsType(val *MeasurementPeriodUmtsType) *NullableMeasurementPeriodUmtsType {
	return &NullableMeasurementPeriodUmtsType{value: val, isSet: true}
}

func (v NullableMeasurementPeriodUmtsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementPeriodUmtsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
