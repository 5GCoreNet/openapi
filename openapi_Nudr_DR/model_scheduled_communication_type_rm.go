/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
	"fmt"
)

// ScheduledCommunicationTypeRm This enumeration is defined in the same way as the 'ScheduledCommunicationTypen' enumeration, but with the OpenAPI 'nullable: true' property.\"  
type ScheduledCommunicationTypeRm struct {
	NullValue *NullValue
	ScheduledCommunicationType *ScheduledCommunicationType
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *ScheduledCommunicationTypeRm) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal JSON data into ScheduledCommunicationType
	err = json.Unmarshal(data, &dst.ScheduledCommunicationType);
	if err == nil {
		jsonScheduledCommunicationType, _ := json.Marshal(dst.ScheduledCommunicationType)
		if string(jsonScheduledCommunicationType) == "{}" { // empty struct
			dst.ScheduledCommunicationType = nil
		} else {
			return nil // data stored in dst.ScheduledCommunicationType, return on the first match
		}
	} else {
		dst.ScheduledCommunicationType = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(ScheduledCommunicationTypeRm)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *ScheduledCommunicationTypeRm) MarshalJSON() ([]byte, error) {
	if src.NullValue != nil {
		return json.Marshal(&src.NullValue)
	}

	if src.ScheduledCommunicationType != nil {
		return json.Marshal(&src.ScheduledCommunicationType)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableScheduledCommunicationTypeRm struct {
	value *ScheduledCommunicationTypeRm
	isSet bool
}

func (v NullableScheduledCommunicationTypeRm) Get() *ScheduledCommunicationTypeRm {
	return v.value
}

func (v *NullableScheduledCommunicationTypeRm) Set(val *ScheduledCommunicationTypeRm) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduledCommunicationTypeRm) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduledCommunicationTypeRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduledCommunicationTypeRm(val *ScheduledCommunicationTypeRm) *NullableScheduledCommunicationTypeRm {
	return &NullableScheduledCommunicationTypeRm{value: val, isSet: true}
}

func (v NullableScheduledCommunicationTypeRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduledCommunicationTypeRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


