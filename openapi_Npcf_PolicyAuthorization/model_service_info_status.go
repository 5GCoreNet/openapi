/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
	"fmt"
)

// ServiceInfoStatus Represents the preliminary or final service information status.
type ServiceInfoStatus struct {
	ServiceInfoStatusAnyOf *ServiceInfoStatusAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceInfoStatus) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into ServiceInfoStatusAnyOf
	err = json.Unmarshal(data, &dst.ServiceInfoStatusAnyOf);
	if err == nil {
		jsonServiceInfoStatusAnyOf, _ := json.Marshal(dst.ServiceInfoStatusAnyOf)
		if string(jsonServiceInfoStatusAnyOf) == "{}" { // empty struct
			dst.ServiceInfoStatusAnyOf = nil
		} else {
			return nil // data stored in dst.ServiceInfoStatusAnyOf, return on the first match
		}
	} else {
		dst.ServiceInfoStatusAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceInfoStatus)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceInfoStatus) MarshalJSON() ([]byte, error) {
	if src.ServiceInfoStatusAnyOf != nil {
		return json.Marshal(&src.ServiceInfoStatusAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceInfoStatus struct {
	value *ServiceInfoStatus
	isSet bool
}

func (v NullableServiceInfoStatus) Get() *ServiceInfoStatus {
	return v.value
}

func (v *NullableServiceInfoStatus) Set(val *ServiceInfoStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceInfoStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceInfoStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceInfoStatus(val *ServiceInfoStatus) *NullableServiceInfoStatus {
	return &NullableServiceInfoStatus{value: val, isSet: true}
}

func (v NullableServiceInfoStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceInfoStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


