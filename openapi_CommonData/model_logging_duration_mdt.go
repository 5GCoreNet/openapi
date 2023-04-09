/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// LoggingDurationMdt The enumeration LoggingIntervalMdt defines Logging Interval for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.12-1.
type LoggingDurationMdt struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LoggingDurationMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
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

	return fmt.Errorf("data failed to match schemas in anyOf(LoggingDurationMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LoggingDurationMdt) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLoggingDurationMdt struct {
	value *LoggingDurationMdt
	isSet bool
}

func (v NullableLoggingDurationMdt) Get() *LoggingDurationMdt {
	return v.value
}

func (v *NullableLoggingDurationMdt) Set(val *LoggingDurationMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingDurationMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingDurationMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingDurationMdt(val *LoggingDurationMdt) *NullableLoggingDurationMdt {
	return &NullableLoggingDurationMdt{value: val, isSet: true}
}

func (v NullableLoggingDurationMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingDurationMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
