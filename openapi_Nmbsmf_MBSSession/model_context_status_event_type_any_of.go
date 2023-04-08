/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
	"fmt"
)

// ContextStatusEventTypeAnyOf the model 'ContextStatusEventTypeAnyOf'
type ContextStatusEventTypeAnyOf string

// List of ContextStatusEventType_anyOf
const (
	QOS_INFO ContextStatusEventTypeAnyOf = "QOS_INFO"
	STATUS_INFO ContextStatusEventTypeAnyOf = "STATUS_INFO"
	SERVICE_AREA_INFO ContextStatusEventTypeAnyOf = "SERVICE_AREA_INFO"
	SESSION_RELEASE ContextStatusEventTypeAnyOf = "SESSION_RELEASE"
	MULT_TRANS_ADD_CHANGE ContextStatusEventTypeAnyOf = "MULT_TRANS_ADD_CHANGE"
	SECURITY_INFO ContextStatusEventTypeAnyOf = "SECURITY_INFO"
)

// All allowed values of ContextStatusEventTypeAnyOf enum
var AllowedContextStatusEventTypeAnyOfEnumValues = []ContextStatusEventTypeAnyOf{
	"QOS_INFO",
	"STATUS_INFO",
	"SERVICE_AREA_INFO",
	"SESSION_RELEASE",
	"MULT_TRANS_ADD_CHANGE",
	"SECURITY_INFO",
}

func (v *ContextStatusEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ContextStatusEventTypeAnyOf(value)
	for _, existing := range AllowedContextStatusEventTypeAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ContextStatusEventTypeAnyOf", value)
}

// NewContextStatusEventTypeAnyOfFromValue returns a pointer to a valid ContextStatusEventTypeAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewContextStatusEventTypeAnyOfFromValue(v string) (*ContextStatusEventTypeAnyOf, error) {
	ev := ContextStatusEventTypeAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ContextStatusEventTypeAnyOf: valid values are %v", v, AllowedContextStatusEventTypeAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ContextStatusEventTypeAnyOf) IsValid() bool {
	for _, existing := range AllowedContextStatusEventTypeAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ContextStatusEventType_anyOf value
func (v ContextStatusEventTypeAnyOf) Ptr() *ContextStatusEventTypeAnyOf {
	return &v
}

type NullableContextStatusEventTypeAnyOf struct {
	value *ContextStatusEventTypeAnyOf
	isSet bool
}

func (v NullableContextStatusEventTypeAnyOf) Get() *ContextStatusEventTypeAnyOf {
	return v.value
}

func (v *NullableContextStatusEventTypeAnyOf) Set(val *ContextStatusEventTypeAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableContextStatusEventTypeAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableContextStatusEventTypeAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextStatusEventTypeAnyOf(val *ContextStatusEventTypeAnyOf) *NullableContextStatusEventTypeAnyOf {
	return &NullableContextStatusEventTypeAnyOf{value: val, isSet: true}
}

func (v NullableContextStatusEventTypeAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextStatusEventTypeAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

