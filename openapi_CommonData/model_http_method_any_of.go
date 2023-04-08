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

// HttpMethodAnyOf the model 'HttpMethodAnyOf'
type HttpMethodAnyOf string

// List of HttpMethod_anyOf
const (
	GET HttpMethodAnyOf = "GET"
	POST HttpMethodAnyOf = "POST"
	PUT HttpMethodAnyOf = "PUT"
	DELETE HttpMethodAnyOf = "DELETE"
	PATCH HttpMethodAnyOf = "PATCH"
	OPTIONS HttpMethodAnyOf = "OPTIONS"
	HEAD HttpMethodAnyOf = "HEAD"
	CONNECT HttpMethodAnyOf = "CONNECT"
	TRACE HttpMethodAnyOf = "TRACE"
)

// All allowed values of HttpMethodAnyOf enum
var AllowedHttpMethodAnyOfEnumValues = []HttpMethodAnyOf{
	"GET",
	"POST",
	"PUT",
	"DELETE",
	"PATCH",
	"OPTIONS",
	"HEAD",
	"CONNECT",
	"TRACE",
}

func (v *HttpMethodAnyOf) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := HttpMethodAnyOf(value)
	for _, existing := range AllowedHttpMethodAnyOfEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid HttpMethodAnyOf", value)
}

// NewHttpMethodAnyOfFromValue returns a pointer to a valid HttpMethodAnyOf
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewHttpMethodAnyOfFromValue(v string) (*HttpMethodAnyOf, error) {
	ev := HttpMethodAnyOf(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for HttpMethodAnyOf: valid values are %v", v, AllowedHttpMethodAnyOfEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v HttpMethodAnyOf) IsValid() bool {
	for _, existing := range AllowedHttpMethodAnyOfEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to HttpMethod_anyOf value
func (v HttpMethodAnyOf) Ptr() *HttpMethodAnyOf {
	return &v
}

type NullableHttpMethodAnyOf struct {
	value *HttpMethodAnyOf
	isSet bool
}

func (v NullableHttpMethodAnyOf) Get() *HttpMethodAnyOf {
	return v.value
}

func (v *NullableHttpMethodAnyOf) Set(val *HttpMethodAnyOf) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpMethodAnyOf) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpMethodAnyOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpMethodAnyOf(val *HttpMethodAnyOf) *NullableHttpMethodAnyOf {
	return &NullableHttpMethodAnyOf{value: val, isSet: true}
}

func (v NullableHttpMethodAnyOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpMethodAnyOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
