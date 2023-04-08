/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
	"fmt"
)

// AdditionalQosFlowInfo The enumeration AdditionalQosFlowInfo provides additional QoS flow information (see clause  9.3.1.12 3GPP TS 38.413 [11]). It shall comply with the provisions defined in table 5.5.3.12-1. 
type AdditionalQosFlowInfo struct {
	AnyOfstringstring *AnyOfstringstring
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AdditionalQosFlowInfo) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AnyOfstringstring
	err = json.Unmarshal(data, &dst.AnyOfstringstring);
	if err == nil {
		jsonAnyOfstringstring, _ := json.Marshal(dst.AnyOfstringstring)
		if string(jsonAnyOfstringstring) == "{}" { // empty struct
			dst.AnyOfstringstring = nil
		} else {
			return nil // data stored in dst.AnyOfstringstring, return on the first match
		}
	} else {
		dst.AnyOfstringstring = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AdditionalQosFlowInfo)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AdditionalQosFlowInfo) MarshalJSON() ([]byte, error) {
	if src.AnyOfstringstring != nil {
		return json.Marshal(&src.AnyOfstringstring)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAdditionalQosFlowInfo struct {
	value *AdditionalQosFlowInfo
	isSet bool
}

func (v NullableAdditionalQosFlowInfo) Get() *AdditionalQosFlowInfo {
	return v.value
}

func (v *NullableAdditionalQosFlowInfo) Set(val *AdditionalQosFlowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAdditionalQosFlowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAdditionalQosFlowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAdditionalQosFlowInfo(val *AdditionalQosFlowInfo) *NullableAdditionalQosFlowInfo {
	return &NullableAdditionalQosFlowInfo{value: val, isSet: true}
}

func (v NullableAdditionalQosFlowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAdditionalQosFlowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


