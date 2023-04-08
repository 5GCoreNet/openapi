/*
Nnef_EASDeployment

NEF EAS Deployment service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_EASDeployment

import (
	"encoding/json"
	"fmt"
)

// MatchingOperatorAnyOf the model 'MatchingOperatorAnyOf'
type MatchingOperatorAnyOf string

// List of MatchingOperator_anyOf
const (
	FULL_MATCH MatchingOperatorAnyOf = "FULL_MATCH"
	MATCH_ALL MatchingOperatorAnyOf = "MATCH_ALL"
	STARTS_WITH MatchingOperatorAnyOf = "STARTS_WITH"
	NOT_START_WITH MatchingOperatorAnyOf = "NOT_START_WITH"
	ENDS_WITH MatchingOperatorAnyOf = "ENDS_WITH"
	NOT_END_WITH MatchingOperatorAnyOf = "NOT_END_WITH"
	CONTAINS MatchingOperatorAnyOf = "CONTAINS"
	NOT_CONTAIN MatchingOperatorAnyOf = "NOT_CONTAIN"
)

// All allowed values of MatchingOperatorAnyOf enum
var AllowedMatchingOperatorAnyOfEnumValues = []MatchingOperatorAnyOf{
	"FULL_MATCH",
	"MATCH_ALL",
	"STARTS_WITH",
	"NOT_START_WITH",
	"ENDS_WITH",
	"NOT_END_WITH",
	"CONTAINS",
	"NOT_CONTAIN",
}

func (v *MatchingOperatorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MatchingOperatorAnyOf(value)
	for _, existing := range AllowedMatchingOperatorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MatchingOperatorAnyOf", value)
}

// NewMatchingOperatorAnyOfFromValue returns a pointer to a valid MatchingOperatorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMatchingOperatorAnyOfFromValue(v string) (*MatchingOperatorAnyOf, error) {
	ev := MatchingOperatorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MatchingOperatorAnyOf: valid values are %v", v, AllowedMatchingOperatorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MatchingOperatorAnyOf) IsValid() bool {
	for _, existing := range AllowedMatchingOperatorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MatchingOperator_anyOf value
func (v MatchingOperatorAnyOf) Ptr() *MatchingOperatorAnyOf {
	return &v
}

type NullableMatchingOperatorAnyOf struct {
	value *MatchingOperatorAnyOf
	isSet bool
}

func (v NullableMatchingOperatorAnyOf) Get() *MatchingOperatorAnyOf {
	return v.value
}

func (v *NullableMatchingOperatorAnyOf) Set(val *MatchingOperatorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMatchingOperatorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchingOperatorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchingOperatorAnyOf(val *MatchingOperatorAnyOf) *NullableMatchingOperatorAnyOf {
	return &NullableMatchingOperatorAnyOf{value: val, isSet: true}
}

func (v NullableMatchingOperatorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchingOperatorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

