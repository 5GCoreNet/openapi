/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
	"fmt"
)

// BwpContext the model 'BwpContext'
type BwpContext string

// List of BwpContext
const (
	DL BwpContext = "DL"
	UL BwpContext = "UL"
	SUL BwpContext = "SUL"
)

// All allowed values of BwpContext enum
var AllowedBwpContextEnumValues = []BwpContext{
	"DL",
	"UL",
	"SUL",
}

func (v *BwpContext) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BwpContext(value)
	for _, existing := range AllowedBwpContextEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BwpContext", value)
}

// NewBwpContextFromValue returns a pointer to a valid BwpContext
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBwpContextFromValue(v string) (*BwpContext, error) {
	ev := BwpContext(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BwpContext: valid values are %v", v, AllowedBwpContextEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BwpContext) IsValid() bool {
	for _, existing := range AllowedBwpContextEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BwpContext value
func (v BwpContext) Ptr() *BwpContext {
	return &v
}

type NullableBwpContext struct {
	value *BwpContext
	isSet bool
}

func (v NullableBwpContext) Get() *BwpContext {
	return v.value
}

func (v *NullableBwpContext) Set(val *BwpContext) {
	v.value = val
	v.isSet = true
}

func (v NullableBwpContext) IsSet() bool {
	return v.isSet
}

func (v *NullableBwpContext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBwpContext(val *BwpContext) *NullableBwpContext {
	return &NullableBwpContext{value: val, isSet: true}
}

func (v NullableBwpContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBwpContext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

