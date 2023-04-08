/*
Nsmf_PDUSession

SMF PDU Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nsmf_PDUSession

import (
	"encoding/json"
)

// checks if the AfCoordinationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AfCoordinationInfo{}

// AfCoordinationInfo AF Coordination Information
type AfCoordinationInfo struct {
	// DNAI (Data network access identifier), see clause 5.6.7 of 3GPP TS 23.501.
	SourceDnai *string `json:"sourceDnai,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	SourceUeIpv4Addr *string `json:"sourceUeIpv4Addr,omitempty"`
	SourceUeIpv6Prefix *Ipv6Prefix `json:"sourceUeIpv6Prefix,omitempty"`
	NotificationInfoList []NotificationInfo `json:"notificationInfoList,omitempty"`
}

// NewAfCoordinationInfo instantiates a new AfCoordinationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAfCoordinationInfo() *AfCoordinationInfo {
	this := AfCoordinationInfo{}
	return &this
}

// NewAfCoordinationInfoWithDefaults instantiates a new AfCoordinationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAfCoordinationInfoWithDefaults() *AfCoordinationInfo {
	this := AfCoordinationInfo{}
	return &this
}

// GetSourceDnai returns the SourceDnai field value if set, zero value otherwise.
func (o *AfCoordinationInfo) GetSourceDnai() string {
	if o == nil || isNil(o.SourceDnai) {
		var ret string
		return ret
	}
	return *o.SourceDnai
}

// GetSourceDnaiOk returns a tuple with the SourceDnai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfCoordinationInfo) GetSourceDnaiOk() (*string, bool) {
	if o == nil || isNil(o.SourceDnai) {
		return nil, false
	}
	return o.SourceDnai, true
}

// HasSourceDnai returns a boolean if a field has been set.
func (o *AfCoordinationInfo) HasSourceDnai() bool {
	if o != nil && !isNil(o.SourceDnai) {
		return true
	}

	return false
}

// SetSourceDnai gets a reference to the given string and assigns it to the SourceDnai field.
func (o *AfCoordinationInfo) SetSourceDnai(v string) {
	o.SourceDnai = &v
}

// GetSourceUeIpv4Addr returns the SourceUeIpv4Addr field value if set, zero value otherwise.
func (o *AfCoordinationInfo) GetSourceUeIpv4Addr() string {
	if o == nil || isNil(o.SourceUeIpv4Addr) {
		var ret string
		return ret
	}
	return *o.SourceUeIpv4Addr
}

// GetSourceUeIpv4AddrOk returns a tuple with the SourceUeIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfCoordinationInfo) GetSourceUeIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.SourceUeIpv4Addr) {
		return nil, false
	}
	return o.SourceUeIpv4Addr, true
}

// HasSourceUeIpv4Addr returns a boolean if a field has been set.
func (o *AfCoordinationInfo) HasSourceUeIpv4Addr() bool {
	if o != nil && !isNil(o.SourceUeIpv4Addr) {
		return true
	}

	return false
}

// SetSourceUeIpv4Addr gets a reference to the given string and assigns it to the SourceUeIpv4Addr field.
func (o *AfCoordinationInfo) SetSourceUeIpv4Addr(v string) {
	o.SourceUeIpv4Addr = &v
}

// GetSourceUeIpv6Prefix returns the SourceUeIpv6Prefix field value if set, zero value otherwise.
func (o *AfCoordinationInfo) GetSourceUeIpv6Prefix() Ipv6Prefix {
	if o == nil || isNil(o.SourceUeIpv6Prefix) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.SourceUeIpv6Prefix
}

// GetSourceUeIpv6PrefixOk returns a tuple with the SourceUeIpv6Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfCoordinationInfo) GetSourceUeIpv6PrefixOk() (*Ipv6Prefix, bool) {
	if o == nil || isNil(o.SourceUeIpv6Prefix) {
		return nil, false
	}
	return o.SourceUeIpv6Prefix, true
}

// HasSourceUeIpv6Prefix returns a boolean if a field has been set.
func (o *AfCoordinationInfo) HasSourceUeIpv6Prefix() bool {
	if o != nil && !isNil(o.SourceUeIpv6Prefix) {
		return true
	}

	return false
}

// SetSourceUeIpv6Prefix gets a reference to the given Ipv6Prefix and assigns it to the SourceUeIpv6Prefix field.
func (o *AfCoordinationInfo) SetSourceUeIpv6Prefix(v Ipv6Prefix) {
	o.SourceUeIpv6Prefix = &v
}

// GetNotificationInfoList returns the NotificationInfoList field value if set, zero value otherwise.
func (o *AfCoordinationInfo) GetNotificationInfoList() []NotificationInfo {
	if o == nil || isNil(o.NotificationInfoList) {
		var ret []NotificationInfo
		return ret
	}
	return o.NotificationInfoList
}

// GetNotificationInfoListOk returns a tuple with the NotificationInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AfCoordinationInfo) GetNotificationInfoListOk() ([]NotificationInfo, bool) {
	if o == nil || isNil(o.NotificationInfoList) {
		return nil, false
	}
	return o.NotificationInfoList, true
}

// HasNotificationInfoList returns a boolean if a field has been set.
func (o *AfCoordinationInfo) HasNotificationInfoList() bool {
	if o != nil && !isNil(o.NotificationInfoList) {
		return true
	}

	return false
}

// SetNotificationInfoList gets a reference to the given []NotificationInfo and assigns it to the NotificationInfoList field.
func (o *AfCoordinationInfo) SetNotificationInfoList(v []NotificationInfo) {
	o.NotificationInfoList = v
}

func (o AfCoordinationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AfCoordinationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SourceDnai) {
		toSerialize["sourceDnai"] = o.SourceDnai
	}
	if !isNil(o.SourceUeIpv4Addr) {
		toSerialize["sourceUeIpv4Addr"] = o.SourceUeIpv4Addr
	}
	if !isNil(o.SourceUeIpv6Prefix) {
		toSerialize["sourceUeIpv6Prefix"] = o.SourceUeIpv6Prefix
	}
	if !isNil(o.NotificationInfoList) {
		toSerialize["notificationInfoList"] = o.NotificationInfoList
	}
	return toSerialize, nil
}

type NullableAfCoordinationInfo struct {
	value *AfCoordinationInfo
	isSet bool
}

func (v NullableAfCoordinationInfo) Get() *AfCoordinationInfo {
	return v.value
}

func (v *NullableAfCoordinationInfo) Set(val *AfCoordinationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAfCoordinationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAfCoordinationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAfCoordinationInfo(val *AfCoordinationInfo) *NullableAfCoordinationInfo {
	return &NullableAfCoordinationInfo{value: val, isSet: true}
}

func (v NullableAfCoordinationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAfCoordinationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


