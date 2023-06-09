/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// StationaryIndicationRm This enumeration is defined in the same way as the 'StationaryIndication' enumeration, but with the OpenAPI 'nullable: true' property.\"
type StationaryIndicationRm struct {
	NullValue            *NullValue
	StationaryIndication *StationaryIndication
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *StationaryIndicationRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue)
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	// try to unmarshal JSON data into StationaryIndication
	err = json.Unmarshal(data, &dst.StationaryIndication)
	if err == nil {
		jsonStationaryIndication, _ := json.Marshal(dst.StationaryIndication)
		if string(jsonStationaryIndication) == "{}" { // empty struct
			dst.StationaryIndication = nil
		} else {
			return nil // data stored in dst.StationaryIndication, return on the first match
		}
	} else {
		dst.StationaryIndication = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(StationaryIndicationRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *StationaryIndicationRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.StationaryIndication != nil {
		return json.Marshal(&src.StationaryIndication)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableStationaryIndicationRm struct {
	value *StationaryIndicationRm
	isSet bool
}

func (v NullableStationaryIndicationRm) Get() *StationaryIndicationRm {
	return v.value
}

func (v *NullableStationaryIndicationRm) Set(val *StationaryIndicationRm) {
	v.value = val
	v.isSet = true
}

func (v NullableStationaryIndicationRm) IsSet() bool {
	return v.isSet
}

func (v *NullableStationaryIndicationRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStationaryIndicationRm(val *StationaryIndicationRm) *NullableStationaryIndicationRm {
	return &NullableStationaryIndicationRm{value: val, isSet: true}
}

func (v NullableStationaryIndicationRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStationaryIndicationRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
