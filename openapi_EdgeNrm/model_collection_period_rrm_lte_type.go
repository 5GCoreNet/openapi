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

// CollectionPeriodRrmLteType See details in 3GPP TS 32.422 clause 5.10.20.
type CollectionPeriodRrmLteType string

// List of collectionPeriodRrmLte-Type
const (
	_100MS CollectionPeriodRrmLteType = "100ms"
	_1000MS CollectionPeriodRrmLteType = "1000ms"
	_1024MS CollectionPeriodRrmLteType = "1024ms"
	_1280MS CollectionPeriodRrmLteType = "1280ms"
	_2048MS CollectionPeriodRrmLteType = "2048ms"
	_2560MS CollectionPeriodRrmLteType = "2560ms"
	_5120MS CollectionPeriodRrmLteType = "5120ms"
	_10000MS CollectionPeriodRrmLteType = "10000ms"
	_10240MS CollectionPeriodRrmLteType = "10240ms"
	_60000MS CollectionPeriodRrmLteType = "60000ms"
)

// All allowed values of CollectionPeriodRrmLteType enum
var AllowedCollectionPeriodRrmLteTypeEnumValues = []CollectionPeriodRrmLteType{
	"100ms",
	"1000ms",
	"1024ms",
	"1280ms",
	"2048ms",
	"2560ms",
	"5120ms",
	"10000ms",
	"10240ms",
	"60000ms",
}

func (v *CollectionPeriodRrmLteType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CollectionPeriodRrmLteType(value)
	for _, existing := range AllowedCollectionPeriodRrmLteTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CollectionPeriodRrmLteType", value)
}

// NewCollectionPeriodRrmLteTypeFromValue returns a pointer to a valid CollectionPeriodRrmLteType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCollectionPeriodRrmLteTypeFromValue(v string) (*CollectionPeriodRrmLteType, error) {
	ev := CollectionPeriodRrmLteType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CollectionPeriodRrmLteType: valid values are %v", v, AllowedCollectionPeriodRrmLteTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CollectionPeriodRrmLteType) IsValid() bool {
	for _, existing := range AllowedCollectionPeriodRrmLteTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to collectionPeriodRrmLte-Type value
func (v CollectionPeriodRrmLteType) Ptr() *CollectionPeriodRrmLteType {
	return &v
}

type NullableCollectionPeriodRrmLteType struct {
	value *CollectionPeriodRrmLteType
	isSet bool
}

func (v NullableCollectionPeriodRrmLteType) Get() *CollectionPeriodRrmLteType {
	return v.value
}

func (v *NullableCollectionPeriodRrmLteType) Set(val *CollectionPeriodRrmLteType) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionPeriodRrmLteType) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionPeriodRrmLteType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionPeriodRrmLteType(val *CollectionPeriodRrmLteType) *NullableCollectionPeriodRrmLteType {
	return &NullableCollectionPeriodRrmLteType{value: val, isSet: true}
}

func (v NullableCollectionPeriodRrmLteType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionPeriodRrmLteType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
