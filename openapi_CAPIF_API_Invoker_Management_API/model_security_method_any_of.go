/*
CAPIF_API_Invoker_Management_API

API for API invoker management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_API_Invoker_Management_API

import (
	"encoding/json"
	"fmt"
)

// SecurityMethodAnyOf the model 'SecurityMethodAnyOf'
type SecurityMethodAnyOf string

// List of SecurityMethod_anyOf
const (
	PSK SecurityMethodAnyOf = "PSK"
	PKI SecurityMethodAnyOf = "PKI"
	OAUTH SecurityMethodAnyOf = "OAUTH"
)

// All allowed values of SecurityMethodAnyOf enum
var AllowedSecurityMethodAnyOfEnumValues = []SecurityMethodAnyOf{
	"PSK",
	"PKI",
	"OAUTH",
}

func (v *SecurityMethodAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SecurityMethodAnyOf(value)
	for _, existing := range AllowedSecurityMethodAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SecurityMethodAnyOf", value)
}

// NewSecurityMethodAnyOfFromValue returns a pointer to a valid SecurityMethodAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSecurityMethodAnyOfFromValue(v string) (*SecurityMethodAnyOf, error) {
	ev := SecurityMethodAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SecurityMethodAnyOf: valid values are %v", v, AllowedSecurityMethodAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SecurityMethodAnyOf) IsValid() bool {
	for _, existing := range AllowedSecurityMethodAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SecurityMethod_anyOf value
func (v SecurityMethodAnyOf) Ptr() *SecurityMethodAnyOf {
	return &v
}

type NullableSecurityMethodAnyOf struct {
	value *SecurityMethodAnyOf
	isSet bool
}

func (v NullableSecurityMethodAnyOf) Get() *SecurityMethodAnyOf {
	return v.value
}

func (v *NullableSecurityMethodAnyOf) Set(val *SecurityMethodAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMethodAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMethodAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMethodAnyOf(val *SecurityMethodAnyOf) *NullableSecurityMethodAnyOf {
	return &NullableSecurityMethodAnyOf{value: val, isSet: true}
}

func (v NullableSecurityMethodAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMethodAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

