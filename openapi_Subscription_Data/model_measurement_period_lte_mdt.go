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

// MeasurementPeriodLteMdt The enumeration MeasurementPeriodLteMdt defines Measurement period LTE for MDT in the trace.  See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.16-1. 
type MeasurementPeriodLteMdt struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MeasurementPeriodLteMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MeasurementPeriodLteMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MeasurementPeriodLteMdt) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMeasurementPeriodLteMdt struct {
	value *MeasurementPeriodLteMdt
	isSet bool
}

func (v NullableMeasurementPeriodLteMdt) Get() *MeasurementPeriodLteMdt {
	return v.value
}

func (v *NullableMeasurementPeriodLteMdt) Set(val *MeasurementPeriodLteMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementPeriodLteMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementPeriodLteMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementPeriodLteMdt(val *MeasurementPeriodLteMdt) *NullableMeasurementPeriodLteMdt {
	return &NullableMeasurementPeriodLteMdt{value: val, isSet: true}
}

func (v NullableMeasurementPeriodLteMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementPeriodLteMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


