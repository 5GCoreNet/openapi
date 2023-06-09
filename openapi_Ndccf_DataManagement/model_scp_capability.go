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

// ScpCapability Indicates the capabilities supported by an SCP
type ScpCapability struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ScpCapability) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ScpCapability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ScpCapability) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableScpCapability struct {
	value *ScpCapability
	isSet bool
}

func (v NullableScpCapability) Get() *ScpCapability {
	return v.value
}

func (v *NullableScpCapability) Set(val *ScpCapability) {
	v.value = val
	v.isSet = true
}

func (v NullableScpCapability) IsSet() bool {
	return v.isSet
}

func (v *NullableScpCapability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScpCapability(val *ScpCapability) *NullableScpCapability {
	return &NullableScpCapability{value: val, isSet: true}
}

func (v NullableScpCapability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScpCapability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
