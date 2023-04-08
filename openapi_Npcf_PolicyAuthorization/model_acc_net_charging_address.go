/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AccNetChargingAddress Describes the network entity within the access network performing charging
type AccNetChargingAddress struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccNetChargingAddress) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AccNetChargingAddress)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccNetChargingAddress) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccNetChargingAddress struct {
	value *AccNetChargingAddress
	isSet bool
}

func (v NullableAccNetChargingAddress) Get() *AccNetChargingAddress {
	return v.value
}

func (v *NullableAccNetChargingAddress) Set(val *AccNetChargingAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableAccNetChargingAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableAccNetChargingAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccNetChargingAddress(val *AccNetChargingAddress) *NullableAccNetChargingAddress {
	return &NullableAccNetChargingAddress{value: val, isSet: true}
}

func (v NullableAccNetChargingAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccNetChargingAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


