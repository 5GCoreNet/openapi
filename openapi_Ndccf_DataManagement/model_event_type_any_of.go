/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// EventTypeAnyOf the model 'EventTypeAnyOf'
type EventTypeAnyOf string

// List of EventType_anyOf
const (
	LOSS_OF_CONNECTIVITY EventTypeAnyOf = "LOSS_OF_CONNECTIVITY"
	UE_REACHABILITY_FOR_DATA EventTypeAnyOf = "UE_REACHABILITY_FOR_DATA"
	UE_REACHABILITY_FOR_SMS EventTypeAnyOf = "UE_REACHABILITY_FOR_SMS"
	LOCATION_REPORTING EventTypeAnyOf = "LOCATION_REPORTING"
	CHANGE_OF_SUPI_PEI_ASSOCIATION EventTypeAnyOf = "CHANGE_OF_SUPI_PEI_ASSOCIATION"
	ROAMING_STATUS EventTypeAnyOf = "ROAMING_STATUS"
	COMMUNICATION_FAILURE EventTypeAnyOf = "COMMUNICATION_FAILURE"
	AVAILABILITY_AFTER_DDN_FAILURE EventTypeAnyOf = "AVAILABILITY_AFTER_DDN_FAILURE"
	CN_TYPE_CHANGE EventTypeAnyOf = "CN_TYPE_CHANGE"
	DL_DATA_DELIVERY_STATUS EventTypeAnyOf = "DL_DATA_DELIVERY_STATUS"
	PDN_CONNECTIVITY_STATUS EventTypeAnyOf = "PDN_CONNECTIVITY_STATUS"
	UE_CONNECTION_MANAGEMENT_STATE EventTypeAnyOf = "UE_CONNECTION_MANAGEMENT_STATE"
	ACCESS_TYPE_REPORT EventTypeAnyOf = "ACCESS_TYPE_REPORT"
	REGISTRATION_STATE_REPORT EventTypeAnyOf = "REGISTRATION_STATE_REPORT"
	CONNECTIVITY_STATE_REPORT EventTypeAnyOf = "CONNECTIVITY_STATE_REPORT"
	TYPE_ALLOCATION_CODE_REPORT EventTypeAnyOf = "TYPE_ALLOCATION_CODE_REPORT"
	FREQUENT_MOBILITY_REGISTRATION_REPORT EventTypeAnyOf = "FREQUENT_MOBILITY_REGISTRATION_REPORT"
	PDU_SES_REL EventTypeAnyOf = "PDU_SES_REL"
	PDU_SES_EST EventTypeAnyOf = "PDU_SES_EST"
	UE_MEMORY_AVAILABLE_FOR_SMS EventTypeAnyOf = "UE_MEMORY_AVAILABLE_FOR_SMS"
)

// All allowed values of EventTypeAnyOf enum
var AllowedEventTypeAnyOfEnumValues = []EventTypeAnyOf{
	"LOSS_OF_CONNECTIVITY",
	"UE_REACHABILITY_FOR_DATA",
	"UE_REACHABILITY_FOR_SMS",
	"LOCATION_REPORTING",
	"CHANGE_OF_SUPI_PEI_ASSOCIATION",
	"ROAMING_STATUS",
	"COMMUNICATION_FAILURE",
	"AVAILABILITY_AFTER_DDN_FAILURE",
	"CN_TYPE_CHANGE",
	"DL_DATA_DELIVERY_STATUS",
	"PDN_CONNECTIVITY_STATUS",
	"UE_CONNECTION_MANAGEMENT_STATE",
	"ACCESS_TYPE_REPORT",
	"REGISTRATION_STATE_REPORT",
	"CONNECTIVITY_STATE_REPORT",
	"TYPE_ALLOCATION_CODE_REPORT",
	"FREQUENT_MOBILITY_REGISTRATION_REPORT",
	"PDU_SES_REL",
	"PDU_SES_EST",
	"UE_MEMORY_AVAILABLE_FOR_SMS",
}

func (v *EventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := EventTypeAnyOf(value)
	for _, existing := range AllowedEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid EventTypeAnyOf", value)
}

// NewEventTypeAnyOfFromValue returns a pointer to a valid EventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewEventTypeAnyOfFromValue(v string) (*EventTypeAnyOf, error) {
	ev := EventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for EventTypeAnyOf: valid values are %v", v, AllowedEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v EventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to EventType_anyOf value
func (v EventTypeAnyOf) Ptr() *EventTypeAnyOf {
	return &v
}

type NullableEventTypeAnyOf struct {
	value *EventTypeAnyOf
	isSet bool
}

func (v NullableEventTypeAnyOf) Get() *EventTypeAnyOf {
	return v.value
}

func (v *NullableEventTypeAnyOf) Set(val *EventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventTypeAnyOf(val *EventTypeAnyOf) *NullableEventTypeAnyOf {
	return &NullableEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

