/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
	"fmt"
)

// ServiceExperienceType Possible values are: - VOICE: Indicates that the service experience analytics is for voice service. - VIDEO: Indicates that the service experience analytics is for video service. - OTHER: Indicates that the service experience analytics is for other service.
type ServiceExperienceType struct {
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ServiceExperienceType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(ServiceExperienceType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ServiceExperienceType) MarshalJSON() ([]byte, error) {
	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableServiceExperienceType struct {
	value *ServiceExperienceType
	isSet bool
}

func (v NullableServiceExperienceType) Get() *ServiceExperienceType {
	return v.value
}

func (v *NullableServiceExperienceType) Set(val *ServiceExperienceType) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceExperienceType) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceExperienceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceExperienceType(val *ServiceExperienceType) *NullableServiceExperienceType {
	return &NullableServiceExperienceType{value: val, isSet: true}
}

func (v NullableServiceExperienceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceExperienceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
