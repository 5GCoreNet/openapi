/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
	"fmt"
)

// TrafficDescriptorComponents Traffic descriptor components for the requested URSP.
type TrafficDescriptorComponents struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TrafficDescriptorComponents) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface);
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			return nil // data stored in dst.Interface, return on the first match
		}
	} else {
		dst.Interface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TrafficDescriptorComponents)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TrafficDescriptorComponents) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTrafficDescriptorComponents struct {
	value *TrafficDescriptorComponents
	isSet bool
}

func (v NullableTrafficDescriptorComponents) Get() *TrafficDescriptorComponents {
	return v.value
}

func (v *NullableTrafficDescriptorComponents) Set(val *TrafficDescriptorComponents) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficDescriptorComponents) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficDescriptorComponents) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficDescriptorComponents(val *TrafficDescriptorComponents) *NullableTrafficDescriptorComponents {
	return &NullableTrafficDescriptorComponents{value: val, isSet: true}
}

func (v NullableTrafficDescriptorComponents) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficDescriptorComponents) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


