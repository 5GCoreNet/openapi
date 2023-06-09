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

// PduSessionTypeRm PduSessionType indicates the type of a PDU session. It shall comply with the provisions defined in table 5.4.3.3-1 but with the OpenAPI \"nullable: true\" property.
type PduSessionTypeRm struct {
	NullValue      *NullValue
	PduSessionType *PduSessionType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PduSessionTypeRm) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into PduSessionType
	err = json.Unmarshal(data, &dst.PduSessionType)
	if err == nil {
		jsonPduSessionType, _ := json.Marshal(dst.PduSessionType)
		if string(jsonPduSessionType) == "{}" { // empty struct
			dst.PduSessionType = nil
		} else {
			return nil // data stored in dst.PduSessionType, return on the first match
		}
	} else {
		dst.PduSessionType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(PduSessionTypeRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PduSessionTypeRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.PduSessionType != nil {
		return json.Marshal(&src.PduSessionType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePduSessionTypeRm struct {
	value *PduSessionTypeRm
	isSet bool
}

func (v NullablePduSessionTypeRm) Get() *PduSessionTypeRm {
	return v.value
}

func (v *NullablePduSessionTypeRm) Set(val *PduSessionTypeRm) {
	v.value = val
	v.isSet = true
}

func (v NullablePduSessionTypeRm) IsSet() bool {
	return v.isSet
}

func (v *NullablePduSessionTypeRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePduSessionTypeRm(val *PduSessionTypeRm) *NullablePduSessionTypeRm {
	return &NullablePduSessionTypeRm{value: val, isSet: true}
}

func (v NullablePduSessionTypeRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePduSessionTypeRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
