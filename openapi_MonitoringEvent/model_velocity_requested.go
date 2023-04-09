/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
	"fmt"
)

// VelocityRequested Indicates velocity requirement.
type VelocityRequested struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *VelocityRequested) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(VelocityRequested)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *VelocityRequested) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableVelocityRequested struct {
	value *VelocityRequested
	isSet bool
}

func (v NullableVelocityRequested) Get() *VelocityRequested {
	return v.value
}

func (v *NullableVelocityRequested) Set(val *VelocityRequested) {
	v.value = val
	v.isSet = true
}

func (v NullableVelocityRequested) IsSet() bool {
	return v.isSet
}

func (v *NullableVelocityRequested) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelocityRequested(val *VelocityRequested) *NullableVelocityRequested {
	return &NullableVelocityRequested{value: val, isSet: true}
}

func (v NullableVelocityRequested) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelocityRequested) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
