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

// OriginatorPartyTypeAnyOf the model 'OriginatorPartyTypeAnyOf'
type OriginatorPartyTypeAnyOf string

// List of OriginatorPartyType_anyOf
const (
	CALLING OriginatorPartyTypeAnyOf = "CALLING"
	CALLED OriginatorPartyTypeAnyOf = "CALLED"
)

// All allowed values of OriginatorPartyTypeAnyOf enum
var AllowedOriginatorPartyTypeAnyOfEnumValues = []OriginatorPartyTypeAnyOf{
	"CALLING",
	"CALLED",
}

func (v *OriginatorPartyTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OriginatorPartyTypeAnyOf(value)
	for _, existing := range AllowedOriginatorPartyTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OriginatorPartyTypeAnyOf", value)
}

// NewOriginatorPartyTypeAnyOfFromValue returns a pointer to a valid OriginatorPartyTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOriginatorPartyTypeAnyOfFromValue(v string) (*OriginatorPartyTypeAnyOf, error) {
	ev := OriginatorPartyTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OriginatorPartyTypeAnyOf: valid values are %v", v, AllowedOriginatorPartyTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OriginatorPartyTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedOriginatorPartyTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OriginatorPartyType_anyOf value
func (v OriginatorPartyTypeAnyOf) Ptr() *OriginatorPartyTypeAnyOf {
	return &v
}

type NullableOriginatorPartyTypeAnyOf struct {
	value *OriginatorPartyTypeAnyOf
	isSet bool
}

func (v NullableOriginatorPartyTypeAnyOf) Get() *OriginatorPartyTypeAnyOf {
	return v.value
}

func (v *NullableOriginatorPartyTypeAnyOf) Set(val *OriginatorPartyTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableOriginatorPartyTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableOriginatorPartyTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOriginatorPartyTypeAnyOf(val *OriginatorPartyTypeAnyOf) *NullableOriginatorPartyTypeAnyOf {
	return &NullableOriginatorPartyTypeAnyOf{value: val, isSet: true}
}

func (v NullableOriginatorPartyTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOriginatorPartyTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

