/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// MbsSessionId MBS Session Identifier
type MbsSessionId struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsSessionId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface);
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			return nil // data stored in dst.Interface, return on the first match
		}
	} else {
		dst.Interface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MbsSessionId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsSessionId) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsSessionId struct {
	value *MbsSessionId
	isSet bool
}

func (v NullableMbsSessionId) Get() *MbsSessionId {
	return v.value
}

func (v *NullableMbsSessionId) Set(val *MbsSessionId) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSessionId) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSessionId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSessionId(val *MbsSessionId) *NullableMbsSessionId {
	return &NullableMbsSessionId{value: val, isSet: true}
}

func (v NullableMbsSessionId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSessionId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


