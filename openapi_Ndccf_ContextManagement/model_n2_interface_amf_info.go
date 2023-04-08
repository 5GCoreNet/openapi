/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// N2InterfaceAmfInfo AMF N2 interface information
type N2InterfaceAmfInfo struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N2InterfaceAmfInfo) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(N2InterfaceAmfInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N2InterfaceAmfInfo) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN2InterfaceAmfInfo struct {
	value *N2InterfaceAmfInfo
	isSet bool
}

func (v NullableN2InterfaceAmfInfo) Get() *N2InterfaceAmfInfo {
	return v.value
}

func (v *NullableN2InterfaceAmfInfo) Set(val *N2InterfaceAmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InterfaceAmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InterfaceAmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InterfaceAmfInfo(val *N2InterfaceAmfInfo) *NullableN2InterfaceAmfInfo {
	return &NullableN2InterfaceAmfInfo{value: val, isSet: true}
}

func (v NullableN2InterfaceAmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InterfaceAmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


