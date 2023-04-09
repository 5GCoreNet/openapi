/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
	"fmt"
)

// StationaryIndication Possible values are: - STATIONARY: Identifies the UE is stationary - MOBILE: Identifies the UE is mobile
type StationaryIndication struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *StationaryIndication) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(StationaryIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *StationaryIndication) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableStationaryIndication struct {
	value *StationaryIndication
	isSet bool
}

func (v NullableStationaryIndication) Get() *StationaryIndication {
	return v.value
}

func (v *NullableStationaryIndication) Set(val *StationaryIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableStationaryIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableStationaryIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStationaryIndication(val *StationaryIndication) *NullableStationaryIndication {
	return &NullableStationaryIndication{value: val, isSet: true}
}

func (v NullableStationaryIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStationaryIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
