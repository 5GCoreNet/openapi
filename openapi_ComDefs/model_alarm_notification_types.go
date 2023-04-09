/*
Common Type Definitions

OAS 3.0.1 specification of common type definitions in the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ComDefs

import (
	"encoding/json"
	"fmt"
)

// AlarmNotificationTypes the model 'AlarmNotificationTypes'
type AlarmNotificationTypes string

// List of AlarmNotificationTypes
const (
	NOTIFY_NEW_ALARM                       AlarmNotificationTypes = "notifyNewAlarm"
	NOTIFY_CHANGED_ALARM                   AlarmNotificationTypes = "notifyChangedAlarm"
	NOTIFY_CHANGED_ALARM_GENERAL           AlarmNotificationTypes = "notifyChangedAlarmGeneral"
	NOTIFY_ACK_STATE_CHANGED               AlarmNotificationTypes = "notifyAckStateChanged"
	NOTIFY_CORRELATED_NOTIFICATION_CHANGED AlarmNotificationTypes = "notifyCorrelatedNotificationChanged"
	NOTIFY_COMMENTS                        AlarmNotificationTypes = "notifyComments"
	NOTIFY_CLEARED_ALARM                   AlarmNotificationTypes = "notifyClearedAlarm"
	NOTIFY_ALARM_LIST_REBUILT              AlarmNotificationTypes = "notifyAlarmListRebuilt"
	NOTIFY_POTENTIAL_FAULTY_ALARM_LIST     AlarmNotificationTypes = "notifyPotentialFaultyAlarmList"
)

// All allowed values of AlarmNotificationTypes enum
var AllowedAlarmNotificationTypesEnumValues = []AlarmNotificationTypes{
	"notifyNewAlarm",
	"notifyChangedAlarm",
	"notifyChangedAlarmGeneral",
	"notifyAckStateChanged",
	"notifyCorrelatedNotificationChanged",
	"notifyComments",
	"notifyClearedAlarm",
	"notifyAlarmListRebuilt",
	"notifyPotentialFaultyAlarmList",
}

func (v *AlarmNotificationTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlarmNotificationTypes(value)
	for _, existing := range AllowedAlarmNotificationTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlarmNotificationTypes", value)
}

// NewAlarmNotificationTypesFromValue returns a pointer to a valid AlarmNotificationTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlarmNotificationTypesFromValue(v string) (*AlarmNotificationTypes, error) {
	ev := AlarmNotificationTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlarmNotificationTypes: valid values are %v", v, AllowedAlarmNotificationTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlarmNotificationTypes) IsValid() bool {
	for _, existing := range AllowedAlarmNotificationTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlarmNotificationTypes value
func (v AlarmNotificationTypes) Ptr() *AlarmNotificationTypes {
	return &v
}

type NullableAlarmNotificationTypes struct {
	value *AlarmNotificationTypes
	isSet bool
}

func (v NullableAlarmNotificationTypes) Get() *AlarmNotificationTypes {
	return v.value
}

func (v *NullableAlarmNotificationTypes) Set(val *AlarmNotificationTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmNotificationTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmNotificationTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmNotificationTypes(val *AlarmNotificationTypes) *NullableAlarmNotificationTypes {
	return &NullableAlarmNotificationTypes{value: val, isSet: true}
}

func (v NullableAlarmNotificationTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmNotificationTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
