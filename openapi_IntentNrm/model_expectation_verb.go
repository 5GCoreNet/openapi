/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
	"fmt"
)

// ExpectationVerb the model 'ExpectationVerb'
type ExpectationVerb string

// List of ExpectationVerb
const (
	DELIVER ExpectationVerb = "DELIVER"
	ENSURE  ExpectationVerb = "ENSURE"
)

// All allowed values of ExpectationVerb enum
var AllowedExpectationVerbEnumValues = []ExpectationVerb{
	"DELIVER",
	"ENSURE",
}

func (v *ExpectationVerb) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExpectationVerb(value)
	for _, existing := range AllowedExpectationVerbEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExpectationVerb", value)
}

// NewExpectationVerbFromValue returns a pointer to a valid ExpectationVerb
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExpectationVerbFromValue(v string) (*ExpectationVerb, error) {
	ev := ExpectationVerb(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExpectationVerb: valid values are %v", v, AllowedExpectationVerbEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExpectationVerb) IsValid() bool {
	for _, existing := range AllowedExpectationVerbEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ExpectationVerb value
func (v ExpectationVerb) Ptr() *ExpectationVerb {
	return &v
}

type NullableExpectationVerb struct {
	value *ExpectationVerb
	isSet bool
}

func (v NullableExpectationVerb) Get() *ExpectationVerb {
	return v.value
}

func (v *NullableExpectationVerb) Set(val *ExpectationVerb) {
	v.value = val
	v.isSet = true
}

func (v NullableExpectationVerb) IsSet() bool {
	return v.isSet
}

func (v *NullableExpectationVerb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExpectationVerb(val *ExpectationVerb) *NullableExpectationVerb {
	return &NullableExpectationVerb{value: val, isSet: true}
}

func (v NullableExpectationVerb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExpectationVerb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
