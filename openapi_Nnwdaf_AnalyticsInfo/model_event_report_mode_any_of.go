/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// EventReportModeAnyOf the model 'EventReportModeAnyOf'
type EventReportModeAnyOf string

// List of EventReportMode_anyOf
const (
	PERIODIC EventReportModeAnyOf = "PERIODIC"
	ON_EVENT_DETECTION EventReportModeAnyOf = "ON_EVENT_DETECTION"
)

// All allowed values of EventReportModeAnyOf enum
var AllowedEventReportModeAnyOfEnumValues = []EventReportModeAnyOf{
	"PERIODIC",
	"ON_EVENT_DETECTION",
}

func (v *EventReportModeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventReportModeAnyOf(value)
	for _, existing := range AllowedEventReportModeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventReportModeAnyOf", value)
}

// NewEventReportModeAnyOfFromValue returns a pointer to a valid EventReportModeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventReportModeAnyOfFromValue(v string) (*EventReportModeAnyOf, error) {
	ev := EventReportModeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventReportModeAnyOf: valid values are %v", v, AllowedEventReportModeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventReportModeAnyOf) IsValid() bool {
	for _, existing := range AllowedEventReportModeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventReportMode_anyOf value
func (v EventReportModeAnyOf) Ptr() *EventReportModeAnyOf {
	return &v
}

type NullableEventReportModeAnyOf struct {
	value *EventReportModeAnyOf
	isSet bool
}

func (v NullableEventReportModeAnyOf) Get() *EventReportModeAnyOf {
	return v.value
}

func (v *NullableEventReportModeAnyOf) Set(val *EventReportModeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventReportModeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventReportModeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventReportModeAnyOf(val *EventReportModeAnyOf) *NullableEventReportModeAnyOf {
	return &NullableEventReportModeAnyOf{value: val, isSet: true}
}

func (v NullableEventReportModeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventReportModeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

