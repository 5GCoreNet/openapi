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

// TriggerCategory struct for TriggerCategory
type TriggerCategory struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TriggerCategory) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TriggerCategory)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TriggerCategory) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTriggerCategory struct {
	value *TriggerCategory
	isSet bool
}

func (v NullableTriggerCategory) Get() *TriggerCategory {
	return v.value
}

func (v *NullableTriggerCategory) Set(val *TriggerCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerCategory(val *TriggerCategory) *NullableTriggerCategory {
	return &NullableTriggerCategory{value: val, isSet: true}
}

func (v NullableTriggerCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
