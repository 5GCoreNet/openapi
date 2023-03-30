/*
Ngmlc_Location

GMLC Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ngmlc_Location

import (
	"encoding/json"
	"fmt"
)

// LcsQosClass Specifies LCS QoS class.
type LcsQosClass struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LcsQosClass) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(LcsQosClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LcsQosClass) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLcsQosClass struct {
	value *LcsQosClass
	isSet bool
}

func (v NullableLcsQosClass) Get() *LcsQosClass {
	return v.value
}

func (v *NullableLcsQosClass) Set(val *LcsQosClass) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsQosClass) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsQosClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsQosClass(val *LcsQosClass) *NullableLcsQosClass {
	return &NullableLcsQosClass{value: val, isSet: true}
}

func (v NullableLcsQosClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsQosClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


