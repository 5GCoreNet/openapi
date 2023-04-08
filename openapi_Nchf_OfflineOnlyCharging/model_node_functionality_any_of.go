/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// NodeFunctionalityAnyOf the model 'NodeFunctionalityAnyOf'
type NodeFunctionalityAnyOf string

// List of NodeFunctionality_anyOf
const (
	SMF NodeFunctionalityAnyOf = "SMF"
	SMSF NodeFunctionalityAnyOf = "SMSF"
	I_SMF NodeFunctionalityAnyOf = "I-SMF"
)

// All allowed values of NodeFunctionalityAnyOf enum
var AllowedNodeFunctionalityAnyOfEnumValues = []NodeFunctionalityAnyOf{
	"SMF",
	"SMSF",
	"I-SMF",
}

func (v *NodeFunctionalityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NodeFunctionalityAnyOf(value)
	for _, existing := range AllowedNodeFunctionalityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NodeFunctionalityAnyOf", value)
}

// NewNodeFunctionalityAnyOfFromValue returns a pointer to a valid NodeFunctionalityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNodeFunctionalityAnyOfFromValue(v string) (*NodeFunctionalityAnyOf, error) {
	ev := NodeFunctionalityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NodeFunctionalityAnyOf: valid values are %v", v, AllowedNodeFunctionalityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NodeFunctionalityAnyOf) IsValid() bool {
	for _, existing := range AllowedNodeFunctionalityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NodeFunctionality_anyOf value
func (v NodeFunctionalityAnyOf) Ptr() *NodeFunctionalityAnyOf {
	return &v
}

type NullableNodeFunctionalityAnyOf struct {
	value *NodeFunctionalityAnyOf
	isSet bool
}

func (v NullableNodeFunctionalityAnyOf) Get() *NodeFunctionalityAnyOf {
	return v.value
}

func (v *NullableNodeFunctionalityAnyOf) Set(val *NodeFunctionalityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableNodeFunctionalityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableNodeFunctionalityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNodeFunctionalityAnyOf(val *NodeFunctionalityAnyOf) *NullableNodeFunctionalityAnyOf {
	return &NullableNodeFunctionalityAnyOf{value: val, isSet: true}
}

func (v NullableNodeFunctionalityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNodeFunctionalityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

