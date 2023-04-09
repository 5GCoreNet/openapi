/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// PresenceState Possible values are: -IN_AREA: Indicates that the UE is inside or enters the presence reporting area. -OUT_OF_AREA: Indicates that the UE is outside or leaves the presence reporting area -UNKNOW: Indicates it is unknown whether the UE is in the presence reporting area or not -INACTIVE: Indicates that the presence reporting area is inactive in the serving node.
type PresenceState struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PresenceState) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PresenceState)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PresenceState) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePresenceState struct {
	value *PresenceState
	isSet bool
}

func (v NullablePresenceState) Get() *PresenceState {
	return v.value
}

func (v *NullablePresenceState) Set(val *PresenceState) {
	v.value = val
	v.isSet = true
}

func (v NullablePresenceState) IsSet() bool {
	return v.isSet
}

func (v *NullablePresenceState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePresenceState(val *PresenceState) *NullablePresenceState {
	return &NullablePresenceState{value: val, isSet: true}
}

func (v NullablePresenceState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePresenceState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
