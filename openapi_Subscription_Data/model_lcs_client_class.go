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

// LcsClientClass struct for LcsClientClass
type LcsClientClass struct {
	LcsClientClassAnyOf *LcsClientClassAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *LcsClientClass) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into LcsClientClassAnyOf
	err = json.Unmarshal(data, &dst.LcsClientClassAnyOf);
	if err == nil {
		jsonLcsClientClassAnyOf, _ := json.Marshal(dst.LcsClientClassAnyOf)
		if string(jsonLcsClientClassAnyOf) == "{}" { // empty struct
			dst.LcsClientClassAnyOf = nil
		} else {
			return nil // data stored in dst.LcsClientClassAnyOf, return on the first match
		}
	} else {
		dst.LcsClientClassAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(LcsClientClass)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *LcsClientClass) MarshalJSON() ([]byte, error) {
	if src.LcsClientClassAnyOf != nil {
		return json.Marshal(&src.LcsClientClassAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableLcsClientClass struct {
	value *LcsClientClass
	isSet bool
}

func (v NullableLcsClientClass) Get() *LcsClientClass {
	return v.value
}

func (v *NullableLcsClientClass) Set(val *LcsClientClass) {
	v.value = val
	v.isSet = true
}

func (v NullableLcsClientClass) IsSet() bool {
	return v.isSet
}

func (v *NullableLcsClientClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLcsClientClass(val *LcsClientClass) *NullableLcsClientClass {
	return &NullableLcsClientClass{value: val, isSet: true}
}

func (v NullableLcsClientClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLcsClientClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


