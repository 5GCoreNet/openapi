/*
Ntsctsf_TimeSynchronization Service API

TSCTSF Time Synchronization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ntsctsf_TimeSynchronization

import (
	"encoding/json"
	"fmt"
)

// GmCapable Possible values are: - GPTP: gPTP grandmaster is supported. - PTP: PTP grandmaste is supported. 
type GmCapable struct {
	GmCapableAnyOf *GmCapableAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GmCapable) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into GmCapableAnyOf
	err = json.Unmarshal(data, &dst.GmCapableAnyOf);
	if err == nil {
		jsonGmCapableAnyOf, _ := json.Marshal(dst.GmCapableAnyOf)
		if string(jsonGmCapableAnyOf) == "{}" { // empty struct
			dst.GmCapableAnyOf = nil
		} else {
			return nil // data stored in dst.GmCapableAnyOf, return on the first match
		}
	} else {
		dst.GmCapableAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(GmCapable)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GmCapable) MarshalJSON() ([]byte, error) {
	if src.GmCapableAnyOf != nil {
		return json.Marshal(&src.GmCapableAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGmCapable struct {
	value *GmCapable
	isSet bool
}

func (v NullableGmCapable) Get() *GmCapable {
	return v.value
}

func (v *NullableGmCapable) Set(val *GmCapable) {
	v.value = val
	v.isSet = true
}

func (v NullableGmCapable) IsSet() bool {
	return v.isSet
}

func (v *NullableGmCapable) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGmCapable(val *GmCapable) *NullableGmCapable {
	return &NullableGmCapable{value: val, isSet: true}
}

func (v NullableGmCapable) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGmCapable) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


