/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
	"fmt"
)

// TrendIndication the model 'TrendIndication'
type TrendIndication string

// List of TrendIndication
const (
	MORE_SEVERE TrendIndication = "MORE_SEVERE"
	NO_CHANGE TrendIndication = "NO_CHANGE"
	LESS_SEVERE TrendIndication = "LESS_SEVERE"
)

// All allowed values of TrendIndication enum
var AllowedTrendIndicationEnumValues = []TrendIndication{
	"MORE_SEVERE",
	"NO_CHANGE",
	"LESS_SEVERE",
}

func (v *TrendIndication) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := TrendIndication(value)
	for _, existing := range AllowedTrendIndicationEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid TrendIndication", value)
}

// NewTrendIndicationFromValue returns a pointer to a valid TrendIndication
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewTrendIndicationFromValue(v string) (*TrendIndication, error) {
	ev := TrendIndication(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for TrendIndication: valid values are %v", v, AllowedTrendIndicationEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v TrendIndication) IsValid() bool {
	for _, existing := range AllowedTrendIndicationEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to TrendIndication value
func (v TrendIndication) Ptr() *TrendIndication {
	return &v
}

type NullableTrendIndication struct {
	value *TrendIndication
	isSet bool
}

func (v NullableTrendIndication) Get() *TrendIndication {
	return v.value
}

func (v *NullableTrendIndication) Set(val *TrendIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableTrendIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableTrendIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrendIndication(val *TrendIndication) *NullableTrendIndication {
	return &NullableTrendIndication{value: val, isSet: true}
}

func (v NullableTrendIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrendIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

