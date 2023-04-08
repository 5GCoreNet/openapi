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

// SACEventTrigger Describes how NSACF should generate the report for the event
type SACEventTrigger struct {
	SACEventTriggerAnyOf *SACEventTriggerAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SACEventTrigger) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SACEventTriggerAnyOf
	err = json.Unmarshal(data, &dst.SACEventTriggerAnyOf);
	if err == nil {
		jsonSACEventTriggerAnyOf, _ := json.Marshal(dst.SACEventTriggerAnyOf)
		if string(jsonSACEventTriggerAnyOf) == "{}" { // empty struct
			dst.SACEventTriggerAnyOf = nil
		} else {
			return nil // data stored in dst.SACEventTriggerAnyOf, return on the first match
		}
	} else {
		dst.SACEventTriggerAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
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

	return fmt.Errorf("data failed to match schemas in anyOf(SACEventTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SACEventTrigger) MarshalJSON() ([]byte, error) {
	if src.SACEventTriggerAnyOf != nil {
		return json.Marshal(&src.SACEventTriggerAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSACEventTrigger struct {
	value *SACEventTrigger
	isSet bool
}

func (v NullableSACEventTrigger) Get() *SACEventTrigger {
	return v.value
}

func (v *NullableSACEventTrigger) Set(val *SACEventTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableSACEventTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableSACEventTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSACEventTrigger(val *SACEventTrigger) *NullableSACEventTrigger {
	return &NullableSACEventTrigger{value: val, isSet: true}
}

func (v NullableSACEventTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSACEventTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


