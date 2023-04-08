/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// OutputStrategy Possible values are: - BINARY: Indicates that the analytics shall only be reported when the requested level of accuracy is reached within a cycle of periodic notification. - GRADIENT: Indicates that the analytics shall be reported according with the periodicity irrespective of whether the requested level of accuracy has been reached or not. 
type OutputStrategy struct {
	OutputStrategyAnyOf *OutputStrategyAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *OutputStrategy) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into OutputStrategyAnyOf
	err = json.Unmarshal(data, &dst.OutputStrategyAnyOf);
	if err == nil {
		jsonOutputStrategyAnyOf, _ := json.Marshal(dst.OutputStrategyAnyOf)
		if string(jsonOutputStrategyAnyOf) == "{}" { // empty struct
			dst.OutputStrategyAnyOf = nil
		} else {
			return nil // data stored in dst.OutputStrategyAnyOf, return on the first match
		}
	} else {
		dst.OutputStrategyAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(OutputStrategy)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *OutputStrategy) MarshalJSON() ([]byte, error) {
	if src.OutputStrategyAnyOf != nil {
		return json.Marshal(&src.OutputStrategyAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableOutputStrategy struct {
	value *OutputStrategy
	isSet bool
}

func (v NullableOutputStrategy) Get() *OutputStrategy {
	return v.value
}

func (v *NullableOutputStrategy) Set(val *OutputStrategy) {
	v.value = val
	v.isSet = true
}

func (v NullableOutputStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableOutputStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutputStrategy(val *OutputStrategy) *NullableOutputStrategy {
	return &NullableOutputStrategy{value: val, isSet: true}
}

func (v NullableOutputStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutputStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


