/*
3gpp-traffic-influence

API for AF traffic influence   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_TrafficInfluence

import (
	"encoding/json"
	"fmt"
)

// TrafficInfluSub Represents a traffic influence subscription.
type TrafficInfluSub struct {
	OneOfAnyTypeAnyTypeAnyType                      *OneOfAnyTypeAnyTypeAnyType
	OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType *OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TrafficInfluSub) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into OneOfAnyTypeAnyTypeAnyType
	err = json.Unmarshal(data, &dst.OneOfAnyTypeAnyTypeAnyType)
	if err == nil {
		jsonOneOfAnyTypeAnyTypeAnyType, _ := json.Marshal(dst.OneOfAnyTypeAnyTypeAnyType)
		if string(jsonOneOfAnyTypeAnyTypeAnyType) == "{}" { // empty struct
			dst.OneOfAnyTypeAnyTypeAnyType = nil
		} else {
			return nil // data stored in dst.OneOfAnyTypeAnyTypeAnyType, return on the first match
		}
	} else {
		dst.OneOfAnyTypeAnyTypeAnyType = nil
	}

	// try to unmarshal JSON data into OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType
	err = json.Unmarshal(data, &dst.OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType)
	if err == nil {
		jsonOneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType, _ := json.Marshal(dst.OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType)
		if string(jsonOneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType) == "{}" { // empty struct
			dst.OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType = nil
		} else {
			return nil // data stored in dst.OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType, return on the first match
		}
	} else {
		dst.OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(TrafficInfluSub)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TrafficInfluSub) MarshalJSON() ([]byte, error) {
	if src.OneOfAnyTypeAnyTypeAnyType != nil {
		return json.Marshal(&src.OneOfAnyTypeAnyTypeAnyType)
	}

	if src.OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType != nil {
		return json.Marshal(&src.OneOfAnyTypeAnyTypeAnyTypeAnyTypeAnyTypeAnyType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTrafficInfluSub struct {
	value *TrafficInfluSub
	isSet bool
}

func (v NullableTrafficInfluSub) Get() *TrafficInfluSub {
	return v.value
}

func (v *NullableTrafficInfluSub) Set(val *TrafficInfluSub) {
	v.value = val
	v.isSet = true
}

func (v NullableTrafficInfluSub) IsSet() bool {
	return v.isSet
}

func (v *NullableTrafficInfluSub) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrafficInfluSub(val *TrafficInfluSub) *NullableTrafficInfluSub {
	return &NullableTrafficInfluSub{value: val, isSet: true}
}

func (v NullableTrafficInfluSub) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrafficInfluSub) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
