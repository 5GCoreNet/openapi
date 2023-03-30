/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
	"fmt"
)

// DnaiChangeType Possible values are: - EARLY: Early notification of UP path reconfiguration. - EARLY_LATE: Early and late notification of UP path reconfiguration. This value shall   only be present in the subscription to the DNAI change event. - LATE: Late notification of UP path reconfiguration.  
type DnaiChangeType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnaiChangeType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnaiChangeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnaiChangeType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
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

