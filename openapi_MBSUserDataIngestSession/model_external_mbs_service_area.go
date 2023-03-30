/*
3gpp-mbs-ud-ingest

API for MBS User Data Ingest Session.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSUserDataIngestSession

import (
	"encoding/json"
	"fmt"
)

// ExternalMbsServiceArea - List of geographic area or list of civic address info for MBS Service Area
type ExternalMbsServiceArea struct {
	Interface{} *interface{}
}

// interface{}AsExternalMbsServiceArea is a convenience function that returns interface{} wrapped in ExternalMbsServiceArea
func Interface{}AsExternalMbsServiceArea(v *interface{}) ExternalMbsServiceArea {
	return ExternalMbsServiceArea{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *ExternalMbsServiceArea) UnmarshalJSON(data []byte) error {
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

		return fmt.Errorf("data matches more than one schema in oneOf(ExternalMbsServiceArea)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ExternalMbsServiceArea)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ExternalMbsServiceArea) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ExternalMbsServiceArea) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableExternalMbsServiceArea struct {
	value *ExternalMbsServiceArea
	isSet bool
}

func (v NullableExternalMbsServiceArea) Get() *ExternalMbsServiceArea {
	return v.value
}

func (v *NullableExternalMbsServiceArea) Set(val *ExternalMbsServiceArea) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalMbsServiceArea) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalMbsServiceArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalMbsServiceArea(val *ExternalMbsServiceArea) *NullableExternalMbsServiceArea {
	return &NullableExternalMbsServiceArea{value: val, isSet: true}
}

func (v NullableExternalMbsServiceArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalMbsServiceArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


