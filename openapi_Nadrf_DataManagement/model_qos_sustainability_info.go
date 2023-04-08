/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"time"
	"fmt"
)

// QosSustainabilityInfo - Represents the QoS Sustainability information.
type QosSustainabilityInfo struct {
	Interface *interface{}
}

// interface{}AsQosSustainabilityInfo is a convenience function that returns interface{} wrapped in QosSustainabilityInfo
func InterfaceAsQosSustainabilityInfo(v *interface{}) QosSustainabilityInfo {
	return QosSustainabilityInfo{
		Interface: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *QosSustainabilityInfo) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(QosSustainabilityInfo)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(QosSustainabilityInfo)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src QosSustainabilityInfo) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *QosSustainabilityInfo) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullableQosSustainabilityInfo struct {
	value *QosSustainabilityInfo
	isSet bool
}

func (v NullableQosSustainabilityInfo) Get() *QosSustainabilityInfo {
	return v.value
}

func (v *NullableQosSustainabilityInfo) Set(val *QosSustainabilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableQosSustainabilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableQosSustainabilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosSustainabilityInfo(val *QosSustainabilityInfo) *NullableQosSustainabilityInfo {
	return &NullableQosSustainabilityInfo{value: val, isSet: true}
}

func (v NullableQosSustainabilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosSustainabilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


