/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// AppliedSmccTypeAnyOf This string indicates the type of applied SM congestion control. 
type AppliedSmccTypeAnyOf string

// List of AppliedSmccType_anyOf
const (
	DNN_CC AppliedSmccTypeAnyOf = "DNN_CC"
	SNSSAI_CC AppliedSmccTypeAnyOf = "SNSSAI_CC"
)

// All allowed values of AppliedSmccTypeAnyOf enum
var AllowedAppliedSmccTypeAnyOfEnumValues = []AppliedSmccTypeAnyOf{
	"DNN_CC",
	"SNSSAI_CC",
}

func (v *AppliedSmccTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppliedSmccTypeAnyOf(value)
	for _, existing := range AllowedAppliedSmccTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppliedSmccTypeAnyOf", value)
}

// NewAppliedSmccTypeAnyOfFromValue returns a pointer to a valid AppliedSmccTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppliedSmccTypeAnyOfFromValue(v string) (*AppliedSmccTypeAnyOf, error) {
	ev := AppliedSmccTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppliedSmccTypeAnyOf: valid values are %v", v, AllowedAppliedSmccTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppliedSmccTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAppliedSmccTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppliedSmccType_anyOf value
func (v AppliedSmccTypeAnyOf) Ptr() *AppliedSmccTypeAnyOf {
	return &v
}

type NullableAppliedSmccTypeAnyOf struct {
	value *AppliedSmccTypeAnyOf
	isSet bool
}

func (v NullableAppliedSmccTypeAnyOf) Get() *AppliedSmccTypeAnyOf {
	return v.value
}

func (v *NullableAppliedSmccTypeAnyOf) Set(val *AppliedSmccTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedSmccTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedSmccTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedSmccTypeAnyOf(val *AppliedSmccTypeAnyOf) *NullableAppliedSmccTypeAnyOf {
	return &NullableAppliedSmccTypeAnyOf{value: val, isSet: true}
}

func (v NullableAppliedSmccTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedSmccTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

