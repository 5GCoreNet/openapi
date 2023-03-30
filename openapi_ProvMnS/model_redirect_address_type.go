/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
	"fmt"
)

// RedirectAddressType Possible values are - IPV4_ADDR: Indicates that the address type is in the form of \"dotted-decimal\" IPv4  address. - IPV6_ADDR: Indicates that the address type is in the form of IPv6 address. - URL: Indicates that the address type is in the form of Uniform Resource Locator. - SIP_URI: Indicates that the address type is in the form of SIP Uniform Resource  Identifier. 
type RedirectAddressType struct {
	string *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RedirectAddressType) UnmarshalJSON(data []byte) error {
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

	return fmt.Errorf("data failed to match schemas in anyOf(RedirectAddressType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RedirectAddressType) MarshalJSON() ([]byte, error) {
	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableRedirectAddressType struct {
	value *RedirectAddressType
	isSet bool
}

func (v NullableRedirectAddressType) Get() *RedirectAddressType {
	return v.value
}

func (v *NullableRedirectAddressType) Set(val *RedirectAddressType) {
	v.value = val
	v.isSet = true
}

func (v NullableRedirectAddressType) IsSet() bool {
	return v.isSet
}

func (v *NullableRedirectAddressType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRedirectAddressType(val *RedirectAddressType) *NullableRedirectAddressType {
	return &NullableRedirectAddressType{value: val, isSet: true}
}

func (v NullableRedirectAddressType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRedirectAddressType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


