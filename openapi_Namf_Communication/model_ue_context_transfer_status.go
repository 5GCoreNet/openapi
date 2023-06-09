/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
	"fmt"
)

// UeContextTransferStatus Describes the status of an individual ueContext resource in UE Context Transfer procedures
type UeContextTransferStatus struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeContextTransferStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeContextTransferStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeContextTransferStatus) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeContextTransferStatus struct {
	value *UeContextTransferStatus
	isSet bool
}

func (v NullableUeContextTransferStatus) Get() *UeContextTransferStatus {
	return v.value
}

func (v *NullableUeContextTransferStatus) Set(val *UeContextTransferStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextTransferStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextTransferStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextTransferStatus(val *UeContextTransferStatus) *NullableUeContextTransferStatus {
	return &NullableUeContextTransferStatus{value: val, isSet: true}
}

func (v NullableUeContextTransferStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextTransferStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
