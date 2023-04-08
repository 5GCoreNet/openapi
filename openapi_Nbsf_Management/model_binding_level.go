/*
Nbsf_Management

Binding Support Management Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.4.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsf_Management

import (
	"encoding/json"
	"fmt"
)

// BindingLevel Possible values are: - \"NF_SET\" - \"NF_INSTANCE\" 
type BindingLevel struct {
	BindingLevelAnyOf *BindingLevelAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BindingLevel) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BindingLevelAnyOf
	err = json.Unmarshal(data, &dst.BindingLevelAnyOf);
	if err == nil {
		jsonBindingLevelAnyOf, _ := json.Marshal(dst.BindingLevelAnyOf)
		if string(jsonBindingLevelAnyOf) == "{}" { // empty struct
			dst.BindingLevelAnyOf = nil
		} else {
			return nil // data stored in dst.BindingLevelAnyOf, return on the first match
		}
	} else {
		dst.BindingLevelAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(BindingLevel)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *BindingLevel) MarshalJSON() ([]byte, error) {
	if src.BindingLevelAnyOf != nil {
		return json.Marshal(&src.BindingLevelAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableBindingLevel struct {
	value *BindingLevel
	isSet bool
}

func (v NullableBindingLevel) Get() *BindingLevel {
	return v.value
}

func (v *NullableBindingLevel) Set(val *BindingLevel) {
	v.value = val
	v.isSet = true
}

func (v NullableBindingLevel) IsSet() bool {
	return v.isSet
}

func (v *NullableBindingLevel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBindingLevel(val *BindingLevel) *NullableBindingLevel {
	return &NullableBindingLevel{value: val, isSet: true}
}

func (v NullableBindingLevel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBindingLevel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


