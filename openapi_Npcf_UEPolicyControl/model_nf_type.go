/*
Npcf_UEPolicyControl

UE Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_UEPolicyControl

import (
	"encoding/json"
	"fmt"
)

// NFType NF types known to NRF
type NFType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NFType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NFType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NFType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNFType struct {
	value *NFType
	isSet bool
}

func (v NullableNFType) Get() *NFType {
	return v.value
}

func (v *NullableNFType) Set(val *NFType) {
	v.value = val
	v.isSet = true
}

func (v NullableNFType) IsSet() bool {
	return v.isSet
}

func (v *NullableNFType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFType(val *NFType) *NullableNFType {
	return &NullableNFType{value: val, isSet: true}
}

func (v NullableNFType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
