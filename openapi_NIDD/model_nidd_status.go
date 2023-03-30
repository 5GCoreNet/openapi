/*
3gpp-nidd

API for non IP data delivery.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NIDD

import (
	"encoding/json"
	"fmt"
)

// NiddStatus Possible values are - ACTIVE: The NIDD configuration is active. - TERMINATED_UE_NOT_AUTHORIZED: The NIDD configuration was terminated because the UE´s authorisation was revoked. - TERMINATED: The NIDD configuration was terminated. - RDS_PORT_UNKNOWN: The RDS port is unknown. 
type NiddStatus struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NiddStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NiddStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NiddStatus) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNiddStatus struct {
	value *NiddStatus
	isSet bool
}

func (v NullableNiddStatus) Get() *NiddStatus {
	return v.value
}

func (v *NullableNiddStatus) Set(val *NiddStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNiddStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNiddStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiddStatus(val *NiddStatus) *NullableNiddStatus {
	return &NullableNiddStatus{value: val, isSet: true}
}

func (v NullableNiddStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiddStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

