/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// PreemptionCapabilityAnyOf the model 'PreemptionCapabilityAnyOf'
type PreemptionCapabilityAnyOf string

// List of PreemptionCapability_anyOf
const (
	NOT_PREEMPT PreemptionCapabilityAnyOf = "NOT_PREEMPT"
	MAY_PREEMPT PreemptionCapabilityAnyOf = "MAY_PREEMPT"
)

// All allowed values of PreemptionCapabilityAnyOf enum
var AllowedPreemptionCapabilityAnyOfEnumValues = []PreemptionCapabilityAnyOf{
	"NOT_PREEMPT",
	"MAY_PREEMPT",
}

func (v *PreemptionCapabilityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PreemptionCapabilityAnyOf(value)
	for _, existing := range AllowedPreemptionCapabilityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PreemptionCapabilityAnyOf", value)
}

// NewPreemptionCapabilityAnyOfFromValue returns a pointer to a valid PreemptionCapabilityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPreemptionCapabilityAnyOfFromValue(v string) (*PreemptionCapabilityAnyOf, error) {
	ev := PreemptionCapabilityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PreemptionCapabilityAnyOf: valid values are %v", v, AllowedPreemptionCapabilityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PreemptionCapabilityAnyOf) IsValid() bool {
	for _, existing := range AllowedPreemptionCapabilityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PreemptionCapability_anyOf value
func (v PreemptionCapabilityAnyOf) Ptr() *PreemptionCapabilityAnyOf {
	return &v
}

type NullablePreemptionCapabilityAnyOf struct {
	value *PreemptionCapabilityAnyOf
	isSet bool
}

func (v NullablePreemptionCapabilityAnyOf) Get() *PreemptionCapabilityAnyOf {
	return v.value
}

func (v *NullablePreemptionCapabilityAnyOf) Set(val *PreemptionCapabilityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePreemptionCapabilityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePreemptionCapabilityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreemptionCapabilityAnyOf(val *PreemptionCapabilityAnyOf) *NullablePreemptionCapabilityAnyOf {
	return &NullablePreemptionCapabilityAnyOf{value: val, isSet: true}
}

func (v NullablePreemptionCapabilityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreemptionCapabilityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

