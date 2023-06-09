/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// MulticastAccessControl Indicates whether the service data flow, corresponding to the service data flow template, is  allowed or not allowed.
type MulticastAccessControl struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MulticastAccessControl) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(MulticastAccessControl)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MulticastAccessControl) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMulticastAccessControl struct {
	value *MulticastAccessControl
	isSet bool
}

func (v NullableMulticastAccessControl) Get() *MulticastAccessControl {
	return v.value
}

func (v *NullableMulticastAccessControl) Set(val *MulticastAccessControl) {
	v.value = val
	v.isSet = true
}

func (v NullableMulticastAccessControl) IsSet() bool {
	return v.isSet
}

func (v *NullableMulticastAccessControl) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMulticastAccessControl(val *MulticastAccessControl) *NullableMulticastAccessControl {
	return &NullableMulticastAccessControl{value: val, isSet: true}
}

func (v NullableMulticastAccessControl) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMulticastAccessControl) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
