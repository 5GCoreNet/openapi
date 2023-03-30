/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
	"fmt"
)

// PerceivedSeverity the model 'PerceivedSeverity'
type PerceivedSeverity string

// List of PerceivedSeverity
const (
	INDETERMINATE PerceivedSeverity = "INDETERMINATE"
	CRITICAL PerceivedSeverity = "CRITICAL"
	MAJOR PerceivedSeverity = "MAJOR"
	MINOR PerceivedSeverity = "MINOR"
	WARNING PerceivedSeverity = "WARNING"
	CLEARED PerceivedSeverity = "CLEARED"
)

// All allowed values of PerceivedSeverity enum
var AllowedPerceivedSeverityEnumValues = []PerceivedSeverity{
	"INDETERMINATE",
	"CRITICAL",
	"MAJOR",
	"MINOR",
	"WARNING",
	"CLEARED",
}

func (v *PerceivedSeverity) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PerceivedSeverity(value)
	for _, existing := range AllowedPerceivedSeverityEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PerceivedSeverity", value)
}

// NewPerceivedSeverityFromValue returns a pointer to a valid PerceivedSeverity
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPerceivedSeverityFromValue(v string) (*PerceivedSeverity, error) {
	ev := PerceivedSeverity(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PerceivedSeverity: valid values are %v", v, AllowedPerceivedSeverityEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PerceivedSeverity) IsValid() bool {
	for _, existing := range AllowedPerceivedSeverityEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PerceivedSeverity value
func (v PerceivedSeverity) Ptr() *PerceivedSeverity {
	return &v
}

type NullablePerceivedSeverity struct {
	value *PerceivedSeverity
	isSet bool
}

func (v NullablePerceivedSeverity) Get() *PerceivedSeverity {
	return v.value
}

func (v *NullablePerceivedSeverity) Set(val *PerceivedSeverity) {
	v.value = val
	v.isSet = true
}

func (v NullablePerceivedSeverity) IsSet() bool {
	return v.isSet
}

func (v *NullablePerceivedSeverity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePerceivedSeverity(val *PerceivedSeverity) *NullablePerceivedSeverity {
	return &NullablePerceivedSeverity{value: val, isSet: true}
}

func (v NullablePerceivedSeverity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePerceivedSeverity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

