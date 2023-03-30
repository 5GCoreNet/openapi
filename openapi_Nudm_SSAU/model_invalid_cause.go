/*
Nudm_SSAU

Nudm Service Specific Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SSAU

import (
	"encoding/json"
	"fmt"
)

// InvalidCause Possible values are - SUBSRIPTION_WITHDRAWAL - DNN_REMOVED - SLICE_REMOVED - AUTHORIZATION_REVOKED 
type InvalidCause struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *InvalidCause) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(InvalidCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *InvalidCause) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableInvalidCause struct {
	value *InvalidCause
	isSet bool
}

func (v NullableInvalidCause) Get() *InvalidCause {
	return v.value
}

func (v *NullableInvalidCause) Set(val *InvalidCause) {
	v.value = val
	v.isSet = true
}

func (v NullableInvalidCause) IsSet() bool {
	return v.isSet
}

func (v *NullableInvalidCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInvalidCause(val *InvalidCause) *NullableInvalidCause {
	return &NullableInvalidCause{value: val, isSet: true}
}

func (v NullableInvalidCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInvalidCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


