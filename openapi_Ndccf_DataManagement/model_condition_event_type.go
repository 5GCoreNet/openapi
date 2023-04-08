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

// ConditionEventType Indicates whether a notification is due to the NF Instance to start or stop being part of a condition for a subscription to a set of NFs 
type ConditionEventType struct {
	ConditionEventTypeAnyOf *ConditionEventTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ConditionEventType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ConditionEventTypeAnyOf
	err = json.Unmarshal(data, &dst.ConditionEventTypeAnyOf);
	if err == nil {
		jsonConditionEventTypeAnyOf, _ := json.Marshal(dst.ConditionEventTypeAnyOf)
		if string(jsonConditionEventTypeAnyOf) == "{}" { // empty struct
			dst.ConditionEventTypeAnyOf = nil
		} else {
			return nil // data stored in dst.ConditionEventTypeAnyOf, return on the first match
		}
	} else {
		dst.ConditionEventTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			return nil // data stored in dst.String, return on the first match
		}
	} else {
		dst.String = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ConditionEventType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ConditionEventType) MarshalJSON() ([]byte, error) {
	if src.ConditionEventTypeAnyOf != nil {
		return json.Marshal(&src.ConditionEventTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableConditionEventType struct {
	value *ConditionEventType
	isSet bool
}

func (v NullableConditionEventType) Get() *ConditionEventType {
	return v.value
}

func (v *NullableConditionEventType) Set(val *ConditionEventType) {
	v.value = val
	v.isSet = true
}

func (v NullableConditionEventType) IsSet() bool {
	return v.isSet
}

func (v *NullableConditionEventType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConditionEventType(val *ConditionEventType) *NullableConditionEventType {
	return &NullableConditionEventType{value: val, isSet: true}
}

func (v NullableConditionEventType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConditionEventType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


