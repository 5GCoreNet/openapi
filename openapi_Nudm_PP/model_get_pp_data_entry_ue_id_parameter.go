/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_PP

import (
	"encoding/json"
	"fmt"
)

// GetPPDataEntryUeIdParameter struct for GetPPDataEntryUeIdParameter
type GetPPDataEntryUeIdParameter struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *GetPPDataEntryUeIdParameter) UnmarshalJSON(data []byte) error {
	var err error
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

	return fmt.Errorf("data failed to match schemas in anyOf(GetPPDataEntryUeIdParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *GetPPDataEntryUeIdParameter) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableGetPPDataEntryUeIdParameter struct {
	value *GetPPDataEntryUeIdParameter
	isSet bool
}

func (v NullableGetPPDataEntryUeIdParameter) Get() *GetPPDataEntryUeIdParameter {
	return v.value
}

func (v *NullableGetPPDataEntryUeIdParameter) Set(val *GetPPDataEntryUeIdParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableGetPPDataEntryUeIdParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableGetPPDataEntryUeIdParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetPPDataEntryUeIdParameter(val *GetPPDataEntryUeIdParameter) *NullableGetPPDataEntryUeIdParameter {
	return &NullableGetPPDataEntryUeIdParameter{value: val, isSet: true}
}

func (v NullableGetPPDataEntryUeIdParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetPPDataEntryUeIdParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


