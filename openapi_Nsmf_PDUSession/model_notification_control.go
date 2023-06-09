/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// NotificationControl The enumeration NotificationControl indicates whether notifications are requested from the RAN when the GFBR can no longer  (or again) be fulfilled for a QoS Flow during the lifetime of the QoS Flow (see clause 5.7.2.4 of 3GPP TS 23.501). It shall comply with the provisions defined in table 5.5.3.5-1.
type NotificationControl struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NotificationControl) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(NotificationControl)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NotificationControl) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNotificationControl struct {
	value *NotificationControl
	isSet bool
}

func (v NullableNotificationControl) Get() *NotificationControl {
	return v.value
}

func (v *NullableNotificationControl) Set(val *NotificationControl) {
	v.value = val
	v.isSet = true
}

func (v NullableNotificationControl) IsSet() bool {
	return v.isSet
}

func (v *NullableNotificationControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNotificationControl(val *NotificationControl) *NullableNotificationControl {
	return &NullableNotificationControl{value: val, isSet: true}
}

func (v NullableNotificationControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNotificationControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
