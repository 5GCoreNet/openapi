/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// UriScheme HTTP and HTTPS URI scheme.
type UriScheme struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UriScheme) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(UriScheme)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UriScheme) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUriScheme struct {
	value *UriScheme
	isSet bool
}

func (v NullableUriScheme) Get() *UriScheme {
	return v.value
}

func (v *NullableUriScheme) Set(val *UriScheme) {
	v.value = val
	v.isSet = true
}

func (v NullableUriScheme) IsSet() bool {
	return v.isSet
}

func (v *NullableUriScheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUriScheme(val *UriScheme) *NullableUriScheme {
	return &NullableUriScheme{value: val, isSet: true}
}

func (v NullableUriScheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUriScheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
