/*
AUSF API

AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UEAuthentication

import (
	"encoding/json"
	"fmt"
)

// ServerAddressingInfo Contains addressing information (IP addresses and/or FQDNs) of a server.
type ServerAddressingInfo struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServerAddressingInfo) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServerAddressingInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServerAddressingInfo) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServerAddressingInfo struct {
	value *ServerAddressingInfo
	isSet bool
}

func (v NullableServerAddressingInfo) Get() *ServerAddressingInfo {
	return v.value
}

func (v *NullableServerAddressingInfo) Set(val *ServerAddressingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableServerAddressingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableServerAddressingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServerAddressingInfo(val *ServerAddressingInfo) *NullableServerAddressingInfo {
	return &NullableServerAddressingInfo{value: val, isSet: true}
}

func (v NullableServerAddressingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServerAddressingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


