/*
coslaNrm

OAS 3.0.1 specification of the Cosla NRM © 2020, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.3.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CoslaNrm

import (
	"encoding/json"
	"fmt"
)

// AlarmType the model 'AlarmType'
type AlarmType string

// List of AlarmType
const (
	COMMUNICATIONS_ALARM                    AlarmType = "COMMUNICATIONS_ALARM"
	QUALITY_OF_SERVICE_ALARM                AlarmType = "QUALITY_OF_SERVICE_ALARM"
	PROCESSING_ERROR_ALARM                  AlarmType = "PROCESSING_ERROR_ALARM"
	EQUIPMENT_ALARM                         AlarmType = "EQUIPMENT_ALARM"
	ENVIRONMENTAL_ALARM                     AlarmType = "ENVIRONMENTAL_ALARM"
	INTEGRITY_VIOLATION                     AlarmType = "INTEGRITY_VIOLATION"
	OPERATIONAL_VIOLATION                   AlarmType = "OPERATIONAL_VIOLATION"
	PHYSICAL_VIOLATION                      AlarmType = "PHYSICAL_VIOLATION"
	SECURITY_SERVICE_OR_MECHANISM_VIOLATION AlarmType = "SECURITY_SERVICE_OR_MECHANISM_VIOLATION"
	TIME_DOMAIN_VIOLATION                   AlarmType = "TIME_DOMAIN_VIOLATION"
)

// All allowed values of AlarmType enum
var AllowedAlarmTypeEnumValues = []AlarmType{
	"COMMUNICATIONS_ALARM",
	"QUALITY_OF_SERVICE_ALARM",
	"PROCESSING_ERROR_ALARM",
	"EQUIPMENT_ALARM",
	"ENVIRONMENTAL_ALARM",
	"INTEGRITY_VIOLATION",
	"OPERATIONAL_VIOLATION",
	"PHYSICAL_VIOLATION",
	"SECURITY_SERVICE_OR_MECHANISM_VIOLATION",
	"TIME_DOMAIN_VIOLATION",
}

func (v *AlarmType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AlarmType(value)
	for _, existing := range AllowedAlarmTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AlarmType", value)
}

// NewAlarmTypeFromValue returns a pointer to a valid AlarmType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAlarmTypeFromValue(v string) (*AlarmType, error) {
	ev := AlarmType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AlarmType: valid values are %v", v, AllowedAlarmTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AlarmType) IsValid() bool {
	for _, existing := range AllowedAlarmTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AlarmType value
func (v AlarmType) Ptr() *AlarmType {
	return &v
}

type NullableAlarmType struct {
	value *AlarmType
	isSet bool
}

func (v NullableAlarmType) Get() *AlarmType {
	return v.value
}

func (v *NullableAlarmType) Set(val *AlarmType) {
	v.value = val
	v.isSet = true
}

func (v NullableAlarmType) IsSet() bool {
	return v.isSet
}

func (v *NullableAlarmType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlarmType(val *AlarmType) *NullableAlarmType {
	return &NullableAlarmType{value: val, isSet: true}
}

func (v NullableAlarmType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlarmType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
