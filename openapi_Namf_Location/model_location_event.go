/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
	"fmt"
)

// LocationEvent Type of events initiating location procedures
type LocationEvent struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LocationEvent) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(LocationEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LocationEvent) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocationEvent struct {
	value *LocationEvent
	isSet bool
}

func (v NullableLocationEvent) Get() *LocationEvent {
	return v.value
}

func (v *NullableLocationEvent) Set(val *LocationEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationEvent(val *LocationEvent) *NullableLocationEvent {
	return &NullableLocationEvent{value: val, isSet: true}
}

func (v NullableLocationEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


