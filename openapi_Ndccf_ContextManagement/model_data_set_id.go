/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// DataSetId Types of data sets and subsets stored in UDR
type DataSetId struct {
	DataSetIdAnyOf *DataSetIdAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DataSetId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into DataSetIdAnyOf
	err = json.Unmarshal(data, &dst.DataSetIdAnyOf);
	if err == nil {
		jsonDataSetIdAnyOf, _ := json.Marshal(dst.DataSetIdAnyOf)
		if string(jsonDataSetIdAnyOf) == "{}" { // empty struct
			dst.DataSetIdAnyOf = nil
		} else {
			return nil // data stored in dst.DataSetIdAnyOf, return on the first match
		}
	} else {
		dst.DataSetIdAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
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

	return fmt.Errorf("data failed to match schemas in anyOf(DataSetId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DataSetId) MarshalJSON() ([]byte, error) {
	if src.DataSetIdAnyOf != nil {
		return json.Marshal(&src.DataSetIdAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDataSetId struct {
	value *DataSetId
	isSet bool
}

func (v NullableDataSetId) Get() *DataSetId {
	return v.value
}

func (v *NullableDataSetId) Set(val *DataSetId) {
	v.value = val
	v.isSet = true
}

func (v NullableDataSetId) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSetId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSetId(val *DataSetId) *NullableDataSetId {
	return &NullableDataSetId{value: val, isSet: true}
}

func (v NullableDataSetId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSetId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


