/*
3gpp-analyticsexposure

API for Analytics Exposure.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AnalyticsExposure

import (
	"encoding/json"
	"fmt"
)

// DispersionType - Possible values are:   - DVDA: Data Volume Dispersion Analytics.   - TDA: Transactions Dispersion Analytics.   - DVDA_AND_TDA: Data Volume Dispersion Analytics and Transactions Dispersion Analytics. 
type DispersionType struct {
	String *string
}

// stringAsDispersionType is a convenience function that returns string wrapped in DispersionType
func StringAsDispersionType(v *string) DispersionType {
	return DispersionType{
		String: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *DispersionType) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into String
	err = newStrictDecoder(data).Decode(&dst.String)
	if err == nil {
		jsonString, _ := json.Marshal(dst.String)
		if string(jsonString) == "{}" { // empty struct
			dst.String = nil
		} else {
			match++
		}
	} else {
		dst.String = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.String = nil

		return fmt.Errorf("data matches more than one schema in oneOf(DispersionType)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(DispersionType)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src DispersionType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *DispersionType) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.String != nil {
		return obj.String
	}

	// all schemas are nil
	return nil
}

type NullableDispersionType struct {
	value *DispersionType
	isSet bool
}

func (v NullableDispersionType) Get() *DispersionType {
	return v.value
}

func (v *NullableDispersionType) Set(val *DispersionType) {
	v.value = val
	v.isSet = true
}

func (v NullableDispersionType) IsSet() bool {
	return v.isSet
}

func (v *NullableDispersionType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDispersionType(val *DispersionType) *NullableDispersionType {
	return &NullableDispersionType{value: val, isSet: true}
}

func (v NullableDispersionType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDispersionType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

