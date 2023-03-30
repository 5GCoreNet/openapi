/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// N1MessageClass Enumeration for N1 Message Class
type N1MessageClass struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N1MessageClass) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(N1MessageClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N1MessageClass) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN1MessageClass struct {
	value *N1MessageClass
	isSet bool
}

func (v NullableN1MessageClass) Get() *N1MessageClass {
	return v.value
}

func (v *NullableN1MessageClass) Set(val *N1MessageClass) {
	v.value = val
	v.isSet = true
}

func (v NullableN1MessageClass) IsSet() bool {
	return v.isSet
}

func (v *NullableN1MessageClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN1MessageClass(val *N1MessageClass) *NullableN1MessageClass {
	return &NullableN1MessageClass{value: val, isSet: true}
}

func (v NullableN1MessageClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN1MessageClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


