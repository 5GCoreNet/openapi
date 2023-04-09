/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// IpSmGwRegistration struct for IpSmGwRegistration
type IpSmGwRegistration struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IpSmGwRegistration) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(IpSmGwRegistration)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IpSmGwRegistration) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
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
