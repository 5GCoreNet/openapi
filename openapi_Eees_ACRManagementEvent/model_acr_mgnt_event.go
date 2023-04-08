/*
EES ACR Management Event_API

API for EES ACR Management Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRManagementEvent

import (
	"encoding/json"
	"fmt"
)

// AcrMgntEvent Possible values are: - UP_PATH_CHG: User plane path change event. - ACR_MONITORING: ACR monitoring event. - ACR_FACILITATION: ACR facilitation event. - ACT_START_STOP: ACT start/stop event. 
type AcrMgntEvent struct {
	AcrMgntEventAnyOf *AcrMgntEventAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AcrMgntEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AcrMgntEventAnyOf
	err = json.Unmarshal(data, &dst.AcrMgntEventAnyOf);
	if err == nil {
		jsonAcrMgntEventAnyOf, _ := json.Marshal(dst.AcrMgntEventAnyOf)
		if string(jsonAcrMgntEventAnyOf) == "{}" { // empty struct
			dst.AcrMgntEventAnyOf = nil
		} else {
			return nil // data stored in dst.AcrMgntEventAnyOf, return on the first match
		}
	} else {
		dst.AcrMgntEventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AcrMgntEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AcrMgntEvent) MarshalJSON() ([]byte, error) {
	if src.AcrMgntEventAnyOf != nil {
		return json.Marshal(&src.AcrMgntEventAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAcrMgntEvent struct {
	value *AcrMgntEvent
	isSet bool
}

func (v NullableAcrMgntEvent) Get() *AcrMgntEvent {
	return v.value
}

func (v *NullableAcrMgntEvent) Set(val *AcrMgntEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableAcrMgntEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableAcrMgntEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcrMgntEvent(val *AcrMgntEvent) *NullableAcrMgntEvent {
	return &NullableAcrMgntEvent{value: val, isSet: true}
}

func (v NullableAcrMgntEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcrMgntEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


