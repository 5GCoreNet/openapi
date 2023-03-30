/*
Ndcaf_DataReporting

Data Collection AF: Data Collection and Reporting Configuration API and Data Reporting API © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndcaf_DataReporting

import (
	"encoding/json"
	"fmt"
)

// ReportingConditionType The type of condition that triggers reporting by a data collection client to the Data Collection AF.
type ReportingConditionType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReportingConditionType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReportingConditionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReportingConditionType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReportingConditionType struct {
	value *ReportingConditionType
	isSet bool
}

func (v NullableReportingConditionType) Get() *ReportingConditionType {
	return v.value
}

func (v *NullableReportingConditionType) Set(val *ReportingConditionType) {
	v.value = val
	v.isSet = true
}

func (v NullableReportingConditionType) IsSet() bool {
	return v.isSet
}

func (v *NullableReportingConditionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReportingConditionType(val *ReportingConditionType) *NullableReportingConditionType {
	return &NullableReportingConditionType{value: val, isSet: true}
}

func (v NullableReportingConditionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReportingConditionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


