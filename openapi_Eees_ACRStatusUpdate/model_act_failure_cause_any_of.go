/*
EES ACR Status Update Service

EES ACR Status Update Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRStatusUpdate

import (
	"encoding/json"
	"fmt"
)

// ACTFailureCauseAnyOf the model 'ACTFailureCauseAnyOf'
type ACTFailureCauseAnyOf string

// List of ACTFailureCause_anyOf
const (
	ACR_CANCELLATION ACTFailureCauseAnyOf = "ACR_CANCELLATION"
	OTHER ACTFailureCauseAnyOf = "OTHER"
)

// All allowed values of ACTFailureCauseAnyOf enum
var AllowedACTFailureCauseAnyOfEnumValues = []ACTFailureCauseAnyOf{
	"ACR_CANCELLATION",
	"OTHER",
}

func (v *ACTFailureCauseAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ACTFailureCauseAnyOf(value)
	for _, existing := range AllowedACTFailureCauseAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ACTFailureCauseAnyOf", value)
}

// NewACTFailureCauseAnyOfFromValue returns a pointer to a valid ACTFailureCauseAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewACTFailureCauseAnyOfFromValue(v string) (*ACTFailureCauseAnyOf, error) {
	ev := ACTFailureCauseAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ACTFailureCauseAnyOf: valid values are %v", v, AllowedACTFailureCauseAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ACTFailureCauseAnyOf) IsValid() bool {
	for _, existing := range AllowedACTFailureCauseAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ACTFailureCause_anyOf value
func (v ACTFailureCauseAnyOf) Ptr() *ACTFailureCauseAnyOf {
	return &v
}

type NullableACTFailureCauseAnyOf struct {
	value *ACTFailureCauseAnyOf
	isSet bool
}

func (v NullableACTFailureCauseAnyOf) Get() *ACTFailureCauseAnyOf {
	return v.value
}

func (v *NullableACTFailureCauseAnyOf) Set(val *ACTFailureCauseAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableACTFailureCauseAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableACTFailureCauseAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableACTFailureCauseAnyOf(val *ACTFailureCauseAnyOf) *NullableACTFailureCauseAnyOf {
	return &NullableACTFailureCauseAnyOf{value: val, isSet: true}
}

func (v NullableACTFailureCauseAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableACTFailureCauseAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

