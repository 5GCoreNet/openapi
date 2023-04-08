/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
	"fmt"
)

// SubscribedEvent Possible values are: - UP_PATH_CHANGE: The AF requests to be notified when the UP path changes for the PDU session. 
type SubscribedEvent struct {
	SubscribedEventAnyOf *SubscribedEventAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SubscribedEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into SubscribedEventAnyOf
	err = json.Unmarshal(data, &dst.SubscribedEventAnyOf);
	if err == nil {
		jsonSubscribedEventAnyOf, _ := json.Marshal(dst.SubscribedEventAnyOf)
		if string(jsonSubscribedEventAnyOf) == "{}" { // empty struct
			dst.SubscribedEventAnyOf = nil
		} else {
			return nil // data stored in dst.SubscribedEventAnyOf, return on the first match
		}
	} else {
		dst.SubscribedEventAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
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

	return fmt.Errorf("data failed to match schemas in anyOf(SubscribedEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SubscribedEvent) MarshalJSON() ([]byte, error) {
	if src.SubscribedEventAnyOf != nil {
		return json.Marshal(&src.SubscribedEventAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSubscribedEvent struct {
	value *SubscribedEvent
	isSet bool
}

func (v NullableSubscribedEvent) Get() *SubscribedEvent {
	return v.value
}

func (v *NullableSubscribedEvent) Set(val *SubscribedEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscribedEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscribedEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscribedEvent(val *SubscribedEvent) *NullableSubscribedEvent {
	return &NullableSubscribedEvent{value: val, isSet: true}
}

func (v NullableSubscribedEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscribedEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


