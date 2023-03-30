/*
Npcf_AMPolicyAuthorization Service API

PCF Access and Mobility Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_AMPolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AppAmContextRespData It represents a response to a modification or creation request of an Individual Application AM resource. It may contain the notification of the already met events.
type AppAmContextRespData struct {
	AmEventsNotification *AmEventsNotification
	AppAmContextData *AppAmContextData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AppAmContextRespData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AmEventsNotification
	err = json.Unmarshal(data, &dst.AmEventsNotification);
	if err == nil {
		jsonAmEventsNotification, _ := json.Marshal(dst.AmEventsNotification)
		if string(jsonAmEventsNotification) == "{}" { // empty struct
			dst.AmEventsNotification = nil
		} else {
			return nil // data stored in dst.AmEventsNotification, return on the first match
		}
	} else {
		dst.AmEventsNotification = nil
	}

	// try to unmarshal JSON data into AppAmContextData
	err = json.Unmarshal(data, &dst.AppAmContextData);
	if err == nil {
		jsonAppAmContextData, _ := json.Marshal(dst.AppAmContextData)
		if string(jsonAppAmContextData) == "{}" { // empty struct
			dst.AppAmContextData = nil
		} else {
			return nil // data stored in dst.AppAmContextData, return on the first match
		}
	} else {
		dst.AppAmContextData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AppAmContextRespData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AppAmContextRespData) MarshalJSON() ([]byte, error) {
	if src.AmEventsNotification != nil {
		return json.Marshal(&src.AmEventsNotification)
	}

	if src.AppAmContextData != nil {
		return json.Marshal(&src.AppAmContextData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAppAmContextRespData struct {
	value *AppAmContextRespData
	isSet bool
}

func (v NullableAppAmContextRespData) Get() *AppAmContextRespData {
	return v.value
}

func (v *NullableAppAmContextRespData) Set(val *AppAmContextRespData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAmContextRespData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAmContextRespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAmContextRespData(val *AppAmContextRespData) *NullableAppAmContextRespData {
	return &NullableAppAmContextRespData{value: val, isSet: true}
}

func (v NullableAppAmContextRespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAmContextRespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


