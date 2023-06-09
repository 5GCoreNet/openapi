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

// AmEventsSubscRespData Identifies the events the application subscribes to within an AM Policy Events Subscription subresource data. It may contain the notification of the already met events.
type AmEventsSubscRespData struct {
	AmEventsNotification *AmEventsNotification
	AmEventsSubscData    *AmEventsSubscData
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AmEventsSubscRespData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AmEventsNotification
	err = json.Unmarshal(data, &dst.AmEventsNotification)
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

	// try to unmarshal JSON data into AmEventsSubscData
	err = json.Unmarshal(data, &dst.AmEventsSubscData)
	if err == nil {
		jsonAmEventsSubscData, _ := json.Marshal(dst.AmEventsSubscData)
		if string(jsonAmEventsSubscData) == "{}" { // empty struct
			dst.AmEventsSubscData = nil
		} else {
			return nil // data stored in dst.AmEventsSubscData, return on the first match
		}
	} else {
		dst.AmEventsSubscData = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AmEventsSubscRespData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AmEventsSubscRespData) MarshalJSON() ([]byte, error) {
	if src.AmEventsNotification != nil {
		return json.Marshal(&src.AmEventsNotification)
	}

	if src.AmEventsSubscData != nil {
		return json.Marshal(&src.AmEventsSubscData)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAmEventsSubscRespData struct {
	value *AmEventsSubscRespData
	isSet bool
}

func (v NullableAmEventsSubscRespData) Get() *AmEventsSubscRespData {
	return v.value
}

func (v *NullableAmEventsSubscRespData) Set(val *AmEventsSubscRespData) {
	v.value = val
	v.isSet = true
}

func (v NullableAmEventsSubscRespData) IsSet() bool {
	return v.isSet
}

func (v *NullableAmEventsSubscRespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmEventsSubscRespData(val *AmEventsSubscRespData) *NullableAmEventsSubscRespData {
	return &NullableAmEventsSubscRespData{value: val, isSet: true}
}

func (v NullableAmEventsSubscRespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmEventsSubscRespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
