/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
	"fmt"
)

// ExceptionId Possible values are: - UNEXPECTED_UE_LOCATION: Unexpected UE location - UNEXPECTED_LONG_LIVE_FLOW: Unexpected long-live rate flows - UNEXPECTED_LARGE_RATE_FLOW: Unexpected large rate flows - UNEXPECTED_WAKEUP: Unexpected wakeup - SUSPICION_OF_DDOS_ATTACK: Suspicion of DDoS attack - WRONG_DESTINATION_ADDRESS: Wrong destination address - TOO_FREQUENT_SERVICE_ACCESS: Too frequent Service Access - UNEXPECTED_RADIO_LINK_FAILURES: Unexpected radio link failures - PING_PONG_ACROSS_CELLS: Ping-ponging across neighbouring cells 
type ExceptionId struct {
	ExceptionIdAnyOf *ExceptionIdAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ExceptionId) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ExceptionIdAnyOf
	err = json.Unmarshal(data, &dst.ExceptionIdAnyOf);
	if err == nil {
		jsonExceptionIdAnyOf, _ := json.Marshal(dst.ExceptionIdAnyOf)
		if string(jsonExceptionIdAnyOf) == "{}" { // empty struct
			dst.ExceptionIdAnyOf = nil
		} else {
			return nil // data stored in dst.ExceptionIdAnyOf, return on the first match
		}
	} else {
		dst.ExceptionIdAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ExceptionId)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ExceptionId) MarshalJSON() ([]byte, error) {
	if src.ExceptionIdAnyOf != nil {
		return json.Marshal(&src.ExceptionIdAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableExceptionId struct {
	value *ExceptionId
	isSet bool
}

func (v NullableExceptionId) Get() *ExceptionId {
	return v.value
}

func (v *NullableExceptionId) Set(val *ExceptionId) {
	v.value = val
	v.isSet = true
}

func (v NullableExceptionId) IsSet() bool {
	return v.isSet
}

func (v *NullableExceptionId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExceptionId(val *ExceptionId) *NullableExceptionId {
	return &NullableExceptionId{value: val, isSet: true}
}

func (v NullableExceptionId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExceptionId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


