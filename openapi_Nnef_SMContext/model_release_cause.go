/*
Nnef_SMContext

Nnef SMContext Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_SMContext

import (
	"encoding/json"
	"fmt"
)

// ReleaseCause The cause to release the SM Context. Possible values are - PDU_SESSION_RELEASED: Indicates that the Individual SM Context for NIDD is released. 
type ReleaseCause struct {
	ReleaseCauseAnyOf *ReleaseCauseAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ReleaseCause) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ReleaseCauseAnyOf
	err = json.Unmarshal(data, &dst.ReleaseCauseAnyOf);
	if err == nil {
		jsonReleaseCauseAnyOf, _ := json.Marshal(dst.ReleaseCauseAnyOf)
		if string(jsonReleaseCauseAnyOf) == "{}" { // empty struct
			dst.ReleaseCauseAnyOf = nil
		} else {
			return nil // data stored in dst.ReleaseCauseAnyOf, return on the first match
		}
	} else {
		dst.ReleaseCauseAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ReleaseCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ReleaseCause) MarshalJSON() ([]byte, error) {
	if src.ReleaseCauseAnyOf != nil {
		return json.Marshal(&src.ReleaseCauseAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableReleaseCause struct {
	value *ReleaseCause
	isSet bool
}

func (v NullableReleaseCause) Get() *ReleaseCause {
	return v.value
}

func (v *NullableReleaseCause) Set(val *ReleaseCause) {
	v.value = val
	v.isSet = true
}

func (v NullableReleaseCause) IsSet() bool {
	return v.isSet
}

func (v *NullableReleaseCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReleaseCause(val *ReleaseCause) *NullableReleaseCause {
	return &NullableReleaseCause{value: val, isSet: true}
}

func (v NullableReleaseCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReleaseCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


