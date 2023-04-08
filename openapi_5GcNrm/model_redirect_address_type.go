/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
	"fmt"
)

// RedirectAddressType Possible values are - IPV4_ADDR: Indicates that the address type is in the form of \"dotted-decimal\" IPv4  address. - IPV6_ADDR: Indicates that the address type is in the form of IPv6 address. - URL: Indicates that the address type is in the form of Uniform Resource Locator. - SIP_URI: Indicates that the address type is in the form of SIP Uniform Resource  Identifier. 
type RedirectAddressType struct {
	RedirectAddressTypeAnyOf *RedirectAddressTypeAnyOf
	String *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *RedirectAddressType) UnmarshalJSON(data []byte) error {
	var err error
	// try to unmarshal JSON data into RedirectAddressTypeAnyOf
	err = json.Unmarshal(data, &dst.RedirectAddressTypeAnyOf);
	if err == nil {
		jsonRedirectAddressTypeAnyOf, _ := json.Marshal(dst.RedirectAddressTypeAnyOf)
		if string(jsonRedirectAddressTypeAnyOf) == "{}" { // empty struct
			dst.RedirectAddressTypeAnyOf = nil
		} else {
			return nil // data stored in dst.RedirectAddressTypeAnyOf, return on the first match
		}
	} else {
		dst.RedirectAddressTypeAnyOf = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(RedirectAddressType)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *RedirectAddressType) MarshalJSON() ([]byte, error) {
	if src.RedirectAddressTypeAnyOf != nil {
		return json.Marshal(&src.RedirectAddressTypeAnyOf)
	}

	if src.String != nil {
		return json.Marshal(&src.String)
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


