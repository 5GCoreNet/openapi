/*
3gpp-service-parameter

API for AF service paramter   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ServiceParameter

import (
	"encoding/json"
	"fmt"
)

// AfNotification Notifications upon AF Service Parameter Authorization Update e.g. to revoke the authorization, and/or AF subscribed event notification of the outcome related to the invocation of service parameter provisioning.
type AfNotification struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfNotification) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AfNotification)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AfNotification) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAfNotification struct {
	value *AfNotification
	isSet bool
}

func (v NullableAfNotification) Get() *AfNotification {
	return v.value
}

func (v *NullableAfNotification) Set(val *AfNotification) {
	v.value = val
	v.isSet = true
}

func (v NullableAfNotification) IsSet() bool {
	return v.isSet
}

func (v *NullableAfNotification) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfNotification(val *AfNotification) *NullableAfNotification {
	return &NullableAfNotification{value: val, isSet: true}
}

func (v NullableAfNotification) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfNotification) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
