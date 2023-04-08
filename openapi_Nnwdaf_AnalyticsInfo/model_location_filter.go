/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
	"fmt"
)

// LocationFilter Describes the supported filters of LOCATION_REPORT event type
type LocationFilter struct {
	LocationFilterAnyOf *LocationFilterAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LocationFilter) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LocationFilterAnyOf
	err = json.Unmarshal(data, &dst.LocationFilterAnyOf);
	if err == nil {
		jsonLocationFilterAnyOf, _ := json.Marshal(dst.LocationFilterAnyOf)
		if string(jsonLocationFilterAnyOf) == "{}" { // empty struct
			dst.LocationFilterAnyOf = nil
		} else {
			return nil // data stored in dst.LocationFilterAnyOf, return on the first match
		}
	} else {
		dst.LocationFilterAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LocationFilter)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LocationFilter) MarshalJSON() ([]byte, error) {
	if src.LocationFilterAnyOf != nil {
		return json.Marshal(&src.LocationFilterAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLocationFilter struct {
	value *LocationFilter
	isSet bool
}

func (v NullableLocationFilter) Get() *LocationFilter {
	return v.value
}

func (v *NullableLocationFilter) Set(val *LocationFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableLocationFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableLocationFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocationFilter(val *LocationFilter) *NullableLocationFilter {
	return &NullableLocationFilter{value: val, isSet: true}
}

func (v NullableLocationFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocationFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


