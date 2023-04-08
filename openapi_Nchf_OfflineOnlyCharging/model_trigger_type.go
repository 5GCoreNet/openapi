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

// TriggerType struct for TriggerType
type TriggerType struct {
	TriggerTypeAnyOf *TriggerTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TriggerType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TriggerTypeAnyOf
	err = json.Unmarshal(data, &dst.TriggerTypeAnyOf);
	if err == nil {
		jsonTriggerTypeAnyOf, _ := json.Marshal(dst.TriggerTypeAnyOf)
		if string(jsonTriggerTypeAnyOf) == "{}" { // empty struct
			dst.TriggerTypeAnyOf = nil
		} else {
			return nil // data stored in dst.TriggerTypeAnyOf, return on the first match
		}
	} else {
		dst.TriggerTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
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

	return fmt.Errorf("data failed to match schemas in anyOf(TriggerType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TriggerType) MarshalJSON() ([]byte, error) {
	if src.TriggerTypeAnyOf != nil {
		return json.Marshal(&src.TriggerTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTriggerType struct {
	value *TriggerType
	isSet bool
}

func (v NullableTriggerType) Get() *TriggerType {
	return v.value
}

func (v *NullableTriggerType) Set(val *TriggerType) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerType) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerType(val *TriggerType) *NullableTriggerType {
	return &NullableTriggerType{value: val, isSet: true}
}

func (v NullableTriggerType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


