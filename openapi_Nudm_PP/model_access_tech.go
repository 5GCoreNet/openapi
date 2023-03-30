/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_PP

import (
	"encoding/json"
	"fmt"
)

// AccessTech Represents the access technology
type AccessTech struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessTech) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccessTech)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessTech) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessTech struct {
	value *AccessTech
	isSet bool
}

func (v NullableAccessTech) Get() *AccessTech {
	return v.value
}

func (v *NullableAccessTech) Set(val *AccessTech) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTech) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTech) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTech(val *AccessTech) *NullableAccessTech {
	return &NullableAccessTech{value: val, isSet: true}
}

func (v NullableAccessTech) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTech) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


