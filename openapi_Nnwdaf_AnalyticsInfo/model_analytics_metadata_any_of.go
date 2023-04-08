/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// AnalyticsMetadataAnyOf the model 'AnalyticsMetadataAnyOf'
type AnalyticsMetadataAnyOf string

// List of AnalyticsMetadata_anyOf
const (
	NUM_OF_SAMPLES AnalyticsMetadataAnyOf = "NUM_OF_SAMPLES"
	DATA_WINDOW AnalyticsMetadataAnyOf = "DATA_WINDOW"
	DATA_STAT_PROPS AnalyticsMetadataAnyOf = "DATA_STAT_PROPS"
	STRATEGY AnalyticsMetadataAnyOf = "STRATEGY"
	ACCURACY AnalyticsMetadataAnyOf = "ACCURACY"
)

// All allowed values of AnalyticsMetadataAnyOf enum
var AllowedAnalyticsMetadataAnyOfEnumValues = []AnalyticsMetadataAnyOf{
	"NUM_OF_SAMPLES",
	"DATA_WINDOW",
	"DATA_STAT_PROPS",
	"STRATEGY",
	"ACCURACY",
}

func (v *AnalyticsMetadataAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnalyticsMetadataAnyOf(value)
	for _, existing := range AllowedAnalyticsMetadataAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnalyticsMetadataAnyOf", value)
}

// NewAnalyticsMetadataAnyOfFromValue returns a pointer to a valid AnalyticsMetadataAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnalyticsMetadataAnyOfFromValue(v string) (*AnalyticsMetadataAnyOf, error) {
	ev := AnalyticsMetadataAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnalyticsMetadataAnyOf: valid values are %v", v, AllowedAnalyticsMetadataAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnalyticsMetadataAnyOf) IsValid() bool {
	for _, existing := range AllowedAnalyticsMetadataAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AnalyticsMetadata_anyOf value
func (v AnalyticsMetadataAnyOf) Ptr() *AnalyticsMetadataAnyOf {
	return &v
}

type NullableAnalyticsMetadataAnyOf struct {
	value *AnalyticsMetadataAnyOf
	isSet bool
}

func (v NullableAnalyticsMetadataAnyOf) Get() *AnalyticsMetadataAnyOf {
	return v.value
}

func (v *NullableAnalyticsMetadataAnyOf) Set(val *AnalyticsMetadataAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAnalyticsMetadataAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAnalyticsMetadataAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnalyticsMetadataAnyOf(val *AnalyticsMetadataAnyOf) *NullableAnalyticsMetadataAnyOf {
	return &NullableAnalyticsMetadataAnyOf{value: val, isSet: true}
}

func (v NullableAnalyticsMetadataAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnalyticsMetadataAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
