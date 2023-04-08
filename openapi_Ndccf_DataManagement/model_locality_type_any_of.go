/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// LocalityTypeAnyOf the model 'LocalityTypeAnyOf'
type LocalityTypeAnyOf string

// List of LocalityType_anyOf
const (
	DATA_CENTER LocalityTypeAnyOf = "DATA_CENTER"
	CITY LocalityTypeAnyOf = "CITY"
	COUNTY LocalityTypeAnyOf = "COUNTY"
	DISTRICT LocalityTypeAnyOf = "DISTRICT"
	STATE LocalityTypeAnyOf = "STATE"
	CANTON LocalityTypeAnyOf = "CANTON"
	REGION LocalityTypeAnyOf = "REGION"
	PROVINCE LocalityTypeAnyOf = "PROVINCE"
	PREFECTURE LocalityTypeAnyOf = "PREFECTURE"
	COUNTRY LocalityTypeAnyOf = "COUNTRY"
)

// All allowed values of LocalityTypeAnyOf enum
var AllowedLocalityTypeAnyOfEnumValues = []LocalityTypeAnyOf{
	"DATA_CENTER",
	"CITY",
	"COUNTY",
	"DISTRICT",
	"STATE",
	"CANTON",
	"REGION",
	"PROVINCE",
	"PREFECTURE",
	"COUNTRY",
}

func (v *LocalityTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LocalityTypeAnyOf(value)
	for _, existing := range AllowedLocalityTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LocalityTypeAnyOf", value)
}

// NewLocalityTypeAnyOfFromValue returns a pointer to a valid LocalityTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLocalityTypeAnyOfFromValue(v string) (*LocalityTypeAnyOf, error) {
	ev := LocalityTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LocalityTypeAnyOf: valid values are %v", v, AllowedLocalityTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LocalityTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedLocalityTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LocalityType_anyOf value
func (v LocalityTypeAnyOf) Ptr() *LocalityTypeAnyOf {
	return &v
}

type NullableLocalityTypeAnyOf struct {
	value *LocalityTypeAnyOf
	isSet bool
}

func (v NullableLocalityTypeAnyOf) Get() *LocalityTypeAnyOf {
	return v.value
}

func (v *NullableLocalityTypeAnyOf) Set(val *LocalityTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalityTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalityTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalityTypeAnyOf(val *LocalityTypeAnyOf) *NullableLocalityTypeAnyOf {
	return &NullableLocalityTypeAnyOf{value: val, isSet: true}
}

func (v NullableLocalityTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalityTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

