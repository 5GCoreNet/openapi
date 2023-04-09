/*
Namf_MBSCommunication

AMF Communication Service for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MBSCommunication

import (
	"encoding/json"
)

// checks if the Ipv6Prefix type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv6Prefix{}

// Ipv6Prefix String identifying an IPv6 address prefix formatted according to clause 4 of RFC 5952. IPv6Prefix data type may contain an individual /128 IPv6 address.
type Ipv6Prefix struct {
}

// NewIpv6Prefix instantiates a new Ipv6Prefix object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6Prefix() *Ipv6Prefix {
	this := Ipv6Prefix{}
	return &this
}

// NewIpv6PrefixWithDefaults instantiates a new Ipv6Prefix object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6PrefixWithDefaults() *Ipv6Prefix {
	this := Ipv6Prefix{}
	return &this
}

func (o Ipv6Prefix) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv6Prefix) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableIpv6Prefix struct {
	value *Ipv6Prefix
	isSet bool
}

func (v NullableIpv6Prefix) Get() *Ipv6Prefix {
	return v.value
}

func (v *NullableIpv6Prefix) Set(val *Ipv6Prefix) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6Prefix) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6Prefix) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6Prefix(val *Ipv6Prefix) *NullableIpv6Prefix {
	return &NullableIpv6Prefix{value: val, isSet: true}
}

func (v NullableIpv6Prefix) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6Prefix) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
