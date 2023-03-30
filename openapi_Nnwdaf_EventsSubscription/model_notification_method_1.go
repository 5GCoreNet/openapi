/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// NotificationMethod1 Possible values are: - PERIODIC - ONE_TIME - ON_EVENT_DETECTION 
type NotificationMethod1 struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationMethod1) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationMethod1)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationMethod1) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationMethod1 struct {
	value *NotificationMethod1
	isSet bool
}

func (v NullableNotificationMethod1) Get() *NotificationMethod1 {
	return v.value
}

func (v *NullableNotificationMethod1) Set(val *NotificationMethod1) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationMethod1) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationMethod1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationMethod1(val *NotificationMethod1) *NullableNotificationMethod1 {
	return &NullableNotificationMethod1{value: val, isSet: true}
}

func (v NullableNotificationMethod1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationMethod1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


