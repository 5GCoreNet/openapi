/*
Npcf_EventExposure

PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// NotificationMethod Possible values are: - PERIODIC - ONE_TIME - ON_EVENT_DETECTION
type NotificationMethod struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationMethod) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationMethod) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationMethod struct {
	value *NotificationMethod
	isSet bool
}

func (v NullableNotificationMethod) Get() *NotificationMethod {
	return v.value
}

func (v *NullableNotificationMethod) Set(val *NotificationMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationMethod(val *NotificationMethod) *NullableNotificationMethod {
	return &NullableNotificationMethod{value: val, isSet: true}
}

func (v NullableNotificationMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
