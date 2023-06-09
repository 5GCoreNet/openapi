/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// ImmediateReport1 - struct for ImmediateReport1
type ImmediateReport1 struct {
	ProvisionedDataSets *ProvisionedDataSets
	ArrayOfSharedData   *[]SharedData
}

// ProvisionedDataSetsAsImmediateReport1 is a convenience function that returns ProvisionedDataSets wrapped in ImmediateReport1
func ProvisionedDataSetsAsImmediateReport1(v *ProvisionedDataSets) ImmediateReport1 {
	return ImmediateReport1{
		ProvisionedDataSets: v,
	}
}

// []SharedDataAsImmediateReport1 is a convenience function that returns []SharedData wrapped in ImmediateReport1
func ArrayOfSharedDataAsImmediateReport1(v *[]SharedData) ImmediateReport1 {
	return ImmediateReport1{
		ArrayOfSharedData: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *ImmediateReport1) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ProvisionedDataSets
	err = newStrictDecoder(data).Decode(&dst.ProvisionedDataSets)
	if err == nil {
		jsonProvisionedDataSets, _ := json.Marshal(dst.ProvisionedDataSets)
		if string(jsonProvisionedDataSets) == "{}" { // empty struct
			dst.ProvisionedDataSets = nil
		} else {
			match++
		}
	} else {
		dst.ProvisionedDataSets = nil
	}

	// try to unmarshal data into ArrayOfSharedData
	err = newStrictDecoder(data).Decode(&dst.ArrayOfSharedData)
	if err == nil {
		jsonArrayOfSharedData, _ := json.Marshal(dst.ArrayOfSharedData)
		if string(jsonArrayOfSharedData) == "{}" { // empty struct
			dst.ArrayOfSharedData = nil
		} else {
			match++
		}
	} else {
		dst.ArrayOfSharedData = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ProvisionedDataSets = nil
		dst.ArrayOfSharedData = nil

		return fmt.Errorf("data matches more than one schema in oneOf(ImmediateReport1)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(ImmediateReport1)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src ImmediateReport1) MarshalJSON() ([]byte, error) {
	if src.ProvisionedDataSets != nil {
		return json.Marshal(&src.ProvisionedDataSets)
	}

	if src.ArrayOfSharedData != nil {
		return json.Marshal(&src.ArrayOfSharedData)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *ImmediateReport1) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ProvisionedDataSets != nil {
		return obj.ProvisionedDataSets
	}

	if obj.ArrayOfSharedData != nil {
		return obj.ArrayOfSharedData
	}

	// all schemas are nil
	return nil
}

type NullableImmediateReport1 struct {
	value *ImmediateReport1
	isSet bool
}

func (v NullableImmediateReport1) Get() *ImmediateReport1 {
	return v.value
}

func (v *NullableImmediateReport1) Set(val *ImmediateReport1) {
	v.value = val
	v.isSet = true
}

func (v NullableImmediateReport1) IsSet() bool {
	return v.isSet
}

func (v *NullableImmediateReport1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImmediateReport1(val *ImmediateReport1) *NullableImmediateReport1 {
	return &NullableImmediateReport1{value: val, isSet: true}
}

func (v NullableImmediateReport1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImmediateReport1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
