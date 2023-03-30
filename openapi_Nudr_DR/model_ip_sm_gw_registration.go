/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// IpSmGwRegistration struct for IpSmGwRegistration
type IpSmGwRegistration struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IpSmGwRegistration) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(IpSmGwRegistration)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IpSmGwRegistration) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIpSmGwRegistration struct {
	value *IpSmGwRegistration
	isSet bool
}

func (v NullableIpSmGwRegistration) Get() *IpSmGwRegistration {
	return v.value
}

func (v *NullableIpSmGwRegistration) Set(val *IpSmGwRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableIpSmGwRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableIpSmGwRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpSmGwRegistration(val *IpSmGwRegistration) *NullableIpSmGwRegistration {
	return &NullableIpSmGwRegistration{value: val, isSet: true}
}

func (v NullableIpSmGwRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpSmGwRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


