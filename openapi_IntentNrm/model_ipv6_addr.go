/*
Intent NRM

OAS 3.0.1 definition of the Intent NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_IntentNrm

import (
	"encoding/json"
)

// checks if the Ipv6Addr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Ipv6Addr{}

// Ipv6Addr struct for Ipv6Addr
type Ipv6Addr struct {
}

// NewIpv6Addr instantiates a new Ipv6Addr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIpv6Addr() *Ipv6Addr {
	this := Ipv6Addr{}
	return &this
}

// NewIpv6AddrWithDefaults instantiates a new Ipv6Addr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIpv6AddrWithDefaults() *Ipv6Addr {
	this := Ipv6Addr{}
	return &this
}

func (o Ipv6Addr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Ipv6Addr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableIpv6Addr struct {
	value *Ipv6Addr
	isSet bool
}

func (v NullableIpv6Addr) Get() *Ipv6Addr {
	return v.value
}

func (v *NullableIpv6Addr) Set(val *Ipv6Addr) {
	v.value = val
	v.isSet = true
}

func (v NullableIpv6Addr) IsSet() bool {
	return v.isSet
}

func (v *NullableIpv6Addr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIpv6Addr(val *Ipv6Addr) *NullableIpv6Addr {
	return &NullableIpv6Addr{value: val, isSet: true}
}

func (v NullableIpv6Addr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIpv6Addr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


