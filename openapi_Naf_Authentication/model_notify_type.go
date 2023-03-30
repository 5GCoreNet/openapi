/*
Naf_Authentication

AF Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_Authentication

import (
	"encoding/json"
	"fmt"
)

// NotifyType Possible values are: - REAUTHENTICATE: The UAV needs to be reauthenticated. - REAUTHORIZE: Authorization data needs to be updated to UAV. - REVOKE: Revoke UAV authentication and authorization. 
type NotifyType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotifyType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotifyType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotifyType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotifyType struct {
	value *NotifyType
	isSet bool
}

func (v NullableNotifyType) Get() *NotifyType {
	return v.value
}

func (v *NullableNotifyType) Set(val *NotifyType) {
	v.value = val
	v.isSet = true
}

func (v NullableNotifyType) IsSet() bool {
	return v.isSet
}

func (v *NullableNotifyType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotifyType(val *NotifyType) *NullableNotifyType {
	return &NullableNotifyType{value: val, isSet: true}
}

func (v NullableNotifyType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotifyType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

