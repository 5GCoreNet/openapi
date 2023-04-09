/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
	"fmt"
)

// GnssId Global Navigation Satellite System (GNSS) ID.
type GnssId struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GnssId) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(GnssId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GnssId) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGnssId struct {
	value *GnssId
	isSet bool
}

func (v NullableGnssId) Get() *GnssId {
	return v.value
}

func (v *NullableGnssId) Set(val *GnssId) {
	v.value = val
	v.isSet = true
}

func (v NullableGnssId) IsSet() bool {
	return v.isSet
}

func (v *NullableGnssId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGnssId(val *GnssId) *NullableGnssId {
	return &NullableGnssId{value: val, isSet: true}
}

func (v NullableGnssId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGnssId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
