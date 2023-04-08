/*
Naf_ProSe API

Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_ProSe

import (
	"encoding/json"
	"fmt"
)

// RevocationResult Possible values are: - REVOCATION_SUCCESSFUL: Indicates the successful revocation for a set of Banned RPAUID - Banned PDUID for Restricted ProSe Direct Discovery. - REVOCATION_NOT_SUCCESSFUL: Indicates that unsuccessful revocation for a set of Banned RPAUID - Banned PDUID for Restricted ProSe Direct Discovery. 
type RevocationResult struct {
	RevocationResultAnyOf *RevocationResultAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RevocationResult) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RevocationResultAnyOf
	err = json.Unmarshal(data, &dst.RevocationResultAnyOf);
	if err == nil {
		jsonRevocationResultAnyOf, _ := json.Marshal(dst.RevocationResultAnyOf)
		if string(jsonRevocationResultAnyOf) == "{}" { // empty struct
			dst.RevocationResultAnyOf = nil
		} else {
			return nil // data stored in dst.RevocationResultAnyOf, return on the first match
		}
	} else {
		dst.RevocationResultAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RevocationResult)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RevocationResult) MarshalJSON() ([]byte, error) {
	if src.RevocationResultAnyOf != nil {
		return json.Marshal(&src.RevocationResultAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRevocationResult struct {
	value *RevocationResult
	isSet bool
}

func (v NullableRevocationResult) Get() *RevocationResult {
	return v.value
}

func (v *NullableRevocationResult) Set(val *RevocationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableRevocationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableRevocationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevocationResult(val *RevocationResult) *NullableRevocationResult {
	return &NullableRevocationResult{value: val, isSet: true}
}

func (v NullableRevocationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevocationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


