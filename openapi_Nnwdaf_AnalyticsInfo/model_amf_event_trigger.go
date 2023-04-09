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

// AmfEventTrigger Describes how AMF should generate the report for the event
type AmfEventTrigger struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmfEventTrigger) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AmfEventTrigger)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmfEventTrigger) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmfEventTrigger struct {
	value *AmfEventTrigger
	isSet bool
}

func (v NullableAmfEventTrigger) Get() *AmfEventTrigger {
	return v.value
}

func (v *NullableAmfEventTrigger) Set(val *AmfEventTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventTrigger(val *AmfEventTrigger) *NullableAmfEventTrigger {
	return &NullableAmfEventTrigger{value: val, isSet: true}
}

func (v NullableAmfEventTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
