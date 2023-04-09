/*
MSGG_L3GDelivery

API for MSGG L3G Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSGG_L3GDelivery

import (
	"encoding/json"
	"fmt"
)

// AddressType Possible values are: - UE: The address type is UE. - AS: The address type is AS. - GROUP: The address type is GROUP. - BC: The address type is BC. - TOPIC: The address type is TOPIC.
type AddressType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AddressType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AddressType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AddressType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAddressType struct {
	value *AddressType
	isSet bool
}

func (v NullableAddressType) Get() *AddressType {
	return v.value
}

func (v *NullableAddressType) Set(val *AddressType) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressType) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressType(val *AddressType) *NullableAddressType {
	return &NullableAddressType{value: val, isSet: true}
}

func (v NullableAddressType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
