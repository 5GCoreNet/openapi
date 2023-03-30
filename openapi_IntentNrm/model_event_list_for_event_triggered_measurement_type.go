/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
	"fmt"
)

// EventListForEventTriggeredMeasurementType See details in 3GPP TS 32.422 clause 5.10.28.
type EventListForEventTriggeredMeasurementType string

// List of eventListForEventTriggeredMeasurement-Type
const (
	OUT_OF_COVERAGE EventListForEventTriggeredMeasurementType = "OUT_OF_COVERAGE"
	A2_EVENT EventListForEventTriggeredMeasurementType = "A2_EVENT"
)

// All allowed values of EventListForEventTriggeredMeasurementType enum
var AllowedEventListForEventTriggeredMeasurementTypeEnumValues = []EventListForEventTriggeredMeasurementType{
	"OUT_OF_COVERAGE",
	"A2_EVENT",
}

func (v *EventListForEventTriggeredMeasurementType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventListForEventTriggeredMeasurementType(value)
	for _, existing := range AllowedEventListForEventTriggeredMeasurementTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventListForEventTriggeredMeasurementType", value)
}

// NewEventListForEventTriggeredMeasurementTypeFromValue returns a pointer to a valid EventListForEventTriggeredMeasurementType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventListForEventTriggeredMeasurementTypeFromValue(v string) (*EventListForEventTriggeredMeasurementType, error) {
	ev := EventListForEventTriggeredMeasurementType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventListForEventTriggeredMeasurementType: valid values are %v", v, AllowedEventListForEventTriggeredMeasurementTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventListForEventTriggeredMeasurementType) IsValid() bool {
	for _, existing := range AllowedEventListForEventTriggeredMeasurementTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to eventListForEventTriggeredMeasurement-Type value
func (v EventListForEventTriggeredMeasurementType) Ptr() *EventListForEventTriggeredMeasurementType {
	return &v
}

type NullableEventListForEventTriggeredMeasurementType struct {
	value *EventListForEventTriggeredMeasurementType
	isSet bool
}

func (v NullableEventListForEventTriggeredMeasurementType) Get() *EventListForEventTriggeredMeasurementType {
	return v.value
}

func (v *NullableEventListForEventTriggeredMeasurementType) Set(val *EventListForEventTriggeredMeasurementType) {
	v.value = val
	v.isSet = true
}

func (v NullableEventListForEventTriggeredMeasurementType) IsSet() bool {
	return v.isSet
}

func (v *NullableEventListForEventTriggeredMeasurementType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventListForEventTriggeredMeasurementType(val *EventListForEventTriggeredMeasurementType) *NullableEventListForEventTriggeredMeasurementType {
	return &NullableEventListForEventTriggeredMeasurementType{value: val, isSet: true}
}

func (v NullableEventListForEventTriggeredMeasurementType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventListForEventTriggeredMeasurementType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

