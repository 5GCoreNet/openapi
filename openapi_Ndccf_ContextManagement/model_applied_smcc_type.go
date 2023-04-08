/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// AppliedSmccType Possible values are: - DNN_CC: Indicates the DNN based congestion control. - SNSSAI_CC: Indicates the S-NSSAI based congestion control. 
type AppliedSmccType struct {
	AppliedSmccTypeAnyOf *AppliedSmccTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AppliedSmccType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AppliedSmccTypeAnyOf
	err = json.Unmarshal(data, &dst.AppliedSmccTypeAnyOf);
	if err == nil {
		jsonAppliedSmccTypeAnyOf, _ := json.Marshal(dst.AppliedSmccTypeAnyOf)
		if string(jsonAppliedSmccTypeAnyOf) == "{}" { // empty struct
			dst.AppliedSmccTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AppliedSmccTypeAnyOf, return on the first match
		}
	} else {
		dst.AppliedSmccTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AppliedSmccType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AppliedSmccType) MarshalJSON() ([]byte, error) {
	if src.AppliedSmccTypeAnyOf != nil {
		return json.Marshal(&src.AppliedSmccTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAppliedSmccType struct {
	value *AppliedSmccType
	isSet bool
}

func (v NullableAppliedSmccType) Get() *AppliedSmccType {
	return v.value
}

func (v *NullableAppliedSmccType) Set(val *AppliedSmccType) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedSmccType) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedSmccType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedSmccType(val *AppliedSmccType) *NullableAppliedSmccType {
	return &NullableAppliedSmccType{value: val, isSet: true}
}

func (v NullableAppliedSmccType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedSmccType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


