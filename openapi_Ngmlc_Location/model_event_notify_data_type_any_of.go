/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// EventNotifyDataTypeAnyOf the model 'EventNotifyDataTypeAnyOf'
type EventNotifyDataTypeAnyOf string

// List of EventNotifyDataType_anyOf
const (
	UE_AVAILABLE EventNotifyDataTypeAnyOf = "UE_AVAILABLE"
	PERIODIC EventNotifyDataTypeAnyOf = "PERIODIC"
	ENTERING_INTO_AREA EventNotifyDataTypeAnyOf = "ENTERING_INTO_AREA"
	LEAVING_FROM_AREA EventNotifyDataTypeAnyOf = "LEAVING_FROM_AREA"
	BEING_INSIDE_AREA EventNotifyDataTypeAnyOf = "BEING_INSIDE_AREA"
	MOTION EventNotifyDataTypeAnyOf = "MOTION"
	MAXIMUM_INTERVAL_EXPIRATION_EVENT EventNotifyDataTypeAnyOf = "MAXIMUM_INTERVAL_EXPIRATION_EVENT"
	LOCATION_CANCELLATION_EVENT EventNotifyDataTypeAnyOf = "LOCATION_CANCELLATION_EVENT"
	ACTIVATION_OF_DEFERRED_LOCATION EventNotifyDataTypeAnyOf = "ACTIVATION_OF_DEFERRED_LOCATION"
	UE_MOBILITY_FOR_DEFERRED_LOCATION EventNotifyDataTypeAnyOf = "UE_MOBILITY_FOR_DEFERRED_LOCATION"
	_5_GC_MT_LR EventNotifyDataTypeAnyOf = "5GC_MT_LR"
)

// All allowed values of EventNotifyDataTypeAnyOf enum
var AllowedEventNotifyDataTypeAnyOfEnumValues = []EventNotifyDataTypeAnyOf{
	"UE_AVAILABLE",
	"PERIODIC",
	"ENTERING_INTO_AREA",
	"LEAVING_FROM_AREA",
	"BEING_INSIDE_AREA",
	"MOTION",
	"MAXIMUM_INTERVAL_EXPIRATION_EVENT",
	"LOCATION_CANCELLATION_EVENT",
	"ACTIVATION_OF_DEFERRED_LOCATION",
	"UE_MOBILITY_FOR_DEFERRED_LOCATION",
	"5GC_MT_LR",
}

func (v *EventNotifyDataTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventNotifyDataTypeAnyOf(value)
	for _, existing := range AllowedEventNotifyDataTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventNotifyDataTypeAnyOf", value)
}

// NewEventNotifyDataTypeAnyOfFromValue returns a pointer to a valid EventNotifyDataTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventNotifyDataTypeAnyOfFromValue(v string) (*EventNotifyDataTypeAnyOf, error) {
	ev := EventNotifyDataTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventNotifyDataTypeAnyOf: valid values are %v", v, AllowedEventNotifyDataTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventNotifyDataTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedEventNotifyDataTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventNotifyDataType_anyOf value
func (v EventNotifyDataTypeAnyOf) Ptr() *EventNotifyDataTypeAnyOf {
	return &v
}

type NullableEventNotifyDataTypeAnyOf struct {
	value *EventNotifyDataTypeAnyOf
	isSet bool
}

func (v NullableEventNotifyDataTypeAnyOf) Get() *EventNotifyDataTypeAnyOf {
	return v.value
}

func (v *NullableEventNotifyDataTypeAnyOf) Set(val *EventNotifyDataTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventNotifyDataTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventNotifyDataTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventNotifyDataTypeAnyOf(val *EventNotifyDataTypeAnyOf) *NullableEventNotifyDataTypeAnyOf {
	return &NullableEventNotifyDataTypeAnyOf{value: val, isSet: true}
}

func (v NullableEventNotifyDataTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventNotifyDataTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

