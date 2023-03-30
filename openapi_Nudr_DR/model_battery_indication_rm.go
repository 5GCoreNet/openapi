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

// BatteryIndicationRm This data type is defined in the same way as the 'BatteryIndication' data type, but with the OpenAPI 'nullable: true' property. 
type BatteryIndicationRm struct {
	BatteryIndication *BatteryIndication
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *BatteryIndicationRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into BatteryIndication
	err = json.Unmarshal(data, &dst.BatteryIndication);
	if err == nil {
		jsonBatteryIndication, _ := json.Marshal(dst.BatteryIndication)
		if string(jsonBatteryIndication) == "{}" { // empty struct
			dst.BatteryIndication = nil
		} else {
			return nil // data stored in dst.BatteryIndication, return on the first match
		}
	} else {
		dst.BatteryIndication = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
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

	return fmt.Errorf("data failed to match schemas in anyOf(BatteryIndicationRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *BatteryIndicationRm) MarshalJSON() ([]byte, error) {
	if src.BatteryIndication != nil {
		return json.Marshal(&src.BatteryIndication)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableBatteryIndicationRm struct {
	value *BatteryIndicationRm
	isSet bool
}

func (v NullableBatteryIndicationRm) Get() *BatteryIndicationRm {
	return v.value
}

func (v *NullableBatteryIndicationRm) Set(val *BatteryIndicationRm) {
	v.value = val
	v.isSet = true
}

func (v NullableBatteryIndicationRm) IsSet() bool {
	return v.isSet
}

func (v *NullableBatteryIndicationRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatteryIndicationRm(val *BatteryIndicationRm) *NullableBatteryIndicationRm {
	return &NullableBatteryIndicationRm{value: val, isSet: true}
}

func (v NullableBatteryIndicationRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatteryIndicationRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

