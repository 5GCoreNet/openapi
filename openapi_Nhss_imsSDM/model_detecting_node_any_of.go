/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// DetectingNodeAnyOf the model 'DetectingNodeAnyOf'
type DetectingNodeAnyOf string

// List of DetectingNode_anyOf
const (
	SGSN DetectingNodeAnyOf = "SGSN"
	MME DetectingNodeAnyOf = "MME"
	AMF DetectingNodeAnyOf = "AMF"
)

// All allowed values of DetectingNodeAnyOf enum
var AllowedDetectingNodeAnyOfEnumValues = []DetectingNodeAnyOf{
	"SGSN",
	"MME",
	"AMF",
}

func (v *DetectingNodeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DetectingNodeAnyOf(value)
	for _, existing := range AllowedDetectingNodeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DetectingNodeAnyOf", value)
}

// NewDetectingNodeAnyOfFromValue returns a pointer to a valid DetectingNodeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDetectingNodeAnyOfFromValue(v string) (*DetectingNodeAnyOf, error) {
	ev := DetectingNodeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DetectingNodeAnyOf: valid values are %v", v, AllowedDetectingNodeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DetectingNodeAnyOf) IsValid() bool {
	for _, existing := range AllowedDetectingNodeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DetectingNode_anyOf value
func (v DetectingNodeAnyOf) Ptr() *DetectingNodeAnyOf {
	return &v
}

type NullableDetectingNodeAnyOf struct {
	value *DetectingNodeAnyOf
	isSet bool
}

func (v NullableDetectingNodeAnyOf) Get() *DetectingNodeAnyOf {
	return v.value
}

func (v *NullableDetectingNodeAnyOf) Set(val *DetectingNodeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDetectingNodeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDetectingNodeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDetectingNodeAnyOf(val *DetectingNodeAnyOf) *NullableDetectingNodeAnyOf {
	return &NullableDetectingNodeAnyOf{value: val, isSet: true}
}

func (v NullableDetectingNodeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDetectingNodeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

