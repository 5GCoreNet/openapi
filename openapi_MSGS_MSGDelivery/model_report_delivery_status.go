/*
MSGS_MSGDelivery

API for MSGG MSGin5G Server Message Delivery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSGS_MSGDelivery

import (
	"encoding/json"
	"fmt"
)

// ReportDeliveryStatus Possible values are: - REPT_DELY_SUCCESS: Indicates that the report delivery is successful. - REPT_DELY_FAILED: Indicates that the report delivery is failed. 
type ReportDeliveryStatus struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportDeliveryStatus) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportDeliveryStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportDeliveryStatus) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportDeliveryStatus struct {
	value *ReportDeliveryStatus
	isSet bool
}

func (v NullableReportDeliveryStatus) Get() *ReportDeliveryStatus {
	return v.value
}

func (v *NullableReportDeliveryStatus) Set(val *ReportDeliveryStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableReportDeliveryStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableReportDeliveryStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportDeliveryStatus(val *ReportDeliveryStatus) *NullableReportDeliveryStatus {
	return &NullableReportDeliveryStatus{value: val, isSet: true}
}

func (v NullableReportDeliveryStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportDeliveryStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


