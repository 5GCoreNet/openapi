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

// IpIndex Represents the IP Index to be sent from UDM to the SMF (its value can be either an integer or a string)
type IpIndex struct {
	int32 *int32
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *IpIndex) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into int32
	err = json.Unmarshal(data, &dst.int32);
	if err == nil {
		jsonint32, _ := json.Marshal(dst.int32)
		if string(jsonint32) == "{}" { // empty struct
			dst.int32 = nil
		} else {
			return nil // data stored in dst.int32, return on the first match
		}
	} else {
		dst.int32 = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(IpIndex)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *IpIndex) MarshalJSON() ([]byte, error) {
	if src.int32 != nil {
		return json.Marshal(&src.int32)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableIpIndex struct {
	value *IpIndex
	isSet bool
}

func (v NullableIpIndex) Get() *IpIndex {
	return v.value
}

func (v *NullableIpIndex) Set(val *IpIndex) {
	v.value = val
	v.isSet = true
}

func (v NullableIpIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableIpIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpIndex(val *IpIndex) *NullableIpIndex {
	return &NullableIpIndex{value: val, isSet: true}
}

func (v NullableIpIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


