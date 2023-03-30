/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// PreemptionControlInformation Represents Pre-emption control information.
type PreemptionControlInformation struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PreemptionControlInformation) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PreemptionControlInformation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PreemptionControlInformation) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePreemptionControlInformation struct {
	value *PreemptionControlInformation
	isSet bool
}

func (v NullablePreemptionControlInformation) Get() *PreemptionControlInformation {
	return v.value
}

func (v *NullablePreemptionControlInformation) Set(val *PreemptionControlInformation) {
	v.value = val
	v.isSet = true
}

func (v NullablePreemptionControlInformation) IsSet() bool {
	return v.isSet
}

func (v *NullablePreemptionControlInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreemptionControlInformation(val *PreemptionControlInformation) *NullablePreemptionControlInformation {
	return &NullablePreemptionControlInformation{value: val, isSet: true}
}

func (v NullablePreemptionControlInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreemptionControlInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


