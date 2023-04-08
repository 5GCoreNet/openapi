/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
	"fmt"
)

// N1MessageClass Enumeration for N1 Message Class
type N1MessageClass struct {
	N1MessageClassAnyOf *N1MessageClassAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N1MessageClass) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into N1MessageClassAnyOf
	err = json.Unmarshal(data, &dst.N1MessageClassAnyOf);
	if err == nil {
		jsonN1MessageClassAnyOf, _ := json.Marshal(dst.N1MessageClassAnyOf)
		if string(jsonN1MessageClassAnyOf) == "{}" { // empty struct
			dst.N1MessageClassAnyOf = nil
		} else {
			return nil // data stored in dst.N1MessageClassAnyOf, return on the first match
		}
	} else {
		dst.N1MessageClassAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(N1MessageClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N1MessageClass) MarshalJSON() ([]byte, error) {
	if src.N1MessageClassAnyOf != nil {
		return json.Marshal(&src.N1MessageClassAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
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


