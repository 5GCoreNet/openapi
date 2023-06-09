/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// AmbrRm This data type is defined in the same way as the 'Ambr' data type, but with the OpenAPI 'nullable: true' property.\"
type AmbrRm struct {
	Ambr      *Ambr
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmbrRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into Ambr
	err = json.Unmarshal(data, &dst.Ambr)
	if err == nil {
		jsonAmbr, _ := json.Marshal(dst.Ambr)
		if string(jsonAmbr) == "{}" { // empty struct
			dst.Ambr = nil
		} else {
			return nil // data stored in dst.Ambr, return on the first match
		}
	} else {
		dst.Ambr = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue)
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AmbrRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmbrRm) MarshalJSON() ([]byte, error) {
	if src.Ambr != nil {
		return json.Marshal(&src.Ambr)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmbrRm struct {
	value *AmbrRm
	isSet bool
}

func (v NullableAmbrRm) Get() *AmbrRm {
	return v.value
}

func (v *NullableAmbrRm) Set(val *AmbrRm) {
	v.value = val
	v.isSet = true
}

func (v NullableAmbrRm) IsSet() bool {
	return v.isSet
}

func (v *NullableAmbrRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmbrRm(val *AmbrRm) *NullableAmbrRm {
	return &NullableAmbrRm{value: val, isSet: true}
}

func (v NullableAmbrRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmbrRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
