/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the AddressList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddressList{}

// AddressList Represents a list of IPv4 and/or IPv6 addresses.
type AddressList struct {
	Ipv4Addrs []string   `json:"ipv4Addrs,omitempty"`
	Ipv6Addrs []Ipv6Addr `json:"ipv6Addrs,omitempty"`
}

// NewAddressList instantiates a new AddressList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddressList() *AddressList {
	this := AddressList{}
	return &this
}

// NewAddressListWithDefaults instantiates a new AddressList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddressListWithDefaults() *AddressList {
	this := AddressList{}
	return &this
}

// GetIpv4Addrs returns the Ipv4Addrs field value if set, zero value otherwise.
func (o *AddressList) GetIpv4Addrs() []string {
	if o == nil || IsNil(o.Ipv4Addrs) {
		var ret []string
		return ret
	}
	return o.Ipv4Addrs
}

// GetIpv4AddrsOk returns a tuple with the Ipv4Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressList) GetIpv4AddrsOk() ([]string, bool) {
	if o == nil || IsNil(o.Ipv4Addrs) {
		return nil, false
	}
	return o.Ipv4Addrs, true
}

// HasIpv4Addrs returns a boolean if a field has been set.
func (o *AddressList) HasIpv4Addrs() bool {
	if o != nil && !IsNil(o.Ipv4Addrs) {
		return true
	}

	return false
}

// SetIpv4Addrs gets a reference to the given []string and assigns it to the Ipv4Addrs field.
func (o *AddressList) SetIpv4Addrs(v []string) {
	o.Ipv4Addrs = v
}

// GetIpv6Addrs returns the Ipv6Addrs field value if set, zero value otherwise.
func (o *AddressList) GetIpv6Addrs() []Ipv6Addr {
	if o == nil || IsNil(o.Ipv6Addrs) {
		var ret []Ipv6Addr
		return ret
	}
	return o.Ipv6Addrs
}

// GetIpv6AddrsOk returns a tuple with the Ipv6Addrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddressList) GetIpv6AddrsOk() ([]Ipv6Addr, bool) {
	if o == nil || IsNil(o.Ipv6Addrs) {
		return nil, false
	}
	return o.Ipv6Addrs, true
}

// HasIpv6Addrs returns a boolean if a field has been set.
func (o *AddressList) HasIpv6Addrs() bool {
	if o != nil && !IsNil(o.Ipv6Addrs) {
		return true
	}

	return false
}

// SetIpv6Addrs gets a reference to the given []Ipv6Addr and assigns it to the Ipv6Addrs field.
func (o *AddressList) SetIpv6Addrs(v []Ipv6Addr) {
	o.Ipv6Addrs = v
}

func (o AddressList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddressList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4Addrs) {
		toSerialize["ipv4Addrs"] = o.Ipv4Addrs
	}
	if !IsNil(o.Ipv6Addrs) {
		toSerialize["ipv6Addrs"] = o.Ipv6Addrs
	}
	return toSerialize, nil
}

type NullableAddressList struct {
	value *AddressList
	isSet bool
}

func (v NullableAddressList) Get() *AddressList {
	return v.value
}

func (v *NullableAddressList) Set(val *AddressList) {
	v.value = val
	v.isSet = true
}

func (v NullableAddressList) IsSet() bool {
	return v.isSet
}

func (v *NullableAddressList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAddressList(val *AddressList) *NullableAddressList {
	return &NullableAddressList{value: val, isSet: true}
}

func (v NullableAddressList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAddressList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
