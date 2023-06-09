/*
EES UE Location Information_API

API for EES UE Location Information.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_UELocation

import (
	"encoding/json"
	"fmt"
)

// Accuracy Possible values are - CGI_ECGI: The SCS/AS requests to be notified using cell level location accuracy. - ENODEB: The SCS/AS requests to be notified using eNodeB level location accuracy. - TA_RA: The SCS/AS requests to be notified using TA/RA level location accuracy. - PLMN: The SCS/AS requests to be notified using PLMN level location accuracy. - TWAN_ID: The SCS/AS requests to be notified using TWAN identifier level location accuracy. - GEO_AREA: The SCS/AS requests to be notified using the geographical area accuracy. - CIVIC_ADDR: The SCS/AS requests to be notified using the civic address accuracy.
type Accuracy struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *Accuracy) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(Accuracy)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *Accuracy) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccuracy struct {
	value *Accuracy
	isSet bool
}

func (v NullableAccuracy) Get() *Accuracy {
	return v.value
}

func (v *NullableAccuracy) Set(val *Accuracy) {
	v.value = val
	v.isSet = true
}

func (v NullableAccuracy) IsSet() bool {
	return v.isSet
}

func (v *NullableAccuracy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccuracy(val *Accuracy) *NullableAccuracy {
	return &NullableAccuracy{value: val, isSet: true}
}

func (v NullableAccuracy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccuracy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
