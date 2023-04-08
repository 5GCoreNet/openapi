/*
Namf_EventExposure

AMF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// LossOfConnectivityReason Describes the reason for loss of connectivity
type LossOfConnectivityReason struct {
	LossOfConnectivityReasonAnyOf *LossOfConnectivityReasonAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LossOfConnectivityReason) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LossOfConnectivityReasonAnyOf
	err = json.Unmarshal(data, &dst.LossOfConnectivityReasonAnyOf);
	if err == nil {
		jsonLossOfConnectivityReasonAnyOf, _ := json.Marshal(dst.LossOfConnectivityReasonAnyOf)
		if string(jsonLossOfConnectivityReasonAnyOf) == "{}" { // empty struct
			dst.LossOfConnectivityReasonAnyOf = nil
		} else {
			return nil // data stored in dst.LossOfConnectivityReasonAnyOf, return on the first match
		}
	} else {
		dst.LossOfConnectivityReasonAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LossOfConnectivityReason)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LossOfConnectivityReason) MarshalJSON() ([]byte, error) {
	if src.LossOfConnectivityReasonAnyOf != nil {
		return json.Marshal(&src.LossOfConnectivityReasonAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLossOfConnectivityReason struct {
	value *LossOfConnectivityReason
	isSet bool
}

func (v NullableLossOfConnectivityReason) Get() *LossOfConnectivityReason {
	return v.value
}

func (v *NullableLossOfConnectivityReason) Set(val *LossOfConnectivityReason) {
	v.value = val
	v.isSet = true
}

func (v NullableLossOfConnectivityReason) IsSet() bool {
	return v.isSet
}

func (v *NullableLossOfConnectivityReason) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLossOfConnectivityReason(val *LossOfConnectivityReason) *NullableLossOfConnectivityReason {
	return &NullableLossOfConnectivityReason{value: val, isSet: true}
}

func (v NullableLossOfConnectivityReason) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLossOfConnectivityReason) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


