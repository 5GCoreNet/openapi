/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
	"fmt"
)

// PositioningMethodType See details in 3GPP TS 32.422 clause 5.10.19.
type PositioningMethodType string

// List of positioningMethod-Type
const (
	GNSS PositioningMethodType = "GNSS"
	E_CELL_ID PositioningMethodType = "E-CELL_ID"
)

// All allowed values of PositioningMethodType enum
var AllowedPositioningMethodTypeEnumValues = []PositioningMethodType{
	"GNSS",
	"E-CELL_ID",
}

func (v *PositioningMethodType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PositioningMethodType(value)
	for _, existing := range AllowedPositioningMethodTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PositioningMethodType", value)
}

// NewPositioningMethodTypeFromValue returns a pointer to a valid PositioningMethodType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPositioningMethodTypeFromValue(v string) (*PositioningMethodType, error) {
	ev := PositioningMethodType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PositioningMethodType: valid values are %v", v, AllowedPositioningMethodTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PositioningMethodType) IsValid() bool {
	for _, existing := range AllowedPositioningMethodTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to positioningMethod-Type value
func (v PositioningMethodType) Ptr() *PositioningMethodType {
	return &v
}

type NullablePositioningMethodType struct {
	value *PositioningMethodType
	isSet bool
}

func (v NullablePositioningMethodType) Get() *PositioningMethodType {
	return v.value
}

func (v *NullablePositioningMethodType) Set(val *PositioningMethodType) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningMethodType) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningMethodType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningMethodType(val *PositioningMethodType) *NullablePositioningMethodType {
	return &NullablePositioningMethodType{value: val, isSet: true}
}

func (v NullablePositioningMethodType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningMethodType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

