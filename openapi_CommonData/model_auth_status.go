/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// AuthStatus Possible values are: - \"EAP_SUCCESS\": The NSSAA status is EAP-Success. - \"EAP_FAILURE\": The NSSAA status is EAP-Failure. - \"PENDING\": The NSSAA status is Pending.  
type AuthStatus struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AuthStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AuthStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AuthStatus) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAuthStatus struct {
	value *AuthStatus
	isSet bool
}

func (v NullableAuthStatus) Get() *AuthStatus {
	return v.value
}

func (v *NullableAuthStatus) Set(val *AuthStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthStatus(val *AuthStatus) *NullableAuthStatus {
	return &NullableAuthStatus{value: val, isSet: true}
}

func (v NullableAuthStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


