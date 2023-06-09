/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// AssociationType struct for AssociationType
type AssociationType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AssociationType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
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

	return fmt.Errorf("data failed to match schemas in anyOf(AssociationType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AssociationType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAssociationType struct {
	value *AssociationType
	isSet bool
}

func (v NullableAssociationType) Get() *AssociationType {
	return v.value
}

func (v *NullableAssociationType) Set(val *AssociationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationType(val *AssociationType) *NullableAssociationType {
	return &NullableAssociationType{value: val, isSet: true}
}

func (v NullableAssociationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
