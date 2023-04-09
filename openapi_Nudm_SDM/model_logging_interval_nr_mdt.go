/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
	"fmt"
)

// LoggingIntervalNrMdt The enumeration LoggingIntervalNrMdt defines Logging Interval in NR for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.18-1.
type LoggingIntervalNrMdt struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LoggingIntervalNrMdt) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(LoggingIntervalNrMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LoggingIntervalNrMdt) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLoggingIntervalNrMdt struct {
	value *LoggingIntervalNrMdt
	isSet bool
}

func (v NullableLoggingIntervalNrMdt) Get() *LoggingIntervalNrMdt {
	return v.value
}

func (v *NullableLoggingIntervalNrMdt) Set(val *LoggingIntervalNrMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableLoggingIntervalNrMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableLoggingIntervalNrMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoggingIntervalNrMdt(val *LoggingIntervalNrMdt) *NullableLoggingIntervalNrMdt {
	return &NullableLoggingIntervalNrMdt{value: val, isSet: true}
}

func (v NullableLoggingIntervalNrMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoggingIntervalNrMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
