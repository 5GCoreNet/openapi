/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_DNSContext

import (
	"encoding/json"
	"fmt"
)

// FqdnPatternMatchingRule - a matching rule for a FQDN pattern
type FqdnPatternMatchingRule struct {
	Interface *interface{}
}

// interface{}AsFqdnPatternMatchingRule is a convenience function that returns interface{} wrapped in FqdnPatternMatchingRule
func InterfaceAsFqdnPatternMatchingRule(v *interface{}) FqdnPatternMatchingRule {
	return FqdnPatternMatchingRule{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *FqdnPatternMatchingRule) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(FqdnPatternMatchingRule)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(FqdnPatternMatchingRule)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src FqdnPatternMatchingRule) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *FqdnPatternMatchingRule) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableFqdnPatternMatchingRule struct {
	value *FqdnPatternMatchingRule
	isSet bool
}

func (v NullableFqdnPatternMatchingRule) Get() *FqdnPatternMatchingRule {
	return v.value
}

func (v *NullableFqdnPatternMatchingRule) Set(val *FqdnPatternMatchingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableFqdnPatternMatchingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableFqdnPatternMatchingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFqdnPatternMatchingRule(val *FqdnPatternMatchingRule) *NullableFqdnPatternMatchingRule {
	return &NullableFqdnPatternMatchingRule{value: val, isSet: true}
}

func (v NullableFqdnPatternMatchingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFqdnPatternMatchingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


