/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
	"fmt"
)

// PositioningMethodMdt The enumeration LoggingDurationMdt defines Logging Duration for MDT in the trace. See 3GPP TS 32.422 for further description of the values. It shall comply with the provisions defined in table 5.6.3.13-1. 
type PositioningMethodMdt struct {
	PositioningMethodMdtAnyOf *PositioningMethodMdtAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *PositioningMethodMdt) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into PositioningMethodMdtAnyOf
	err = json.Unmarshal(data, &dst.PositioningMethodMdtAnyOf);
	if err == nil {
		jsonPositioningMethodMdtAnyOf, _ := json.Marshal(dst.PositioningMethodMdtAnyOf)
		if string(jsonPositioningMethodMdtAnyOf) == "{}" { // empty struct
			dst.PositioningMethodMdtAnyOf = nil
		} else {
			return nil // data stored in dst.PositioningMethodMdtAnyOf, return on the first match
		}
	} else {
		dst.PositioningMethodMdtAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(PositioningMethodMdt)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *PositioningMethodMdt) MarshalJSON() ([]byte, error) {
	if src.PositioningMethodMdtAnyOf != nil {
		return json.Marshal(&src.PositioningMethodMdtAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
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


