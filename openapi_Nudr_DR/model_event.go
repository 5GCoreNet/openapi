/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// Event Possible values are: - SUCCESS_UE_POL_DEL_SP: Successful UE Policy Delivery related to    the invocation of AF provisioned Service Parameters. - UNSUCCESS_UE_POL_DEL_SP: Unsuccessful UE Policy Delivery related to the invocation of AF    provisioned Service Parameters. 
type Event struct {
	EventAnyOf *EventAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Event) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into EventAnyOf
	err = json.Unmarshal(data, &dst.EventAnyOf);
	if err == nil {
		jsonEventAnyOf, _ := json.Marshal(dst.EventAnyOf)
		if string(jsonEventAnyOf) == "{}" { // empty struct
			dst.EventAnyOf = nil
		} else {
			return nil // data stored in dst.EventAnyOf, return on the first match
		}
	} else {
		dst.EventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(Event)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Event) MarshalJSON() ([]byte, error) {
	if src.EventAnyOf != nil {
		return json.Marshal(&src.EventAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableEvent struct {
	value *Event
	isSet bool
}

func (v NullableEvent) Get() *Event {
	return v.value
}

func (v *NullableEvent) Set(val *Event) {
	v.value = val
	v.isSet = true
}

func (v NullableEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEvent(val *Event) *NullableEvent {
	return &NullableEvent{value: val, isSet: true}
}

func (v NullableEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


