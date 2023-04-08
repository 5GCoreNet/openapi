/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// SubTerminationReason Subscription Termination Reason.
type SubTerminationReason struct {
	SubTerminationReasonAnyOf *SubTerminationReasonAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SubTerminationReason) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SubTerminationReasonAnyOf
	err = json.Unmarshal(data, &dst.SubTerminationReasonAnyOf);
	if err == nil {
		jsonSubTerminationReasonAnyOf, _ := json.Marshal(dst.SubTerminationReasonAnyOf)
		if string(jsonSubTerminationReasonAnyOf) == "{}" { // empty struct
			dst.SubTerminationReasonAnyOf = nil
		} else {
			return nil // data stored in dst.SubTerminationReasonAnyOf, return on the first match
		}
	} else {
		dst.SubTerminationReasonAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(SubTerminationReason)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SubTerminationReason) MarshalJSON() ([]byte, error) {
	if src.SubTerminationReasonAnyOf != nil {
		return json.Marshal(&src.SubTerminationReasonAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSubTerminationReason struct {
	value *SubTerminationReason
	isSet bool
}

func (v NullableSubTerminationReason) Get() *SubTerminationReason {
	return v.value
}

func (v *NullableSubTerminationReason) Set(val *SubTerminationReason) {
	v.value = val
	v.isSet = true
}

func (v NullableSubTerminationReason) IsSet() bool {
	return v.isSet
}

func (v *NullableSubTerminationReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubTerminationReason(val *SubTerminationReason) *NullableSubTerminationReason {
	return &NullableSubTerminationReason{value: val, isSet: true}
}

func (v NullableSubTerminationReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubTerminationReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


