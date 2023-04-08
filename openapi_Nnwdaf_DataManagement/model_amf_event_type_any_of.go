/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// AmfEventTypeAnyOf the model 'AmfEventTypeAnyOf'
type AmfEventTypeAnyOf string

// List of AmfEventType_anyOf
const (
	LOCATION_REPORT AmfEventTypeAnyOf = "LOCATION_REPORT"
	PRESENCE_IN_AOI_REPORT AmfEventTypeAnyOf = "PRESENCE_IN_AOI_REPORT"
	TIMEZONE_REPORT AmfEventTypeAnyOf = "TIMEZONE_REPORT"
	ACCESS_TYPE_REPORT AmfEventTypeAnyOf = "ACCESS_TYPE_REPORT"
	REGISTRATION_STATE_REPORT AmfEventTypeAnyOf = "REGISTRATION_STATE_REPORT"
	CONNECTIVITY_STATE_REPORT AmfEventTypeAnyOf = "CONNECTIVITY_STATE_REPORT"
	REACHABILITY_REPORT AmfEventTypeAnyOf = "REACHABILITY_REPORT"
	COMMUNICATION_FAILURE_REPORT AmfEventTypeAnyOf = "COMMUNICATION_FAILURE_REPORT"
	UES_IN_AREA_REPORT AmfEventTypeAnyOf = "UES_IN_AREA_REPORT"
	SUBSCRIPTION_ID_CHANGE AmfEventTypeAnyOf = "SUBSCRIPTION_ID_CHANGE"
	SUBSCRIPTION_ID_ADDITION AmfEventTypeAnyOf = "SUBSCRIPTION_ID_ADDITION"
	SUBSCRIPTION_TERMINATION AmfEventTypeAnyOf = "SUBSCRIPTION_TERMINATION"
	LOSS_OF_CONNECTIVITY AmfEventTypeAnyOf = "LOSS_OF_CONNECTIVITY"
	_5_GS_USER_STATE_REPORT AmfEventTypeAnyOf = "5GS_USER_STATE_REPORT"
	AVAILABILITY_AFTER_DDN_FAILURE AmfEventTypeAnyOf = "AVAILABILITY_AFTER_DDN_FAILURE"
	TYPE_ALLOCATION_CODE_REPORT AmfEventTypeAnyOf = "TYPE_ALLOCATION_CODE_REPORT"
	FREQUENT_MOBILITY_REGISTRATION_REPORT AmfEventTypeAnyOf = "FREQUENT_MOBILITY_REGISTRATION_REPORT"
	SNSSAI_TA_MAPPING_REPORT AmfEventTypeAnyOf = "SNSSAI_TA_MAPPING_REPORT"
	UE_LOCATION_TRENDS AmfEventTypeAnyOf = "UE_LOCATION_TRENDS"
	UE_ACCESS_BEHAVIOR_TRENDS AmfEventTypeAnyOf = "UE_ACCESS_BEHAVIOR_TRENDS"
	UE_MM_TRANSACTION_REPORT AmfEventTypeAnyOf = "UE_MM_TRANSACTION_REPORT"
)

// All allowed values of AmfEventTypeAnyOf enum
var AllowedAmfEventTypeAnyOfEnumValues = []AmfEventTypeAnyOf{
	"LOCATION_REPORT",
	"PRESENCE_IN_AOI_REPORT",
	"TIMEZONE_REPORT",
	"ACCESS_TYPE_REPORT",
	"REGISTRATION_STATE_REPORT",
	"CONNECTIVITY_STATE_REPORT",
	"REACHABILITY_REPORT",
	"COMMUNICATION_FAILURE_REPORT",
	"UES_IN_AREA_REPORT",
	"SUBSCRIPTION_ID_CHANGE",
	"SUBSCRIPTION_ID_ADDITION",
	"SUBSCRIPTION_TERMINATION",
	"LOSS_OF_CONNECTIVITY",
	"5GS_USER_STATE_REPORT",
	"AVAILABILITY_AFTER_DDN_FAILURE",
	"TYPE_ALLOCATION_CODE_REPORT",
	"FREQUENT_MOBILITY_REGISTRATION_REPORT",
	"SNSSAI_TA_MAPPING_REPORT",
	"UE_LOCATION_TRENDS",
	"UE_ACCESS_BEHAVIOR_TRENDS",
	"UE_MM_TRANSACTION_REPORT",
}

func (v *AmfEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AmfEventTypeAnyOf(value)
	for _, existing := range AllowedAmfEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AmfEventTypeAnyOf", value)
}

// NewAmfEventTypeAnyOfFromValue returns a pointer to a valid AmfEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAmfEventTypeAnyOfFromValue(v string) (*AmfEventTypeAnyOf, error) {
	ev := AmfEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AmfEventTypeAnyOf: valid values are %v", v, AllowedAmfEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AmfEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedAmfEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AmfEventType_anyOf value
func (v AmfEventTypeAnyOf) Ptr() *AmfEventTypeAnyOf {
	return &v
}

type NullableAmfEventTypeAnyOf struct {
	value *AmfEventTypeAnyOf
	isSet bool
}

func (v NullableAmfEventTypeAnyOf) Get() *AmfEventTypeAnyOf {
	return v.value
}

func (v *NullableAmfEventTypeAnyOf) Set(val *AmfEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableAmfEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableAmfEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAmfEventTypeAnyOf(val *AmfEventTypeAnyOf) *NullableAmfEventTypeAnyOf {
	return &NullableAmfEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableAmfEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAmfEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

