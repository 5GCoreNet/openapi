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

// DataInd Possible values are - PFD - IPTV - BDT - SVC_PARAM - AM 
type DataInd struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *DataInd) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(DataInd)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *DataInd) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableDataInd struct {
	value *DataInd
	isSet bool
}

func (v NullableDataInd) Get() *DataInd {
	return v.value
}

func (v *NullableDataInd) Set(val *DataInd) {
	v.value = val
	v.isSet = true
}

func (v NullableDataInd) IsSet() bool {
	return v.isSet
}

func (v *NullableDataInd) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataInd(val *DataInd) *NullableDataInd {
	return &NullableDataInd{value: val, isSet: true}
}

func (v NullableDataInd) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataInd) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


