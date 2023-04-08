/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// DnnSelectionMode struct for DnnSelectionMode
type DnnSelectionMode struct {
	DnnSelectionModeAnyOf *DnnSelectionModeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DnnSelectionMode) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DnnSelectionModeAnyOf
	err = json.Unmarshal(data, &dst.DnnSelectionModeAnyOf);
	if err == nil {
		jsonDnnSelectionModeAnyOf, _ := json.Marshal(dst.DnnSelectionModeAnyOf)
		if string(jsonDnnSelectionModeAnyOf) == "{}" { // empty struct
			dst.DnnSelectionModeAnyOf = nil
		} else {
			return nil // data stored in dst.DnnSelectionModeAnyOf, return on the first match
		}
	} else {
		dst.DnnSelectionModeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(DnnSelectionMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DnnSelectionMode) MarshalJSON() ([]byte, error) {
	if src.DnnSelectionModeAnyOf != nil {
		return json.Marshal(&src.DnnSelectionModeAnyOf)
	}

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


