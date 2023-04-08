/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// QosRequirement - Represents the QoS requirements.
type QosRequirement struct {
	Interface *interface{}
}

// interface{}AsQosRequirement is a convenience function that returns interface{} wrapped in QosRequirement
func InterfaceAsQosRequirement(v *interface{}) QosRequirement {
	return QosRequirement{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QosRequirement) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(QosRequirement)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QosRequirement)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QosRequirement) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QosRequirement) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableQosRequirement struct {
	value *QosRequirement
	isSet bool
}

func (v NullableQosRequirement) Get() *QosRequirement {
	return v.value
}

func (v *NullableQosRequirement) Set(val *QosRequirement) {
	v.value = val
	v.isSet = true
}

func (v NullableQosRequirement) IsSet() bool {
	return v.isSet
}

func (v *NullableQosRequirement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosRequirement(val *QosRequirement) *NullableQosRequirement {
	return &NullableQosRequirement{value: val, isSet: true}
}

func (v NullableQosRequirement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosRequirement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


