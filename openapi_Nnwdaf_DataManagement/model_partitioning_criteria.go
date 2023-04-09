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

// PartitioningCriteria Possible values are: - \"TAC\": Type Allocation Code - \"SUBPLMN\": Subscriber PLMN ID - \"GEOAREA\": Geographical area, i.e. list(s) of TAI(s) - \"SNSSAI\": S-NSSAI - \"DNN\": DNN
type PartitioningCriteria struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PartitioningCriteria) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PartitioningCriteria)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PartitioningCriteria) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePartitioningCriteria struct {
	value *PartitioningCriteria
	isSet bool
}

func (v NullablePartitioningCriteria) Get() *PartitioningCriteria {
	return v.value
}

func (v *NullablePartitioningCriteria) Set(val *PartitioningCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullablePartitioningCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullablePartitioningCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePartitioningCriteria(val *PartitioningCriteria) *NullablePartitioningCriteria {
	return &NullablePartitioningCriteria{value: val, isSet: true}
}

func (v NullablePartitioningCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePartitioningCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
