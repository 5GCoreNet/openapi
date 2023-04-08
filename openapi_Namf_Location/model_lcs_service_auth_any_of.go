/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
	"fmt"
)

// LcsServiceAuthAnyOf the model 'LcsServiceAuthAnyOf'
type LcsServiceAuthAnyOf string

// List of LcsServiceAuth_anyOf
const (
	LOCATION_ALLOWED_WITH_NOTIFICATION LcsServiceAuthAnyOf = "LOCATION_ALLOWED_WITH_NOTIFICATION"
	LOCATION_ALLOWED_WITHOUT_NOTIFICATION LcsServiceAuthAnyOf = "LOCATION_ALLOWED_WITHOUT_NOTIFICATION"
	LOCATION_ALLOWED_WITHOUT_RESPONSE LcsServiceAuthAnyOf = "LOCATION_ALLOWED_WITHOUT_RESPONSE"
	LOCATION_RESTRICTED_WITHOUT_RESPONSE LcsServiceAuthAnyOf = "LOCATION_RESTRICTED_WITHOUT_RESPONSE"
	NOTIFICATION_ONLY LcsServiceAuthAnyOf = "NOTIFICATION_ONLY"
	NOTIFICATION_AND_VERIFICATION_ONLY LcsServiceAuthAnyOf = "NOTIFICATION_AND_VERIFICATION_ONLY"
)

// All allowed values of LcsServiceAuthAnyOf enum
var AllowedLcsServiceAuthAnyOfEnumValues = []LcsServiceAuthAnyOf{
	"LOCATION_ALLOWED_WITH_NOTIFICATION",
	"LOCATION_ALLOWED_WITHOUT_NOTIFICATION",
	"LOCATION_ALLOWED_WITHOUT_RESPONSE",
	"LOCATION_RESTRICTED_WITHOUT_RESPONSE",
	"NOTIFICATION_ONLY",
	"NOTIFICATION_AND_VERIFICATION_ONLY",
}

func (v *LcsServiceAuthAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := LcsServiceAuthAnyOf(value)
	for _, existing := range AllowedLcsServiceAuthAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid LcsServiceAuthAnyOf", value)
}

// NewLcsServiceAuthAnyOfFromValue returns a pointer to a valid LcsServiceAuthAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewLcsServiceAuthAnyOfFromValue(v string) (*LcsServiceAuthAnyOf, error) {
	ev := LcsServiceAuthAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for LcsServiceAuthAnyOf: valid values are %v", v, AllowedLcsServiceAuthAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v LcsServiceAuthAnyOf) IsValid() bool {
	for _, existing := range AllowedLcsServiceAuthAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to LcsServiceAuth_anyOf value
func (v LcsServiceAuthAnyOf) Ptr() *LcsServiceAuthAnyOf {
	return &v
}

type NullableLcsServiceAuthAnyOf struct {
	value *LcsServiceAuthAnyOf
	isSet bool
}

func (v NullableLcsServiceAuthAnyOf) Get() *LcsServiceAuthAnyOf {
	return v.value
}

func (v *NullableLcsServiceAuthAnyOf) Set(val *LcsServiceAuthAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsServiceAuthAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsServiceAuthAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsServiceAuthAnyOf(val *LcsServiceAuthAnyOf) *NullableLcsServiceAuthAnyOf {
	return &NullableLcsServiceAuthAnyOf{value: val, isSet: true}
}

func (v NullableLcsServiceAuthAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsServiceAuthAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

