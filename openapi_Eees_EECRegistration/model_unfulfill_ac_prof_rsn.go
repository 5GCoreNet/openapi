/*
Eees_EECRegistration

API for EEC registration. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_EECRegistration

import (
	"encoding/json"
	"fmt"
)

// UnfulfillACProfRsn represents reason for unfulfilled AC profile requirements.
type UnfulfillACProfRsn struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UnfulfillACProfRsn) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(UnfulfillACProfRsn)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UnfulfillACProfRsn) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUnfulfillACProfRsn struct {
	value *UnfulfillACProfRsn
	isSet bool
}

func (v NullableUnfulfillACProfRsn) Get() *UnfulfillACProfRsn {
	return v.value
}

func (v *NullableUnfulfillACProfRsn) Set(val *UnfulfillACProfRsn) {
	v.value = val
	v.isSet = true
}

func (v NullableUnfulfillACProfRsn) IsSet() bool {
	return v.isSet
}

func (v *NullableUnfulfillACProfRsn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnfulfillACProfRsn(val *UnfulfillACProfRsn) *NullableUnfulfillACProfRsn {
	return &NullableUnfulfillACProfRsn{value: val, isSet: true}
}

func (v NullableUnfulfillACProfRsn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnfulfillACProfRsn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
