/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
	"fmt"
)

// NotificationMethod Possible values are: - PERIODIC - ONE_TIME - ON_EVENT_DETECTION 
type NotificationMethod struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationMethod) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationMethod) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
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


