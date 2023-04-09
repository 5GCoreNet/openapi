/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// PduSessionContinuityInd struct for PduSessionContinuityInd
type PduSessionContinuityInd struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PduSessionContinuityInd) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PduSessionContinuityInd)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PduSessionContinuityInd) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePduSessionContinuityInd struct {
	value *PduSessionContinuityInd
	isSet bool
}

func (v NullablePduSessionContinuityInd) Get() *PduSessionContinuityInd {
	return v.value
}

func (v *NullablePduSessionContinuityInd) Set(val *PduSessionContinuityInd) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionContinuityInd) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionContinuityInd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionContinuityInd(val *PduSessionContinuityInd) *NullablePduSessionContinuityInd {
	return &NullablePduSessionContinuityInd{value: val, isSet: true}
}

func (v NullablePduSessionContinuityInd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionContinuityInd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
