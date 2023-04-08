/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_PP

import (
	"encoding/json"
	"fmt"
)

// AccessTechAnyOf the model 'AccessTechAnyOf'
type AccessTechAnyOf string

// List of AccessTech_anyOf
const (
	NR AccessTechAnyOf = "NR"
	EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE AccessTechAnyOf = "EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE"
	EUTRAN_IN_NBS1_MODE_ONLY AccessTechAnyOf = "EUTRAN_IN_NBS1_MODE_ONLY"
	EUTRAN_IN_WBS1_MODE_ONLY AccessTechAnyOf = "EUTRAN_IN_WBS1_MODE_ONLY"
	UTRAN AccessTechAnyOf = "UTRAN"
	GSM_AND_ECGSM_IO_T AccessTechAnyOf = "GSM_AND_ECGSM_IoT"
	GSM_WITHOUT_ECGSM_IO_T AccessTechAnyOf = "GSM_WITHOUT_ECGSM_IoT"
	ECGSM_IO_T_ONLY AccessTechAnyOf = "ECGSM_IoT_ONLY"
	CDMA_1X_RTT AccessTechAnyOf = "CDMA_1xRTT"
	CDMA_HRPD AccessTechAnyOf = "CDMA_HRPD"
	GSM_COMPACT AccessTechAnyOf = "GSM_COMPACT"
)

// All allowed values of AccessTechAnyOf enum
var AllowedAccessTechAnyOfEnumValues = []AccessTechAnyOf{
	"NR",
	"EUTRAN_IN_WBS1_MODE_AND_NBS1_MODE",
	"EUTRAN_IN_NBS1_MODE_ONLY",
	"EUTRAN_IN_WBS1_MODE_ONLY",
	"UTRAN",
	"GSM_AND_ECGSM_IoT",
	"GSM_WITHOUT_ECGSM_IoT",
	"ECGSM_IoT_ONLY",
	"CDMA_1xRTT",
	"CDMA_HRPD",
	"GSM_COMPACT",
}

func (v *AccessTechAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccessTechAnyOf(value)
	for _, existing := range AllowedAccessTechAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccessTechAnyOf", value)
}

// NewAccessTechAnyOfFromValue returns a pointer to a valid AccessTechAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccessTechAnyOfFromValue(v string) (*AccessTechAnyOf, error) {
	ev := AccessTechAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccessTechAnyOf: valid values are %v", v, AllowedAccessTechAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccessTechAnyOf) IsValid() bool {
	for _, existing := range AllowedAccessTechAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccessTech_anyOf value
func (v AccessTechAnyOf) Ptr() *AccessTechAnyOf {
	return &v
}

type NullableAccessTechAnyOf struct {
	value *AccessTechAnyOf
	isSet bool
}

func (v NullableAccessTechAnyOf) Get() *AccessTechAnyOf {
	return v.value
}

func (v *NullableAccessTechAnyOf) Set(val *AccessTechAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTechAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTechAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTechAnyOf(val *AccessTechAnyOf) *NullableAccessTechAnyOf {
	return &NullableAccessTechAnyOf{value: val, isSet: true}
}

func (v NullableAccessTechAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTechAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

