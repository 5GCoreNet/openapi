/*
Nudsf_Timer

Nudsf Timer Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_Timer

import (
	"encoding/json"
	"fmt"
)

// PatchOperation Operations as defined in IETF RFC 6902.
type PatchOperation struct {
	PatchOperationAnyOf *PatchOperationAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PatchOperation) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PatchOperationAnyOf
	err = json.Unmarshal(data, &dst.PatchOperationAnyOf);
	if err == nil {
		jsonPatchOperationAnyOf, _ := json.Marshal(dst.PatchOperationAnyOf)
		if string(jsonPatchOperationAnyOf) == "{}" { // empty struct
			dst.PatchOperationAnyOf = nil
		} else {
			return nil // data stored in dst.PatchOperationAnyOf, return on the first match
		}
	} else {
		dst.PatchOperationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PatchOperation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PatchOperation) MarshalJSON() ([]byte, error) {
	if src.PatchOperationAnyOf != nil {
		return json.Marshal(&src.PatchOperationAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePatchOperation struct {
	value *PatchOperation
	isSet bool
}

func (v NullablePatchOperation) Get() *PatchOperation {
	return v.value
}

func (v *NullablePatchOperation) Set(val *PatchOperation) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchOperation) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchOperation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchOperation(val *PatchOperation) *NullablePatchOperation {
	return &NullablePatchOperation{value: val, isSet: true}
}

func (v NullablePatchOperation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchOperation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


