/*
MDA NRM

OAS 3.0.1 specification of the MDA NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MdaNrm

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRrmNrType See details in 3GPP TS 32.422 clause 5.10.30.
type CollectionPeriodRrmNrType string

// List of collectionPeriodRrmNr-Type
const (
	_1024MS CollectionPeriodRrmNrType = "1024ms"
	_2048MS CollectionPeriodRrmNrType = "2048ms"
	_5120MS CollectionPeriodRrmNrType = "5120ms"
	_10240MS CollectionPeriodRrmNrType = "10240ms"
	_60000MS CollectionPeriodRrmNrType = "60000ms"
)

// All allowed values of CollectionPeriodRrmNrType enum
var AllowedCollectionPeriodRrmNrTypeEnumValues = []CollectionPeriodRrmNrType{
	"1024ms",
	"2048ms",
	"5120ms",
	"10240ms",
	"60000ms",
}

func (v *CollectionPeriodRrmNrType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodRrmNrType(value)
	for _, existing := range AllowedCollectionPeriodRrmNrTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodRrmNrType", value)
}

// NewCollectionPeriodRrmNrTypeFromValue returns a pointer to a valid CollectionPeriodRrmNrType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodRrmNrTypeFromValue(v string) (*CollectionPeriodRrmNrType, error) {
	ev := CollectionPeriodRrmNrType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodRrmNrType: valid values are %v", v, AllowedCollectionPeriodRrmNrTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodRrmNrType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodRrmNrTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodRrmNr-Type value
func (v CollectionPeriodRrmNrType) Ptr() *CollectionPeriodRrmNrType {
	return &v
}

type NullableCollectionPeriodRrmNrType struct {
	value *CollectionPeriodRrmNrType
	isSet bool
}

func (v NullableCollectionPeriodRrmNrType) Get() *CollectionPeriodRrmNrType {
	return v.value
}

func (v *NullableCollectionPeriodRrmNrType) Set(val *CollectionPeriodRrmNrType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRrmNrType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRrmNrType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRrmNrType(val *CollectionPeriodRrmNrType) *NullableCollectionPeriodRrmNrType {
	return &NullableCollectionPeriodRrmNrType{value: val, isSet: true}
}

func (v NullableCollectionPeriodRrmNrType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRrmNrType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

