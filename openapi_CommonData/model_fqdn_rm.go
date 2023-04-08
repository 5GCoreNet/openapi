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

// FqdnRm Fully Qualified Domain Name, but it also allows the null value
type FqdnRm struct {
	NullValue *NullValue
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FqdnRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FqdnRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FqdnRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFqdnRm struct {
	value *FqdnRm
	isSet bool
}

func (v NullableFqdnRm) Get() *FqdnRm {
	return v.value
}

func (v *NullableFqdnRm) Set(val *FqdnRm) {
	v.value = val
	v.isSet = true
}

func (v NullableFqdnRm) IsSet() bool {
	return v.isSet
}

func (v *NullableFqdnRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFqdnRm(val *FqdnRm) *NullableFqdnRm {
	return &NullableFqdnRm{value: val, isSet: true}
}

func (v NullableFqdnRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFqdnRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


