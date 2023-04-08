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

// LocationPrivacyInd struct for LocationPrivacyInd
type LocationPrivacyInd struct {
	LocationPrivacyIndAnyOf *LocationPrivacyIndAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LocationPrivacyInd) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LocationPrivacyIndAnyOf
	err = json.Unmarshal(data, &dst.LocationPrivacyIndAnyOf);
	if err == nil {
		jsonLocationPrivacyIndAnyOf, _ := json.Marshal(dst.LocationPrivacyIndAnyOf)
		if string(jsonLocationPrivacyIndAnyOf) == "{}" { // empty struct
			dst.LocationPrivacyIndAnyOf = nil
		} else {
			return nil // data stored in dst.LocationPrivacyIndAnyOf, return on the first match
		}
	} else {
		dst.LocationPrivacyIndAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LocationPrivacyInd)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LocationPrivacyInd) MarshalJSON() ([]byte, error) {
	if src.LocationPrivacyIndAnyOf != nil {
		return json.Marshal(&src.LocationPrivacyIndAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocationPrivacyInd struct {
	value *LocationPrivacyInd
	isSet bool
}

func (v NullableLocationPrivacyInd) Get() *LocationPrivacyInd {
	return v.value
}

func (v *NullableLocationPrivacyInd) Set(val *LocationPrivacyInd) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationPrivacyInd) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationPrivacyInd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationPrivacyInd(val *LocationPrivacyInd) *NullableLocationPrivacyInd {
	return &NullableLocationPrivacyInd{value: val, isSet: true}
}

func (v NullableLocationPrivacyInd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationPrivacyInd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


