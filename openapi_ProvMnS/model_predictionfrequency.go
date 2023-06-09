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

// Predictionfrequency the model 'Predictionfrequency'
type Predictionfrequency string

// List of Predictionfrequency
const (
	PERSEC  Predictionfrequency = "PERSEC"
	PERMIN  Predictionfrequency = "PERMIN"
	PERHOUR Predictionfrequency = "PERHOUR"
)

// All allowed values of Predictionfrequency enum
var AllowedPredictionfrequencyEnumValues = []Predictionfrequency{
	"PERSEC",
	"PERMIN",
	"PERHOUR",
}

func (v *Predictionfrequency) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Predictionfrequency(value)
	for _, existing := range AllowedPredictionfrequencyEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Predictionfrequency", value)
}

// NewPredictionfrequencyFromValue returns a pointer to a valid Predictionfrequency
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPredictionfrequencyFromValue(v string) (*Predictionfrequency, error) {
	ev := Predictionfrequency(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Predictionfrequency: valid values are %v", v, AllowedPredictionfrequencyEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Predictionfrequency) IsValid() bool {
	for _, existing := range AllowedPredictionfrequencyEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Predictionfrequency value
func (v Predictionfrequency) Ptr() *Predictionfrequency {
	return &v
}

type NullablePredictionfrequency struct {
	value *Predictionfrequency
	isSet bool
}

func (v NullablePredictionfrequency) Get() *Predictionfrequency {
	return v.value
}

func (v *NullablePredictionfrequency) Set(val *Predictionfrequency) {
	v.value = val
	v.isSet = true
}

func (v NullablePredictionfrequency) IsSet() bool {
	return v.isSet
}

func (v *NullablePredictionfrequency) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePredictionfrequency(val *Predictionfrequency) *NullablePredictionfrequency {
	return &NullablePredictionfrequency{value: val, isSet: true}
}

func (v NullablePredictionfrequency) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePredictionfrequency) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
