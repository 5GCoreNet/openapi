/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// Model5GsUserState Describes the 5GS User State of a UE
type Model5GsUserState struct {
	5GsUserStateAnyOf *5GsUserStateAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Model5GsUserState) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into 5GsUserStateAnyOf
	err = json.Unmarshal(data, &dst.5GsUserStateAnyOf);
	if err == nil {
		json5GsUserStateAnyOf, _ := json.Marshal(dst.5GsUserStateAnyOf)
		if string(json5GsUserStateAnyOf) == "{}" { // empty struct
			dst.5GsUserStateAnyOf = nil
		} else {
			return nil // data stored in dst.5GsUserStateAnyOf, return on the first match
		}
	} else {
		dst.5GsUserStateAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Model5GsUserState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Model5GsUserState) MarshalJSON() ([]byte, error) {
	if src.5GsUserStateAnyOf != nil {
		return json.Marshal(&src.5GsUserStateAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableModel5GsUserState struct {
	value *Model5GsUserState
	isSet bool
}

func (v NullableModel5GsUserState) Get() *Model5GsUserState {
	return v.value
}

func (v *NullableModel5GsUserState) Set(val *Model5GsUserState) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GsUserState) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GsUserState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GsUserState(val *Model5GsUserState) *NullableModel5GsUserState {
	return &NullableModel5GsUserState{value: val, isSet: true}
}

func (v NullableModel5GsUserState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GsUserState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


