/*
Fault Supervision MnS

OAS 3.0.1 definition of the Fault Supervision MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_FaultMnS

import (
	"encoding/json"
	"fmt"
)

// AckState the model 'AckState'
type AckState string

// List of AckState
const (
	ACKNOWLEDGED AckState = "ACKNOWLEDGED"
	UNACKNOWLEDGED AckState = "UNACKNOWLEDGED"
)

// All allowed values of AckState enum
var AllowedAckStateEnumValues = []AckState{
	"ACKNOWLEDGED",
	"UNACKNOWLEDGED",
}

func (v *AckState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AckState(value)
	for _, existing := range AllowedAckStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AckState", value)
}

// NewAckStateFromValue returns a pointer to a valid AckState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAckStateFromValue(v string) (*AckState, error) {
	ev := AckState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AckState: valid values are %v", v, AllowedAckStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AckState) IsValid() bool {
	for _, existing := range AllowedAckStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AckState value
func (v AckState) Ptr() *AckState {
	return &v
}

type NullableAckState struct {
	value *AckState
	isSet bool
}

func (v NullableAckState) Get() *AckState {
	return v.value
}

func (v *NullableAckState) Set(val *AckState) {
	v.value = val
	v.isSet = true
}

func (v NullableAckState) IsSet() bool {
	return v.isSet
}

func (v *NullableAckState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAckState(val *AckState) *NullableAckState {
	return &NullableAckState{value: val, isSet: true}
}

func (v NullableAckState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAckState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

