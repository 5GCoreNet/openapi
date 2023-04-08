/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// LdrType Indicates LDR types.
type LdrType struct {
	LdrTypeAnyOf *LdrTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LdrType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LdrTypeAnyOf
	err = json.Unmarshal(data, &dst.LdrTypeAnyOf);
	if err == nil {
		jsonLdrTypeAnyOf, _ := json.Marshal(dst.LdrTypeAnyOf)
		if string(jsonLdrTypeAnyOf) == "{}" { // empty struct
			dst.LdrTypeAnyOf = nil
		} else {
			return nil // data stored in dst.LdrTypeAnyOf, return on the first match
		}
	} else {
		dst.LdrTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LdrType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LdrType) MarshalJSON() ([]byte, error) {
	if src.LdrTypeAnyOf != nil {
		return json.Marshal(&src.LdrTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLdrType struct {
	value *LdrType
	isSet bool
}

func (v NullableLdrType) Get() *LdrType {
	return v.value
}

func (v *NullableLdrType) Set(val *LdrType) {
	v.value = val
	v.isSet = true
}

func (v NullableLdrType) IsSet() bool {
	return v.isSet
}

func (v *NullableLdrType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLdrType(val *LdrType) *NullableLdrType {
	return &NullableLdrType{value: val, isSet: true}
}

func (v NullableLdrType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLdrType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


