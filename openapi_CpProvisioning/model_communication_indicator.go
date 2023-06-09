/*
3gpp-cp-parameter-provisioning

API for provisioning communication pattern parameters.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CpProvisioning

import (
	"encoding/json"
	"fmt"
)

// CommunicationIndicator Possible values are - PERIODICALLY: Identifies the UE communicates periodically - ON_DEMAND: Identifies the UE communicates on demand
type CommunicationIndicator struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CommunicationIndicator) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(CommunicationIndicator)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CommunicationIndicator) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCommunicationIndicator struct {
	value *CommunicationIndicator
	isSet bool
}

func (v NullableCommunicationIndicator) Get() *CommunicationIndicator {
	return v.value
}

func (v *NullableCommunicationIndicator) Set(val *CommunicationIndicator) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationIndicator) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationIndicator) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationIndicator(val *CommunicationIndicator) *NullableCommunicationIndicator {
	return &NullableCommunicationIndicator{value: val, isSet: true}
}

func (v NullableCommunicationIndicator) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationIndicator) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
