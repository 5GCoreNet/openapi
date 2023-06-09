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

// RequestIndication Request Indication in Update (SM context) service operation. Possible values are - UE_REQ_PDU_SES_MOD - UE_REQ_PDU_SES_REL - PDU_SES_MOB - NW_REQ_PDU_SES_AUTH - NW_REQ_PDU_SES_MOD - NW_REQ_PDU_SES_REL - EBI_ASSIGNMENT_REQ - REL_DUE_TO_5G_AN_REQUEST
type RequestIndication struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RequestIndication) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RequestIndication)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RequestIndication) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRequestIndication struct {
	value *RequestIndication
	isSet bool
}

func (v NullableRequestIndication) Get() *RequestIndication {
	return v.value
}

func (v *NullableRequestIndication) Set(val *RequestIndication) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestIndication) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestIndication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestIndication(val *RequestIndication) *NullableRequestIndication {
	return &NullableRequestIndication{value: val, isSet: true}
}

func (v NullableRequestIndication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestIndication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
