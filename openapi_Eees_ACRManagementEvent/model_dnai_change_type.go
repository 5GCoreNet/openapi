/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
	"fmt"
)

// DnaiChangeType Possible values are: - EARLY: Early notification of UP path reconfiguration. - EARLY_LATE: Early and late notification of UP path reconfiguration. This value shall   only be present in the subscription to the DNAI change event. - LATE: Late notification of UP path reconfiguration.
type DnaiChangeType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnaiChangeType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnaiChangeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnaiChangeType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnaiChangeType struct {
	value *DnaiChangeType
	isSet bool
}

func (v NullableDnaiChangeType) Get() *DnaiChangeType {
	return v.value
}

func (v *NullableDnaiChangeType) Set(val *DnaiChangeType) {
	v.value = val
	v.isSet = true
}

func (v NullableDnaiChangeType) IsSet() bool {
	return v.isSet
}

func (v *NullableDnaiChangeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnaiChangeType(val *DnaiChangeType) *NullableDnaiChangeType {
	return &NullableDnaiChangeType{value: val, isSet: true}
}

func (v NullableDnaiChangeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnaiChangeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
