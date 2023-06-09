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

// DataAvailability Indicates availability of ciphering key data at an LMF.
type DataAvailability struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DataAvailability) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DataAvailability)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DataAvailability) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDataAvailability struct {
	value *DataAvailability
	isSet bool
}

func (v NullableDataAvailability) Get() *DataAvailability {
	return v.value
}

func (v *NullableDataAvailability) Set(val *DataAvailability) {
	v.value = val
	v.isSet = true
}

func (v NullableDataAvailability) IsSet() bool {
	return v.isSet
}

func (v *NullableDataAvailability) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataAvailability(val *DataAvailability) *NullableDataAvailability {
	return &NullableDataAvailability{value: val, isSet: true}
}

func (v NullableDataAvailability) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataAvailability) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
