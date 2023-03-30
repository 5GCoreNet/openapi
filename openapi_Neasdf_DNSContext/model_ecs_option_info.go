/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_DNSContext

import (
	"encoding/json"
	"fmt"
)

// EcsOptionInfo ECS Option Information
type EcsOptionInfo struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *EcsOptionInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.interface{});
	if err == nil {
		jsoninterface{}, _ := json.Marshal(dst.interface{})
		if string(jsoninterface{}) == "{}" { // empty struct
			dst.interface{} = nil
		} else {
			return nil // data stored in dst.interface{}, return on the first match
		}
	} else {
		dst.interface{} = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(EcsOptionInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *EcsOptionInfo) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEcsOptionInfo struct {
	value *EcsOptionInfo
	isSet bool
}

func (v NullableEcsOptionInfo) Get() *EcsOptionInfo {
	return v.value
}

func (v *NullableEcsOptionInfo) Set(val *EcsOptionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableEcsOptionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableEcsOptionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcsOptionInfo(val *EcsOptionInfo) *NullableEcsOptionInfo {
	return &NullableEcsOptionInfo{value: val, isSet: true}
}

func (v NullableEcsOptionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcsOptionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


