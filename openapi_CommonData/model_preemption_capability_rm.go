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

// PreemptionCapabilityRm This enumeration is defined in the same way as the 'PreemptionCapability' enumeration, but with the OpenAPI 'nullable: true' property.
type PreemptionCapabilityRm struct {
	NullValue            *NullValue
	PreemptionCapability *PreemptionCapability
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PreemptionCapabilityRm) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into PreemptionCapability
	err = json.Unmarshal(data, &dst.PreemptionCapability)
	if err == nil {
		jsonPreemptionCapability, _ := json.Marshal(dst.PreemptionCapability)
		if string(jsonPreemptionCapability) == "{}" { // empty struct
			dst.PreemptionCapability = nil
		} else {
			return nil // data stored in dst.PreemptionCapability, return on the first match
		}
	} else {
		dst.PreemptionCapability = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PreemptionCapabilityRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PreemptionCapabilityRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.PreemptionCapability != nil {
		return json.Marshal(&src.PreemptionCapability)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePreemptionCapabilityRm struct {
	value *PreemptionCapabilityRm
	isSet bool
}

func (v NullablePreemptionCapabilityRm) Get() *PreemptionCapabilityRm {
	return v.value
}

func (v *NullablePreemptionCapabilityRm) Set(val *PreemptionCapabilityRm) {
	v.value = val
	v.isSet = true
}

func (v NullablePreemptionCapabilityRm) IsSet() bool {
	return v.isSet
}

func (v *NullablePreemptionCapabilityRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreemptionCapabilityRm(val *PreemptionCapabilityRm) *NullablePreemptionCapabilityRm {
	return &NullablePreemptionCapabilityRm{value: val, isSet: true}
}

func (v NullablePreemptionCapabilityRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreemptionCapabilityRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
