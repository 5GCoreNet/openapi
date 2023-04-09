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

// NFServiceType the model 'NFServiceType'
type NFServiceType string

// List of NFServiceType
const (
	NAMF_COMMUNICATION  NFServiceType = "Namf_Communication"
	NAMF_EVENT_EXPOSURE NFServiceType = "Namf_EventExposure"
	NAMF_MT             NFServiceType = "Namf_MT"
	NAMF_LOCATION       NFServiceType = "Namf_Location"
	NSMF_PDU_SESSION    NFServiceType = "Nsmf_PDUSession"
	NSMF_EVENT_EXPOSURE NFServiceType = "Nsmf_EventExposure"
	OTHERS              NFServiceType = "Others"
)

// All allowed values of NFServiceType enum
var AllowedNFServiceTypeEnumValues = []NFServiceType{
	"Namf_Communication",
	"Namf_EventExposure",
	"Namf_MT",
	"Namf_Location",
	"Nsmf_PDUSession",
	"Nsmf_EventExposure",
	"Others",
}

func (v *NFServiceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := NFServiceType(value)
	for _, existing := range AllowedNFServiceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid NFServiceType", value)
}

// NewNFServiceTypeFromValue returns a pointer to a valid NFServiceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewNFServiceTypeFromValue(v string) (*NFServiceType, error) {
	ev := NFServiceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for NFServiceType: valid values are %v", v, AllowedNFServiceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v NFServiceType) IsValid() bool {
	for _, existing := range AllowedNFServiceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to NFServiceType value
func (v NFServiceType) Ptr() *NFServiceType {
	return &v
}

type NullableNFServiceType struct {
	value *NFServiceType
	isSet bool
}

func (v NullableNFServiceType) Get() *NFServiceType {
	return v.value
}

func (v *NullableNFServiceType) Set(val *NFServiceType) {
	v.value = val
	v.isSet = true
}

func (v NullableNFServiceType) IsSet() bool {
	return v.isSet
}

func (v *NullableNFServiceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNFServiceType(val *NFServiceType) *NullableNFServiceType {
	return &NullableNFServiceType{value: val, isSet: true}
}

func (v NullableNFServiceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNFServiceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
