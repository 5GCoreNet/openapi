/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// TriggerCategory struct for TriggerCategory
type TriggerCategory struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TriggerCategory) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TriggerCategory)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TriggerCategory) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
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


