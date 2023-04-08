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

// ConnectionCapabilitiesAnyOf the model 'ConnectionCapabilitiesAnyOf'
type ConnectionCapabilitiesAnyOf string

// List of ConnectionCapabilities_anyOf
const (
	IMS ConnectionCapabilitiesAnyOf = "IMS"
	MMS ConnectionCapabilitiesAnyOf = "MMS"
	SUPL ConnectionCapabilitiesAnyOf = "SUPL"
	INTERNET ConnectionCapabilitiesAnyOf = "INTERNET"
)

// All allowed values of ConnectionCapabilitiesAnyOf enum
var AllowedConnectionCapabilitiesAnyOfEnumValues = []ConnectionCapabilitiesAnyOf{
	"IMS",
	"MMS",
	"SUPL",
	"INTERNET",
}

func (v *ConnectionCapabilitiesAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ConnectionCapabilitiesAnyOf(value)
	for _, existing := range AllowedConnectionCapabilitiesAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ConnectionCapabilitiesAnyOf", value)
}

// NewConnectionCapabilitiesAnyOfFromValue returns a pointer to a valid ConnectionCapabilitiesAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewConnectionCapabilitiesAnyOfFromValue(v string) (*ConnectionCapabilitiesAnyOf, error) {
	ev := ConnectionCapabilitiesAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ConnectionCapabilitiesAnyOf: valid values are %v", v, AllowedConnectionCapabilitiesAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ConnectionCapabilitiesAnyOf) IsValid() bool {
	for _, existing := range AllowedConnectionCapabilitiesAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ConnectionCapabilities_anyOf value
func (v ConnectionCapabilitiesAnyOf) Ptr() *ConnectionCapabilitiesAnyOf {
	return &v
}

type NullableConnectionCapabilitiesAnyOf struct {
	value *ConnectionCapabilitiesAnyOf
	isSet bool
}

func (v NullableConnectionCapabilitiesAnyOf) Get() *ConnectionCapabilitiesAnyOf {
	return v.value
}

func (v *NullableConnectionCapabilitiesAnyOf) Set(val *ConnectionCapabilitiesAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectionCapabilitiesAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectionCapabilitiesAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectionCapabilitiesAnyOf(val *ConnectionCapabilitiesAnyOf) *NullableConnectionCapabilitiesAnyOf {
	return &NullableConnectionCapabilitiesAnyOf{value: val, isSet: true}
}

func (v NullableConnectionCapabilitiesAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectionCapabilitiesAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

