/*
3gpp-applying-bdt-policy

API for applying BDT policy   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ApplyingBdtPolicy

import (
	"encoding/json"
	"fmt"
)

// AppliedBdtPolicy - Represents an applied BDT policy.
type AppliedBdtPolicy struct {
	Interface *interface{}
}

// interface{}AsAppliedBdtPolicy is a convenience function that returns interface{} wrapped in AppliedBdtPolicy
func InterfaceAsAppliedBdtPolicy(v *interface{}) AppliedBdtPolicy {
	return AppliedBdtPolicy{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppliedBdtPolicy) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(AppliedBdtPolicy)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppliedBdtPolicy)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppliedBdtPolicy) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppliedBdtPolicy) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableAppliedBdtPolicy struct {
	value *AppliedBdtPolicy
	isSet bool
}

func (v NullableAppliedBdtPolicy) Get() *AppliedBdtPolicy {
	return v.value
}

func (v *NullableAppliedBdtPolicy) Set(val *AppliedBdtPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableAppliedBdtPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableAppliedBdtPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppliedBdtPolicy(val *AppliedBdtPolicy) *NullableAppliedBdtPolicy {
	return &NullableAppliedBdtPolicy{value: val, isSet: true}
}

func (v NullableAppliedBdtPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppliedBdtPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
