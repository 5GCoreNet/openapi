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

// OccurrenceInfo Specifies occurrence of event reporting.
type OccurrenceInfo struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OccurrenceInfo) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(OccurrenceInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *OccurrenceInfo) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOccurrenceInfo struct {
	value *OccurrenceInfo
	isSet bool
}

func (v NullableOccurrenceInfo) Get() *OccurrenceInfo {
	return v.value
}

func (v *NullableOccurrenceInfo) Set(val *OccurrenceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableOccurrenceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableOccurrenceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOccurrenceInfo(val *OccurrenceInfo) *NullableOccurrenceInfo {
	return &NullableOccurrenceInfo{value: val, isSet: true}
}

func (v NullableOccurrenceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOccurrenceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
