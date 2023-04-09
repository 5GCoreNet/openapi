/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// SensorMeasurement The enumeration SensorMeasurement defines sensor measurement type for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.7-1.
type SensorMeasurement struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SensorMeasurement) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(SensorMeasurement)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SensorMeasurement) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSensorMeasurement struct {
	value *SensorMeasurement
	isSet bool
}

func (v NullableSensorMeasurement) Get() *SensorMeasurement {
	return v.value
}

func (v *NullableSensorMeasurement) Set(val *SensorMeasurement) {
	v.value = val
	v.isSet = true
}

func (v NullableSensorMeasurement) IsSet() bool {
	return v.isSet
}

func (v *NullableSensorMeasurement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSensorMeasurement(val *SensorMeasurement) *NullableSensorMeasurement {
	return &NullableSensorMeasurement{value: val, isSet: true}
}

func (v NullableSensorMeasurement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSensorMeasurement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
