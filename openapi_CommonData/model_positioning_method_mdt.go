/*
Common Data Types

Common Data Types for Service Based Interfaces.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.   

API version: 1.5.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CommonData

import (
	"encoding/json"
	"fmt"
)

// PositioningMethodMdt The enumeration LoggingDurationMdt defines Logging Duration for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.13-1. 
type PositioningMethodMdt struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PositioningMethodMdt) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(PositioningMethodMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PositioningMethodMdt) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullablePositioningMethodMdt struct {
	value *PositioningMethodMdt
	isSet bool
}

func (v NullablePositioningMethodMdt) Get() *PositioningMethodMdt {
	return v.value
}

func (v *NullablePositioningMethodMdt) Set(val *PositioningMethodMdt) {
	v.value = val
	v.isSet = true
}

func (v NullablePositioningMethodMdt) IsSet() bool {
	return v.isSet
}

func (v *NullablePositioningMethodMdt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePositioningMethodMdt(val *PositioningMethodMdt) *NullablePositioningMethodMdt {
	return &NullablePositioningMethodMdt{value: val, isSet: true}
}

func (v NullablePositioningMethodMdt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePositioningMethodMdt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


