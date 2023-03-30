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

// TermCause Possible values are:     - USER_CONSENT_REVOKED: The user consent has been revoked.   - NWDAF_OVERLOAD: The NWDAF is overloaded.   - UE_LEFT_AREA: The UE has moved out of the NWDAF serving area. 
type TermCause struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *TermCause) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(TermCause)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *TermCause) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableTermCause struct {
	value *TermCause
	isSet bool
}

func (v NullableTermCause) Get() *TermCause {
	return v.value
}

func (v *NullableTermCause) Set(val *TermCause) {
	v.value = val
	v.isSet = true
}

func (v NullableTermCause) IsSet() bool {
	return v.isSet
}

func (v *NullableTermCause) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTermCause(val *TermCause) *NullableTermCause {
	return &NullableTermCause{value: val, isSet: true}
}

func (v NullableTermCause) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTermCause) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


