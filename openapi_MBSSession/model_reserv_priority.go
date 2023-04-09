/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
	"fmt"
)

// ReservPriority Indicates the reservation priority.
type ReservPriority struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReservPriority) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReservPriority)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReservPriority) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReservPriority struct {
	value *ReservPriority
	isSet bool
}

func (v NullableReservPriority) Get() *ReservPriority {
	return v.value
}

func (v *NullableReservPriority) Set(val *ReservPriority) {
	v.value = val
	v.isSet = true
}

func (v NullableReservPriority) IsSet() bool {
	return v.isSet
}

func (v *NullableReservPriority) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReservPriority(val *ReservPriority) *NullableReservPriority {
	return &NullableReservPriority{value: val, isSet: true}
}

func (v NullableReservPriority) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReservPriority) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
