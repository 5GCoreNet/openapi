/*
VAE_ApplicationRequirement

API for VAE Application Requirement Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_VAE_ApplicationRequirement

import (
	"encoding/json"
	"fmt"
)

// ReservationResult Represents the result of the network resource adaptation corresponding to the V2X  application requirement. 
type ReservationResult struct {
	ReservationResultAnyOf *ReservationResultAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReservationResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReservationResultAnyOf
	err = json.Unmarshal(data, &dst.ReservationResultAnyOf);
	if err == nil {
		jsonReservationResultAnyOf, _ := json.Marshal(dst.ReservationResultAnyOf)
		if string(jsonReservationResultAnyOf) == "{}" { // empty struct
			dst.ReservationResultAnyOf = nil
		} else {
			return nil // data stored in dst.ReservationResultAnyOf, return on the first match
		}
	} else {
		dst.ReservationResultAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReservationResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReservationResult) MarshalJSON() ([]byte, error) {
	if src.ReservationResultAnyOf != nil {
		return json.Marshal(&src.ReservationResultAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReservationResult struct {
	value *ReservationResult
	isSet bool
}

func (v NullableReservationResult) Get() *ReservationResult {
	return v.value
}

func (v *NullableReservationResult) Set(val *ReservationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableReservationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableReservationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservationResult(val *ReservationResult) *NullableReservationResult {
	return &NullableReservationResult{value: val, isSet: true}
}

func (v NullableReservationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


