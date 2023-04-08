/*
Nhss_imsUECM

Nhss UE Context Management Service for IMS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_imsUECM

import (
	"encoding/json"
	"fmt"
)

// AuthorizationType Represents the type of authorization requested by the UE
type AuthorizationType struct {
	AuthorizationTypeAnyOf *AuthorizationTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AuthorizationType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into AuthorizationTypeAnyOf
	err = json.Unmarshal(data, &dst.AuthorizationTypeAnyOf);
	if err == nil {
		jsonAuthorizationTypeAnyOf, _ := json.Marshal(dst.AuthorizationTypeAnyOf)
		if string(jsonAuthorizationTypeAnyOf) == "{}" { // empty struct
			dst.AuthorizationTypeAnyOf = nil
		} else {
			return nil // data stored in dst.AuthorizationTypeAnyOf, return on the first match
		}
	} else {
		dst.AuthorizationTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(AuthorizationType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AuthorizationType) MarshalJSON() ([]byte, error) {
	if src.AuthorizationTypeAnyOf != nil {
		return json.Marshal(&src.AuthorizationTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAuthorizationType struct {
	value *AuthorizationType
	isSet bool
}

func (v NullableAuthorizationType) Get() *AuthorizationType {
	return v.value
}

func (v *NullableAuthorizationType) Set(val *AuthorizationType) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthorizationType) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthorizationType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthorizationType(val *AuthorizationType) *NullableAuthorizationType {
	return &NullableAuthorizationType{value: val, isSet: true}
}

func (v NullableAuthorizationType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthorizationType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


