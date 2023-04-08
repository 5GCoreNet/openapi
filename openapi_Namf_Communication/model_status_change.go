/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// StatusChange Enumeration for AMF status
type StatusChange struct {
	StatusChangeAnyOf *StatusChangeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *StatusChange) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into StatusChangeAnyOf
	err = json.Unmarshal(data, &dst.StatusChangeAnyOf);
	if err == nil {
		jsonStatusChangeAnyOf, _ := json.Marshal(dst.StatusChangeAnyOf)
		if string(jsonStatusChangeAnyOf) == "{}" { // empty struct
			dst.StatusChangeAnyOf = nil
		} else {
			return nil // data stored in dst.StatusChangeAnyOf, return on the first match
		}
	} else {
		dst.StatusChangeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(StatusChange)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *StatusChange) MarshalJSON() ([]byte, error) {
	if src.StatusChangeAnyOf != nil {
		return json.Marshal(&src.StatusChangeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableStatusChange struct {
	value *StatusChange
	isSet bool
}

func (v NullableStatusChange) Get() *StatusChange {
	return v.value
}

func (v *NullableStatusChange) Set(val *StatusChange) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusChange) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusChange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusChange(val *StatusChange) *NullableStatusChange {
	return &NullableStatusChange{value: val, isSet: true}
}

func (v NullableStatusChange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusChange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


