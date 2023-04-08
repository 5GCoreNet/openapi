/*
Unified Data Repository Service API file for structured data for exposure

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Exposure_Data

import (
	"encoding/json"
	"fmt"
)

// PduSessionStatus Possible values are - \"ACTIVE\" - \"RELEASED\" 
type PduSessionStatus struct {
	PduSessionStatusAnyOf *PduSessionStatusAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PduSessionStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PduSessionStatusAnyOf
	err = json.Unmarshal(data, &dst.PduSessionStatusAnyOf);
	if err == nil {
		jsonPduSessionStatusAnyOf, _ := json.Marshal(dst.PduSessionStatusAnyOf)
		if string(jsonPduSessionStatusAnyOf) == "{}" { // empty struct
			dst.PduSessionStatusAnyOf = nil
		} else {
			return nil // data stored in dst.PduSessionStatusAnyOf, return on the first match
		}
	} else {
		dst.PduSessionStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PduSessionStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PduSessionStatus) MarshalJSON() ([]byte, error) {
	if src.PduSessionStatusAnyOf != nil {
		return json.Marshal(&src.PduSessionStatusAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePduSessionStatus struct {
	value *PduSessionStatus
	isSet bool
}

func (v NullablePduSessionStatus) Get() *PduSessionStatus {
	return v.value
}

func (v *NullablePduSessionStatus) Set(val *PduSessionStatus) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionStatus) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionStatus(val *PduSessionStatus) *NullablePduSessionStatus {
	return &NullablePduSessionStatus{value: val, isSet: true}
}

func (v NullablePduSessionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


