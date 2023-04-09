/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// TrafficForwardingWay struct for TrafficForwardingWay
type TrafficForwardingWay struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TrafficForwardingWay) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TrafficForwardingWay)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TrafficForwardingWay) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTrafficForwardingWay struct {
	value *TrafficForwardingWay
	isSet bool
}

func (v NullableTrafficForwardingWay) Get() *TrafficForwardingWay {
	return v.value
}

func (v *NullableTrafficForwardingWay) Set(val *TrafficForwardingWay) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficForwardingWay) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficForwardingWay) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficForwardingWay(val *TrafficForwardingWay) *NullableTrafficForwardingWay {
	return &NullableTrafficForwardingWay{value: val, isSet: true}
}

func (v NullableTrafficForwardingWay) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficForwardingWay) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
