/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// ExternalClientTypeAnyOf the model 'ExternalClientTypeAnyOf'
type ExternalClientTypeAnyOf string

// List of ExternalClientType_anyOf
const (
	EMERGENCY_SERVICES ExternalClientTypeAnyOf = "EMERGENCY_SERVICES"
	VALUE_ADDED_SERVICES ExternalClientTypeAnyOf = "VALUE_ADDED_SERVICES"
	PLMN_OPERATOR_SERVICES ExternalClientTypeAnyOf = "PLMN_OPERATOR_SERVICES"
	LAWFUL_INTERCEPT_SERVICES ExternalClientTypeAnyOf = "LAWFUL_INTERCEPT_SERVICES"
	PLMN_OPERATOR_BROADCAST_SERVICES ExternalClientTypeAnyOf = "PLMN_OPERATOR_BROADCAST_SERVICES"
	PLMN_OPERATOR_OM ExternalClientTypeAnyOf = "PLMN_OPERATOR_OM"
	PLMN_OPERATOR_ANONYMOUS_STATISTICS ExternalClientTypeAnyOf = "PLMN_OPERATOR_ANONYMOUS_STATISTICS"
	PLMN_OPERATOR_TARGET_MS_SERVICE_SUPPORT ExternalClientTypeAnyOf = "PLMN_OPERATOR_TARGET_MS_SERVICE_SUPPORT"
)

// All allowed values of ExternalClientTypeAnyOf enum
var AllowedExternalClientTypeAnyOfEnumValues = []ExternalClientTypeAnyOf{
	"EMERGENCY_SERVICES",
	"VALUE_ADDED_SERVICES",
	"PLMN_OPERATOR_SERVICES",
	"LAWFUL_INTERCEPT_SERVICES",
	"PLMN_OPERATOR_BROADCAST_SERVICES",
	"PLMN_OPERATOR_OM",
	"PLMN_OPERATOR_ANONYMOUS_STATISTICS",
	"PLMN_OPERATOR_TARGET_MS_SERVICE_SUPPORT",
}

func (v *ExternalClientTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExternalClientTypeAnyOf(value)
	for _, existing := range AllowedExternalClientTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExternalClientTypeAnyOf", value)
}

// NewExternalClientTypeAnyOfFromValue returns a pointer to a valid ExternalClientTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExternalClientTypeAnyOfFromValue(v string) (*ExternalClientTypeAnyOf, error) {
	ev := ExternalClientTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExternalClientTypeAnyOf: valid values are %v", v, AllowedExternalClientTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExternalClientTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedExternalClientTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExternalClientType_anyOf value
func (v ExternalClientTypeAnyOf) Ptr() *ExternalClientTypeAnyOf {
	return &v
}

type NullableExternalClientTypeAnyOf struct {
	value *ExternalClientTypeAnyOf
	isSet bool
}

func (v NullableExternalClientTypeAnyOf) Get() *ExternalClientTypeAnyOf {
	return v.value
}

func (v *NullableExternalClientTypeAnyOf) Set(val *ExternalClientTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalClientTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalClientTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalClientTypeAnyOf(val *ExternalClientTypeAnyOf) *NullableExternalClientTypeAnyOf {
	return &NullableExternalClientTypeAnyOf{value: val, isSet: true}
}

func (v NullableExternalClientTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalClientTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

