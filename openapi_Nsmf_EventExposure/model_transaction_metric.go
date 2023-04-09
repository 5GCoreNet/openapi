/*
Nsmf_EventExposure

Session Management Event Exposure Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_EventExposure

import (
	"encoding/json"
	"fmt"
)

// TransactionMetric Possible values are: - PDU_SES_EST: PDU Session Establishment - PDU_SES_AUTH: PDU Session Authentication - PDU_SES_MODIF: PDU Session Modification - PDU_SES_REL: PDU Session Release
type TransactionMetric struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TransactionMetric) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TransactionMetric)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TransactionMetric) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTransactionMetric struct {
	value *TransactionMetric
	isSet bool
}

func (v NullableTransactionMetric) Get() *TransactionMetric {
	return v.value
}

func (v *NullableTransactionMetric) Set(val *TransactionMetric) {
	v.value = val
	v.isSet = true
}

func (v NullableTransactionMetric) IsSet() bool {
	return v.isSet
}

func (v *NullableTransactionMetric) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransactionMetric(val *TransactionMetric) *NullableTransactionMetric {
	return &NullableTransactionMetric{value: val, isSet: true}
}

func (v NullableTransactionMetric) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransactionMetric) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
