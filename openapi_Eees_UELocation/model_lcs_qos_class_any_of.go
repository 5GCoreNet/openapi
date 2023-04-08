/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
	"fmt"
)

// LcsQosClassAnyOf the model 'LcsQosClassAnyOf'
type LcsQosClassAnyOf string

// List of LcsQosClass_anyOf
const (
	BEST_EFFORT LcsQosClassAnyOf = "BEST_EFFORT"
	ASSURED LcsQosClassAnyOf = "ASSURED"
	MULTIPLE_QOS LcsQosClassAnyOf = "MULTIPLE_QOS"
)

// All allowed values of LcsQosClassAnyOf enum
var AllowedLcsQosClassAnyOfEnumValues = []LcsQosClassAnyOf{
	"BEST_EFFORT",
	"ASSURED",
	"MULTIPLE_QOS",
}

func (v *LcsQosClassAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LcsQosClassAnyOf(value)
	for _, existing := range AllowedLcsQosClassAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LcsQosClassAnyOf", value)
}

// NewLcsQosClassAnyOfFromValue returns a pointer to a valid LcsQosClassAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLcsQosClassAnyOfFromValue(v string) (*LcsQosClassAnyOf, error) {
	ev := LcsQosClassAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LcsQosClassAnyOf: valid values are %v", v, AllowedLcsQosClassAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LcsQosClassAnyOf) IsValid() bool {
	for _, existing := range AllowedLcsQosClassAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LcsQosClass_anyOf value
func (v LcsQosClassAnyOf) Ptr() *LcsQosClassAnyOf {
	return &v
}

type NullableLcsQosClassAnyOf struct {
	value *LcsQosClassAnyOf
	isSet bool
}

func (v NullableLcsQosClassAnyOf) Get() *LcsQosClassAnyOf {
	return v.value
}

func (v *NullableLcsQosClassAnyOf) Set(val *LcsQosClassAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsQosClassAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsQosClassAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsQosClassAnyOf(val *LcsQosClassAnyOf) *NullableLcsQosClassAnyOf {
	return &NullableLcsQosClassAnyOf{value: val, isSet: true}
}

func (v NullableLcsQosClassAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsQosClassAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

