/*
Npcf_EventExposure

PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// PduSessionInformation - Represents PDU session identification information.
type PduSessionInformation struct {
	AnyOfAnyTypeAnyType *AnyOfAnyTypeAnyType
	Interface           *interface{}
}

// AnyOfAnyTypeAnyTypeAsPduSessionInformation is a convenience function that returns AnyOfAnyTypeAnyType wrapped in PduSessionInformation
func AnyOfAnyTypeAnyTypeAsPduSessionInformation(v *AnyOfAnyTypeAnyType) PduSessionInformation {
	return PduSessionInformation{
		AnyOfAnyTypeAnyType: v,
	}
}

// interface{}AsPduSessionInformation is a convenience function that returns interface{} wrapped in PduSessionInformation
func InterfaceAsPduSessionInformation(v *interface{}) PduSessionInformation {
	return PduSessionInformation{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PduSessionInformation) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AnyOfAnyTypeAnyType
	err = newStrictDecoder(data).Decode(&dst.AnyOfAnyTypeAnyType)
	if err == nil {
		jsonAnyOfAnyTypeAnyType, _ := json.Marshal(dst.AnyOfAnyTypeAnyType)
		if string(jsonAnyOfAnyTypeAnyType) == "{}" { // empty struct
			dst.AnyOfAnyTypeAnyType = nil
		} else {
			match++
		}
	} else {
		dst.AnyOfAnyTypeAnyType = nil
	}

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
		dst.AnyOfAnyTypeAnyType = nil
		dst.Interface = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PduSessionInformation)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PduSessionInformation)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PduSessionInformation) MarshalJSON() ([]byte, error) {
	if src.AnyOfAnyTypeAnyType != nil {
		return json.Marshal(&src.AnyOfAnyTypeAnyType)
	}

	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PduSessionInformation) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AnyOfAnyTypeAnyType != nil {
		return obj.AnyOfAnyTypeAnyType
	}

	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullablePduSessionInformation struct {
	value *PduSessionInformation
	isSet bool
}

func (v NullablePduSessionInformation) Get() *PduSessionInformation {
	return v.value
}

func (v *NullablePduSessionInformation) Set(val *PduSessionInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionInformation(val *PduSessionInformation) *NullablePduSessionInformation {
	return &NullablePduSessionInformation{value: val, isSet: true}
}

func (v NullablePduSessionInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
