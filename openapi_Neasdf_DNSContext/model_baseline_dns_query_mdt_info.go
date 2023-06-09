/*
Neasdf_DNSContext

EASDF DNS Context Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Neasdf_DNSContext

import (
	"encoding/json"
)

// checks if the BaselineDnsQueryMdtInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BaselineDnsQueryMdtInfo{}

// BaselineDnsQueryMdtInfo Baseline DNS Query MDT Information
type BaselineDnsQueryMdtInfo struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166.
	SourceIpv4Addr   *string            `json:"sourceIpv4Addr,omitempty"`
	SourceIpv6Prefix *Ipv6Prefix        `json:"sourceIpv6Prefix,omitempty"`
	BaseDnsMdtList   []BaselineDnsMdtId `json:"baseDnsMdtList"`
}

// NewBaselineDnsQueryMdtInfo instantiates a new BaselineDnsQueryMdtInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBaselineDnsQueryMdtInfo(baseDnsMdtList []BaselineDnsMdtId) *BaselineDnsQueryMdtInfo {
	this := BaselineDnsQueryMdtInfo{}
	this.BaseDnsMdtList = baseDnsMdtList
	return &this
}

// NewBaselineDnsQueryMdtInfoWithDefaults instantiates a new BaselineDnsQueryMdtInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBaselineDnsQueryMdtInfoWithDefaults() *BaselineDnsQueryMdtInfo {
	this := BaselineDnsQueryMdtInfo{}
	return &this
}

// GetSourceIpv4Addr returns the SourceIpv4Addr field value if set, zero value otherwise.
func (o *BaselineDnsQueryMdtInfo) GetSourceIpv4Addr() string {
	if o == nil || IsNil(o.SourceIpv4Addr) {
		var ret string
		return ret
	}
	return *o.SourceIpv4Addr
}

// GetSourceIpv4AddrOk returns a tuple with the SourceIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaselineDnsQueryMdtInfo) GetSourceIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.SourceIpv4Addr) {
		return nil, false
	}
	return o.SourceIpv4Addr, true
}

// HasSourceIpv4Addr returns a boolean if a field has been set.
func (o *BaselineDnsQueryMdtInfo) HasSourceIpv4Addr() bool {
	if o != nil && !IsNil(o.SourceIpv4Addr) {
		return true
	}

	return false
}

// SetSourceIpv4Addr gets a reference to the given string and assigns it to the SourceIpv4Addr field.
func (o *BaselineDnsQueryMdtInfo) SetSourceIpv4Addr(v string) {
	o.SourceIpv4Addr = &v
}

// GetSourceIpv6Prefix returns the SourceIpv6Prefix field value if set, zero value otherwise.
func (o *BaselineDnsQueryMdtInfo) GetSourceIpv6Prefix() Ipv6Prefix {
	if o == nil || IsNil(o.SourceIpv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.SourceIpv6Prefix
}

// GetSourceIpv6PrefixOk returns a tuple with the SourceIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BaselineDnsQueryMdtInfo) GetSourceIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || IsNil(o.SourceIpv6Prefix) {
		return nil, false
	}
	return o.SourceIpv6Prefix, true
}

// HasSourceIpv6Prefix returns a boolean if a field has been set.
func (o *BaselineDnsQueryMdtInfo) HasSourceIpv6Prefix() bool {
	if o != nil && !IsNil(o.SourceIpv6Prefix) {
		return true
	}

	return false
}

// SetSourceIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the SourceIpv6Prefix field.
func (o *BaselineDnsQueryMdtInfo) SetSourceIpv6Prefix(v Ipv6Prefix) {
	o.SourceIpv6Prefix = &v
}

// GetBaseDnsMdtList returns the BaseDnsMdtList field value
func (o *BaselineDnsQueryMdtInfo) GetBaseDnsMdtList() []BaselineDnsMdtId {
	if o == nil {
		var ret []BaselineDnsMdtId
		return ret
	}

	return o.BaseDnsMdtList
}

// GetBaseDnsMdtListOk returns a tuple with the BaseDnsMdtList field value
// and a boolean to check if the value has been set.
func (o *BaselineDnsQueryMdtInfo) GetBaseDnsMdtListOk() ([]BaselineDnsMdtId, bool) {
	if o == nil {
		return nil, false
	}
	return o.BaseDnsMdtList, true
}

// SetBaseDnsMdtList sets field value
func (o *BaselineDnsQueryMdtInfo) SetBaseDnsMdtList(v []BaselineDnsMdtId) {
	o.BaseDnsMdtList = v
}

func (o BaselineDnsQueryMdtInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BaselineDnsQueryMdtInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SourceIpv4Addr) {
		toSerialize["sourceIpv4Addr"] = o.SourceIpv4Addr
	}
	if !IsNil(o.SourceIpv6Prefix) {
		toSerialize["sourceIpv6Prefix"] = o.SourceIpv6Prefix
	}
	toSerialize["baseDnsMdtList"] = o.BaseDnsMdtList
	return toSerialize, nil
}

type NullableBaselineDnsQueryMdtInfo struct {
	value *BaselineDnsQueryMdtInfo
	isSet bool
}

func (v NullableBaselineDnsQueryMdtInfo) Get() *BaselineDnsQueryMdtInfo {
	return v.value
}

func (v *NullableBaselineDnsQueryMdtInfo) Set(val *BaselineDnsQueryMdtInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableBaselineDnsQueryMdtInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableBaselineDnsQueryMdtInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBaselineDnsQueryMdtInfo(val *BaselineDnsQueryMdtInfo) *NullableBaselineDnsQueryMdtInfo {
	return &NullableBaselineDnsQueryMdtInfo{value: val, isSet: true}
}

func (v NullableBaselineDnsQueryMdtInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBaselineDnsQueryMdtInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
