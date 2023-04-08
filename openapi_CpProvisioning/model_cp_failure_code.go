/*
3gpp-cp-parameter-provisioning

API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CpProvisioning

import (
	"encoding/json"
	"fmt"
)

// CpFailureCode Possible values are - MALFUNCTION: This value indicates that something functions wrongly in CP parameter provisioning or the CP parameter provisioning does not function at all. - SET_ID_DUPLICATED: The received CP set identifier(s) are already provisioned. - OTHER_REASON: Other reason unspecified. 
type CpFailureCode struct {
	CpFailureCodeAnyOf *CpFailureCodeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CpFailureCode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CpFailureCodeAnyOf
	err = json.Unmarshal(data, &dst.CpFailureCodeAnyOf);
	if err == nil {
		jsonCpFailureCodeAnyOf, _ := json.Marshal(dst.CpFailureCodeAnyOf)
		if string(jsonCpFailureCodeAnyOf) == "{}" { // empty struct
			dst.CpFailureCodeAnyOf = nil
		} else {
			return nil // data stored in dst.CpFailureCodeAnyOf, return on the first match
		}
	} else {
		dst.CpFailureCodeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CpFailureCode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CpFailureCode) MarshalJSON() ([]byte, error) {
	if src.CpFailureCodeAnyOf != nil {
		return json.Marshal(&src.CpFailureCodeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCpFailureCode struct {
	value *CpFailureCode
	isSet bool
}

func (v NullableCpFailureCode) Get() *CpFailureCode {
	return v.value
}

func (v *NullableCpFailureCode) Set(val *CpFailureCode) {
	v.value = val
	v.isSet = true
}

func (v NullableCpFailureCode) IsSet() bool {
	return v.isSet
}

func (v *NullableCpFailureCode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpFailureCode(val *CpFailureCode) *NullableCpFailureCode {
	return &NullableCpFailureCode{value: val, isSet: true}
}

func (v NullableCpFailureCode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpFailureCode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


