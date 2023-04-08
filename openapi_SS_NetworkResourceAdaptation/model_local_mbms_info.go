/*
SS_NetworkResourceAdaptation

SS Network Resource Adaptation Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_NetworkResourceAdaptation

import (
	"encoding/json"
)

// checks if the LocalMbmsInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocalMbmsInfo{}

// LocalMbmsInfo Contains the local MBMS information.
type LocalMbmsInfo struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	MbmsEnbIpv4MulAddr *string `json:"mbmsEnbIpv4MulAddr,omitempty"`
	MbmsEnbIpv6MulAddr *Ipv6Prefix `json:"mbmsEnbIpv6MulAddr,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	MbmsGwIpv4SsmAddr *string `json:"mbmsGwIpv4SsmAddr,omitempty"`
	MbmsGwIpv6SsmAddr *Ipv6Addr `json:"mbmsGwIpv6SsmAddr,omitempty"`
	Cteid *string `json:"cteid,omitempty"`
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	BmscIpv4Addr *string `json:"bmscIpv4Addr,omitempty"`
	BmscIpv6Addr *Ipv6Addr `json:"bmscIpv6Addr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	BmscPort *int32 `json:"bmscPort,omitempty"`
}

// NewLocalMbmsInfo instantiates a new LocalMbmsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocalMbmsInfo() *LocalMbmsInfo {
	this := LocalMbmsInfo{}
	return &this
}

// NewLocalMbmsInfoWithDefaults instantiates a new LocalMbmsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocalMbmsInfoWithDefaults() *LocalMbmsInfo {
	this := LocalMbmsInfo{}
	return &this
}

// GetMbmsEnbIpv4MulAddr returns the MbmsEnbIpv4MulAddr field value if set, zero value otherwise.
func (o *LocalMbmsInfo) GetMbmsEnbIpv4MulAddr() string {
	if o == nil || isNil(o.MbmsEnbIpv4MulAddr) {
		var ret string
		return ret
	}
	return *o.MbmsEnbIpv4MulAddr
}

// GetMbmsEnbIpv4MulAddrOk returns a tuple with the MbmsEnbIpv4MulAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalMbmsInfo) GetMbmsEnbIpv4MulAddrOk() (*string, bool) {
	if o == nil || isNil(o.MbmsEnbIpv4MulAddr) {
		return nil, false
	}
	return o.MbmsEnbIpv4MulAddr, true
}

// HasMbmsEnbIpv4MulAddr returns a boolean if a field has been set.
func (o *LocalMbmsInfo) HasMbmsEnbIpv4MulAddr() bool {
	if o != nil && !isNil(o.MbmsEnbIpv4MulAddr) {
		return true
	}

	return false
}

// SetMbmsEnbIpv4MulAddr gets a reference to the given string and assigns it to the MbmsEnbIpv4MulAddr field.
func (o *LocalMbmsInfo) SetMbmsEnbIpv4MulAddr(v string) {
	o.MbmsEnbIpv4MulAddr = &v
}

// GetMbmsEnbIpv6MulAddr returns the MbmsEnbIpv6MulAddr field value if set, zero value otherwise.
func (o *LocalMbmsInfo) GetMbmsEnbIpv6MulAddr() Ipv6Prefix {
	if o == nil || isNil(o.MbmsEnbIpv6MulAddr) {
		var ret Ipv6Prefix
		return ret
	}
	return *o.MbmsEnbIpv6MulAddr
}

// GetMbmsEnbIpv6MulAddrOk returns a tuple with the MbmsEnbIpv6MulAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalMbmsInfo) GetMbmsEnbIpv6MulAddrOk() (*Ipv6Prefix, bool) {
	if o == nil || isNil(o.MbmsEnbIpv6MulAddr) {
		return nil, false
	}
	return o.MbmsEnbIpv6MulAddr, true
}

// HasMbmsEnbIpv6MulAddr returns a boolean if a field has been set.
func (o *LocalMbmsInfo) HasMbmsEnbIpv6MulAddr() bool {
	if o != nil && !isNil(o.MbmsEnbIpv6MulAddr) {
		return true
	}

	return false
}

// SetMbmsEnbIpv6MulAddr gets a reference to the given Ipv6Prefix and assigns it to the MbmsEnbIpv6MulAddr field.
func (o *LocalMbmsInfo) SetMbmsEnbIpv6MulAddr(v Ipv6Prefix) {
	o.MbmsEnbIpv6MulAddr = &v
}

// GetMbmsGwIpv4SsmAddr returns the MbmsGwIpv4SsmAddr field value if set, zero value otherwise.
func (o *LocalMbmsInfo) GetMbmsGwIpv4SsmAddr() string {
	if o == nil || isNil(o.MbmsGwIpv4SsmAddr) {
		var ret string
		return ret
	}
	return *o.MbmsGwIpv4SsmAddr
}

// GetMbmsGwIpv4SsmAddrOk returns a tuple with the MbmsGwIpv4SsmAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalMbmsInfo) GetMbmsGwIpv4SsmAddrOk() (*string, bool) {
	if o == nil || isNil(o.MbmsGwIpv4SsmAddr) {
		return nil, false
	}
	return o.MbmsGwIpv4SsmAddr, true
}

// HasMbmsGwIpv4SsmAddr returns a boolean if a field has been set.
func (o *LocalMbmsInfo) HasMbmsGwIpv4SsmAddr() bool {
	if o != nil && !isNil(o.MbmsGwIpv4SsmAddr) {
		return true
	}

	return false
}

// SetMbmsGwIpv4SsmAddr gets a reference to the given string and assigns it to the MbmsGwIpv4SsmAddr field.
func (o *LocalMbmsInfo) SetMbmsGwIpv4SsmAddr(v string) {
	o.MbmsGwIpv4SsmAddr = &v
}

// GetMbmsGwIpv6SsmAddr returns the MbmsGwIpv6SsmAddr field value if set, zero value otherwise.
func (o *LocalMbmsInfo) GetMbmsGwIpv6SsmAddr() Ipv6Addr {
	if o == nil || isNil(o.MbmsGwIpv6SsmAddr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.MbmsGwIpv6SsmAddr
}

// GetMbmsGwIpv6SsmAddrOk returns a tuple with the MbmsGwIpv6SsmAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalMbmsInfo) GetMbmsGwIpv6SsmAddrOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.MbmsGwIpv6SsmAddr) {
		return nil, false
	}
	return o.MbmsGwIpv6SsmAddr, true
}

// HasMbmsGwIpv6SsmAddr returns a boolean if a field has been set.
func (o *LocalMbmsInfo) HasMbmsGwIpv6SsmAddr() bool {
	if o != nil && !isNil(o.MbmsGwIpv6SsmAddr) {
		return true
	}

	return false
}

// SetMbmsGwIpv6SsmAddr gets a reference to the given Ipv6Addr and assigns it to the MbmsGwIpv6SsmAddr field.
func (o *LocalMbmsInfo) SetMbmsGwIpv6SsmAddr(v Ipv6Addr) {
	o.MbmsGwIpv6SsmAddr = &v
}

// GetCteid returns the Cteid field value if set, zero value otherwise.
func (o *LocalMbmsInfo) GetCteid() string {
	if o == nil || isNil(o.Cteid) {
		var ret string
		return ret
	}
	return *o.Cteid
}

// GetCteidOk returns a tuple with the Cteid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalMbmsInfo) GetCteidOk() (*string, bool) {
	if o == nil || isNil(o.Cteid) {
		return nil, false
	}
	return o.Cteid, true
}

// HasCteid returns a boolean if a field has been set.
func (o *LocalMbmsInfo) HasCteid() bool {
	if o != nil && !isNil(o.Cteid) {
		return true
	}

	return false
}

// SetCteid gets a reference to the given string and assigns it to the Cteid field.
func (o *LocalMbmsInfo) SetCteid(v string) {
	o.Cteid = &v
}

// GetBmscIpv4Addr returns the BmscIpv4Addr field value if set, zero value otherwise.
func (o *LocalMbmsInfo) GetBmscIpv4Addr() string {
	if o == nil || isNil(o.BmscIpv4Addr) {
		var ret string
		return ret
	}
	return *o.BmscIpv4Addr
}

// GetBmscIpv4AddrOk returns a tuple with the BmscIpv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalMbmsInfo) GetBmscIpv4AddrOk() (*string, bool) {
	if o == nil || isNil(o.BmscIpv4Addr) {
		return nil, false
	}
	return o.BmscIpv4Addr, true
}

// HasBmscIpv4Addr returns a boolean if a field has been set.
func (o *LocalMbmsInfo) HasBmscIpv4Addr() bool {
	if o != nil && !isNil(o.BmscIpv4Addr) {
		return true
	}

	return false
}

// SetBmscIpv4Addr gets a reference to the given string and assigns it to the BmscIpv4Addr field.
func (o *LocalMbmsInfo) SetBmscIpv4Addr(v string) {
	o.BmscIpv4Addr = &v
}

// GetBmscIpv6Addr returns the BmscIpv6Addr field value if set, zero value otherwise.
func (o *LocalMbmsInfo) GetBmscIpv6Addr() Ipv6Addr {
	if o == nil || isNil(o.BmscIpv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.BmscIpv6Addr
}

// GetBmscIpv6AddrOk returns a tuple with the BmscIpv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalMbmsInfo) GetBmscIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || isNil(o.BmscIpv6Addr) {
		return nil, false
	}
	return o.BmscIpv6Addr, true
}

// HasBmscIpv6Addr returns a boolean if a field has been set.
func (o *LocalMbmsInfo) HasBmscIpv6Addr() bool {
	if o != nil && !isNil(o.BmscIpv6Addr) {
		return true
	}

	return false
}

// SetBmscIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the BmscIpv6Addr field.
func (o *LocalMbmsInfo) SetBmscIpv6Addr(v Ipv6Addr) {
	o.BmscIpv6Addr = &v
}

// GetBmscPort returns the BmscPort field value if set, zero value otherwise.
func (o *LocalMbmsInfo) GetBmscPort() int32 {
	if o == nil || isNil(o.BmscPort) {
		var ret int32
		return ret
	}
	return *o.BmscPort
}

// GetBmscPortOk returns a tuple with the BmscPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LocalMbmsInfo) GetBmscPortOk() (*int32, bool) {
	if o == nil || isNil(o.BmscPort) {
		return nil, false
	}
	return o.BmscPort, true
}

// HasBmscPort returns a boolean if a field has been set.
func (o *LocalMbmsInfo) HasBmscPort() bool {
	if o != nil && !isNil(o.BmscPort) {
		return true
	}

	return false
}

// SetBmscPort gets a reference to the given int32 and assigns it to the BmscPort field.
func (o *LocalMbmsInfo) SetBmscPort(v int32) {
	o.BmscPort = &v
}

func (o LocalMbmsInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocalMbmsInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.MbmsEnbIpv4MulAddr) {
		toSerialize["mbmsEnbIpv4MulAddr"] = o.MbmsEnbIpv4MulAddr
	}
	if !isNil(o.MbmsEnbIpv6MulAddr) {
		toSerialize["mbmsEnbIpv6MulAddr"] = o.MbmsEnbIpv6MulAddr
	}
	if !isNil(o.MbmsGwIpv4SsmAddr) {
		toSerialize["mbmsGwIpv4SsmAddr"] = o.MbmsGwIpv4SsmAddr
	}
	if !isNil(o.MbmsGwIpv6SsmAddr) {
		toSerialize["mbmsGwIpv6SsmAddr"] = o.MbmsGwIpv6SsmAddr
	}
	if !isNil(o.Cteid) {
		toSerialize["cteid"] = o.Cteid
	}
	if !isNil(o.BmscIpv4Addr) {
		toSerialize["bmscIpv4Addr"] = o.BmscIpv4Addr
	}
	if !isNil(o.BmscIpv6Addr) {
		toSerialize["bmscIpv6Addr"] = o.BmscIpv6Addr
	}
	if !isNil(o.BmscPort) {
		toSerialize["bmscPort"] = o.BmscPort
	}
	return toSerialize, nil
}

type NullableLocalMbmsInfo struct {
	value *LocalMbmsInfo
	isSet bool
}

func (v NullableLocalMbmsInfo) Get() *LocalMbmsInfo {
	return v.value
}

func (v *NullableLocalMbmsInfo) Set(val *LocalMbmsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableLocalMbmsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableLocalMbmsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocalMbmsInfo(val *LocalMbmsInfo) *NullableLocalMbmsInfo {
	return &NullableLocalMbmsInfo{value: val, isSet: true}
}

func (v NullableLocalMbmsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocalMbmsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


