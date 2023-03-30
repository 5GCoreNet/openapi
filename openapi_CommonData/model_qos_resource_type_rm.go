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

// QosResourceTypeRm This enumeration is defined in the same way as the 'QosResourceType' enumeration, but with the OpenAPI 'nullable: true' property. \" 
type QosResourceTypeRm struct {
	NullValue *NullValue
	QosResourceType *QosResourceType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *QosResourceTypeRm) UnmarshalJSON(data []byte) error {
	var err error
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

	// try to unmarshal JSON data into QosResourceType
	err = json.Unmarshal(data, &dst.QosResourceType);
	if err == nil {
		jsonQosResourceType, _ := json.Marshal(dst.QosResourceType)
		if string(jsonQosResourceType) == "{}" { // empty struct
			dst.QosResourceType = nil
		} else {
			return nil // data stored in dst.QosResourceType, return on the first match
		}
	} else {
		dst.QosResourceType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(QosResourceTypeRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *QosResourceTypeRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.QosResourceType != nil {
		return json.Marshal(&src.QosResourceType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableQosResourceTypeRm struct {
	value *QosResourceTypeRm
	isSet bool
}

func (v NullableQosResourceTypeRm) Get() *QosResourceTypeRm {
	return v.value
}

func (v *NullableQosResourceTypeRm) Set(val *QosResourceTypeRm) {
	v.value = val
	v.isSet = true
}

func (v NullableQosResourceTypeRm) IsSet() bool {
	return v.isSet
}

func (v *NullableQosResourceTypeRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQosResourceTypeRm(val *QosResourceTypeRm) *NullableQosResourceTypeRm {
	return &NullableQosResourceTypeRm{value: val, isSet: true}
}

func (v NullableQosResourceTypeRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQosResourceTypeRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

