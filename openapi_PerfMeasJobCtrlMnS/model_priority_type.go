/*
TS 28.550 Performance Measurement Job Control Service

OAS 3.0.1 specification of the Performance Measurement Job Control Service @ 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 16.8.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_PerfMeasJobCtrlMnS

import (
	"encoding/json"
	"fmt"
)

// PriorityType the model 'PriorityType'
type PriorityType string

// List of priority-Type
const (
	LOW PriorityType = "low"
	MEDIUM PriorityType = "medium"
	HIGH PriorityType = "high"
)

// All allowed values of PriorityType enum
var AllowedPriorityTypeEnumValues = []PriorityType{
	"low",
	"medium",
	"high",
}

func (v *PriorityType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PriorityType(value)
	for _, existing := range AllowedPriorityTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PriorityType", value)
}

// NewPriorityTypeFromValue returns a pointer to a valid PriorityType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPriorityTypeFromValue(v string) (*PriorityType, error) {
	ev := PriorityType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PriorityType: valid values are %v", v, AllowedPriorityTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PriorityType) IsValid() bool {
	for _, existing := range AllowedPriorityTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to priority-Type value
func (v PriorityType) Ptr() *PriorityType {
	return &v
}

type NullablePriorityType struct {
	value *PriorityType
	isSet bool
}

func (v NullablePriorityType) Get() *PriorityType {
	return v.value
}

func (v *NullablePriorityType) Set(val *PriorityType) {
	v.value = val
	v.isSet = true
}

func (v NullablePriorityType) IsSet() bool {
	return v.isSet
}

func (v *NullablePriorityType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePriorityType(val *PriorityType) *NullablePriorityType {
	return &NullablePriorityType{value: val, isSet: true}
}

func (v NullablePriorityType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePriorityType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

