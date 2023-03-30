/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// AfNotifMethod Represents the notification methods that can be subscribed for an event.
type AfNotifMethod struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AfNotifMethod) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(AfNotifMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AfNotifMethod) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAfNotifMethod struct {
	value *AfNotifMethod
	isSet bool
}

func (v NullableAfNotifMethod) Get() *AfNotifMethod {
	return v.value
}

func (v *NullableAfNotifMethod) Set(val *AfNotifMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableAfNotifMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableAfNotifMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfNotifMethod(val *AfNotifMethod) *NullableAfNotifMethod {
	return &NullableAfNotifMethod{value: val, isSet: true}
}

func (v NullableAfNotifMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfNotifMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

