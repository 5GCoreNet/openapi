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

// MeasurementDataType Possible values are: - DL_DELAY: Downlink packet delay. - UL_DELAY: Uplink packet delay. - RT_DELAY: Round trip packet delay. - AVG_PLR: Average packet loss rate. - AVG_DATA_RATE: Average data rate. - MAX_DATA_RATE: Maximum data rate. - AVG_DL_TRAFFIC_VOLUME: Average downlink traffic volume. - AVG_UL_TRAFFIC_VOLUME: Average uplink traffic volume.
type MeasurementDataType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *MeasurementDataType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(MeasurementDataType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *MeasurementDataType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableMeasurementDataType struct {
	value *MeasurementDataType
	isSet bool
}

func (v NullableMeasurementDataType) Get() *MeasurementDataType {
	return v.value
}

func (v *NullableMeasurementDataType) Set(val *MeasurementDataType) {
	v.value = val
	v.isSet = true
}

func (v NullableMeasurementDataType) IsSet() bool {
	return v.isSet
}

func (v *NullableMeasurementDataType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeasurementDataType(val *MeasurementDataType) *NullableMeasurementDataType {
	return &NullableMeasurementDataType{value: val, isSet: true}
}

func (v NullableMeasurementDataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeasurementDataType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
