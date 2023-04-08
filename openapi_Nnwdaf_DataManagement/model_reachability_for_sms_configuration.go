/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// ReachabilityForSmsConfiguration struct for ReachabilityForSmsConfiguration
type ReachabilityForSmsConfiguration struct {
	ReachabilityForSmsConfigurationAnyOf *ReachabilityForSmsConfigurationAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReachabilityForSmsConfiguration) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReachabilityForSmsConfigurationAnyOf
	err = json.Unmarshal(data, &dst.ReachabilityForSmsConfigurationAnyOf);
	if err == nil {
		jsonReachabilityForSmsConfigurationAnyOf, _ := json.Marshal(dst.ReachabilityForSmsConfigurationAnyOf)
		if string(jsonReachabilityForSmsConfigurationAnyOf) == "{}" { // empty struct
			dst.ReachabilityForSmsConfigurationAnyOf = nil
		} else {
			return nil // data stored in dst.ReachabilityForSmsConfigurationAnyOf, return on the first match
		}
	} else {
		dst.ReachabilityForSmsConfigurationAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReachabilityForSmsConfiguration)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReachabilityForSmsConfiguration) MarshalJSON() ([]byte, error) {
	if src.ReachabilityForSmsConfigurationAnyOf != nil {
		return json.Marshal(&src.ReachabilityForSmsConfigurationAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReachabilityForSmsConfiguration struct {
	value *ReachabilityForSmsConfiguration
	isSet bool
}

func (v NullableReachabilityForSmsConfiguration) Get() *ReachabilityForSmsConfiguration {
	return v.value
}

func (v *NullableReachabilityForSmsConfiguration) Set(val *ReachabilityForSmsConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityForSmsConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityForSmsConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityForSmsConfiguration(val *ReachabilityForSmsConfiguration) *NullableReachabilityForSmsConfiguration {
	return &NullableReachabilityForSmsConfiguration{value: val, isSet: true}
}

func (v NullableReachabilityForSmsConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityForSmsConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


