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

// QosFlowAccessType Access type associated with a QoS Flow. Possible values are   - 3GPP   - NON_3GPP   - 3GPP_AND_NON_3GPP
type QosFlowAccessType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *QosFlowAccessType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(QosFlowAccessType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *QosFlowAccessType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableQosFlowAccessType struct {
	value *QosFlowAccessType
	isSet bool
}

func (v NullableQosFlowAccessType) Get() *QosFlowAccessType {
	return v.value
}

func (v *NullableQosFlowAccessType) Set(val *QosFlowAccessType) {
	v.value = val
	v.isSet = true
}

func (v NullableQosFlowAccessType) IsSet() bool {
	return v.isSet
}

func (v *NullableQosFlowAccessType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosFlowAccessType(val *QosFlowAccessType) *NullableQosFlowAccessType {
	return &NullableQosFlowAccessType{value: val, isSet: true}
}

func (v NullableQosFlowAccessType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosFlowAccessType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
