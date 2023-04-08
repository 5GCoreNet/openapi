/*
CAPIF_Routing_Info_API

API for Routing information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Routing_Info_API

import (
	"encoding/json"
	"fmt"
)

// OperationAnyOf the model 'OperationAnyOf'
type OperationAnyOf string

// List of Operation_anyOf
const (
	GET OperationAnyOf = "GET"
	POST OperationAnyOf = "POST"
	PUT OperationAnyOf = "PUT"
	PATCH OperationAnyOf = "PATCH"
	DELETE OperationAnyOf = "DELETE"
)

// All allowed values of OperationAnyOf enum
var AllowedOperationAnyOfEnumValues = []OperationAnyOf{
	"GET",
	"POST",
	"PUT",
	"PATCH",
	"DELETE",
}

func (v *OperationAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationAnyOf(value)
	for _, existing := range AllowedOperationAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationAnyOf", value)
}

// NewOperationAnyOfFromValue returns a pointer to a valid OperationAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationAnyOfFromValue(v string) (*OperationAnyOf, error) {
	ev := OperationAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationAnyOf: valid values are %v", v, AllowedOperationAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationAnyOf) IsValid() bool {
	for _, existing := range AllowedOperationAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Operation_anyOf value
func (v OperationAnyOf) Ptr() *OperationAnyOf {
	return &v
}

type NullableOperationAnyOf struct {
	value *OperationAnyOf
	isSet bool
}

func (v NullableOperationAnyOf) Get() *OperationAnyOf {
	return v.value
}

func (v *NullableOperationAnyOf) Set(val *OperationAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationAnyOf(val *OperationAnyOf) *NullableOperationAnyOf {
	return &NullableOperationAnyOf{value: val, isSet: true}
}

func (v NullableOperationAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

