/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
	"fmt"
)

// RevokedCause struct for RevokedCause
type RevokedCause struct {
	RevokedCauseAnyOf *RevokedCauseAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RevokedCause) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RevokedCauseAnyOf
	err = json.Unmarshal(data, &dst.RevokedCauseAnyOf);
	if err == nil {
		jsonRevokedCauseAnyOf, _ := json.Marshal(dst.RevokedCauseAnyOf)
		if string(jsonRevokedCauseAnyOf) == "{}" { // empty struct
			dst.RevokedCauseAnyOf = nil
		} else {
			return nil // data stored in dst.RevokedCauseAnyOf, return on the first match
		}
	} else {
		dst.RevokedCauseAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RevokedCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RevokedCause) MarshalJSON() ([]byte, error) {
	if src.RevokedCauseAnyOf != nil {
		return json.Marshal(&src.RevokedCauseAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRevokedCause struct {
	value *RevokedCause
	isSet bool
}

func (v NullableRevokedCause) Get() *RevokedCause {
	return v.value
}

func (v *NullableRevokedCause) Set(val *RevokedCause) {
	v.value = val
	v.isSet = true
}

func (v NullableRevokedCause) IsSet() bool {
	return v.isSet
}

func (v *NullableRevokedCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRevokedCause(val *RevokedCause) *NullableRevokedCause {
	return &NullableRevokedCause{value: val, isSet: true}
}

func (v NullableRevokedCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRevokedCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


