/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
	"fmt"
)

// ChargingCharacteristicsSelectionMode struct for ChargingCharacteristicsSelectionMode
type ChargingCharacteristicsSelectionMode struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ChargingCharacteristicsSelectionMode) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ChargingCharacteristicsSelectionMode)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ChargingCharacteristicsSelectionMode) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableChargingCharacteristicsSelectionMode struct {
	value *ChargingCharacteristicsSelectionMode
	isSet bool
}

func (v NullableChargingCharacteristicsSelectionMode) Get() *ChargingCharacteristicsSelectionMode {
	return v.value
}

func (v *NullableChargingCharacteristicsSelectionMode) Set(val *ChargingCharacteristicsSelectionMode) {
	v.value = val
	v.isSet = true
}

func (v NullableChargingCharacteristicsSelectionMode) IsSet() bool {
	return v.isSet
}

func (v *NullableChargingCharacteristicsSelectionMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableChargingCharacteristicsSelectionMode(val *ChargingCharacteristicsSelectionMode) *NullableChargingCharacteristicsSelectionMode {
	return &NullableChargingCharacteristicsSelectionMode{value: val, isSet: true}
}

func (v NullableChargingCharacteristicsSelectionMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableChargingCharacteristicsSelectionMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


