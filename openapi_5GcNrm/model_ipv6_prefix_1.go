/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the Ipv6Prefix1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv6Prefix1{}

// Ipv6Prefix1 String identifying an IPv6 address prefix formatted according to clause 4 of RFC 5952. IPv6Prefix data type may contain an individual /128 IPv6 address. 
type Ipv6Prefix1 struct {
}

// NewIpv6Prefix1 instantiates a new Ipv6Prefix1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6Prefix1() *Ipv6Prefix1 {
	this := Ipv6Prefix1{}
	return &this
}

// NewIpv6Prefix1WithDefaults instantiates a new Ipv6Prefix1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6Prefix1WithDefaults() *Ipv6Prefix1 {
	this := Ipv6Prefix1{}
	return &this
}

func (o Ipv6Prefix1) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv6Prefix1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableIpv6Prefix1 struct {
	value *Ipv6Prefix1
	isSet bool
}

func (v NullableIpv6Prefix1) Get() *Ipv6Prefix1 {
	return v.value
}

func (v *NullableIpv6Prefix1) Set(val *Ipv6Prefix1) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6Prefix1) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6Prefix1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6Prefix1(val *Ipv6Prefix1) *NullableIpv6Prefix1 {
	return &NullableIpv6Prefix1{value: val, isSet: true}
}

func (v NullableIpv6Prefix1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6Prefix1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


