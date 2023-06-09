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

// N4MessageType N4 Message Type. Possible values are   - PFCP_SES_EST_REQ   - PFCP_SES_EST_RSP   - PFCP_SES_MOD_REQ   - PFCP_SES_MOD_RSP   - PFCP_SES_DEL_REQ   - PFCP_SES_DEL_RSP   - PFCP_SES_REP_REQ   - PFCP_SES_REP_RSP
type N4MessageType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *N4MessageType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(N4MessageType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *N4MessageType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableN4MessageType struct {
	value *N4MessageType
	isSet bool
}

func (v NullableN4MessageType) Get() *N4MessageType {
	return v.value
}

func (v *NullableN4MessageType) Set(val *N4MessageType) {
	v.value = val
	v.isSet = true
}

func (v NullableN4MessageType) IsSet() bool {
	return v.isSet
}

func (v *NullableN4MessageType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableN4MessageType(val *N4MessageType) *NullableN4MessageType {
	return &NullableN4MessageType{value: val, isSet: true}
}

func (v NullableN4MessageType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableN4MessageType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
