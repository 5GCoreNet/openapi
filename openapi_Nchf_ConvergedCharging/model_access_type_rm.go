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

// AccessTypeRm Indicates wether the access is via 3GPP or via non-3GPP but with the OpenAPI  'nullable: true' property.\" 
type AccessTypeRm struct {
	AccessType *AccessType
	NullValue *NullValue
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessTypeRm) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AccessType
	err = json.Unmarshal(data, &dst.AccessType);
	if err == nil {
		jsonAccessType, _ := json.Marshal(dst.AccessType)
		if string(jsonAccessType) == "{}" { // empty struct
			dst.AccessType = nil
		} else {
			return nil // data stored in dst.AccessType, return on the first match
		}
	} else {
		dst.AccessType = nil
	}

	// try to unmarshal JSON data into NullValue
	err = json.Unmarshal(data, &dst.NullValue);
	if err == nil {
		jsonNullValue, _ := json.Marshal(dst.NullValue)
		if string(jsonNullValue) == "{}" { // empty struct
			dst.NullValue = nil
		} else {
			return nil // data stored in dst.NullValue, return on the first match
		}
	} else {
		dst.NullValue = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AccessTypeRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessTypeRm) MarshalJSON() ([]byte, error) {
	if src.AccessType != nil {
		return json.Marshal(&src.AccessType)
	}

	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessTypeRm struct {
	value *AccessTypeRm
	isSet bool
}

func (v NullableAccessTypeRm) Get() *AccessTypeRm {
	return v.value
}

func (v *NullableAccessTypeRm) Set(val *AccessTypeRm) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTypeRm) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTypeRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTypeRm(val *AccessTypeRm) *NullableAccessTypeRm {
	return &NullableAccessTypeRm{value: val, isSet: true}
}

func (v NullableAccessTypeRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTypeRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


