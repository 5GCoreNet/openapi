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

// UcPurpose Indicates the purpose of the user consent.
type UcPurpose struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UcPurpose) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(UcPurpose)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UcPurpose) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUcPurpose struct {
	value *UcPurpose
	isSet bool
}

func (v NullableUcPurpose) Get() *UcPurpose {
	return v.value
}

func (v *NullableUcPurpose) Set(val *UcPurpose) {
	v.value = val
	v.isSet = true
}

func (v NullableUcPurpose) IsSet() bool {
	return v.isSet
}

func (v *NullableUcPurpose) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUcPurpose(val *UcPurpose) *NullableUcPurpose {
	return &NullableUcPurpose{value: val, isSet: true}
}

func (v NullableUcPurpose) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUcPurpose) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


