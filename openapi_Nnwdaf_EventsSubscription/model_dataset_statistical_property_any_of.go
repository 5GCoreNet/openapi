/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// DatasetStatisticalPropertyAnyOf the model 'DatasetStatisticalPropertyAnyOf'
type DatasetStatisticalPropertyAnyOf string

// List of DatasetStatisticalProperty_anyOf
const (
	UNIFORM_DIST_DATA DatasetStatisticalPropertyAnyOf = "UNIFORM_DIST_DATA"
	NO_OUTLIERS DatasetStatisticalPropertyAnyOf = "NO_OUTLIERS"
)

// All allowed values of DatasetStatisticalPropertyAnyOf enum
var AllowedDatasetStatisticalPropertyAnyOfEnumValues = []DatasetStatisticalPropertyAnyOf{
	"UNIFORM_DIST_DATA",
	"NO_OUTLIERS",
}

func (v *DatasetStatisticalPropertyAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := DatasetStatisticalPropertyAnyOf(value)
	for _, existing := range AllowedDatasetStatisticalPropertyAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid DatasetStatisticalPropertyAnyOf", value)
}

// NewDatasetStatisticalPropertyAnyOfFromValue returns a pointer to a valid DatasetStatisticalPropertyAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewDatasetStatisticalPropertyAnyOfFromValue(v string) (*DatasetStatisticalPropertyAnyOf, error) {
	ev := DatasetStatisticalPropertyAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for DatasetStatisticalPropertyAnyOf: valid values are %v", v, AllowedDatasetStatisticalPropertyAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v DatasetStatisticalPropertyAnyOf) IsValid() bool {
	for _, existing := range AllowedDatasetStatisticalPropertyAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to DatasetStatisticalProperty_anyOf value
func (v DatasetStatisticalPropertyAnyOf) Ptr() *DatasetStatisticalPropertyAnyOf {
	return &v
}

type NullableDatasetStatisticalPropertyAnyOf struct {
	value *DatasetStatisticalPropertyAnyOf
	isSet bool
}

func (v NullableDatasetStatisticalPropertyAnyOf) Get() *DatasetStatisticalPropertyAnyOf {
	return v.value
}

func (v *NullableDatasetStatisticalPropertyAnyOf) Set(val *DatasetStatisticalPropertyAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableDatasetStatisticalPropertyAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableDatasetStatisticalPropertyAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDatasetStatisticalPropertyAnyOf(val *DatasetStatisticalPropertyAnyOf) *NullableDatasetStatisticalPropertyAnyOf {
	return &NullableDatasetStatisticalPropertyAnyOf{value: val, isSet: true}
}

func (v NullableDatasetStatisticalPropertyAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDatasetStatisticalPropertyAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

