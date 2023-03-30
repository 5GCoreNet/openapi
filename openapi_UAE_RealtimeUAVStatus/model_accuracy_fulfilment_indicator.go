/*
UAE Server Real-time UAV Status Service

UAE Server Real-time UAV Status Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_UAE_RealtimeUAVStatus

import (
	"encoding/json"
	"fmt"
)

// AccuracyFulfilmentIndicator Indicates fulfilment of requested accuracy.
type AccuracyFulfilmentIndicator struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccuracyFulfilmentIndicator) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccuracyFulfilmentIndicator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccuracyFulfilmentIndicator) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccuracyFulfilmentIndicator struct {
	value *AccuracyFulfilmentIndicator
	isSet bool
}

func (v NullableAccuracyFulfilmentIndicator) Get() *AccuracyFulfilmentIndicator {
	return v.value
}

func (v *NullableAccuracyFulfilmentIndicator) Set(val *AccuracyFulfilmentIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableAccuracyFulfilmentIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableAccuracyFulfilmentIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccuracyFulfilmentIndicator(val *AccuracyFulfilmentIndicator) *NullableAccuracyFulfilmentIndicator {
	return &NullableAccuracyFulfilmentIndicator{value: val, isSet: true}
}

func (v NullableAccuracyFulfilmentIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccuracyFulfilmentIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


