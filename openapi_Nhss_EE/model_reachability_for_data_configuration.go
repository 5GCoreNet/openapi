/*
Nhss_EE

HSS Event Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_EE

import (
	"encoding/json"
	"fmt"
)

// ReachabilityForDataConfiguration Contains data needed for a Monitoring Configuration, specific to the 'Reachability for Data' event type
type ReachabilityForDataConfiguration struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReachabilityForDataConfiguration) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface)
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			return nil // data stored in dst.Interface, return on the first match
		}
	} else {
		dst.Interface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ReachabilityForDataConfiguration)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReachabilityForDataConfiguration) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReachabilityForDataConfiguration struct {
	value *ReachabilityForDataConfiguration
	isSet bool
}

func (v NullableReachabilityForDataConfiguration) Get() *ReachabilityForDataConfiguration {
	return v.value
}

func (v *NullableReachabilityForDataConfiguration) Set(val *ReachabilityForDataConfiguration) {
	v.value = val
	v.isSet = true
}

func (v NullableReachabilityForDataConfiguration) IsSet() bool {
	return v.isSet
}

func (v *NullableReachabilityForDataConfiguration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReachabilityForDataConfiguration(val *ReachabilityForDataConfiguration) *NullableReachabilityForDataConfiguration {
	return &NullableReachabilityForDataConfiguration{value: val, isSet: true}
}

func (v NullableReachabilityForDataConfiguration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReachabilityForDataConfiguration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
