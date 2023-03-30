/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// AcsInfoRm This data type is defined in the same way as the 'AcsInfo' data type, but with the  OpenAPI 'nullable: true' property. 
type AcsInfoRm struct {
	AcsInfo *AcsInfo
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AcsInfoRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AcsInfo
	err = json.Unmarshal(data, &dst.AcsInfo);
	if err == nil {
		jsonAcsInfo, _ := json.Marshal(dst.AcsInfo)
		if string(jsonAcsInfo) == "{}" { // empty struct
			dst.AcsInfo = nil
		} else {
			return nil // data stored in dst.AcsInfo, return on the first match
		}
	} else {
		dst.AcsInfo = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
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

	return fmt.Errorf("data failed to match schemas in anyOf(AcsInfoRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AcsInfoRm) MarshalJSON() ([]byte, error) {
	if src.AcsInfo != nil {
		return json.Marshal(&src.AcsInfo)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAcsInfoRm struct {
	value *AcsInfoRm
	isSet bool
}

func (v NullableAcsInfoRm) Get() *AcsInfoRm {
	return v.value
}

func (v *NullableAcsInfoRm) Set(val *AcsInfoRm) {
	v.value = val
	v.isSet = true
}

func (v NullableAcsInfoRm) IsSet() bool {
	return v.isSet
}

func (v *NullableAcsInfoRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcsInfoRm(val *AcsInfoRm) *NullableAcsInfoRm {
	return &NullableAcsInfoRm{value: val, isSet: true}
}

func (v NullableAcsInfoRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcsInfoRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


