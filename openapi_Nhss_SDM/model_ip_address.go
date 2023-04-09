/*
Nhss_SDM

HSS Subscriber Data Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_SDM

import (
	"encoding/json"
	"fmt"
)

// IpAddress - struct for IpAddress
type IpAddress struct {
	Interface *interface{}
}

// interface{}AsIpAddress is a convenience function that returns interface{} wrapped in IpAddress
func InterfaceAsIpAddress(v *interface{}) IpAddress {
	return IpAddress{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IpAddress) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(IpAddress)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IpAddress)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IpAddress) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IpAddress) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableIpAddress struct {
	value *IpAddress
	isSet bool
}

func (v NullableIpAddress) Get() *IpAddress {
	return v.value
}

func (v *NullableIpAddress) Set(val *IpAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableIpAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableIpAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpAddress(val *IpAddress) *NullableIpAddress {
	return &NullableIpAddress{value: val, isSet: true}
}

func (v NullableIpAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
