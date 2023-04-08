/*
3gpp-as-session-with-qos

API for setting us an AS session with required QoS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AsSessionWithQoS

import (
	"encoding/json"
	"fmt"
)

// RequestedQosMonitoringParameterAnyOf the model 'RequestedQosMonitoringParameterAnyOf'
type RequestedQosMonitoringParameterAnyOf string

// List of RequestedQosMonitoringParameter_anyOf
const (
	DOWNLINK RequestedQosMonitoringParameterAnyOf = "DOWNLINK"
	UPLINK RequestedQosMonitoringParameterAnyOf = "UPLINK"
	ROUND_TRIP RequestedQosMonitoringParameterAnyOf = "ROUND_TRIP"
)

// All allowed values of RequestedQosMonitoringParameterAnyOf enum
var AllowedRequestedQosMonitoringParameterAnyOfEnumValues = []RequestedQosMonitoringParameterAnyOf{
	"DOWNLINK",
	"UPLINK",
	"ROUND_TRIP",
}

func (v *RequestedQosMonitoringParameterAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RequestedQosMonitoringParameterAnyOf(value)
	for _, existing := range AllowedRequestedQosMonitoringParameterAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RequestedQosMonitoringParameterAnyOf", value)
}

// NewRequestedQosMonitoringParameterAnyOfFromValue returns a pointer to a valid RequestedQosMonitoringParameterAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRequestedQosMonitoringParameterAnyOfFromValue(v string) (*RequestedQosMonitoringParameterAnyOf, error) {
	ev := RequestedQosMonitoringParameterAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RequestedQosMonitoringParameterAnyOf: valid values are %v", v, AllowedRequestedQosMonitoringParameterAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RequestedQosMonitoringParameterAnyOf) IsValid() bool {
	for _, existing := range AllowedRequestedQosMonitoringParameterAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RequestedQosMonitoringParameter_anyOf value
func (v RequestedQosMonitoringParameterAnyOf) Ptr() *RequestedQosMonitoringParameterAnyOf {
	return &v
}

type NullableRequestedQosMonitoringParameterAnyOf struct {
	value *RequestedQosMonitoringParameterAnyOf
	isSet bool
}

func (v NullableRequestedQosMonitoringParameterAnyOf) Get() *RequestedQosMonitoringParameterAnyOf {
	return v.value
}

func (v *NullableRequestedQosMonitoringParameterAnyOf) Set(val *RequestedQosMonitoringParameterAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestedQosMonitoringParameterAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestedQosMonitoringParameterAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestedQosMonitoringParameterAnyOf(val *RequestedQosMonitoringParameterAnyOf) *NullableRequestedQosMonitoringParameterAnyOf {
	return &NullableRequestedQosMonitoringParameterAnyOf{value: val, isSet: true}
}

func (v NullableRequestedQosMonitoringParameterAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestedQosMonitoringParameterAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

