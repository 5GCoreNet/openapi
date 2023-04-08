/*
Nnwdaf_EventsSubscription

Nnwdaf_EventsSubscription Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_EventsSubscription

import (
	"encoding/json"
	"fmt"
)

// CongestionType Possible values are: - USER_PLANE: The congestion analytics type is User Plane.  - CONTROL_PLANE: The congestion analytics type is Control Plane. - USER_AND_CONTROL_PLANE: The congestion analytics type is User Plane and Control Plane. 
type CongestionType struct {
	CongestionTypeAnyOf *CongestionTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CongestionType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into CongestionTypeAnyOf
	err = json.Unmarshal(data, &dst.CongestionTypeAnyOf);
	if err == nil {
		jsonCongestionTypeAnyOf, _ := json.Marshal(dst.CongestionTypeAnyOf)
		if string(jsonCongestionTypeAnyOf) == "{}" { // empty struct
			dst.CongestionTypeAnyOf = nil
		} else {
			return nil // data stored in dst.CongestionTypeAnyOf, return on the first match
		}
	} else {
		dst.CongestionTypeAnyOf = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String);
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

	return fmt.Errorf("data failed to match schemas in anyOf(CongestionType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CongestionType) MarshalJSON() ([]byte, error) {
	if src.CongestionTypeAnyOf != nil {
		return json.Marshal(&src.CongestionTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCongestionType struct {
	value *CongestionType
	isSet bool
}

func (v NullableCongestionType) Get() *CongestionType {
	return v.value
}

func (v *NullableCongestionType) Set(val *CongestionType) {
	v.value = val
	v.isSet = true
}

func (v NullableCongestionType) IsSet() bool {
	return v.isSet
}

func (v *NullableCongestionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCongestionType(val *CongestionType) *NullableCongestionType {
	return &NullableCongestionType{value: val, isSet: true}
}

func (v NullableCongestionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCongestionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


