/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-ddnmf_Discovery

import (
	"encoding/json"
	"fmt"
)

// DiscoveryType Possible values are - OPEN: Discovery type is \"open\". - RESTRICTED: Discovery type is \"restricted\". 
type DiscoveryType struct {
	DiscoveryTypeAnyOf *DiscoveryTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DiscoveryType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DiscoveryTypeAnyOf
	err = json.Unmarshal(data, &dst.DiscoveryTypeAnyOf);
	if err == nil {
		jsonDiscoveryTypeAnyOf, _ := json.Marshal(dst.DiscoveryTypeAnyOf)
		if string(jsonDiscoveryTypeAnyOf) == "{}" { // empty struct
			dst.DiscoveryTypeAnyOf = nil
		} else {
			return nil // data stored in dst.DiscoveryTypeAnyOf, return on the first match
		}
	} else {
		dst.DiscoveryTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DiscoveryType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DiscoveryType) MarshalJSON() ([]byte, error) {
	if src.DiscoveryTypeAnyOf != nil {
		return json.Marshal(&src.DiscoveryTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDiscoveryType struct {
	value *DiscoveryType
	isSet bool
}

func (v NullableDiscoveryType) Get() *DiscoveryType {
	return v.value
}

func (v *NullableDiscoveryType) Set(val *DiscoveryType) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryType) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryType(val *DiscoveryType) *NullableDiscoveryType {
	return &NullableDiscoveryType{value: val, isSet: true}
}

func (v NullableDiscoveryType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


