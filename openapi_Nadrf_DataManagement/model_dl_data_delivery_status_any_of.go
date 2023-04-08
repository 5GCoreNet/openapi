/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// DlDataDeliveryStatusAnyOf the model 'DlDataDeliveryStatusAnyOf'
type DlDataDeliveryStatusAnyOf string

// List of DlDataDeliveryStatus_anyOf
const (
	BUFFERED DlDataDeliveryStatusAnyOf = "BUFFERED"
	TRANSMITTED DlDataDeliveryStatusAnyOf = "TRANSMITTED"
	DISCARDED DlDataDeliveryStatusAnyOf = "DISCARDED"
)

// All allowed values of DlDataDeliveryStatusAnyOf enum
var AllowedDlDataDeliveryStatusAnyOfEnumValues = []DlDataDeliveryStatusAnyOf{
	"BUFFERED",
	"TRANSMITTED",
	"DISCARDED",
}

func (v *DlDataDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DlDataDeliveryStatusAnyOf(value)
	for _, existing := range AllowedDlDataDeliveryStatusAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DlDataDeliveryStatusAnyOf", value)
}

// NewDlDataDeliveryStatusAnyOfFromValue returns a pointer to a valid DlDataDeliveryStatusAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDlDataDeliveryStatusAnyOfFromValue(v string) (*DlDataDeliveryStatusAnyOf, error) {
	ev := DlDataDeliveryStatusAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DlDataDeliveryStatusAnyOf: valid values are %v", v, AllowedDlDataDeliveryStatusAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DlDataDeliveryStatusAnyOf) IsValid() bool {
	for _, existing := range AllowedDlDataDeliveryStatusAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DlDataDeliveryStatus_anyOf value
func (v DlDataDeliveryStatusAnyOf) Ptr() *DlDataDeliveryStatusAnyOf {
	return &v
}

type NullableDlDataDeliveryStatusAnyOf struct {
	value *DlDataDeliveryStatusAnyOf
	isSet bool
}

func (v NullableDlDataDeliveryStatusAnyOf) Get() *DlDataDeliveryStatusAnyOf {
	return v.value
}

func (v *NullableDlDataDeliveryStatusAnyOf) Set(val *DlDataDeliveryStatusAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDlDataDeliveryStatusAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDlDataDeliveryStatusAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDlDataDeliveryStatusAnyOf(val *DlDataDeliveryStatusAnyOf) *NullableDlDataDeliveryStatusAnyOf {
	return &NullableDlDataDeliveryStatusAnyOf{value: val, isSet: true}
}

func (v NullableDlDataDeliveryStatusAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDlDataDeliveryStatusAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

