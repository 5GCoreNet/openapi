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

// MeasurementNrForMdt The enumeration MeasurementNrForMdt defines Measurements used for MDT in NR in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.6-1.
type MeasurementNrForMdt struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MeasurementNrForMdt) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(MeasurementNrForMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MeasurementNrForMdt) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMeasurementNrForMdt struct {
	value *MeasurementNrForMdt
	isSet bool
}

func (v NullableMeasurementNrForMdt) Get() *MeasurementNrForMdt {
	return v.value
}

func (v *NullableMeasurementNrForMdt) Set(val *MeasurementNrForMdt) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementNrForMdt) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementNrForMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementNrForMdt(val *MeasurementNrForMdt) *NullableMeasurementNrForMdt {
	return &NullableMeasurementNrForMdt{value: val, isSet: true}
}

func (v NullableMeasurementNrForMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementNrForMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
