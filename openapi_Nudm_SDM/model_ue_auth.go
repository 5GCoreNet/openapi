/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// UeAuth Possible values are: - AUTHORIZED: Indicates that the UE is authorized. - NOT_AUTHORIZED: Indicates that the UE is not authorized. 
type UeAuth struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UeAuth) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(UeAuth)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UeAuth) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUeAuth struct {
	value *UeAuth
	isSet bool
}

func (v NullableUeAuth) Get() *UeAuth {
	return v.value
}

func (v *NullableUeAuth) Set(val *UeAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableUeAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableUeAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeAuth(val *UeAuth) *NullableUeAuth {
	return &NullableUeAuth{value: val, isSet: true}
}

func (v NullableUeAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

