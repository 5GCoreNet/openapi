/*
Nudm_SSAU

Nudm Service Specific Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SSAU

import (
	"encoding/json"
	"fmt"
)

// ServiceTypeAnyOf the model 'ServiceTypeAnyOf'
type ServiceTypeAnyOf string

// List of ServiceType_anyOf
const (
	AF_GUIDANCE_FOR_URSP ServiceTypeAnyOf = "AF_GUIDANCE_FOR_URSP"
)

// All allowed values of ServiceTypeAnyOf enum
var AllowedServiceTypeAnyOfEnumValues = []ServiceTypeAnyOf{
	"AF_GUIDANCE_FOR_URSP",
}

func (v *ServiceTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ServiceTypeAnyOf(value)
	for _, existing := range AllowedServiceTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ServiceTypeAnyOf", value)
}

// NewServiceTypeAnyOfFromValue returns a pointer to a valid ServiceTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewServiceTypeAnyOfFromValue(v string) (*ServiceTypeAnyOf, error) {
	ev := ServiceTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ServiceTypeAnyOf: valid values are %v", v, AllowedServiceTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ServiceTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedServiceTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ServiceType_anyOf value
func (v ServiceTypeAnyOf) Ptr() *ServiceTypeAnyOf {
	return &v
}

type NullableServiceTypeAnyOf struct {
	value *ServiceTypeAnyOf
	isSet bool
}

func (v NullableServiceTypeAnyOf) Get() *ServiceTypeAnyOf {
	return v.value
}

func (v *NullableServiceTypeAnyOf) Set(val *ServiceTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceTypeAnyOf(val *ServiceTypeAnyOf) *NullableServiceTypeAnyOf {
	return &NullableServiceTypeAnyOf{value: val, isSet: true}
}

func (v NullableServiceTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

