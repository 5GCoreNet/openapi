/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
	"fmt"
)

// EventThresholdTypeEventThresholdRSRQ - struct for EventThresholdTypeEventThresholdRSRQ
type EventThresholdTypeEventThresholdRSRQ struct {
	Int32 *int32
}

// int32AsEventThresholdTypeEventThresholdRSRQ is a convenience function that returns int32 wrapped in EventThresholdTypeEventThresholdRSRQ
func Int32AsEventThresholdTypeEventThresholdRSRQ(v *int32) EventThresholdTypeEventThresholdRSRQ {
	return EventThresholdTypeEventThresholdRSRQ{
		Int32: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *EventThresholdTypeEventThresholdRSRQ) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Int32
	err = newStrictDecoder(data).Decode(&dst.Int32)
	if err == nil {
		jsonInt32, _ := json.Marshal(dst.Int32)
		if string(jsonInt32) == "{}" { // empty struct
			dst.Int32 = nil
		} else {
			match++
		}
	} else {
		dst.Int32 = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Int32 = nil

		return fmt.Errorf("data matches more than one schema in oneOf(EventThresholdTypeEventThresholdRSRQ)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(EventThresholdTypeEventThresholdRSRQ)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src EventThresholdTypeEventThresholdRSRQ) MarshalJSON() ([]byte, error) {
	if src.Int32 != nil {
		return json.Marshal(&src.Int32)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *EventThresholdTypeEventThresholdRSRQ) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Int32 != nil {
		return obj.Int32
	}

	// all schemas are nil
	return nil
}

type NullableEventThresholdTypeEventThresholdRSRQ struct {
	value *EventThresholdTypeEventThresholdRSRQ
	isSet bool
}

func (v NullableEventThresholdTypeEventThresholdRSRQ) Get() *EventThresholdTypeEventThresholdRSRQ {
	return v.value
}

func (v *NullableEventThresholdTypeEventThresholdRSRQ) Set(val *EventThresholdTypeEventThresholdRSRQ) {
	v.value = val
	v.isSet = true
}

func (v NullableEventThresholdTypeEventThresholdRSRQ) IsSet() bool {
	return v.isSet
}

func (v *NullableEventThresholdTypeEventThresholdRSRQ) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventThresholdTypeEventThresholdRSRQ(val *EventThresholdTypeEventThresholdRSRQ) *NullableEventThresholdTypeEventThresholdRSRQ {
	return &NullableEventThresholdTypeEventThresholdRSRQ{value: val, isSet: true}
}

func (v NullableEventThresholdTypeEventThresholdRSRQ) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventThresholdTypeEventThresholdRSRQ) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
