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

// LineType Possible values are: - DSL: Identifies a DSL line - PON: Identifies a PON line
type LineType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LineType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(LineType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LineType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLineType struct {
	value *LineType
	isSet bool
}

func (v NullableLineType) Get() *LineType {
	return v.value
}

func (v *NullableLineType) Set(val *LineType) {
	v.value = val
	v.isSet = true
}

func (v NullableLineType) IsSet() bool {
	return v.isSet
}

func (v *NullableLineType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLineType(val *LineType) *NullableLineType {
	return &NullableLineType{value: val, isSet: true}
}

func (v NullableLineType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLineType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
