/*
5G-EIR Equipment Identity Check

5G-EIR Equipment Identity Check Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-eir_EquipmentIdentityCheck

import (
	"encoding/json"
	"fmt"
)

// EquipmentStatus Represents equipment status of the PEI. This data type is a string.
type EquipmentStatus string

// List of EquipmentStatus
const (
	WHITELISTED EquipmentStatus = "WHITELISTED"
	BLACKLISTED EquipmentStatus = "BLACKLISTED"
	GREYLISTED EquipmentStatus = "GREYLISTED"
)

// All allowed values of EquipmentStatus enum
var AllowedEquipmentStatusEnumValues = []EquipmentStatus{
	"WHITELISTED",
	"BLACKLISTED",
	"GREYLISTED",
}

func (v *EquipmentStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EquipmentStatus(value)
	for _, existing := range AllowedEquipmentStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EquipmentStatus", value)
}

// NewEquipmentStatusFromValue returns a pointer to a valid EquipmentStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEquipmentStatusFromValue(v string) (*EquipmentStatus, error) {
	ev := EquipmentStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EquipmentStatus: valid values are %v", v, AllowedEquipmentStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EquipmentStatus) IsValid() bool {
	for _, existing := range AllowedEquipmentStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EquipmentStatus value
func (v EquipmentStatus) Ptr() *EquipmentStatus {
	return &v
}

type NullableEquipmentStatus struct {
	value *EquipmentStatus
	isSet bool
}

func (v NullableEquipmentStatus) Get() *EquipmentStatus {
	return v.value
}

func (v *NullableEquipmentStatus) Set(val *EquipmentStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableEquipmentStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableEquipmentStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEquipmentStatus(val *EquipmentStatus) *NullableEquipmentStatus {
	return &NullableEquipmentStatus{value: val, isSet: true}
}

func (v NullableEquipmentStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEquipmentStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

