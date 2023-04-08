/*
CAPIF_Discover_Service_API

API for discovering service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Discover_Service_API

import (
	"encoding/json"
	"fmt"
)

// AefProfile - Represents the AEF profile data.
type AefProfile struct {
	Interface *interface{}
}

// interface{}AsAefProfile is a convenience function that returns interface{} wrapped in AefProfile
func InterfaceAsAefProfile(v *interface{}) AefProfile {
	return AefProfile{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AefProfile) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface
	err = newStrictDecoder(data).Decode(&dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			match++
		}
	} else {
		dst.Interface = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AefProfile)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AefProfile)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AefProfile) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AefProfile) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableAefProfile struct {
	value *AefProfile
	isSet bool
}

func (v NullableAefProfile) Get() *AefProfile {
	return v.value
}

func (v *NullableAefProfile) Set(val *AefProfile) {
	v.value = val
	v.isSet = true
}

func (v NullableAefProfile) IsSet() bool {
	return v.isSet
}

func (v *NullableAefProfile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAefProfile(val *AefProfile) *NullableAefProfile {
	return &NullableAefProfile{value: val, isSet: true}
}

func (v NullableAefProfile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAefProfile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


