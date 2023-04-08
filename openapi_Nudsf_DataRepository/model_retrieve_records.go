/*
Nudsf_DataRepository

Nudsf Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudsf_DataRepository

import (
	"encoding/json"
	"fmt"
)

// RetrieveRecords Indicates the data to be retrieved.
type RetrieveRecords struct {
	RetrieveRecordsAnyOf *RetrieveRecordsAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RetrieveRecords) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RetrieveRecordsAnyOf
	err = json.Unmarshal(data, &dst.RetrieveRecordsAnyOf);
	if err == nil {
		jsonRetrieveRecordsAnyOf, _ := json.Marshal(dst.RetrieveRecordsAnyOf)
		if string(jsonRetrieveRecordsAnyOf) == "{}" { // empty struct
			dst.RetrieveRecordsAnyOf = nil
		} else {
			return nil // data stored in dst.RetrieveRecordsAnyOf, return on the first match
		}
	} else {
		dst.RetrieveRecordsAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RetrieveRecords)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RetrieveRecords) MarshalJSON() ([]byte, error) {
	if src.RetrieveRecordsAnyOf != nil {
		return json.Marshal(&src.RetrieveRecordsAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRetrieveRecords struct {
	value *RetrieveRecords
	isSet bool
}

func (v NullableRetrieveRecords) Get() *RetrieveRecords {
	return v.value
}

func (v *NullableRetrieveRecords) Set(val *RetrieveRecords) {
	v.value = val
	v.isSet = true
}

func (v NullableRetrieveRecords) IsSet() bool {
	return v.isSet
}

func (v *NullableRetrieveRecords) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRetrieveRecords(val *RetrieveRecords) *NullableRetrieveRecords {
	return &NullableRetrieveRecords{value: val, isSet: true}
}

func (v NullableRetrieveRecords) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRetrieveRecords) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


