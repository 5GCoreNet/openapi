/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// PduSessionTypeAnyOf the model 'PduSessionTypeAnyOf'
type PduSessionTypeAnyOf string

// List of PduSessionType_anyOf
const (
	IPV4 PduSessionTypeAnyOf = "IPV4"
	IPV6 PduSessionTypeAnyOf = "IPV6"
	IPV4_V6 PduSessionTypeAnyOf = "IPV4V6"
	UNSTRUCTURED PduSessionTypeAnyOf = "UNSTRUCTURED"
	ETHERNET PduSessionTypeAnyOf = "ETHERNET"
)

// All allowed values of PduSessionTypeAnyOf enum
var AllowedPduSessionTypeAnyOfEnumValues = []PduSessionTypeAnyOf{
	"IPV4",
	"IPV6",
	"IPV4V6",
	"UNSTRUCTURED",
	"ETHERNET",
}

func (v *PduSessionTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PduSessionTypeAnyOf(value)
	for _, existing := range AllowedPduSessionTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PduSessionTypeAnyOf", value)
}

// NewPduSessionTypeAnyOfFromValue returns a pointer to a valid PduSessionTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPduSessionTypeAnyOfFromValue(v string) (*PduSessionTypeAnyOf, error) {
	ev := PduSessionTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PduSessionTypeAnyOf: valid values are %v", v, AllowedPduSessionTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PduSessionTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedPduSessionTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PduSessionType_anyOf value
func (v PduSessionTypeAnyOf) Ptr() *PduSessionTypeAnyOf {
	return &v
}

type NullablePduSessionTypeAnyOf struct {
	value *PduSessionTypeAnyOf
	isSet bool
}

func (v NullablePduSessionTypeAnyOf) Get() *PduSessionTypeAnyOf {
	return v.value
}

func (v *NullablePduSessionTypeAnyOf) Set(val *PduSessionTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionTypeAnyOf(val *PduSessionTypeAnyOf) *NullablePduSessionTypeAnyOf {
	return &NullablePduSessionTypeAnyOf{value: val, isSet: true}
}

func (v NullablePduSessionTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

