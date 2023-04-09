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

// NotificationControlRm This enumeration is defined in the same way as the 'NotificationControl' enumeration, but with the OpenAPI 'nullable: true' property.
type NotificationControlRm struct {
	NotificationControl *NotificationControl
	NullValue           *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationControlRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NotificationControl
	err = json.Unmarshal(data, &dst.NotificationControl)
	if err == nil {
		jsonNotificationControl, _ := json.Marshal(dst.NotificationControl)
		if string(jsonNotificationControl) == "{}" { // empty struct
			dst.NotificationControl = nil
		} else {
			return nil // data stored in dst.NotificationControl, return on the first match
		}
	} else {
		dst.NotificationControl = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue)
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationControlRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationControlRm) MarshalJSON() ([]byte, error) {
	if src.NotificationControl != nil {
		return json.Marshal(&src.NotificationControl)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationControlRm struct {
	value *NotificationControlRm
	isSet bool
}

func (v NullableNotificationControlRm) Get() *NotificationControlRm {
	return v.value
}

func (v *NullableNotificationControlRm) Set(val *NotificationControlRm) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationControlRm) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationControlRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationControlRm(val *NotificationControlRm) *NullableNotificationControlRm {
	return &NullableNotificationControlRm{value: val, isSet: true}
}

func (v NullableNotificationControlRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationControlRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
