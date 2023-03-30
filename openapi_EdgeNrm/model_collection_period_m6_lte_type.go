/*
3GPP Edge NRM

OAS 3.0.1 specification of the Edge NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_EdgeNrm

import (
	"encoding/json"
	"fmt"
)

// CollectionPeriodM6LteType See details in 3GPP TS 32.422 clause 5.10.32.
type CollectionPeriodM6LteType string

// List of collectionPeriodM6Lte-Type
const (
	_1024MS CollectionPeriodM6LteType = "1024ms"
	_2048MS CollectionPeriodM6LteType = "2048ms"
	_5120MS CollectionPeriodM6LteType = "5120ms"
	_10240MS CollectionPeriodM6LteType = "10240ms"
)

// All allowed values of CollectionPeriodM6LteType enum
var AllowedCollectionPeriodM6LteTypeEnumValues = []CollectionPeriodM6LteType{
	"1024ms",
	"2048ms",
	"5120ms",
	"10240ms",
}

func (v *CollectionPeriodM6LteType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodM6LteType(value)
	for _, existing := range AllowedCollectionPeriodM6LteTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodM6LteType", value)
}

// NewCollectionPeriodM6LteTypeFromValue returns a pointer to a valid CollectionPeriodM6LteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodM6LteTypeFromValue(v string) (*CollectionPeriodM6LteType, error) {
	ev := CollectionPeriodM6LteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodM6LteType: valid values are %v", v, AllowedCollectionPeriodM6LteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodM6LteType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodM6LteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodM6Lte-Type value
func (v CollectionPeriodM6LteType) Ptr() *CollectionPeriodM6LteType {
	return &v
}

type NullableCollectionPeriodM6LteType struct {
	value *CollectionPeriodM6LteType
	isSet bool
}

func (v NullableCollectionPeriodM6LteType) Get() *CollectionPeriodM6LteType {
	return v.value
}

func (v *NullableCollectionPeriodM6LteType) Set(val *CollectionPeriodM6LteType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodM6LteType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodM6LteType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodM6LteType(val *CollectionPeriodM6LteType) *NullableCollectionPeriodM6LteType {
	return &NullableCollectionPeriodM6LteType{value: val, isSet: true}
}

func (v NullableCollectionPeriodM6LteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodM6LteType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

