/*
Naf_Authentication

AF Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_Authentication

import (
	"encoding/json"
	"fmt"
)

// AuthResult Possible values are: - AUTH_SUCCESS: The UUAA or C2 authorization has succeeded. - AUTH_FAIL: The UUAA or C2 authorization has failed. 
type AuthResult struct {
	AuthResultAnyOf *AuthResultAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AuthResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AuthResultAnyOf
	err = json.Unmarshal(data, &dst.AuthResultAnyOf);
	if err == nil {
		jsonAuthResultAnyOf, _ := json.Marshal(dst.AuthResultAnyOf)
		if string(jsonAuthResultAnyOf) == "{}" { // empty struct
			dst.AuthResultAnyOf = nil
		} else {
			return nil // data stored in dst.AuthResultAnyOf, return on the first match
		}
	} else {
		dst.AuthResultAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AuthResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AuthResult) MarshalJSON() ([]byte, error) {
	if src.AuthResultAnyOf != nil {
		return json.Marshal(&src.AuthResultAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAuthResult struct {
	value *AuthResult
	isSet bool
}

func (v NullableAuthResult) Get() *AuthResult {
	return v.value
}

func (v *NullableAuthResult) Set(val *AuthResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthResult(val *AuthResult) *NullableAuthResult {
	return &NullableAuthResult{value: val, isSet: true}
}

func (v NullableAuthResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


