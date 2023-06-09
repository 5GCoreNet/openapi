/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// VariablePartType struct for VariablePartType
type VariablePartType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *VariablePartType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(VariablePartType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *VariablePartType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableVariablePartType struct {
	value *VariablePartType
	isSet bool
}

func (v NullableVariablePartType) Get() *VariablePartType {
	return v.value
}

func (v *NullableVariablePartType) Set(val *VariablePartType) {
	v.value = val
	v.isSet = true
}

func (v NullableVariablePartType) IsSet() bool {
	return v.isSet
}

func (v *NullableVariablePartType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVariablePartType(val *VariablePartType) *NullableVariablePartType {
	return &NullableVariablePartType{value: val, isSet: true}
}

func (v NullableVariablePartType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVariablePartType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
