/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
	"fmt"
)

// NoProfileMatchReasonAnyOf the model 'NoProfileMatchReasonAnyOf'
type NoProfileMatchReasonAnyOf string

// List of NoProfileMatchReason_anyOf
const (
	REQUESTER_PLMN_NOT_ALLOWED NoProfileMatchReasonAnyOf = "REQUESTER_PLMN_NOT_ALLOWED"
	TARGET_NF_SUSPENDED NoProfileMatchReasonAnyOf = "TARGET_NF_SUSPENDED"
	TARGET_NF_UNDISCOVERABLE NoProfileMatchReasonAnyOf = "TARGET_NF_UNDISCOVERABLE"
	QUERY_PARAMS_COMBINATION_NO_MATCH NoProfileMatchReasonAnyOf = "QUERY_PARAMS_COMBINATION_NO_MATCH"
	UNSPECIFIED NoProfileMatchReasonAnyOf = "UNSPECIFIED"
)

// All allowed values of NoProfileMatchReasonAnyOf enum
var AllowedNoProfileMatchReasonAnyOfEnumValues = []NoProfileMatchReasonAnyOf{
	"REQUESTER_PLMN_NOT_ALLOWED",
	"TARGET_NF_SUSPENDED",
	"TARGET_NF_UNDISCOVERABLE",
	"QUERY_PARAMS_COMBINATION_NO_MATCH",
	"UNSPECIFIED",
}

func (v *NoProfileMatchReasonAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NoProfileMatchReasonAnyOf(value)
	for _, existing := range AllowedNoProfileMatchReasonAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NoProfileMatchReasonAnyOf", value)
}

// NewNoProfileMatchReasonAnyOfFromValue returns a pointer to a valid NoProfileMatchReasonAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNoProfileMatchReasonAnyOfFromValue(v string) (*NoProfileMatchReasonAnyOf, error) {
	ev := NoProfileMatchReasonAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NoProfileMatchReasonAnyOf: valid values are %v", v, AllowedNoProfileMatchReasonAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NoProfileMatchReasonAnyOf) IsValid() bool {
	for _, existing := range AllowedNoProfileMatchReasonAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NoProfileMatchReason_anyOf value
func (v NoProfileMatchReasonAnyOf) Ptr() *NoProfileMatchReasonAnyOf {
	return &v
}

type NullableNoProfileMatchReasonAnyOf struct {
	value *NoProfileMatchReasonAnyOf
	isSet bool
}

func (v NullableNoProfileMatchReasonAnyOf) Get() *NoProfileMatchReasonAnyOf {
	return v.value
}

func (v *NullableNoProfileMatchReasonAnyOf) Set(val *NoProfileMatchReasonAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNoProfileMatchReasonAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNoProfileMatchReasonAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNoProfileMatchReasonAnyOf(val *NoProfileMatchReasonAnyOf) *NullableNoProfileMatchReasonAnyOf {
	return &NullableNoProfileMatchReasonAnyOf{value: val, isSet: true}
}

func (v NullableNoProfileMatchReasonAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNoProfileMatchReasonAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

