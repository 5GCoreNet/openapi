/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// IMSNodeFunctionalityAnyOf the model 'IMSNodeFunctionalityAnyOf'
type IMSNodeFunctionalityAnyOf string

// List of IMSNodeFunctionality_anyOf
const (
	S_CSCF IMSNodeFunctionalityAnyOf = "S_CSCF"
	P_CSCF IMSNodeFunctionalityAnyOf = "P_CSCF"
	I_CSCF IMSNodeFunctionalityAnyOf = "I_CSCF"
	MRFC IMSNodeFunctionalityAnyOf = "MRFC"
	MGCF IMSNodeFunctionalityAnyOf = "MGCF"
	BGCF IMSNodeFunctionalityAnyOf = "BGCF"
	AS IMSNodeFunctionalityAnyOf = "AS"
	IBCF IMSNodeFunctionalityAnyOf = "IBCF"
	S_GW IMSNodeFunctionalityAnyOf = "S-GW"
	P_GW IMSNodeFunctionalityAnyOf = "P-GW"
	HSGW IMSNodeFunctionalityAnyOf = "HSGW"
	E_CSCF IMSNodeFunctionalityAnyOf = "E-CSCF"
	MME IMSNodeFunctionalityAnyOf = "MME"
	TRF IMSNodeFunctionalityAnyOf = "TRF"
	TF IMSNodeFunctionalityAnyOf = "TF"
	ATCF IMSNodeFunctionalityAnyOf = "ATCF"
	PROXY IMSNodeFunctionalityAnyOf = "PROXY"
	EPDG IMSNodeFunctionalityAnyOf = "EPDG"
	TDF IMSNodeFunctionalityAnyOf = "TDF"
	TWAG IMSNodeFunctionalityAnyOf = "TWAG"
	SCEF IMSNodeFunctionalityAnyOf = "SCEF"
	IWK_SCEF IMSNodeFunctionalityAnyOf = "IWK_SCEF"
	IMS_GWF IMSNodeFunctionalityAnyOf = "IMS_GWF"
)

// All allowed values of IMSNodeFunctionalityAnyOf enum
var AllowedIMSNodeFunctionalityAnyOfEnumValues = []IMSNodeFunctionalityAnyOf{
	"S_CSCF",
	"P_CSCF",
	"I_CSCF",
	"MRFC",
	"MGCF",
	"BGCF",
	"AS",
	"IBCF",
	"S-GW",
	"P-GW",
	"HSGW",
	"E-CSCF",
	"MME",
	"TRF",
	"TF",
	"ATCF",
	"PROXY",
	"EPDG",
	"TDF",
	"TWAG",
	"SCEF",
	"IWK_SCEF",
	"IMS_GWF",
}

func (v *IMSNodeFunctionalityAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := IMSNodeFunctionalityAnyOf(value)
	for _, existing := range AllowedIMSNodeFunctionalityAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid IMSNodeFunctionalityAnyOf", value)
}

// NewIMSNodeFunctionalityAnyOfFromValue returns a pointer to a valid IMSNodeFunctionalityAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIMSNodeFunctionalityAnyOfFromValue(v string) (*IMSNodeFunctionalityAnyOf, error) {
	ev := IMSNodeFunctionalityAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for IMSNodeFunctionalityAnyOf: valid values are %v", v, AllowedIMSNodeFunctionalityAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v IMSNodeFunctionalityAnyOf) IsValid() bool {
	for _, existing := range AllowedIMSNodeFunctionalityAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to IMSNodeFunctionality_anyOf value
func (v IMSNodeFunctionalityAnyOf) Ptr() *IMSNodeFunctionalityAnyOf {
	return &v
}

type NullableIMSNodeFunctionalityAnyOf struct {
	value *IMSNodeFunctionalityAnyOf
	isSet bool
}

func (v NullableIMSNodeFunctionalityAnyOf) Get() *IMSNodeFunctionalityAnyOf {
	return v.value
}

func (v *NullableIMSNodeFunctionalityAnyOf) Set(val *IMSNodeFunctionalityAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableIMSNodeFunctionalityAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableIMSNodeFunctionalityAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIMSNodeFunctionalityAnyOf(val *IMSNodeFunctionalityAnyOf) *NullableIMSNodeFunctionalityAnyOf {
	return &NullableIMSNodeFunctionalityAnyOf{value: val, isSet: true}
}

func (v NullableIMSNodeFunctionalityAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIMSNodeFunctionalityAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

