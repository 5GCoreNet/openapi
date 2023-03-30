/*
Ntsctsf_QoSandTSCAssistance Service API

TSCTSF QoS and TSC Assistance Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_QoSandTSCAssistance

import (
	"encoding/json"
	"fmt"
)

// TscEvent Represents an event to notify to the AF.
type TscEvent struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TscEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TscEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TscEvent) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTscEvent struct {
	value *TscEvent
	isSet bool
}

func (v NullableTscEvent) Get() *TscEvent {
	return v.value
}

func (v *NullableTscEvent) Set(val *TscEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableTscEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableTscEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTscEvent(val *TscEvent) *NullableTscEvent {
	return &NullableTscEvent{value: val, isSet: true}
}

func (v NullableTscEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTscEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

