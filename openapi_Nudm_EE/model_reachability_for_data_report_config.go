/*
Nudm_EE

Nudm Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_EE

import (
	"encoding/json"
	"fmt"
)

// ReachabilityForDataReportConfig struct for ReachabilityForDataReportConfig
type ReachabilityForDataReportConfig struct {
	ReachabilityForDataReportConfigAnyOf *ReachabilityForDataReportConfigAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReachabilityForDataReportConfig) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReachabilityForDataReportConfigAnyOf
	err = json.Unmarshal(data, &dst.ReachabilityForDataReportConfigAnyOf);
	if err == nil {
		jsonReachabilityForDataReportConfigAnyOf, _ := json.Marshal(dst.ReachabilityForDataReportConfigAnyOf)
		if string(jsonReachabilityForDataReportConfigAnyOf) == "{}" { // empty struct
			dst.ReachabilityForDataReportConfigAnyOf = nil
		} else {
			return nil // data stored in dst.ReachabilityForDataReportConfigAnyOf, return on the first match
		}
	} else {
		dst.ReachabilityForDataReportConfigAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReachabilityForDataReportConfig)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReachabilityForDataReportConfig) MarshalJSON() ([]byte, error) {
	if src.ReachabilityForDataReportConfigAnyOf != nil {
		return json.Marshal(&src.ReachabilityForDataReportConfigAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReachabilityForDataReportConfig struct {
	value *ReachabilityForDataReportConfig
	isSet bool
}

func (v NullableReachabilityForDataReportConfig) Get() *ReachabilityForDataReportConfig {
	return v.value
}

func (v *NullableReachabilityForDataReportConfig) Set(val *ReachabilityForDataReportConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityForDataReportConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityForDataReportConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityForDataReportConfig(val *ReachabilityForDataReportConfig) *NullableReachabilityForDataReportConfig {
	return &NullableReachabilityForDataReportConfig{value: val, isSet: true}
}

func (v NullableReachabilityForDataReportConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityForDataReportConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


