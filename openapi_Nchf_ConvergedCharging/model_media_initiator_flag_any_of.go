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

// MediaInitiatorFlagAnyOf the model 'MediaInitiatorFlagAnyOf'
type MediaInitiatorFlagAnyOf string

// List of MediaInitiatorFlag_anyOf
const (
	CALLED_PARTY MediaInitiatorFlagAnyOf = "CALLED_PARTY"
	CALLING_PARTY MediaInitiatorFlagAnyOf = "CALLING_PARTY"
	UNKNOWN MediaInitiatorFlagAnyOf = "UNKNOWN"
)

// All allowed values of MediaInitiatorFlagAnyOf enum
var AllowedMediaInitiatorFlagAnyOfEnumValues = []MediaInitiatorFlagAnyOf{
	"CALLED_PARTY",
	"CALLING_PARTY",
	"UNKNOWN",
}

func (v *MediaInitiatorFlagAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MediaInitiatorFlagAnyOf(value)
	for _, existing := range AllowedMediaInitiatorFlagAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MediaInitiatorFlagAnyOf", value)
}

// NewMediaInitiatorFlagAnyOfFromValue returns a pointer to a valid MediaInitiatorFlagAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMediaInitiatorFlagAnyOfFromValue(v string) (*MediaInitiatorFlagAnyOf, error) {
	ev := MediaInitiatorFlagAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MediaInitiatorFlagAnyOf: valid values are %v", v, AllowedMediaInitiatorFlagAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MediaInitiatorFlagAnyOf) IsValid() bool {
	for _, existing := range AllowedMediaInitiatorFlagAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MediaInitiatorFlag_anyOf value
func (v MediaInitiatorFlagAnyOf) Ptr() *MediaInitiatorFlagAnyOf {
	return &v
}

type NullableMediaInitiatorFlagAnyOf struct {
	value *MediaInitiatorFlagAnyOf
	isSet bool
}

func (v NullableMediaInitiatorFlagAnyOf) Get() *MediaInitiatorFlagAnyOf {
	return v.value
}

func (v *NullableMediaInitiatorFlagAnyOf) Set(val *MediaInitiatorFlagAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMediaInitiatorFlagAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMediaInitiatorFlagAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMediaInitiatorFlagAnyOf(val *MediaInitiatorFlagAnyOf) *NullableMediaInitiatorFlagAnyOf {
	return &NullableMediaInitiatorFlagAnyOf{value: val, isSet: true}
}

func (v NullableMediaInitiatorFlagAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMediaInitiatorFlagAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
