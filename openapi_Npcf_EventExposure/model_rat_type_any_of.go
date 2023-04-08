/*
Npcf_EventExposure

PCF Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// RatTypeAnyOf the model 'RatTypeAnyOf'
type RatTypeAnyOf string

// List of RatType_anyOf
const (
	NR RatTypeAnyOf = "NR"
	EUTRA RatTypeAnyOf = "EUTRA"
	WLAN RatTypeAnyOf = "WLAN"
	VIRTUAL RatTypeAnyOf = "VIRTUAL"
	NBIOT RatTypeAnyOf = "NBIOT"
	WIRELINE RatTypeAnyOf = "WIRELINE"
	WIRELINE_CABLE RatTypeAnyOf = "WIRELINE_CABLE"
	WIRELINE_BBF RatTypeAnyOf = "WIRELINE_BBF"
	LTE_M RatTypeAnyOf = "LTE-M"
	NR_U RatTypeAnyOf = "NR_U"
	EUTRA_U RatTypeAnyOf = "EUTRA_U"
	TRUSTED_N3_GA RatTypeAnyOf = "TRUSTED_N3GA"
	TRUSTED_WLAN RatTypeAnyOf = "TRUSTED_WLAN"
	UTRA RatTypeAnyOf = "UTRA"
	GERA RatTypeAnyOf = "GERA"
	NR_LEO RatTypeAnyOf = "NR_LEO"
	NR_MEO RatTypeAnyOf = "NR_MEO"
	NR_GEO RatTypeAnyOf = "NR_GEO"
	NR_OTHER_SAT RatTypeAnyOf = "NR_OTHER_SAT"
	NR_REDCAP RatTypeAnyOf = "NR_REDCAP"
)

// All allowed values of RatTypeAnyOf enum
var AllowedRatTypeAnyOfEnumValues = []RatTypeAnyOf{
	"NR",
	"EUTRA",
	"WLAN",
	"VIRTUAL",
	"NBIOT",
	"WIRELINE",
	"WIRELINE_CABLE",
	"WIRELINE_BBF",
	"LTE-M",
	"NR_U",
	"EUTRA_U",
	"TRUSTED_N3GA",
	"TRUSTED_WLAN",
	"UTRA",
	"GERA",
	"NR_LEO",
	"NR_MEO",
	"NR_GEO",
	"NR_OTHER_SAT",
	"NR_REDCAP",
}

func (v *RatTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RatTypeAnyOf(value)
	for _, existing := range AllowedRatTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RatTypeAnyOf", value)
}

// NewRatTypeAnyOfFromValue returns a pointer to a valid RatTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRatTypeAnyOfFromValue(v string) (*RatTypeAnyOf, error) {
	ev := RatTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RatTypeAnyOf: valid values are %v", v, AllowedRatTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RatTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedRatTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RatType_anyOf value
func (v RatTypeAnyOf) Ptr() *RatTypeAnyOf {
	return &v
}

type NullableRatTypeAnyOf struct {
	value *RatTypeAnyOf
	isSet bool
}

func (v NullableRatTypeAnyOf) Get() *RatTypeAnyOf {
	return v.value
}

func (v *NullableRatTypeAnyOf) Set(val *RatTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableRatTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableRatTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRatTypeAnyOf(val *RatTypeAnyOf) *NullableRatTypeAnyOf {
	return &NullableRatTypeAnyOf{value: val, isSet: true}
}

func (v NullableRatTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRatTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

