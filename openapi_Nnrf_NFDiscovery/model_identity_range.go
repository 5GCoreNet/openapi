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

// IdentityRange - A range of GPSIs (subscriber identities), either based on a numeric range, or based on regular-expression matching
type IdentityRange struct {
	Interface *interface{}
}

// interface{}AsIdentityRange is a convenience function that returns interface{} wrapped in IdentityRange
func InterfaceAsIdentityRange(v *interface{}) IdentityRange {
	return IdentityRange{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IdentityRange) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface
	err = newStrictDecoder(data).Decode(&dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			match++
		}
	} else {
		dst.Interface = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IdentityRange)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IdentityRange)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IdentityRange) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IdentityRange) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableIdentityRange struct {
	value *IdentityRange
	isSet bool
}

func (v NullableIdentityRange) Get() *IdentityRange {
	return v.value
}

func (v *NullableIdentityRange) Set(val *IdentityRange) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityRange) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityRange(val *IdentityRange) *NullableIdentityRange {
	return &NullableIdentityRange{value: val, isSet: true}
}

func (v NullableIdentityRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
