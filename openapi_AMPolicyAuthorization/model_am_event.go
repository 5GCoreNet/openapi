/*
3gpp-am-policyauthorization

API for AM policy authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AMPolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AmEvent Possible values are: - SAC_CH: Service Area Coverage Change - PDUID_CH: The PDUID assigned to a UE for the UE ProSe Policies changed
type AmEvent struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmEvent) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AmEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmEvent) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmEvent struct {
	value *AmEvent
	isSet bool
}

func (v NullableAmEvent) Get() *AmEvent {
	return v.value
}

func (v *NullableAmEvent) Set(val *AmEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAmEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAmEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmEvent(val *AmEvent) *NullableAmEvent {
	return &NullableAmEvent{value: val, isSet: true}
}

func (v NullableAmEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
