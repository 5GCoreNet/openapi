/*
Naf_Authentication

AF Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_Authentication

import (
	"encoding/json"
	"fmt"
)

// AuthMsgTypeAnyOf the model 'AuthMsgTypeAnyOf'
type AuthMsgTypeAnyOf string

// List of AuthMsgType_anyOf
const (
	UUAA AuthMsgTypeAnyOf = "UUAA"
	C2_AUTH AuthMsgTypeAnyOf = "C2AUTH"
)

// All allowed values of AuthMsgTypeAnyOf enum
var AllowedAuthMsgTypeAnyOfEnumValues = []AuthMsgTypeAnyOf{
	"UUAA",
	"C2AUTH",
}

func (v *AuthMsgTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AuthMsgTypeAnyOf(value)
	for _, existing := range AllowedAuthMsgTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AuthMsgTypeAnyOf", value)
}

// NewAuthMsgTypeAnyOfFromValue returns a pointer to a valid AuthMsgTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAuthMsgTypeAnyOfFromValue(v string) (*AuthMsgTypeAnyOf, error) {
	ev := AuthMsgTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AuthMsgTypeAnyOf: valid values are %v", v, AllowedAuthMsgTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AuthMsgTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAuthMsgTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AuthMsgType_anyOf value
func (v AuthMsgTypeAnyOf) Ptr() *AuthMsgTypeAnyOf {
	return &v
}

type NullableAuthMsgTypeAnyOf struct {
	value *AuthMsgTypeAnyOf
	isSet bool
}

func (v NullableAuthMsgTypeAnyOf) Get() *AuthMsgTypeAnyOf {
	return v.value
}

func (v *NullableAuthMsgTypeAnyOf) Set(val *AuthMsgTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthMsgTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthMsgTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthMsgTypeAnyOf(val *AuthMsgTypeAnyOf) *NullableAuthMsgTypeAnyOf {
	return &NullableAuthMsgTypeAnyOf{value: val, isSet: true}
}

func (v NullableAuthMsgTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthMsgTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
