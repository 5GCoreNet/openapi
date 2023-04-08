/*
Nhss_imsSDM

Nhss Subscriber Data Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsSDM

import (
	"encoding/json"
	"fmt"
)

// RegistrationType Represents the type of registration associated to the REGISTER request
type RegistrationType struct {
	RegistrationTypeAnyOf *RegistrationTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RegistrationType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RegistrationTypeAnyOf
	err = json.Unmarshal(data, &dst.RegistrationTypeAnyOf);
	if err == nil {
		jsonRegistrationTypeAnyOf, _ := json.Marshal(dst.RegistrationTypeAnyOf)
		if string(jsonRegistrationTypeAnyOf) == "{}" { // empty struct
			dst.RegistrationTypeAnyOf = nil
		} else {
			return nil // data stored in dst.RegistrationTypeAnyOf, return on the first match
		}
	} else {
		dst.RegistrationTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RegistrationType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RegistrationType) MarshalJSON() ([]byte, error) {
	if src.RegistrationTypeAnyOf != nil {
		return json.Marshal(&src.RegistrationTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRegistrationType struct {
	value *RegistrationType
	isSet bool
}

func (v NullableRegistrationType) Get() *RegistrationType {
	return v.value
}

func (v *NullableRegistrationType) Set(val *RegistrationType) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationType) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationType(val *RegistrationType) *NullableRegistrationType {
	return &NullableRegistrationType{value: val, isSet: true}
}

func (v NullableRegistrationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


