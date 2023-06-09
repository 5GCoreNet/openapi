/*
NRF OAuth2

NRF OAuth2 Authorization.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_AccessToken

import (
	"encoding/json"
	"fmt"
)

// AccessTokenClaimsAud struct for AccessTokenClaimsAud
type AccessTokenClaimsAud struct {
	NFType        *NFType
	ArrayOfString *[]string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *AccessTokenClaimsAud) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into NFType
	err = json.Unmarshal(data, &dst.NFType)
	if err == nil {
		jsonNFType, _ := json.Marshal(dst.NFType)
		if string(jsonNFType) == "{}" { // empty struct
			dst.NFType = nil
		} else {
			return nil // data stored in dst.NFType, return on the first match
		}
	} else {
		dst.NFType = nil
	}

	// try to unmarshal JSON data into []string
	err = json.Unmarshal(data, &dst.ArrayOfString)
	if err == nil {
		jsonArrayOfString, _ := json.Marshal(dst.ArrayOfString)
		if string(jsonArrayOfString) == "{}" { // empty struct
			dst.ArrayOfString = nil
		} else {
			return nil // data stored in dst.ArrayOfString, return on the first match
		}
	} else {
		dst.ArrayOfString = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(AccessTokenClaimsAud)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *AccessTokenClaimsAud) MarshalJSON() ([]byte, error) {
	if src.NFType != nil {
		return json.Marshal(&src.NFType)
	}

	if src.ArrayOfString != nil {
		return json.Marshal(&src.ArrayOfString)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableAccessTokenClaimsAud struct {
	value *AccessTokenClaimsAud
	isSet bool
}

func (v NullableAccessTokenClaimsAud) Get() *AccessTokenClaimsAud {
	return v.value
}

func (v *NullableAccessTokenClaimsAud) Set(val *AccessTokenClaimsAud) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessTokenClaimsAud) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessTokenClaimsAud) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessTokenClaimsAud(val *AccessTokenClaimsAud) *NullableAccessTokenClaimsAud {
	return &NullableAccessTokenClaimsAud{value: val, isSet: true}
}

func (v NullableAccessTokenClaimsAud) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessTokenClaimsAud) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
