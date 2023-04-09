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

// DnsContextCreateData Data within Create request
type DnsContextCreateData struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnsContextCreateData) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnsContextCreateData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnsContextCreateData) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnsContextCreateData struct {
	value *DnsContextCreateData
	isSet bool
}

func (v NullableDnsContextCreateData) Get() *DnsContextCreateData {
	return v.value
}

func (v *NullableDnsContextCreateData) Set(val *DnsContextCreateData) {
	v.value = val
	v.isSet = true
}

func (v NullableDnsContextCreateData) IsSet() bool {
	return v.isSet
}

func (v *NullableDnsContextCreateData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnsContextCreateData(val *DnsContextCreateData) *NullableDnsContextCreateData {
	return &NullableDnsContextCreateData{value: val, isSet: true}
}

func (v NullableDnsContextCreateData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnsContextCreateData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
