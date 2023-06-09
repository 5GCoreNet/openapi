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

// CellState the model 'CellState'
type CellState string

// List of CellState
const (
	IDLE     CellState = "IDLE"
	INACTIVE CellState = "INACTIVE"
	ACTIVE   CellState = "ACTIVE"
)

// All allowed values of CellState enum
var AllowedCellStateEnumValues = []CellState{
	"IDLE",
	"INACTIVE",
	"ACTIVE",
}

func (v *CellState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CellState(value)
	for _, existing := range AllowedCellStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CellState", value)
}

// NewCellStateFromValue returns a pointer to a valid CellState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCellStateFromValue(v string) (*CellState, error) {
	ev := CellState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CellState: valid values are %v", v, AllowedCellStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CellState) IsValid() bool {
	for _, existing := range AllowedCellStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CellState value
func (v CellState) Ptr() *CellState {
	return &v
}

type NullableCellState struct {
	value *CellState
	isSet bool
}

func (v NullableCellState) Get() *CellState {
	return v.value
}

func (v *NullableCellState) Set(val *CellState) {
	v.value = val
	v.isSet = true
}

func (v NullableCellState) IsSet() bool {
	return v.isSet
}

func (v *NullableCellState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCellState(val *CellState) *NullableCellState {
	return &NullableCellState{value: val, isSet: true}
}

func (v NullableCellState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCellState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
