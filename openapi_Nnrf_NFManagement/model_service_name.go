/*
NRF NFManagement Service

NRF NFManagement Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFManagement

import (
	"encoding/json"
	"fmt"
)

// ServiceName Service names known to NRF
type ServiceName struct {
	ServiceNameAnyOf *ServiceNameAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceName) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ServiceNameAnyOf
	err = json.Unmarshal(data, &dst.ServiceNameAnyOf);
	if err == nil {
		jsonServiceNameAnyOf, _ := json.Marshal(dst.ServiceNameAnyOf)
		if string(jsonServiceNameAnyOf) == "{}" { // empty struct
			dst.ServiceNameAnyOf = nil
		} else {
			return nil // data stored in dst.ServiceNameAnyOf, return on the first match
		}
	} else {
		dst.ServiceNameAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceName)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceName) MarshalJSON() ([]byte, error) {
	if src.ServiceNameAnyOf != nil {
		return json.Marshal(&src.ServiceNameAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceName struct {
	value *ServiceName
	isSet bool
}

func (v NullableServiceName) Get() *ServiceName {
	return v.value
}

func (v *NullableServiceName) Set(val *ServiceName) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceName) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceName(val *ServiceName) *NullableServiceName {
	return &NullableServiceName{value: val, isSet: true}
}

func (v NullableServiceName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


