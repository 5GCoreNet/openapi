/*
SS_NetworkResourceMonitoring

API for SEAL Network Resource Monitoring.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceMonitoring

import (
	"encoding/json"
	"fmt"
)

// MeasurementData Presents the aggregated measurement data.
type MeasurementData struct {
	Interface *interface{}
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MeasurementData) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into interface{}
	err = json.Unmarshal(data, &dst.Interface);
	if err == nil {
		jsonInterface, _ := json.Marshal(dst.Interface)
		if string(jsonInterface) == "{}" { // empty struct
			dst.Interface = nil
		} else {
			return nil // data stored in dst.Interface, return on the first match
		}
	} else {
		dst.Interface = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(MeasurementData)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MeasurementData) MarshalJSON() ([]byte, error) {
	if src.Interface != nil {
		return json.Marshal(&src.Interface)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMeasurementData struct {
	value *MeasurementData
	isSet bool
}

func (v NullableMeasurementData) Get() *MeasurementData {
	return v.value
}

func (v *NullableMeasurementData) Set(val *MeasurementData) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementData) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementData(val *MeasurementData) *NullableMeasurementData {
	return &NullableMeasurementData{value: val, isSet: true}
}

func (v NullableMeasurementData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


