/*
GBA BSF Nbsp_GBA Service

GBA BSF Nbsp_GBA Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nbsp_GBA

import (
	"encoding/json"
	"fmt"
)

// GbaType Authentication type used by the UE for GBA
type GbaType struct {
	GbaTypeAnyOf *GbaTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GbaType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into GbaTypeAnyOf
	err = json.Unmarshal(data, &dst.GbaTypeAnyOf);
	if err == nil {
		jsonGbaTypeAnyOf, _ := json.Marshal(dst.GbaTypeAnyOf)
		if string(jsonGbaTypeAnyOf) == "{}" { // empty struct
			dst.GbaTypeAnyOf = nil
		} else {
			return nil // data stored in dst.GbaTypeAnyOf, return on the first match
		}
	} else {
		dst.GbaTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(GbaType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GbaType) MarshalJSON() ([]byte, error) {
	if src.GbaTypeAnyOf != nil {
		return json.Marshal(&src.GbaTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGbaType struct {
	value *GbaType
	isSet bool
}

func (v NullableGbaType) Get() *GbaType {
	return v.value
}

func (v *NullableGbaType) Set(val *GbaType) {
	v.value = val
	v.isSet = true
}

func (v NullableGbaType) IsSet() bool {
	return v.isSet
}

func (v *NullableGbaType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGbaType(val *GbaType) *NullableGbaType {
	return &NullableGbaType{value: val, isSet: true}
}

func (v NullableGbaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGbaType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


