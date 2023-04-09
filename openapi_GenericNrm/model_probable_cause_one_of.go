/*
Generic NRM

OAS 3.0.1 definition of the Generic NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_GenericNrm

import (
	"encoding/json"
	"fmt"
)

// ProbableCauseOneOf struct for ProbableCauseOneOf
type ProbableCauseOneOf struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ProbableCauseOneOf) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.String)
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

	return fmt.Errorf("data failed to match schemas in anyOf(ProbableCauseOneOf)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ProbableCauseOneOf) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableProbableCauseOneOf struct {
	value *ProbableCauseOneOf
	isSet bool
}

func (v NullableProbableCauseOneOf) Get() *ProbableCauseOneOf {
	return v.value
}

func (v *NullableProbableCauseOneOf) Set(val *ProbableCauseOneOf) {
	v.value = val
	v.isSet = true
}

func (v NullableProbableCauseOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableProbableCauseOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProbableCauseOneOf(val *ProbableCauseOneOf) *NullableProbableCauseOneOf {
	return &NullableProbableCauseOneOf{value: val, isSet: true}
}

func (v NullableProbableCauseOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProbableCauseOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
