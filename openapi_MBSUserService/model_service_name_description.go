/*
3gpp-mbs-us

API for MBS User Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserService

import (
	"encoding/json"
	"fmt"
)

// ServiceNameDescription Represents a set of per language service names and/or service descriptions.
type ServiceNameDescription struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceNameDescription) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface)
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceNameDescription)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceNameDescription) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceNameDescription struct {
	value *ServiceNameDescription
	isSet bool
}

func (v NullableServiceNameDescription) Get() *ServiceNameDescription {
	return v.value
}

func (v *NullableServiceNameDescription) Set(val *ServiceNameDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceNameDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceNameDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceNameDescription(val *ServiceNameDescription) *NullableServiceNameDescription {
	return &NullableServiceNameDescription{value: val, isSet: true}
}

func (v NullableServiceNameDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceNameDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
