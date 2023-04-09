/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// N2InformationClass Enumeration for N2 Information Class
type N2InformationClass struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N2InformationClass) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(N2InformationClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N2InformationClass) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN2InformationClass struct {
	value *N2InformationClass
	isSet bool
}

func (v NullableN2InformationClass) Get() *N2InformationClass {
	return v.value
}

func (v *NullableN2InformationClass) Set(val *N2InformationClass) {
	v.value = val
	v.isSet = true
}

func (v NullableN2InformationClass) IsSet() bool {
	return v.isSet
}

func (v *NullableN2InformationClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN2InformationClass(val *N2InformationClass) *NullableN2InformationClass {
	return &NullableN2InformationClass{value: val, isSet: true}
}

func (v NullableN2InformationClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN2InformationClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
