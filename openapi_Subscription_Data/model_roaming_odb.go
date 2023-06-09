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

// RoamingOdb The enumeration RoamingOdb defines the Barring of Roaming as. See 3GPP TS 23.015 for further description. It shall comply with the provisions defined in table 5.7.3.1-1.
type RoamingOdb struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RoamingOdb) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RoamingOdb)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RoamingOdb) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRoamingOdb struct {
	value *RoamingOdb
	isSet bool
}

func (v NullableRoamingOdb) Get() *RoamingOdb {
	return v.value
}

func (v *NullableRoamingOdb) Set(val *RoamingOdb) {
	v.value = val
	v.isSet = true
}

func (v NullableRoamingOdb) IsSet() bool {
	return v.isSet
}

func (v *NullableRoamingOdb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoamingOdb(val *RoamingOdb) *NullableRoamingOdb {
	return &NullableRoamingOdb{value: val, isSet: true}
}

func (v NullableRoamingOdb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoamingOdb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
