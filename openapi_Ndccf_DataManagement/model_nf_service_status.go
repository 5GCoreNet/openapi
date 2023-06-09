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

// NFServiceStatus Status of a given NF Service Instance of an NF Instance stored in NRF
type NFServiceStatus struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NFServiceStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NFServiceStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NFServiceStatus) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNFServiceStatus struct {
	value *NFServiceStatus
	isSet bool
}

func (v NullableNFServiceStatus) Get() *NFServiceStatus {
	return v.value
}

func (v *NullableNFServiceStatus) Set(val *NFServiceStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNFServiceStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNFServiceStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFServiceStatus(val *NFServiceStatus) *NullableNFServiceStatus {
	return &NullableNFServiceStatus{value: val, isSet: true}
}

func (v NullableNFServiceStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFServiceStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
