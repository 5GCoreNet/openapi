/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// NFStatus any of enumrated value
type NFStatus string

// List of NFStatus
const (
	REGISTERED NFStatus = "REGISTERED"
	SUSPENDED  NFStatus = "SUSPENDED"
)

// All allowed values of NFStatus enum
var AllowedNFStatusEnumValues = []NFStatus{
	"REGISTERED",
	"SUSPENDED",
}

func (v *NFStatus) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NFStatus(value)
	for _, existing := range AllowedNFStatusEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NFStatus", value)
}

// NewNFStatusFromValue returns a pointer to a valid NFStatus
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNFStatusFromValue(v string) (*NFStatus, error) {
	ev := NFStatus(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NFStatus: valid values are %v", v, AllowedNFStatusEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NFStatus) IsValid() bool {
	for _, existing := range AllowedNFStatusEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NFStatus value
func (v NFStatus) Ptr() *NFStatus {
	return &v
}

type NullableNFStatus struct {
	value *NFStatus
	isSet bool
}

func (v NullableNFStatus) Get() *NFStatus {
	return v.value
}

func (v *NullableNFStatus) Set(val *NFStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableNFStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableNFStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFStatus(val *NFStatus) *NullableNFStatus {
	return &NullableNFStatus{value: val, isSet: true}
}

func (v NullableNFStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
