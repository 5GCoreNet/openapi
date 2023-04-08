/*
Naf_ProSe API

Naf_ProSe Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Naf_ProSe

import (
	"encoding/json"
	"fmt"
)

// MetadataIndicAnyOf the model 'MetadataIndicAnyOf'
type MetadataIndicAnyOf string

// List of MetadataIndic_anyOf
const (
	NO_METADATA MetadataIndicAnyOf = "NO_METADATA"
	METADATA_UPDATE_DISALLOWED MetadataIndicAnyOf = "METADATA_UPDATE_DISALLOWED"
	METADATA_UPDATE_ALLOWED MetadataIndicAnyOf = "METADATA_UPDATE_ALLOWED"
)

// All allowed values of MetadataIndicAnyOf enum
var AllowedMetadataIndicAnyOfEnumValues = []MetadataIndicAnyOf{
	"NO_METADATA",
	"METADATA_UPDATE_DISALLOWED",
	"METADATA_UPDATE_ALLOWED",
}

func (v *MetadataIndicAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetadataIndicAnyOf(value)
	for _, existing := range AllowedMetadataIndicAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetadataIndicAnyOf", value)
}

// NewMetadataIndicAnyOfFromValue returns a pointer to a valid MetadataIndicAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetadataIndicAnyOfFromValue(v string) (*MetadataIndicAnyOf, error) {
	ev := MetadataIndicAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetadataIndicAnyOf: valid values are %v", v, AllowedMetadataIndicAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetadataIndicAnyOf) IsValid() bool {
	for _, existing := range AllowedMetadataIndicAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetadataIndic_anyOf value
func (v MetadataIndicAnyOf) Ptr() *MetadataIndicAnyOf {
	return &v
}

type NullableMetadataIndicAnyOf struct {
	value *MetadataIndicAnyOf
	isSet bool
}

func (v NullableMetadataIndicAnyOf) Get() *MetadataIndicAnyOf {
	return v.value
}

func (v *NullableMetadataIndicAnyOf) Set(val *MetadataIndicAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableMetadataIndicAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableMetadataIndicAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetadataIndicAnyOf(val *MetadataIndicAnyOf) *NullableMetadataIndicAnyOf {
	return &NullableMetadataIndicAnyOf{value: val, isSet: true}
}

func (v NullableMetadataIndicAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetadataIndicAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
