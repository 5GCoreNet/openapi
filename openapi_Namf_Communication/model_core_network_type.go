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

// CoreNetworkType It contains the Core Network type 5GC or EPC.
type CoreNetworkType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CoreNetworkType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CoreNetworkType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CoreNetworkType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCoreNetworkType struct {
	value *CoreNetworkType
	isSet bool
}

func (v NullableCoreNetworkType) Get() *CoreNetworkType {
	return v.value
}

func (v *NullableCoreNetworkType) Set(val *CoreNetworkType) {
	v.value = val
	v.isSet = true
}

func (v NullableCoreNetworkType) IsSet() bool {
	return v.isSet
}

func (v *NullableCoreNetworkType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCoreNetworkType(val *CoreNetworkType) *NullableCoreNetworkType {
	return &NullableCoreNetworkType{value: val, isSet: true}
}

func (v NullableCoreNetworkType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCoreNetworkType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

