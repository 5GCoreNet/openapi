/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// PduSessionStatus Possible values are: - ACTIVATED: PDU Session status is activated. - DEACTIVATED: PDU Session status is deactivated.
type PduSessionStatus struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PduSessionStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PduSessionStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PduSessionStatus) MarshalJSON() ([]byte, error) {
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
