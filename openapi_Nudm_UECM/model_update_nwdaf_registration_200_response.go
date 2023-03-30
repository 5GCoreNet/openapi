/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UECM

import (
	"encoding/json"
	"fmt"
)

// UpdateNwdafRegistration200Response - struct for UpdateNwdafRegistration200Response
type UpdateNwdafRegistration200Response struct {
	NwdafRegistration *NwdafRegistration
	PatchResult *PatchResult
}

// NwdafRegistrationAsUpdateNwdafRegistration200Response is a convenience function that returns NwdafRegistration wrapped in UpdateNwdafRegistration200Response
func NwdafRegistrationAsUpdateNwdafRegistration200Response(v *NwdafRegistration) UpdateNwdafRegistration200Response {
	return UpdateNwdafRegistration200Response{
		NwdafRegistration: v,
	}
}

// PatchResultAsUpdateNwdafRegistration200Response is a convenience function that returns PatchResult wrapped in UpdateNwdafRegistration200Response
func PatchResultAsUpdateNwdafRegistration200Response(v *PatchResult) UpdateNwdafRegistration200Response {
	return UpdateNwdafRegistration200Response{
		PatchResult: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *UpdateNwdafRegistration200Response) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into NwdafRegistration
	err = newStrictDecoder(data).Decode(&dst.NwdafRegistration)
	if err == nil {
		jsonNwdafRegistration, _ := json.Marshal(dst.NwdafRegistration)
		if string(jsonNwdafRegistration) == "{}" { // empty struct
			dst.NwdafRegistration = nil
		} else {
			match++
		}
	} else {
		dst.NwdafRegistration = nil
	}

	// try to unmarshal data into PatchResult
	err = newStrictDecoder(data).Decode(&dst.PatchResult)
	if err == nil {
		jsonPatchResult, _ := json.Marshal(dst.PatchResult)
		if string(jsonPatchResult) == "{}" { // empty struct
			dst.PatchResult = nil
		} else {
			match++
		}
	} else {
		dst.PatchResult = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.NwdafRegistration = nil
		dst.PatchResult = nil

		return fmt.Errorf("data matches more than one schema in oneOf(UpdateNwdafRegistration200Response)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(UpdateNwdafRegistration200Response)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src UpdateNwdafRegistration200Response) MarshalJSON() ([]byte, error) {
	if src.NwdafRegistration != nil {
		return json.Marshal(&src.NwdafRegistration)
	}

	if src.PatchResult != nil {
		return json.Marshal(&src.PatchResult)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *UpdateNwdafRegistration200Response) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.NwdafRegistration != nil {
		return obj.NwdafRegistration
	}

	if obj.PatchResult != nil {
		return obj.PatchResult
	}

	// all schemas are nil
	return nil
}

type NullableUpdateNwdafRegistration200Response struct {
	value *UpdateNwdafRegistration200Response
	isSet bool
}

func (v NullableUpdateNwdafRegistration200Response) Get() *UpdateNwdafRegistration200Response {
	return v.value
}

func (v *NullableUpdateNwdafRegistration200Response) Set(val *UpdateNwdafRegistration200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateNwdafRegistration200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateNwdafRegistration200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateNwdafRegistration200Response(val *UpdateNwdafRegistration200Response) *NullableUpdateNwdafRegistration200Response {
	return &NullableUpdateNwdafRegistration200Response{value: val, isSet: true}
}

func (v NullableUpdateNwdafRegistration200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateNwdafRegistration200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

