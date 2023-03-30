/*
NhssUEAU

HSS UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_UEAU

import (
	"encoding/json"
	"fmt"
)

// AvGenerationResponse - It represents the response body in the AV response sent by HSS to UDM, containing the 5G-AKA or EAP-AKA-prime authentication vector
type AvGenerationResponse struct {
	Interface{} *interface{}
}

// interface{}AsAvGenerationResponse is a convenience function that returns interface{} wrapped in AvGenerationResponse
func Interface{}AsAvGenerationResponse(v *interface{}) AvGenerationResponse {
	return AvGenerationResponse{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AvGenerationResponse) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AvGenerationResponse)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AvGenerationResponse)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AvGenerationResponse) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AvGenerationResponse) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableAvGenerationResponse struct {
	value *AvGenerationResponse
	isSet bool
}

func (v NullableAvGenerationResponse) Get() *AvGenerationResponse {
	return v.value
}

func (v *NullableAvGenerationResponse) Set(val *AvGenerationResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAvGenerationResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAvGenerationResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAvGenerationResponse(val *AvGenerationResponse) *NullableAvGenerationResponse {
	return &NullableAvGenerationResponse{value: val, isSet: true}
}

func (v NullableAvGenerationResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAvGenerationResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


