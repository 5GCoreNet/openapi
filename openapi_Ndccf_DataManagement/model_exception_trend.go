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

// ExceptionTrend Possible values are: - UP: Up trend of the exception level. - DOWN: Down trend of the exception level. - UNKNOW: Unknown trend of the exception level. - STABLE: Stable trend of the exception level. 
type ExceptionTrend struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ExceptionTrend) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ExceptionTrend)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ExceptionTrend) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableExceptionTrend struct {
	value *ExceptionTrend
	isSet bool
}

func (v NullableExceptionTrend) Get() *ExceptionTrend {
	return v.value
}

func (v *NullableExceptionTrend) Set(val *ExceptionTrend) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionTrend) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionTrend) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionTrend(val *ExceptionTrend) *NullableExceptionTrend {
	return &NullableExceptionTrend{value: val, isSet: true}
}

func (v NullableExceptionTrend) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionTrend) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


