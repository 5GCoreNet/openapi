/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbstf_DistSession

import (
	"encoding/json"
	"fmt"
)

// TunnelAddress Tunnel address
type TunnelAddress struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TunnelAddress) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TunnelAddress)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TunnelAddress) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTunnelAddress struct {
	value *TunnelAddress
	isSet bool
}

func (v NullableTunnelAddress) Get() *TunnelAddress {
	return v.value
}

func (v *NullableTunnelAddress) Set(val *TunnelAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableTunnelAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableTunnelAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTunnelAddress(val *TunnelAddress) *NullableTunnelAddress {
	return &NullableTunnelAddress{value: val, isSet: true}
}

func (v NullableTunnelAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTunnelAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


