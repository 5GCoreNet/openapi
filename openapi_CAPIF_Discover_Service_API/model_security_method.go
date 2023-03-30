/*
CAPIF_Discover_Service_API

API for discovering service APIs.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_Discover_Service_API

import (
	"encoding/json"
	"fmt"
)

// SecurityMethod Possible values are: - PSK: Security method 1 (Using TLS-PSK) as described in 3GPP TS 33.122 - PKI: Security method 2 (Using PKI) as described in 3GPP TS 33.122 - OAUTH: Security method 3 (TLS with OAuth token) as described in 3GPP TS 33.122 
type SecurityMethod struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *SecurityMethod) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string);
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(SecurityMethod)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *SecurityMethod) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableSecurityMethod struct {
	value *SecurityMethod
	isSet bool
}

func (v NullableSecurityMethod) Get() *SecurityMethod {
	return v.value
}

func (v *NullableSecurityMethod) Set(val *SecurityMethod) {
	v.value = val
	v.isSet = true
}

func (v NullableSecurityMethod) IsSet() bool {
	return v.isSet
}

func (v *NullableSecurityMethod) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSecurityMethod(val *SecurityMethod) *NullableSecurityMethod {
	return &NullableSecurityMethod{value: val, isSet: true}
}

func (v NullableSecurityMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSecurityMethod) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


