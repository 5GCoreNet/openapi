/*
CAPIF_Publish_Service_API

API for publishing service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Publish_Service_API

import (
	"encoding/json"
	"fmt"
)

// DataFormat Possible values are: - JSON: JavaScript Object Notation
type DataFormat struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DataFormat) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(DataFormat)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DataFormat) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDataFormat struct {
	value *DataFormat
	isSet bool
}

func (v NullableDataFormat) Get() *DataFormat {
	return v.value
}

func (v *NullableDataFormat) Set(val *DataFormat) {
	v.value = val
	v.isSet = true
}

func (v NullableDataFormat) IsSet() bool {
	return v.isSet
}

func (v *NullableDataFormat) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataFormat(val *DataFormat) *NullableDataFormat {
	return &NullableDataFormat{value: val, isSet: true}
}

func (v NullableDataFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataFormat) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
