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

// AfEventAnyOf the model 'AfEventAnyOf'
type AfEventAnyOf string

// List of AfEvent_anyOf
const (
	SVC_EXPERIENCE AfEventAnyOf = "SVC_EXPERIENCE"
	UE_MOBILITY AfEventAnyOf = "UE_MOBILITY"
	UE_COMM AfEventAnyOf = "UE_COMM"
	EXCEPTIONS AfEventAnyOf = "EXCEPTIONS"
	USER_DATA_CONGESTION AfEventAnyOf = "USER_DATA_CONGESTION"
	PERF_DATA AfEventAnyOf = "PERF_DATA"
	DISPERSION AfEventAnyOf = "DISPERSION"
	COLLECTIVE_BEHAVIOUR AfEventAnyOf = "COLLECTIVE_BEHAVIOUR"
	MS_QOE_METRICS AfEventAnyOf = "MS_QOE_METRICS"
	MS_CONSUMPTION AfEventAnyOf = "MS_CONSUMPTION"
	MS_NET_ASSIST_INVOCATION AfEventAnyOf = "MS_NET_ASSIST_INVOCATION"
	MS_DYN_POLICY_INVOCATION AfEventAnyOf = "MS_DYN_POLICY_INVOCATION"
	MS_ACCESS_ACTIVITY AfEventAnyOf = "MS_ACCESS_ACTIVITY"
)

// All allowed values of AfEventAnyOf enum
var AllowedAfEventAnyOfEnumValues = []AfEventAnyOf{
	"SVC_EXPERIENCE",
	"UE_MOBILITY",
	"UE_COMM",
	"EXCEPTIONS",
	"USER_DATA_CONGESTION",
	"PERF_DATA",
	"DISPERSION",
	"COLLECTIVE_BEHAVIOUR",
	"MS_QOE_METRICS",
	"MS_CONSUMPTION",
	"MS_NET_ASSIST_INVOCATION",
	"MS_DYN_POLICY_INVOCATION",
	"MS_ACCESS_ACTIVITY",
}

func (v *AfEventAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AfEventAnyOf(value)
	for _, existing := range AllowedAfEventAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AfEventAnyOf", value)
}

// NewAfEventAnyOfFromValue returns a pointer to a valid AfEventAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAfEventAnyOfFromValue(v string) (*AfEventAnyOf, error) {
	ev := AfEventAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AfEventAnyOf: valid values are %v", v, AllowedAfEventAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AfEventAnyOf) IsValid() bool {
	for _, existing := range AllowedAfEventAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AfEvent_anyOf value
func (v AfEventAnyOf) Ptr() *AfEventAnyOf {
	return &v
}

type NullableAfEventAnyOf struct {
	value *AfEventAnyOf
	isSet bool
}

func (v NullableAfEventAnyOf) Get() *AfEventAnyOf {
	return v.value
}

func (v *NullableAfEventAnyOf) Set(val *AfEventAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAfEventAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAfEventAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfEventAnyOf(val *AfEventAnyOf) *NullableAfEventAnyOf {
	return &NullableAfEventAnyOf{value: val, isSet: true}
}

func (v NullableAfEventAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfEventAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

