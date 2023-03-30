/*
SS_LocationAreaInfoRetrieval

API for SEAL Location Area Info Retrieval.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_LocationAreaInfoRetrieval

import (
	"encoding/json"
	"fmt"
)

// ValTargetUe - Represents the information identifying a VAL user ID or a VAL UE ID.
type ValTargetUe struct {
	Interface{} *interface{}
}

// interface{}AsValTargetUe is a convenience function that returns interface{} wrapped in ValTargetUe
func Interface{}AsValTargetUe(v *interface{}) ValTargetUe {
	return ValTargetUe{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ValTargetUe) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ValTargetUe)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ValTargetUe)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ValTargetUe) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ValTargetUe) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableValTargetUe struct {
	value *ValTargetUe
	isSet bool
}

func (v NullableValTargetUe) Get() *ValTargetUe {
	return v.value
}

func (v *NullableValTargetUe) Set(val *ValTargetUe) {
	v.value = val
	v.isSet = true
}

func (v NullableValTargetUe) IsSet() bool {
	return v.isSet
}

func (v *NullableValTargetUe) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableValTargetUe(val *ValTargetUe) *NullableValTargetUe {
	return &NullableValTargetUe{value: val, isSet: true}
}

func (v NullableValTargetUe) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableValTargetUe) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


