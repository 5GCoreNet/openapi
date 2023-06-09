/*
LMF Location

LMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Location

import (
	"encoding/json"
	"fmt"
)

// EventClass Specifies event classes.
type EventClass struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EventClass) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EventClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EventClass) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEventClass struct {
	value *EventClass
	isSet bool
}

func (v NullableEventClass) Get() *EventClass {
	return v.value
}

func (v *NullableEventClass) Set(val *EventClass) {
	v.value = val
	v.isSet = true
}

func (v NullableEventClass) IsSet() bool {
	return v.isSet
}

func (v *NullableEventClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventClass(val *EventClass) *NullableEventClass {
	return &NullableEventClass{value: val, isSet: true}
}

func (v NullableEventClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
