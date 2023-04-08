/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// CnType struct for CnType
type CnType struct {
	CnTypeAnyOf *CnTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CnType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CnTypeAnyOf
	err = json.Unmarshal(data, &dst.CnTypeAnyOf);
	if err == nil {
		jsonCnTypeAnyOf, _ := json.Marshal(dst.CnTypeAnyOf)
		if string(jsonCnTypeAnyOf) == "{}" { // empty struct
			dst.CnTypeAnyOf = nil
		} else {
			return nil // data stored in dst.CnTypeAnyOf, return on the first match
		}
	} else {
		dst.CnTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CnType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CnType) MarshalJSON() ([]byte, error) {
	if src.CnTypeAnyOf != nil {
		return json.Marshal(&src.CnTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCnType struct {
	value *CnType
	isSet bool
}

func (v NullableCnType) Get() *CnType {
	return v.value
}

func (v *NullableCnType) Set(val *CnType) {
	v.value = val
	v.isSet = true
}

func (v NullableCnType) IsSet() bool {
	return v.isSet
}

func (v *NullableCnType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCnType(val *CnType) *NullableCnType {
	return &NullableCnType{value: val, isSet: true}
}

func (v NullableCnType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCnType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


