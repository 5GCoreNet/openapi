/*
3gpp-network-parameter-configuration

API for network parameter configuration.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NpConfiguration

import (
	"encoding/json"
	"fmt"
)

// NpConfiguration - Represents a network parameters configuration.
type NpConfiguration struct {
	Interface *interface{}
}

// interface{}AsNpConfiguration is a convenience function that returns interface{} wrapped in NpConfiguration
func InterfaceAsNpConfiguration(v *interface{}) NpConfiguration {
	return NpConfiguration{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *NpConfiguration) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(NpConfiguration)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(NpConfiguration)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src NpConfiguration) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *NpConfiguration) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableNpConfiguration struct {
	value *NpConfiguration
	isSet bool
}

func (v NullableNpConfiguration) Get() *NpConfiguration {
	return v.value
}

func (v *NullableNpConfiguration) Set(val *NpConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableNpConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableNpConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNpConfiguration(val *NpConfiguration) *NullableNpConfiguration {
	return &NullableNpConfiguration{value: val, isSet: true}
}

func (v NullableNpConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNpConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
