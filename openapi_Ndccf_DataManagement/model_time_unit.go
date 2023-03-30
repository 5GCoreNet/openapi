/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
	"fmt"
)

// TimeUnit Possible values are: - MINUTE: Time unit is per minute. - HOUR: Time unit is per hour. - DAY: Time unit is per day. 
type TimeUnit struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TimeUnit) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TimeUnit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TimeUnit) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTimeUnit struct {
	value *TimeUnit
	isSet bool
}

func (v NullableTimeUnit) Get() *TimeUnit {
	return v.value
}

func (v *NullableTimeUnit) Set(val *TimeUnit) {
	v.value = val
	v.isSet = true
}

func (v NullableTimeUnit) IsSet() bool {
	return v.isSet
}

func (v *NullableTimeUnit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTimeUnit(val *TimeUnit) *NullableTimeUnit {
	return &NullableTimeUnit{value: val, isSet: true}
}

func (v NullableTimeUnit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTimeUnit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


