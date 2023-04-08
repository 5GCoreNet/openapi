/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// ManagementOperationStatusAnyOf the model 'ManagementOperationStatusAnyOf'
type ManagementOperationStatusAnyOf string

// List of ManagementOperationStatus_anyOf
const (
	SUCCEEDED ManagementOperationStatusAnyOf = "OPERATION_SUCCEEDED"
	FAILED ManagementOperationStatusAnyOf = "OPERATION_FAILED"
)

// All allowed values of ManagementOperationStatusAnyOf enum
var AllowedManagementOperationStatusAnyOfEnumValues = []ManagementOperationStatusAnyOf{
	"OPERATION_SUCCEEDED",
	"OPERATION_FAILED",
}

func (v *ManagementOperationStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ManagementOperationStatusAnyOf(value)
	for _, existing := range AllowedManagementOperationStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ManagementOperationStatusAnyOf", value)
}

// NewManagementOperationStatusAnyOfFromValue returns a pointer to a valid ManagementOperationStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewManagementOperationStatusAnyOfFromValue(v string) (*ManagementOperationStatusAnyOf, error) {
	ev := ManagementOperationStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ManagementOperationStatusAnyOf: valid values are %v", v, AllowedManagementOperationStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ManagementOperationStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedManagementOperationStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ManagementOperationStatus_anyOf value
func (v ManagementOperationStatusAnyOf) Ptr() *ManagementOperationStatusAnyOf {
	return &v
}

type NullableManagementOperationStatusAnyOf struct {
	value *ManagementOperationStatusAnyOf
	isSet bool
}

func (v NullableManagementOperationStatusAnyOf) Get() *ManagementOperationStatusAnyOf {
	return v.value
}

func (v *NullableManagementOperationStatusAnyOf) Set(val *ManagementOperationStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableManagementOperationStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableManagementOperationStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableManagementOperationStatusAnyOf(val *ManagementOperationStatusAnyOf) *NullableManagementOperationStatusAnyOf {
	return &NullableManagementOperationStatusAnyOf{value: val, isSet: true}
}

func (v NullableManagementOperationStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableManagementOperationStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

