/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// FailureHandling struct for FailureHandling
type FailureHandling struct {
	FailureHandlingAnyOf *FailureHandlingAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *FailureHandling) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into FailureHandlingAnyOf
	err = json.Unmarshal(data, &dst.FailureHandlingAnyOf);
	if err == nil {
		jsonFailureHandlingAnyOf, _ := json.Marshal(dst.FailureHandlingAnyOf)
		if string(jsonFailureHandlingAnyOf) == "{}" { // empty struct
			dst.FailureHandlingAnyOf = nil
		} else {
			return nil // data stored in dst.FailureHandlingAnyOf, return on the first match
		}
	} else {
		dst.FailureHandlingAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(FailureHandling)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *FailureHandling) MarshalJSON() ([]byte, error) {
	if src.FailureHandlingAnyOf != nil {
		return json.Marshal(&src.FailureHandlingAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableFailureHandling struct {
	value *FailureHandling
	isSet bool
}

func (v NullableFailureHandling) Get() *FailureHandling {
	return v.value
}

func (v *NullableFailureHandling) Set(val *FailureHandling) {
	v.value = val
	v.isSet = true
}

func (v NullableFailureHandling) IsSet() bool {
	return v.isSet
}

func (v *NullableFailureHandling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFailureHandling(val *FailureHandling) *NullableFailureHandling {
	return &NullableFailureHandling{value: val, isSet: true}
}

func (v NullableFailureHandling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFailureHandling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


