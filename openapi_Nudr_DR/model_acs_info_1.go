/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the AcsInfo1 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcsInfo1{}

// AcsInfo1 The ACS information for the 5G-RG is defined in BBF TR-069 [42] or in BBF TR-369
type AcsInfo1 struct {
	// String providing an URI formatted according to RFC 3986.
	AcsUrl *string `json:"acsUrl,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	AcsIpv4Addr *string   `json:"acsIpv4Addr,omitempty"`
	AcsIpv6Addr *Ipv6Addr `json:"acsIpv6Addr,omitempty"`
}

// NewAcsInfo1 instantiates a new AcsInfo1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcsInfo1() *AcsInfo1 {
	this := AcsInfo1{}
	return &this
}

// NewAcsInfo1WithDefaults instantiates a new AcsInfo1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcsInfo1WithDefaults() *AcsInfo1 {
	this := AcsInfo1{}
	return &this
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise.
func (o *AcsInfo1) GetAcsUrl() string {
	if o == nil || IsNil(o.AcsUrl) {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsInfo1) GetAcsUrlOk() (*string, bool) {
	if o == nil || IsNil(o.AcsUrl) {
		return nil, false
	}
	return o.AcsUrl, true
}

// HasAcsUrl returns a boolean if a field has been set.
func (o *AcsInfo1) HasAcsUrl() bool {
	if o != nil && !IsNil(o.AcsUrl) {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *AcsInfo1) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetAcsIpv4Addr returns the AcsIpv4Addr field value if set, zero value otherwise.
func (o *AcsInfo1) GetAcsIpv4Addr() string {
	if o == nil || IsNil(o.AcsIpv4Addr) {
		var ret string
		return ret
	}
	return *o.AcsIpv4Addr
}

// GetAcsIpv4AddrOk returns a tuple with the AcsIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsInfo1) GetAcsIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.AcsIpv4Addr) {
		return nil, false
	}
	return o.AcsIpv4Addr, true
}

// HasAcsIpv4Addr returns a boolean if a field has been set.
func (o *AcsInfo1) HasAcsIpv4Addr() bool {
	if o != nil && !IsNil(o.AcsIpv4Addr) {
		return true
	}

	return false
}

// SetAcsIpv4Addr gets a reference to the given string and assigns it to the AcsIpv4Addr field.
func (o *AcsInfo1) SetAcsIpv4Addr(v string) {
	o.AcsIpv4Addr = &v
}

// GetAcsIpv6Addr returns the AcsIpv6Addr field value if set, zero value otherwise.
func (o *AcsInfo1) GetAcsIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.AcsIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.AcsIpv6Addr
}

// GetAcsIpv6AddrOk returns a tuple with the AcsIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsInfo1) GetAcsIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.AcsIpv6Addr) {
		return nil, false
	}
	return o.AcsIpv6Addr, true
}

// HasAcsIpv6Addr returns a boolean if a field has been set.
func (o *AcsInfo1) HasAcsIpv6Addr() bool {
	if o != nil && !IsNil(o.AcsIpv6Addr) {
		return true
	}

	return false
}

// SetAcsIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the AcsIpv6Addr field.
func (o *AcsInfo1) SetAcsIpv6Addr(v Ipv6Addr) {
	o.AcsIpv6Addr = &v
}

func (o AcsInfo1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcsInfo1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AcsUrl) {
		toSerialize["acsUrl"] = o.AcsUrl
	}
	if !IsNil(o.AcsIpv4Addr) {
		toSerialize["acsIpv4Addr"] = o.AcsIpv4Addr
	}
	if !IsNil(o.AcsIpv6Addr) {
		toSerialize["acsIpv6Addr"] = o.AcsIpv6Addr
	}
	return toSerialize, nil
}

type NullableAcsInfo1 struct {
	value *AcsInfo1
	isSet bool
}

func (v NullableAcsInfo1) Get() *AcsInfo1 {
	return v.value
}

func (v *NullableAcsInfo1) Set(val *AcsInfo1) {
	v.value = val
	v.isSet = true
}

func (v NullableAcsInfo1) IsSet() bool {
	return v.isSet
}

func (v *NullableAcsInfo1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcsInfo1(val *AcsInfo1) *NullableAcsInfo1 {
	return &NullableAcsInfo1{value: val, isSet: true}
}

func (v NullableAcsInfo1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcsInfo1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
