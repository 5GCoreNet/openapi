/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// N32PurposeAnyOf the model 'N32PurposeAnyOf'
type N32PurposeAnyOf string

// List of N32Purpose_anyOf
const (
	ROAMING N32PurposeAnyOf = "ROAMING"
	INTER_PLMN_MOBILITY N32PurposeAnyOf = "INTER_PLMN_MOBILITY"
	SMS_INTERCONNECT N32PurposeAnyOf = "SMS_INTERCONNECT"
	ROAMING_TEST N32PurposeAnyOf = "ROAMING_TEST"
	INTER_PLMN_MOBILITY_TEST N32PurposeAnyOf = "INTER_PLMN_MOBILITY_TEST"
	SMS_INTERCONNECT_TEST N32PurposeAnyOf = "SMS_INTERCONNECT_TEST"
	SNPN_INTERCONNECT N32PurposeAnyOf = "SNPN_INTERCONNECT"
	SNPN_INTERCONNECT_TEST N32PurposeAnyOf = "SNPN_INTERCONNECT_TEST"
	DISASTER_ROAMING N32PurposeAnyOf = "DISASTER_ROAMING"
	DISASTER_ROAMING_TEST N32PurposeAnyOf = "DISASTER_ROAMING_TEST"
)

// All allowed values of N32PurposeAnyOf enum
var AllowedN32PurposeAnyOfEnumValues = []N32PurposeAnyOf{
	"ROAMING",
	"INTER_PLMN_MOBILITY",
	"SMS_INTERCONNECT",
	"ROAMING_TEST",
	"INTER_PLMN_MOBILITY_TEST",
	"SMS_INTERCONNECT_TEST",
	"SNPN_INTERCONNECT",
	"SNPN_INTERCONNECT_TEST",
	"DISASTER_ROAMING",
	"DISASTER_ROAMING_TEST",
}

func (v *N32PurposeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := N32PurposeAnyOf(value)
	for _, existing := range AllowedN32PurposeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid N32PurposeAnyOf", value)
}

// NewN32PurposeAnyOfFromValue returns a pointer to a valid N32PurposeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewN32PurposeAnyOfFromValue(v string) (*N32PurposeAnyOf, error) {
	ev := N32PurposeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for N32PurposeAnyOf: valid values are %v", v, AllowedN32PurposeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v N32PurposeAnyOf) IsValid() bool {
	for _, existing := range AllowedN32PurposeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to N32Purpose_anyOf value
func (v N32PurposeAnyOf) Ptr() *N32PurposeAnyOf {
	return &v
}

type NullableN32PurposeAnyOf struct {
	value *N32PurposeAnyOf
	isSet bool
}

func (v NullableN32PurposeAnyOf) Get() *N32PurposeAnyOf {
	return v.value
}

func (v *NullableN32PurposeAnyOf) Set(val *N32PurposeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableN32PurposeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableN32PurposeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN32PurposeAnyOf(val *N32PurposeAnyOf) *NullableN32PurposeAnyOf {
	return &NullableN32PurposeAnyOf{value: val, isSet: true}
}

func (v NullableN32PurposeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN32PurposeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

