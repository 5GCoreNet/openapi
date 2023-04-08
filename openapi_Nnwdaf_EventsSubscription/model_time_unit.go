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

// TimeUnit Possible values are: - MINUTE: Time unit is per minute. - HOUR: Time unit is per hour. - DAY: Time unit is per day. 
type TimeUnit struct {
	TimeUnitAnyOf *TimeUnitAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TimeUnit) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into TimeUnitAnyOf
	err = json.Unmarshal(data, &dst.TimeUnitAnyOf);
	if err == nil {
		jsonTimeUnitAnyOf, _ := json.Marshal(dst.TimeUnitAnyOf)
		if string(jsonTimeUnitAnyOf) == "{}" { // empty struct
			dst.TimeUnitAnyOf = nil
		} else {
			return nil // data stored in dst.TimeUnitAnyOf, return on the first match
		}
	} else {
		dst.TimeUnitAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(TimeUnit)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TimeUnit) MarshalJSON() ([]byte, error) {
	if src.TimeUnitAnyOf != nil {
		return json.Marshal(&src.TimeUnitAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
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


