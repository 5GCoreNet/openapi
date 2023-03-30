/*
NR NRM

OAS 3.0.1 specification of the NR NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_NrNrm

import (
	"encoding/json"
	"fmt"
)

// AnonymizationOfMdtDataType Specifies level of MDT anonymization. For additional details see 3GPP TS 32.422 clause 5.10.12.
type AnonymizationOfMdtDataType string

// List of anonymizationOfMdtData-Type
const (
	NO_IDENTITY AnonymizationOfMdtDataType = "NO_IDENTITY"
	TAC_OF_IMEI AnonymizationOfMdtDataType = "TAC_OF_IMEI"
)

// All allowed values of AnonymizationOfMdtDataType enum
var AllowedAnonymizationOfMdtDataTypeEnumValues = []AnonymizationOfMdtDataType{
	"NO_IDENTITY",
	"TAC_OF_IMEI",
}

func (v *AnonymizationOfMdtDataType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AnonymizationOfMdtDataType(value)
	for _, existing := range AllowedAnonymizationOfMdtDataTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AnonymizationOfMdtDataType", value)
}

// NewAnonymizationOfMdtDataTypeFromValue returns a pointer to a valid AnonymizationOfMdtDataType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAnonymizationOfMdtDataTypeFromValue(v string) (*AnonymizationOfMdtDataType, error) {
	ev := AnonymizationOfMdtDataType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AnonymizationOfMdtDataType: valid values are %v", v, AllowedAnonymizationOfMdtDataTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AnonymizationOfMdtDataType) IsValid() bool {
	for _, existing := range AllowedAnonymizationOfMdtDataTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to anonymizationOfMdtData-Type value
func (v AnonymizationOfMdtDataType) Ptr() *AnonymizationOfMdtDataType {
	return &v
}

type NullableAnonymizationOfMdtDataType struct {
	value *AnonymizationOfMdtDataType
	isSet bool
}

func (v NullableAnonymizationOfMdtDataType) Get() *AnonymizationOfMdtDataType {
	return v.value
}

func (v *NullableAnonymizationOfMdtDataType) Set(val *AnonymizationOfMdtDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableAnonymizationOfMdtDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableAnonymizationOfMdtDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnonymizationOfMdtDataType(val *AnonymizationOfMdtDataType) *NullableAnonymizationOfMdtDataType {
	return &NullableAnonymizationOfMdtDataType{value: val, isSet: true}
}

func (v NullableAnonymizationOfMdtDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnonymizationOfMdtDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

