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

// ConfigResult - Represents one configuration processing result for a group's members.
type ConfigResult struct {
	Interface *interface{}
}

// interface{}AsConfigResult is a convenience function that returns interface{} wrapped in ConfigResult
func InterfaceAsConfigResult(v *interface{}) ConfigResult {
	return ConfigResult{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ConfigResult) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ConfigResult)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ConfigResult)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ConfigResult) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ConfigResult) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableConfigResult struct {
	value *ConfigResult
	isSet bool
}

func (v NullableConfigResult) Get() *ConfigResult {
	return v.value
}

func (v *NullableConfigResult) Set(val *ConfigResult) {
	v.value = val
	v.isSet = true
}

func (v NullableConfigResult) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigResult(val *ConfigResult) *NullableConfigResult {
	return &NullableConfigResult{value: val, isSet: true}
}

func (v NullableConfigResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
