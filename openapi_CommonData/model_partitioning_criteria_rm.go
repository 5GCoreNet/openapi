/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// PartitioningCriteriaRm This data type is defined in the same way as the ' PartitioningCriteria ' data type, but with the OpenAPI 'nullable: true' property. 
type PartitioningCriteriaRm struct {
	NullValue *NullValue
	PartitioningCriteria *PartitioningCriteria
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PartitioningCriteriaRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	// try to unmarshal JSON data into PartitioningCriteria
	err = json.Unmarshal(data, &dst.PartitioningCriteria);
	if err == nil {
		jsonPartitioningCriteria, _ := json.Marshal(dst.PartitioningCriteria)
		if string(jsonPartitioningCriteria) == "{}" { // empty struct
			dst.PartitioningCriteria = nil
		} else {
			return nil // data stored in dst.PartitioningCriteria, return on the first match
		}
	} else {
		dst.PartitioningCriteria = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PartitioningCriteriaRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PartitioningCriteriaRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.PartitioningCriteria != nil {
		return json.Marshal(&src.PartitioningCriteria)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePartitioningCriteriaRm struct {
	value *PartitioningCriteriaRm
	isSet bool
}

func (v NullablePartitioningCriteriaRm) Get() *PartitioningCriteriaRm {
	return v.value
}

func (v *NullablePartitioningCriteriaRm) Set(val *PartitioningCriteriaRm) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitioningCriteriaRm) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitioningCriteriaRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitioningCriteriaRm(val *PartitioningCriteriaRm) *NullablePartitioningCriteriaRm {
	return &NullablePartitioningCriteriaRm{value: val, isSet: true}
}

func (v NullablePartitioningCriteriaRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitioningCriteriaRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


