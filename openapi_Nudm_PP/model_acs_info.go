/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_PP

import (
	"encoding/json"
)

// checks if the AcsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AcsInfo{}

// AcsInfo The ACS information for the 5G-RG is defined in BBF TR-069 [42] or in BBF TR-369
type AcsInfo struct {
	// String providing an URI formatted according to RFC 3986.
	AcsUrl *string `json:"acsUrl,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	AcsIpv4Addr *string `json:"acsIpv4Addr,omitempty"`
	AcsIpv6Addr *Ipv6Addr `json:"acsIpv6Addr,omitempty"`
}

// NewAcsInfo instantiates a new AcsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAcsInfo() *AcsInfo {
	this := AcsInfo{}
	return &this
}

// NewAcsInfoWithDefaults instantiates a new AcsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAcsInfoWithDefaults() *AcsInfo {
	this := AcsInfo{}
	return &this
}

// GetAcsUrl returns the AcsUrl field value if set, zero value otherwise.
func (o *AcsInfo) GetAcsUrl() string {
	if o == nil || isNil(o.AcsUrl) {
		var ret string
		return ret
	}
	return *o.AcsUrl
}

// GetAcsUrlOk returns a tuple with the AcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsInfo) GetAcsUrlOk() (*string, bool) {
	if o == nil || isNil(o.AcsUrl) {
		return nil, false
	}
	return o.AcsUrl, true
}

// HasAcsUrl returns a boolean if a field has been set.
func (o *AcsInfo) HasAcsUrl() bool {
	if o != nil && !isNil(o.AcsUrl) {
		return true
	}

	return false
}

// SetAcsUrl gets a reference to the given string and assigns it to the AcsUrl field.
func (o *AcsInfo) SetAcsUrl(v string) {
	o.AcsUrl = &v
}

// GetAcsIpv4Addr returns the AcsIpv4Addr field value if set, zero value otherwise.
func (o *AcsInfo) GetAcsIpv4Addr() string {
	if o == nil || isNil(o.AcsIpv4Addr) {
		var ret string
		return ret
	}
	return *o.AcsIpv4Addr
}

// GetAcsIpv4AddrOk returns a tuple with the AcsIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsInfo) GetAcsIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.AcsIpv4Addr) {
		return nil, false
	}
	return o.AcsIpv4Addr, true
}

// HasAcsIpv4Addr returns a boolean if a field has been set.
func (o *AcsInfo) HasAcsIpv4Addr() bool {
	if o != nil && !isNil(o.AcsIpv4Addr) {
		return true
	}

	return false
}

// SetAcsIpv4Addr gets a reference to the given string and assigns it to the AcsIpv4Addr field.
func (o *AcsInfo) SetAcsIpv4Addr(v string) {
	o.AcsIpv4Addr = &v
}

// GetAcsIpv6Addr returns the AcsIpv6Addr field value if set, zero value otherwise.
func (o *AcsInfo) GetAcsIpv6Addr() Ipv6Addr {
	if o == nil || isNil(o.AcsIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.AcsIpv6Addr
}

// GetAcsIpv6AddrOk returns a tuple with the AcsIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AcsInfo) GetAcsIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.AcsIpv6Addr) {
		return nil, false
	}
	return o.AcsIpv6Addr, true
}

// HasAcsIpv6Addr returns a boolean if a field has been set.
func (o *AcsInfo) HasAcsIpv6Addr() bool {
	if o != nil && !isNil(o.AcsIpv6Addr) {
		return true
	}

	return false
}

// SetAcsIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the AcsIpv6Addr field.
func (o *AcsInfo) SetAcsIpv6Addr(v Ipv6Addr) {
	o.AcsIpv6Addr = &v
}

func (o AcsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AcsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AcsUrl) {
		toSerialize["acsUrl"] = o.AcsUrl
	}
	if !isNil(o.AcsIpv4Addr) {
		toSerialize["acsIpv4Addr"] = o.AcsIpv4Addr
	}
	if !isNil(o.AcsIpv6Addr) {
		toSerialize["acsIpv6Addr"] = o.AcsIpv6Addr
	}
	return toSerialize, nil
}

type NullableAcsInfo struct {
	value *AcsInfo
	isSet bool
}

func (v NullableAcsInfo) Get() *AcsInfo {
	return v.value
}

func (v *NullableAcsInfo) Set(val *AcsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAcsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAcsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAcsInfo(val *AcsInfo) *NullableAcsInfo {
	return &NullableAcsInfo{value: val, isSet: true}
}

func (v NullableAcsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAcsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


