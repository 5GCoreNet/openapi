/*
AI/ML NRM

OAS 3.0.1 specification of the AI/ML NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AiMlNrm

import (
	"encoding/json"
	"fmt"
)

// OperationSemantics the model 'OperationSemantics'
type OperationSemantics string

// List of OperationSemantics
const (
	REQUEST_RESPONSE OperationSemantics = "REQUEST_RESPONSE"
	SUBSCRIBE_NOTIFY OperationSemantics = "SUBSCRIBE_NOTIFY"
)

// All allowed values of OperationSemantics enum
var AllowedOperationSemanticsEnumValues = []OperationSemantics{
	"REQUEST_RESPONSE",
	"SUBSCRIBE_NOTIFY",
}

func (v *OperationSemantics) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := OperationSemantics(value)
	for _, existing := range AllowedOperationSemanticsEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid OperationSemantics", value)
}

// NewOperationSemanticsFromValue returns a pointer to a valid OperationSemantics
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewOperationSemanticsFromValue(v string) (*OperationSemantics, error) {
	ev := OperationSemantics(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for OperationSemantics: valid values are %v", v, AllowedOperationSemanticsEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v OperationSemantics) IsValid() bool {
	for _, existing := range AllowedOperationSemanticsEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to OperationSemantics value
func (v OperationSemantics) Ptr() *OperationSemantics {
	return &v
}

type NullableOperationSemantics struct {
	value *OperationSemantics
	isSet bool
}

func (v NullableOperationSemantics) Get() *OperationSemantics {
	return v.value
}

func (v *NullableOperationSemantics) Set(val *OperationSemantics) {
	v.value = val
	v.isSet = true
}

func (v NullableOperationSemantics) IsSet() bool {
	return v.isSet
}

func (v *NullableOperationSemantics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOperationSemantics(val *OperationSemantics) *NullableOperationSemantics {
	return &NullableOperationSemantics{value: val, isSet: true}
}

func (v NullableOperationSemantics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOperationSemantics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
