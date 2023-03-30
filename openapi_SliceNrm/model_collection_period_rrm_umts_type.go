/*
Slice NRM

OAS 3.0.1 specification of the Slice NRM @ 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SliceNrm

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodRrmUmtsType See details in 3GPP TS 32.422 clause 5.10.21.
type CollectionPeriodRrmUmtsType string

// List of collectionPeriodRrmUmts-Type
const (
	_100MS CollectionPeriodRrmUmtsType = "100ms"
	_250MS CollectionPeriodRrmUmtsType = "250ms"
	_500MS CollectionPeriodRrmUmtsType = "500ms"
	_1000MS CollectionPeriodRrmUmtsType = "1000ms"
	_2000MS CollectionPeriodRrmUmtsType = "2000ms"
	_3000MS CollectionPeriodRrmUmtsType = "3000ms"
	_4000MS CollectionPeriodRrmUmtsType = "4000ms"
	_6000MS CollectionPeriodRrmUmtsType = "6000ms"
)

// All allowed values of CollectionPeriodRrmUmtsType enum
var AllowedCollectionPeriodRrmUmtsTypeEnumValues = []CollectionPeriodRrmUmtsType{
	"100ms",
	"250ms",
	"500ms",
	"1000ms",
	"2000ms",
	"3000ms",
	"4000ms",
	"6000ms",
}

func (v *CollectionPeriodRrmUmtsType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodRrmUmtsType(value)
	for _, existing := range AllowedCollectionPeriodRrmUmtsTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodRrmUmtsType", value)
}

// NewCollectionPeriodRrmUmtsTypeFromValue returns a pointer to a valid CollectionPeriodRrmUmtsType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodRrmUmtsTypeFromValue(v string) (*CollectionPeriodRrmUmtsType, error) {
	ev := CollectionPeriodRrmUmtsType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodRrmUmtsType: valid values are %v", v, AllowedCollectionPeriodRrmUmtsTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodRrmUmtsType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodRrmUmtsTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodRrmUmts-Type value
func (v CollectionPeriodRrmUmtsType) Ptr() *CollectionPeriodRrmUmtsType {
	return &v
}

type NullableCollectionPeriodRrmUmtsType struct {
	value *CollectionPeriodRrmUmtsType
	isSet bool
}

func (v NullableCollectionPeriodRrmUmtsType) Get() *CollectionPeriodRrmUmtsType {
	return v.value
}

func (v *NullableCollectionPeriodRrmUmtsType) Set(val *CollectionPeriodRrmUmtsType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRrmUmtsType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRrmUmtsType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRrmUmtsType(val *CollectionPeriodRrmUmtsType) *NullableCollectionPeriodRrmUmtsType {
	return &NullableCollectionPeriodRrmUmtsType{value: val, isSet: true}
}

func (v NullableCollectionPeriodRrmUmtsType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRrmUmtsType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

