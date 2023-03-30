/*
AEF_Security_API

API for AEF security management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AEF_Security_API

import (
	"encoding/json"
	"fmt"
)

// Cause Possible values are: - OVERLIMIT_USAGE:      The revocation of the authorization of the API invoker is due to the overlimit      usage of the service API - UNEXPECTED_REASON:      The revocation of the authorization of the API invoker is due to unexpected reason. 
type Cause struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Cause) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(Cause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Cause) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCause struct {
	value *Cause
	isSet bool
}

func (v NullableCause) Get() *Cause {
	return v.value
}

func (v *NullableCause) Set(val *Cause) {
	v.value = val
	v.isSet = true
}

func (v NullableCause) IsSet() bool {
	return v.isSet
}

func (v *NullableCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCause(val *Cause) *NullableCause {
	return &NullableCause{value: val, isSet: true}
}

func (v NullableCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


