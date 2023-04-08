/*
EES ACR Status Update Service

EES ACR Status Update Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Eees_ACRStatusUpdate

import (
	"encoding/json"
	"fmt"
)

// E3SubscsStatus Possible values are: - SUCCESSFUL: Indicates that the initialization of EDGE-3 subscriptions was successful. - FAILED: Indicates that the initialization of EDGE-3 subscriptions failed. 
type E3SubscsStatus struct {
	ACTResultAnyOf *ACTResultAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *E3SubscsStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ACTResultAnyOf
	err = json.Unmarshal(data, &dst.ACTResultAnyOf);
	if err == nil {
		jsonACTResultAnyOf, _ := json.Marshal(dst.ACTResultAnyOf)
		if string(jsonACTResultAnyOf) == "{}" { // empty struct
			dst.ACTResultAnyOf = nil
		} else {
			return nil // data stored in dst.ACTResultAnyOf, return on the first match
		}
	} else {
		dst.ACTResultAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(E3SubscsStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *E3SubscsStatus) MarshalJSON() ([]byte, error) {
	if src.ACTResultAnyOf != nil {
		return json.Marshal(&src.ACTResultAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableE3SubscsStatus struct {
	value *E3SubscsStatus
	isSet bool
}

func (v NullableE3SubscsStatus) Get() *E3SubscsStatus {
	return v.value
}

func (v *NullableE3SubscsStatus) Set(val *E3SubscsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableE3SubscsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableE3SubscsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableE3SubscsStatus(val *E3SubscsStatus) *NullableE3SubscsStatus {
	return &NullableE3SubscsStatus{value: val, isSet: true}
}

func (v NullableE3SubscsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableE3SubscsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


