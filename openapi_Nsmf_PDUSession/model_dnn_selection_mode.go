/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// DnnSelectionMode DNN Selection Mode. Possible values are - VERIFIED - UE_DNN_NOT_VERIFIED - NW_DNN_NOT_VERIFIED
type DnnSelectionMode struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnnSelectionMode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnnSelectionMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnnSelectionMode) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDnnSelectionMode struct {
	value *DnnSelectionMode
	isSet bool
}

func (v NullableDnnSelectionMode) Get() *DnnSelectionMode {
	return v.value
}

func (v *NullableDnnSelectionMode) Set(val *DnnSelectionMode) {
	v.value = val
	v.isSet = true
}

func (v NullableDnnSelectionMode) IsSet() bool {
	return v.isSet
}

func (v *NullableDnnSelectionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDnnSelectionMode(val *DnnSelectionMode) *NullableDnnSelectionMode {
	return &NullableDnnSelectionMode{value: val, isSet: true}
}

func (v NullableDnnSelectionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDnnSelectionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
