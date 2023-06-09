/*
3gpp-service-parameter

API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ServiceParameter

import (
	"encoding/json"
	"fmt"
)

// AuthorizationResult Possible values are: - AUTH_REVOKED: Indicated the service parameters authorization is revoked.
type AuthorizationResult struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AuthorizationResult) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AuthorizationResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AuthorizationResult) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAuthorizationResult struct {
	value *AuthorizationResult
	isSet bool
}

func (v NullableAuthorizationResult) Get() *AuthorizationResult {
	return v.value
}

func (v *NullableAuthorizationResult) Set(val *AuthorizationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationResult(val *AuthorizationResult) *NullableAuthorizationResult {
	return &NullableAuthorizationResult{value: val, isSet: true}
}

func (v NullableAuthorizationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
