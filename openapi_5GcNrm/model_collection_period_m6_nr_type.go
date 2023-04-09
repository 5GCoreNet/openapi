/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodM6NrType See details in 3GPP TS 32.422 clause 5.10.34.
type CollectionPeriodM6NrType string

// List of collectionPeriodM6Nr-Type
const (
	_120MS   CollectionPeriodM6NrType = "120ms"
	_240MS   CollectionPeriodM6NrType = "240ms"
	_480MS   CollectionPeriodM6NrType = "480ms"
	_640MS   CollectionPeriodM6NrType = "640ms"
	_1024MS  CollectionPeriodM6NrType = "1024ms"
	_2048MS  CollectionPeriodM6NrType = "2048ms"
	_5120MS  CollectionPeriodM6NrType = "5120ms"
	_10240MS CollectionPeriodM6NrType = "10240ms"
	_20480MS CollectionPeriodM6NrType = "20480ms"
	_40960MS CollectionPeriodM6NrType = "40960ms"
	_1MIN    CollectionPeriodM6NrType = "1min"
	_6MIN    CollectionPeriodM6NrType = "6min"
	_12MIN   CollectionPeriodM6NrType = "12min"
	_30MIN   CollectionPeriodM6NrType = "30min"
)

// All allowed values of CollectionPeriodM6NrType enum
var AllowedCollectionPeriodM6NrTypeEnumValues = []CollectionPeriodM6NrType{
	"120ms",
	"240ms",
	"480ms",
	"640ms",
	"1024ms",
	"2048ms",
	"5120ms",
	"10240ms",
	"20480ms",
	"40960ms",
	"1min",
	"6min",
	"12min",
	"30min",
}

func (v *CollectionPeriodM6NrType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodM6NrType(value)
	for _, existing := range AllowedCollectionPeriodM6NrTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodM6NrType", value)
}

// NewCollectionPeriodM6NrTypeFromValue returns a pointer to a valid CollectionPeriodM6NrType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodM6NrTypeFromValue(v string) (*CollectionPeriodM6NrType, error) {
	ev := CollectionPeriodM6NrType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodM6NrType: valid values are %v", v, AllowedCollectionPeriodM6NrTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodM6NrType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodM6NrTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodM6Nr-Type value
func (v CollectionPeriodM6NrType) Ptr() *CollectionPeriodM6NrType {
	return &v
}

type NullableCollectionPeriodM6NrType struct {
	value *CollectionPeriodM6NrType
	isSet bool
}

func (v NullableCollectionPeriodM6NrType) Get() *CollectionPeriodM6NrType {
	return v.value
}

func (v *NullableCollectionPeriodM6NrType) Set(val *CollectionPeriodM6NrType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodM6NrType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodM6NrType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodM6NrType(val *CollectionPeriodM6NrType) *NullableCollectionPeriodM6NrType {
	return &NullableCollectionPeriodM6NrType{value: val, isSet: true}
}

func (v NullableCollectionPeriodM6NrType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodM6NrType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
