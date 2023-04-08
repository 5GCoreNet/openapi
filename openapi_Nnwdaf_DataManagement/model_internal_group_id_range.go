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

// InternalGroupIdRange - A range of Group IDs (internal group identities), either based on a numeric range, or based on regular-expression matching 
type InternalGroupIdRange struct {
	Interface *interface{}
}

// interface{}AsInternalGroupIdRange is a convenience function that returns interface{} wrapped in InternalGroupIdRange
func InterfaceAsInternalGroupIdRange(v *interface{}) InternalGroupIdRange {
	return InternalGroupIdRange{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *InternalGroupIdRange) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(InternalGroupIdRange)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(InternalGroupIdRange)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src InternalGroupIdRange) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *InternalGroupIdRange) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableInternalGroupIdRange struct {
	value *InternalGroupIdRange
	isSet bool
}

func (v NullableInternalGroupIdRange) Get() *InternalGroupIdRange {
	return v.value
}

func (v *NullableInternalGroupIdRange) Set(val *InternalGroupIdRange) {
	v.value = val
	v.isSet = true
}

func (v NullableInternalGroupIdRange) IsSet() bool {
	return v.isSet
}

func (v *NullableInternalGroupIdRange) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInternalGroupIdRange(val *InternalGroupIdRange) *NullableInternalGroupIdRange {
	return &NullableInternalGroupIdRange{value: val, isSet: true}
}

func (v NullableInternalGroupIdRange) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInternalGroupIdRange) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


