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

// LcsMoServiceClass struct for LcsMoServiceClass
type LcsMoServiceClass struct {
	LcsMoServiceClassAnyOf *LcsMoServiceClassAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LcsMoServiceClass) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LcsMoServiceClassAnyOf
	err = json.Unmarshal(data, &dst.LcsMoServiceClassAnyOf);
	if err == nil {
		jsonLcsMoServiceClassAnyOf, _ := json.Marshal(dst.LcsMoServiceClassAnyOf)
		if string(jsonLcsMoServiceClassAnyOf) == "{}" { // empty struct
			dst.LcsMoServiceClassAnyOf = nil
		} else {
			return nil // data stored in dst.LcsMoServiceClassAnyOf, return on the first match
		}
	} else {
		dst.LcsMoServiceClassAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
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

	return fmt.Errorf("data failed to match schemas in anyOf(LcsMoServiceClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LcsMoServiceClass) MarshalJSON() ([]byte, error) {
	if src.LcsMoServiceClassAnyOf != nil {
		return json.Marshal(&src.LcsMoServiceClassAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLcsMoServiceClass struct {
	value *LcsMoServiceClass
	isSet bool
}

func (v NullableLcsMoServiceClass) Get() *LcsMoServiceClass {
	return v.value
}

func (v *NullableLcsMoServiceClass) Set(val *LcsMoServiceClass) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsMoServiceClass) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsMoServiceClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsMoServiceClass(val *LcsMoServiceClass) *NullableLcsMoServiceClass {
	return &NullableLcsMoServiceClass{value: val, isSet: true}
}

func (v NullableLcsMoServiceClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsMoServiceClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


