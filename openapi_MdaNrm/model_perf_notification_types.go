/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
	"fmt"
)

// PerfNotificationTypes the model 'PerfNotificationTypes'
type PerfNotificationTypes string

// List of PerfNotificationTypes
const (
	NOTIFY_THRESHOLD_CROSSING PerfNotificationTypes = "notifyThresholdCrossing"
)

// All allowed values of PerfNotificationTypes enum
var AllowedPerfNotificationTypesEnumValues = []PerfNotificationTypes{
	"notifyThresholdCrossing",
}

func (v *PerfNotificationTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PerfNotificationTypes(value)
	for _, existing := range AllowedPerfNotificationTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PerfNotificationTypes", value)
}

// NewPerfNotificationTypesFromValue returns a pointer to a valid PerfNotificationTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPerfNotificationTypesFromValue(v string) (*PerfNotificationTypes, error) {
	ev := PerfNotificationTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PerfNotificationTypes: valid values are %v", v, AllowedPerfNotificationTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PerfNotificationTypes) IsValid() bool {
	for _, existing := range AllowedPerfNotificationTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PerfNotificationTypes value
func (v PerfNotificationTypes) Ptr() *PerfNotificationTypes {
	return &v
}

type NullablePerfNotificationTypes struct {
	value *PerfNotificationTypes
	isSet bool
}

func (v NullablePerfNotificationTypes) Get() *PerfNotificationTypes {
	return v.value
}

func (v *NullablePerfNotificationTypes) Set(val *PerfNotificationTypes) {
	v.value = val
	v.isSet = true
}

func (v NullablePerfNotificationTypes) IsSet() bool {
	return v.isSet
}

func (v *NullablePerfNotificationTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerfNotificationTypes(val *PerfNotificationTypes) *NullablePerfNotificationTypes {
	return &NullablePerfNotificationTypes{value: val, isSet: true}
}

func (v NullablePerfNotificationTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerfNotificationTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
