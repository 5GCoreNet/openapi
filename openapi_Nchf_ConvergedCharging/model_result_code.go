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

// ResultCode struct for ResultCode
type ResultCode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ResultCode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ResultCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ResultCode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableResultCode struct {
	value *ResultCode
	isSet bool
}

func (v NullableResultCode) Get() *ResultCode {
	return v.value
}

func (v *NullableResultCode) Set(val *ResultCode) {
	v.value = val
	v.isSet = true
}

func (v NullableResultCode) IsSet() bool {
	return v.isSet
}

func (v *NullableResultCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResultCode(val *ResultCode) *NullableResultCode {
	return &NullableResultCode{value: val, isSet: true}
}

func (v NullableResultCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResultCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
