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

// DnsContextCreatedData Data within Create response
type DnsContextCreatedData struct {
	interface{} *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnsContextCreatedData) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnsContextCreatedData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnsContextCreatedData) MarshalJSON() ([]byte, error) {
	if src.interface{} != nil {
		return json.Marshal(&src.interface{})
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnsContextCreatedData struct {
	value *DnsContextCreatedData
	isSet bool
}

func (v NullableDnsContextCreatedData) Get() *DnsContextCreatedData {
	return v.value
}

func (v *NullableDnsContextCreatedData) Set(val *DnsContextCreatedData) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsContextCreatedData) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsContextCreatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsContextCreatedData(val *DnsContextCreatedData) *NullableDnsContextCreatedData {
	return &NullableDnsContextCreatedData{value: val, isSet: true}
}

func (v NullableDnsContextCreatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsContextCreatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


