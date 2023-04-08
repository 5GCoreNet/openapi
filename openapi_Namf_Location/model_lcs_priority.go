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

// LcsPriority Indicates priority of the LCS client.
type LcsPriority struct {
	LcsPriorityAnyOf *LcsPriorityAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LcsPriority) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LcsPriorityAnyOf
	err = json.Unmarshal(data, &dst.LcsPriorityAnyOf);
	if err == nil {
		jsonLcsPriorityAnyOf, _ := json.Marshal(dst.LcsPriorityAnyOf)
		if string(jsonLcsPriorityAnyOf) == "{}" { // empty struct
			dst.LcsPriorityAnyOf = nil
		} else {
			return nil // data stored in dst.LcsPriorityAnyOf, return on the first match
		}
	} else {
		dst.LcsPriorityAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LcsPriority)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LcsPriority) MarshalJSON() ([]byte, error) {
	if src.LcsPriorityAnyOf != nil {
		return json.Marshal(&src.LcsPriorityAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLcsPriority struct {
	value *LcsPriority
	isSet bool
}

func (v NullableLcsPriority) Get() *LcsPriority {
	return v.value
}

func (v *NullableLcsPriority) Set(val *LcsPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsPriority(val *LcsPriority) *NullableLcsPriority {
	return &NullableLcsPriority{value: val, isSet: true}
}

func (v NullableLcsPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


