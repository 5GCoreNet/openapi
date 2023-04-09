/*
Unified Data Repository Service API file for policy data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Policy_Data

import (
	"encoding/json"
	"fmt"
)

// Periodicity Represents the time period.
type Periodicity struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Periodicity) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(Periodicity)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Periodicity) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePeriodicity struct {
	value *Periodicity
	isSet bool
}

func (v NullablePeriodicity) Get() *Periodicity {
	return v.value
}

func (v *NullablePeriodicity) Set(val *Periodicity) {
	v.value = val
	v.isSet = true
}

func (v NullablePeriodicity) IsSet() bool {
	return v.isSet
}

func (v *NullablePeriodicity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePeriodicity(val *Periodicity) *NullablePeriodicity {
	return &NullablePeriodicity{value: val, isSet: true}
}

func (v NullablePeriodicity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePeriodicity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
