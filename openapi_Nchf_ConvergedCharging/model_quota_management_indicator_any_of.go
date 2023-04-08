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

// QuotaManagementIndicatorAnyOf the model 'QuotaManagementIndicatorAnyOf'
type QuotaManagementIndicatorAnyOf string

// List of QuotaManagementIndicator_anyOf
const (
	ONLINE_CHARGING QuotaManagementIndicatorAnyOf = "ONLINE_CHARGING"
	OFFLINE_CHARGING QuotaManagementIndicatorAnyOf = "OFFLINE_CHARGING"
	QUOTA_MANAGEMENT_SUSPENDED QuotaManagementIndicatorAnyOf = "QUOTA_MANAGEMENT_SUSPENDED"
)

// All allowed values of QuotaManagementIndicatorAnyOf enum
var AllowedQuotaManagementIndicatorAnyOfEnumValues = []QuotaManagementIndicatorAnyOf{
	"ONLINE_CHARGING",
	"OFFLINE_CHARGING",
	"QUOTA_MANAGEMENT_SUSPENDED",
}

func (v *QuotaManagementIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := QuotaManagementIndicatorAnyOf(value)
	for _, existing := range AllowedQuotaManagementIndicatorAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid QuotaManagementIndicatorAnyOf", value)
}

// NewQuotaManagementIndicatorAnyOfFromValue returns a pointer to a valid QuotaManagementIndicatorAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewQuotaManagementIndicatorAnyOfFromValue(v string) (*QuotaManagementIndicatorAnyOf, error) {
	ev := QuotaManagementIndicatorAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for QuotaManagementIndicatorAnyOf: valid values are %v", v, AllowedQuotaManagementIndicatorAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v QuotaManagementIndicatorAnyOf) IsValid() bool {
	for _, existing := range AllowedQuotaManagementIndicatorAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to QuotaManagementIndicator_anyOf value
func (v QuotaManagementIndicatorAnyOf) Ptr() *QuotaManagementIndicatorAnyOf {
	return &v
}

type NullableQuotaManagementIndicatorAnyOf struct {
	value *QuotaManagementIndicatorAnyOf
	isSet bool
}

func (v NullableQuotaManagementIndicatorAnyOf) Get() *QuotaManagementIndicatorAnyOf {
	return v.value
}

func (v *NullableQuotaManagementIndicatorAnyOf) Set(val *QuotaManagementIndicatorAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableQuotaManagementIndicatorAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableQuotaManagementIndicatorAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQuotaManagementIndicatorAnyOf(val *QuotaManagementIndicatorAnyOf) *NullableQuotaManagementIndicatorAnyOf {
	return &NullableQuotaManagementIndicatorAnyOf{value: val, isSet: true}
}

func (v NullableQuotaManagementIndicatorAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQuotaManagementIndicatorAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

