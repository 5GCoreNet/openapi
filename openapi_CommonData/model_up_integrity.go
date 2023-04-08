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

// UpIntegrity indicates whether UP integrity protection is required, preferred or not needed for all the traffic on the PDU Session. It shall comply with the provisions defined in  table 5.4.3.4-1.  
type UpIntegrity struct {
	UpIntegrityAnyOf *UpIntegrityAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *UpIntegrity) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into UpIntegrityAnyOf
	err = json.Unmarshal(data, &dst.UpIntegrityAnyOf);
	if err == nil {
		jsonUpIntegrityAnyOf, _ := json.Marshal(dst.UpIntegrityAnyOf)
		if string(jsonUpIntegrityAnyOf) == "{}" { // empty struct
			dst.UpIntegrityAnyOf = nil
		} else {
			return nil // data stored in dst.UpIntegrityAnyOf, return on the first match
		}
	} else {
		dst.UpIntegrityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(UpIntegrity)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *UpIntegrity) MarshalJSON() ([]byte, error) {
	if src.UpIntegrityAnyOf != nil {
		return json.Marshal(&src.UpIntegrityAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableUpIntegrity struct {
	value *UpIntegrity
	isSet bool
}

func (v NullableUpIntegrity) Get() *UpIntegrity {
	return v.value
}

func (v *NullableUpIntegrity) Set(val *UpIntegrity) {
	v.value = val
	v.isSet = true
}

func (v NullableUpIntegrity) IsSet() bool {
	return v.isSet
}

func (v *NullableUpIntegrity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpIntegrity(val *UpIntegrity) *NullableUpIntegrity {
	return &NullableUpIntegrity{value: val, isSet: true}
}

func (v NullableUpIntegrity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpIntegrity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


