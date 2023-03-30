/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
	"fmt"
)

// AlarmListAlignmentRequirement the model 'AlarmListAlignmentRequirement'
type AlarmListAlignmentRequirement string

// List of AlarmListAlignmentRequirement
const (
	REQUIRED AlarmListAlignmentRequirement = "ALIGNMENT_REQUIRED"
	NOT_REQUIRED AlarmListAlignmentRequirement = "ALIGNMENT_NOT_REQUIRED"
)

// All allowed values of AlarmListAlignmentRequirement enum
var AllowedAlarmListAlignmentRequirementEnumValues = []AlarmListAlignmentRequirement{
	"ALIGNMENT_REQUIRED",
	"ALIGNMENT_NOT_REQUIRED",
}

func (v *AlarmListAlignmentRequirement) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlarmListAlignmentRequirement(value)
	for _, existing := range AllowedAlarmListAlignmentRequirementEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlarmListAlignmentRequirement", value)
}

// NewAlarmListAlignmentRequirementFromValue returns a pointer to a valid AlarmListAlignmentRequirement
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlarmListAlignmentRequirementFromValue(v string) (*AlarmListAlignmentRequirement, error) {
	ev := AlarmListAlignmentRequirement(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlarmListAlignmentRequirement: valid values are %v", v, AllowedAlarmListAlignmentRequirementEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlarmListAlignmentRequirement) IsValid() bool {
	for _, existing := range AllowedAlarmListAlignmentRequirementEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlarmListAlignmentRequirement value
func (v AlarmListAlignmentRequirement) Ptr() *AlarmListAlignmentRequirement {
	return &v
}

type NullableAlarmListAlignmentRequirement struct {
	value *AlarmListAlignmentRequirement
	isSet bool
}

func (v NullableAlarmListAlignmentRequirement) Get() *AlarmListAlignmentRequirement {
	return v.value
}

func (v *NullableAlarmListAlignmentRequirement) Set(val *AlarmListAlignmentRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmListAlignmentRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmListAlignmentRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmListAlignmentRequirement(val *AlarmListAlignmentRequirement) *NullableAlarmListAlignmentRequirement {
	return &NullableAlarmListAlignmentRequirement{value: val, isSet: true}
}

func (v NullableAlarmListAlignmentRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmListAlignmentRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

