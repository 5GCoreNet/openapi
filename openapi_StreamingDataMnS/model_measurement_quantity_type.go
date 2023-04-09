/*
TS 28.532 Streaming data reporting service

OAS 3.0.1 specification for the Streaming data reporting service (Streaming MnS) © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_StreamingDataMnS

import (
	"encoding/json"
	"fmt"
)

// MeasurementQuantityType See details in 3GPP TS 32.422 clause 5.10.15.
type MeasurementQuantityType string

// List of measurementQuantity-Type
const (
	CPICH_EC_NO MeasurementQuantityType = "CPICH_EcNo"
	CPICH_RSCP  MeasurementQuantityType = "CPICH_RSCP"
	PATH_LOSS   MeasurementQuantityType = "PathLoss"
)

// All allowed values of MeasurementQuantityType enum
var AllowedMeasurementQuantityTypeEnumValues = []MeasurementQuantityType{
	"CPICH_EcNo",
	"CPICH_RSCP",
	"PathLoss",
}

func (v *MeasurementQuantityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MeasurementQuantityType(value)
	for _, existing := range AllowedMeasurementQuantityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MeasurementQuantityType", value)
}

// NewMeasurementQuantityTypeFromValue returns a pointer to a valid MeasurementQuantityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMeasurementQuantityTypeFromValue(v string) (*MeasurementQuantityType, error) {
	ev := MeasurementQuantityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MeasurementQuantityType: valid values are %v", v, AllowedMeasurementQuantityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MeasurementQuantityType) IsValid() bool {
	for _, existing := range AllowedMeasurementQuantityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to measurementQuantity-Type value
func (v MeasurementQuantityType) Ptr() *MeasurementQuantityType {
	return &v
}

type NullableMeasurementQuantityType struct {
	value *MeasurementQuantityType
	isSet bool
}

func (v NullableMeasurementQuantityType) Get() *MeasurementQuantityType {
	return v.value
}

func (v *NullableMeasurementQuantityType) Set(val *MeasurementQuantityType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementQuantityType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementQuantityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementQuantityType(val *MeasurementQuantityType) *NullableMeasurementQuantityType {
	return &NullableMeasurementQuantityType{value: val, isSet: true}
}

func (v NullableMeasurementQuantityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementQuantityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
