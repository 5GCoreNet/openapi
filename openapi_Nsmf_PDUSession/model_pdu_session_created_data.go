/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// PduSessionCreatedData - Data within Create Response
type PduSessionCreatedData struct {
	Interface *interface{}
}

// interface{}AsPduSessionCreatedData is a convenience function that returns interface{} wrapped in PduSessionCreatedData
func InterfaceAsPduSessionCreatedData(v *interface{}) PduSessionCreatedData {
	return PduSessionCreatedData{
		Interface: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *PduSessionCreatedData) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(PduSessionCreatedData)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PduSessionCreatedData)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PduSessionCreatedData) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PduSessionCreatedData) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Interface != nil {
		return obj.Interface
	}

	// all schemas are nil
	return nil
}

type NullablePduSessionCreatedData struct {
	value *PduSessionCreatedData
	isSet bool
}

func (v NullablePduSessionCreatedData) Get() *PduSessionCreatedData {
	return v.value
}

func (v *NullablePduSessionCreatedData) Set(val *PduSessionCreatedData) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionCreatedData) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionCreatedData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionCreatedData(val *PduSessionCreatedData) *NullablePduSessionCreatedData {
	return &NullablePduSessionCreatedData{value: val, isSet: true}
}

func (v NullablePduSessionCreatedData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionCreatedData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
