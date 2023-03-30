/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// IpEthFlowDescription - Contains the description of an Uplink and/or Downlink Ethernet flow.
type IpEthFlowDescription struct {
	Interface{} *interface{}
}

// interface{}AsIpEthFlowDescription is a convenience function that returns interface{} wrapped in IpEthFlowDescription
func Interface{}AsIpEthFlowDescription(v *interface{}) IpEthFlowDescription {
	return IpEthFlowDescription{
		Interface{}: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *IpEthFlowDescription) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Interface{}
	err = newStrictDecoder(data).Decode(&dst.Interface{})
	if err == nil {
		jsonInterface{}, _ := json.Marshal(dst.Interface{})
		if string(jsonInterface{}) == "{}" { // empty struct
			dst.Interface{} = nil
		} else {
			match++
		}
	} else {
		dst.Interface{} = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Interface{} = nil

		return fmt.Errorf("data matches more than one schema in oneOf(IpEthFlowDescription)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(IpEthFlowDescription)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IpEthFlowDescription) MarshalJSON() ([]byte, error) {
	if src.Interface{} != nil {
		return json.Marshal(&src.Interface{})
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IpEthFlowDescription) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Interface{} != nil {
		return obj.Interface{}
	}

	// all schemas are nil
	return nil
}

type NullableIpEthFlowDescription struct {
	value *IpEthFlowDescription
	isSet bool
}

func (v NullableIpEthFlowDescription) Get() *IpEthFlowDescription {
	return v.value
}

func (v *NullableIpEthFlowDescription) Set(val *IpEthFlowDescription) {
	v.value = val
	v.isSet = true
}

func (v NullableIpEthFlowDescription) IsSet() bool {
	return v.isSet
}

func (v *NullableIpEthFlowDescription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpEthFlowDescription(val *IpEthFlowDescription) *NullableIpEthFlowDescription {
	return &NullableIpEthFlowDescription{value: val, isSet: true}
}

func (v NullableIpEthFlowDescription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpEthFlowDescription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


