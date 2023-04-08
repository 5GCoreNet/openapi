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

// ComparisonOperator TBD
type ComparisonOperator struct {
	ComparisonOperatorAnyOf *ComparisonOperatorAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ComparisonOperator) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ComparisonOperatorAnyOf
	err = json.Unmarshal(data, &dst.ComparisonOperatorAnyOf);
	if err == nil {
		jsonComparisonOperatorAnyOf, _ := json.Marshal(dst.ComparisonOperatorAnyOf)
		if string(jsonComparisonOperatorAnyOf) == "{}" { // empty struct
			dst.ComparisonOperatorAnyOf = nil
		} else {
			return nil // data stored in dst.ComparisonOperatorAnyOf, return on the first match
		}
	} else {
		dst.ComparisonOperatorAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ComparisonOperator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ComparisonOperator) MarshalJSON() ([]byte, error) {
	if src.ComparisonOperatorAnyOf != nil {
		return json.Marshal(&src.ComparisonOperatorAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableComparisonOperator struct {
	value *ComparisonOperator
	isSet bool
}

func (v NullableComparisonOperator) Get() *ComparisonOperator {
	return v.value
}

func (v *NullableComparisonOperator) Set(val *ComparisonOperator) {
	v.value = val
	v.isSet = true
}

func (v NullableComparisonOperator) IsSet() bool {
	return v.isSet
}

func (v *NullableComparisonOperator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComparisonOperator(val *ComparisonOperator) *NullableComparisonOperator {
	return &NullableComparisonOperator{value: val, isSet: true}
}

func (v NullableComparisonOperator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComparisonOperator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


