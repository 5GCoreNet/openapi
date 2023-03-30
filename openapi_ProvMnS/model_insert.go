/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// Insert the model 'Insert'
type Insert string

// List of Insert
const (
	BEFORE Insert = "before"
	AFTER Insert = "after"
)

// All allowed values of Insert enum
var AllowedInsertEnumValues = []Insert{
	"before",
	"after",
}

func (v *Insert) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Insert(value)
	for _, existing := range AllowedInsertEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Insert", value)
}

// NewInsertFromValue returns a pointer to a valid Insert
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewInsertFromValue(v string) (*Insert, error) {
	ev := Insert(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Insert: valid values are %v", v, AllowedInsertEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Insert) IsValid() bool {
	for _, existing := range AllowedInsertEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Insert value
func (v Insert) Ptr() *Insert {
	return &v
}

type NullableInsert struct {
	value *Insert
	isSet bool
}

func (v NullableInsert) Get() *Insert {
	return v.value
}

func (v *NullableInsert) Set(val *Insert) {
	v.value = val
	v.isSet = true
}

func (v NullableInsert) IsSet() bool {
	return v.isSet
}

func (v *NullableInsert) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInsert(val *Insert) *NullableInsert {
	return &NullableInsert{value: val, isSet: true}
}

func (v NullableInsert) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInsert) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
