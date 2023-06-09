/*
LMF Broadcast

LMF Broadcast Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nlmf_Broadcast

import (
	"encoding/json"
	"fmt"
)

// StorageOutcome Indicates the result of Ciphering Data Set storage.
type StorageOutcome struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *StorageOutcome) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(StorageOutcome)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *StorageOutcome) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableStorageOutcome struct {
	value *StorageOutcome
	isSet bool
}

func (v NullableStorageOutcome) Get() *StorageOutcome {
	return v.value
}

func (v *NullableStorageOutcome) Set(val *StorageOutcome) {
	v.value = val
	v.isSet = true
}

func (v NullableStorageOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableStorageOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStorageOutcome(val *StorageOutcome) *NullableStorageOutcome {
	return &NullableStorageOutcome{value: val, isSet: true}
}

func (v NullableStorageOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStorageOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
