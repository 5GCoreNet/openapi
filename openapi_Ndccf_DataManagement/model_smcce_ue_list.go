/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// SmcceUeList Represents the List of UEs classified based on experience level of Session Management  congestion control.
type SmcceUeList struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SmcceUeList) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(SmcceUeList)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SmcceUeList) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSmcceUeList struct {
	value *SmcceUeList
	isSet bool
}

func (v NullableSmcceUeList) Get() *SmcceUeList {
	return v.value
}

func (v *NullableSmcceUeList) Set(val *SmcceUeList) {
	v.value = val
	v.isSet = true
}

func (v NullableSmcceUeList) IsSet() bool {
	return v.isSet
}

func (v *NullableSmcceUeList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSmcceUeList(val *SmcceUeList) *NullableSmcceUeList {
	return &NullableSmcceUeList{value: val, isSet: true}
}

func (v NullableSmcceUeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSmcceUeList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
