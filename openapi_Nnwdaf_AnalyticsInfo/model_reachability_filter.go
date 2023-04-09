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

// ReachabilityFilter Event filter for REACHABILITY_REPORT event type
type ReachabilityFilter struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReachabilityFilter) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReachabilityFilter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReachabilityFilter) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReachabilityFilter struct {
	value *ReachabilityFilter
	isSet bool
}

func (v NullableReachabilityFilter) Get() *ReachabilityFilter {
	return v.value
}

func (v *NullableReachabilityFilter) Set(val *ReachabilityFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityFilter(val *ReachabilityFilter) *NullableReachabilityFilter {
	return &NullableReachabilityFilter{value: val, isSet: true}
}

func (v NullableReachabilityFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
