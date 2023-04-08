/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
	"fmt"
)

// NrmEvent Possible values are: - UP_DELIVERY_MODE: User Plane delivery mode. 
type NrmEvent struct {
	NrmEventAnyOf *NrmEventAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *NrmEvent) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NrmEventAnyOf
	err = json.Unmarshal(data, &dst.NrmEventAnyOf);
	if err == nil {
		jsonNrmEventAnyOf, _ := json.Marshal(dst.NrmEventAnyOf)
		if string(jsonNrmEventAnyOf) == "{}" { // empty struct
			dst.NrmEventAnyOf = nil
		} else {
			return nil // data stored in dst.NrmEventAnyOf, return on the first match
		}
	} else {
		dst.NrmEventAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(NrmEvent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *NrmEvent) MarshalJSON() ([]byte, error) {
	if src.NrmEventAnyOf != nil {
		return json.Marshal(&src.NrmEventAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableNrmEvent struct {
	value *NrmEvent
	isSet bool
}

func (v NullableNrmEvent) Get() *NrmEvent {
	return v.value
}

func (v *NullableNrmEvent) Set(val *NrmEvent) {
	v.value = val
	v.isSet = true
}

func (v NullableNrmEvent) IsSet() bool {
	return v.isSet
}

func (v *NullableNrmEvent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNrmEvent(val *NrmEvent) *NullableNrmEvent {
	return &NullableNrmEvent{value: val, isSet: true}
}

func (v NullableNrmEvent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNrmEvent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


