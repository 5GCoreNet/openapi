/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UEAU

import (
	"encoding/json"
	"fmt"
)

// HssAvType struct for HssAvType
type HssAvType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *HssAvType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(HssAvType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *HssAvType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableHssAvType struct {
	value *HssAvType
	isSet bool
}

func (v NullableHssAvType) Get() *HssAvType {
	return v.value
}

func (v *NullableHssAvType) Set(val *HssAvType) {
	v.value = val
	v.isSet = true
}

func (v NullableHssAvType) IsSet() bool {
	return v.isSet
}

func (v *NullableHssAvType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHssAvType(val *HssAvType) *NullableHssAvType {
	return &NullableHssAvType{value: val, isSet: true}
}

func (v NullableHssAvType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHssAvType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


