/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// SEPPType any of enumrated value
type SEPPType string

// List of SEPPType
const (
	CSEPP SEPPType = "CSEPP"
	PSEPP SEPPType = "PSEPP"
)

// All allowed values of SEPPType enum
var AllowedSEPPTypeEnumValues = []SEPPType{
	"CSEPP",
	"PSEPP",
}

func (v *SEPPType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SEPPType(value)
	for _, existing := range AllowedSEPPTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SEPPType", value)
}

// NewSEPPTypeFromValue returns a pointer to a valid SEPPType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSEPPTypeFromValue(v string) (*SEPPType, error) {
	ev := SEPPType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SEPPType: valid values are %v", v, AllowedSEPPTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SEPPType) IsValid() bool {
	for _, existing := range AllowedSEPPTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SEPPType value
func (v SEPPType) Ptr() *SEPPType {
	return &v
}

type NullableSEPPType struct {
	value *SEPPType
	isSet bool
}

func (v NullableSEPPType) Get() *SEPPType {
	return v.value
}

func (v *NullableSEPPType) Set(val *SEPPType) {
	v.value = val
	v.isSet = true
}

func (v NullableSEPPType) IsSet() bool {
	return v.isSet
}

func (v *NullableSEPPType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSEPPType(val *SEPPType) *NullableSEPPType {
	return &NullableSEPPType{value: val, isSet: true}
}

func (v NullableSEPPType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSEPPType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
