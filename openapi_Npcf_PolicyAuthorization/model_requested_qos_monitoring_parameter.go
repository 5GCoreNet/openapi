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

// RequestedQosMonitoringParameter Indicates the requested QoS monitoring parameters to be measured.
type RequestedQosMonitoringParameter struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RequestedQosMonitoringParameter) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RequestedQosMonitoringParameter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RequestedQosMonitoringParameter) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRequestedQosMonitoringParameter struct {
	value *RequestedQosMonitoringParameter
	isSet bool
}

func (v NullableRequestedQosMonitoringParameter) Get() *RequestedQosMonitoringParameter {
	return v.value
}

func (v *NullableRequestedQosMonitoringParameter) Set(val *RequestedQosMonitoringParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedQosMonitoringParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedQosMonitoringParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedQosMonitoringParameter(val *RequestedQosMonitoringParameter) *NullableRequestedQosMonitoringParameter {
	return &NullableRequestedQosMonitoringParameter{value: val, isSet: true}
}

func (v NullableRequestedQosMonitoringParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedQosMonitoringParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

