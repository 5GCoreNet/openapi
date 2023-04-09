/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbstf_DistSession

import (
	"encoding/json"
	"fmt"
)

// PktIngestMethod Packet Ingest Method
type PktIngestMethod struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PktIngestMethod) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PktIngestMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PktIngestMethod) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePktIngestMethod struct {
	value *PktIngestMethod
	isSet bool
}

func (v NullablePktIngestMethod) Get() *PktIngestMethod {
	return v.value
}

func (v *NullablePktIngestMethod) Set(val *PktIngestMethod) {
	v.value = val
	v.isSet = true
}

func (v NullablePktIngestMethod) IsSet() bool {
	return v.isSet
}

func (v *NullablePktIngestMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePktIngestMethod(val *PktIngestMethod) *NullablePktIngestMethod {
	return &NullablePktIngestMethod{value: val, isSet: true}
}

func (v NullablePktIngestMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePktIngestMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
