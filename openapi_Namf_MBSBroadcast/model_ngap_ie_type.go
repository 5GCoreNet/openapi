/*
Namf_MBSBroadcast

AMF MBSBroadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSBroadcast

import (
	"encoding/json"
	"fmt"
)

// NgapIeType NGAP Information Element Type
type NgapIeType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NgapIeType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NgapIeType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NgapIeType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNgapIeType struct {
	value *NgapIeType
	isSet bool
}

func (v NullableNgapIeType) Get() *NgapIeType {
	return v.value
}

func (v *NullableNgapIeType) Set(val *NgapIeType) {
	v.value = val
	v.isSet = true
}

func (v NullableNgapIeType) IsSet() bool {
	return v.isSet
}

func (v *NullableNgapIeType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgapIeType(val *NgapIeType) *NullableNgapIeType {
	return &NullableNgapIeType{value: val, isSet: true}
}

func (v NullableNgapIeType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgapIeType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

