/*
Nhss_imsUEAU

Nhss UE Authentication Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsUEAU

import (
	"encoding/json"
	"fmt"
)

// SipDigestAlgorithmAnyOf the model 'SipDigestAlgorithmAnyOf'
type SipDigestAlgorithmAnyOf string

// List of SipDigestAlgorithm_anyOf
const (
	MD5 SipDigestAlgorithmAnyOf = "MD5"
	MD5_SESS SipDigestAlgorithmAnyOf = "MD5_SESS"
)

// All allowed values of SipDigestAlgorithmAnyOf enum
var AllowedSipDigestAlgorithmAnyOfEnumValues = []SipDigestAlgorithmAnyOf{
	"MD5",
	"MD5_SESS",
}

func (v *SipDigestAlgorithmAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SipDigestAlgorithmAnyOf(value)
	for _, existing := range AllowedSipDigestAlgorithmAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SipDigestAlgorithmAnyOf", value)
}

// NewSipDigestAlgorithmAnyOfFromValue returns a pointer to a valid SipDigestAlgorithmAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSipDigestAlgorithmAnyOfFromValue(v string) (*SipDigestAlgorithmAnyOf, error) {
	ev := SipDigestAlgorithmAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SipDigestAlgorithmAnyOf: valid values are %v", v, AllowedSipDigestAlgorithmAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SipDigestAlgorithmAnyOf) IsValid() bool {
	for _, existing := range AllowedSipDigestAlgorithmAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SipDigestAlgorithm_anyOf value
func (v SipDigestAlgorithmAnyOf) Ptr() *SipDigestAlgorithmAnyOf {
	return &v
}

type NullableSipDigestAlgorithmAnyOf struct {
	value *SipDigestAlgorithmAnyOf
	isSet bool
}

func (v NullableSipDigestAlgorithmAnyOf) Get() *SipDigestAlgorithmAnyOf {
	return v.value
}

func (v *NullableSipDigestAlgorithmAnyOf) Set(val *SipDigestAlgorithmAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableSipDigestAlgorithmAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableSipDigestAlgorithmAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSipDigestAlgorithmAnyOf(val *SipDigestAlgorithmAnyOf) *NullableSipDigestAlgorithmAnyOf {
	return &NullableSipDigestAlgorithmAnyOf{value: val, isSet: true}
}

func (v NullableSipDigestAlgorithmAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSipDigestAlgorithmAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

