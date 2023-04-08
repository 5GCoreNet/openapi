/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// DataIndAnyOf the model 'DataIndAnyOf'
type DataIndAnyOf string

// List of DataInd_anyOf
const (
	PFD DataIndAnyOf = "PFD"
	IPTV DataIndAnyOf = "IPTV"
	BDT DataIndAnyOf = "BDT"
	SVC_PARAM DataIndAnyOf = "SVC_PARAM"
	AM DataIndAnyOf = "AM"
)

// All allowed values of DataIndAnyOf enum
var AllowedDataIndAnyOfEnumValues = []DataIndAnyOf{
	"PFD",
	"IPTV",
	"BDT",
	"SVC_PARAM",
	"AM",
}

func (v *DataIndAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DataIndAnyOf(value)
	for _, existing := range AllowedDataIndAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DataIndAnyOf", value)
}

// NewDataIndAnyOfFromValue returns a pointer to a valid DataIndAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDataIndAnyOfFromValue(v string) (*DataIndAnyOf, error) {
	ev := DataIndAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DataIndAnyOf: valid values are %v", v, AllowedDataIndAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DataIndAnyOf) IsValid() bool {
	for _, existing := range AllowedDataIndAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DataInd_anyOf value
func (v DataIndAnyOf) Ptr() *DataIndAnyOf {
	return &v
}

type NullableDataIndAnyOf struct {
	value *DataIndAnyOf
	isSet bool
}

func (v NullableDataIndAnyOf) Get() *DataIndAnyOf {
	return v.value
}

func (v *NullableDataIndAnyOf) Set(val *DataIndAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDataIndAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDataIndAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataIndAnyOf(val *DataIndAnyOf) *NullableDataIndAnyOf {
	return &NullableDataIndAnyOf{value: val, isSet: true}
}

func (v NullableDataIndAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataIndAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

