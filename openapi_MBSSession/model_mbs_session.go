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

// MbsSession Individual MBS session
type MbsSession struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MbsSession) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(MbsSession)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MbsSession) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMbsSession struct {
	value *MbsSession
	isSet bool
}

func (v NullableMbsSession) Get() *MbsSession {
	return v.value
}

func (v *NullableMbsSession) Set(val *MbsSession) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsSession) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsSession(val *MbsSession) *NullableMbsSession {
	return &NullableMbsSession{value: val, isSet: true}
}

func (v NullableMbsSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
