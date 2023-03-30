/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the DddTrafficDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DddTrafficDescriptor{}

// DddTrafficDescriptor Contains a Traffic Descriptor.
type DddTrafficDescriptor struct {
	// String identifying a IPv4 address formatted in the 'dotted decimal' notation as defined in RFC 1166. 
	Ipv4Addr *string `json:"ipv4Addr,omitempty"`
	Ipv6Addr *Ipv6Addr `json:"ipv6Addr,omitempty"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber *int32 `json:"portNumber,omitempty"`
	// String identifying a MAC address formatted in the hexadecimal notation according to clause 1.1 and clause 2.1 of RFC 7042. 
	MacAddr *string `json:"macAddr,omitempty"`
}

// NewDddTrafficDescriptor instantiates a new DddTrafficDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDddTrafficDescriptor() *DddTrafficDescriptor {
	this := DddTrafficDescriptor{}
	return &this
}

// NewDddTrafficDescriptorWithDefaults instantiates a new DddTrafficDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDddTrafficDescriptorWithDefaults() *DddTrafficDescriptor {
	this := DddTrafficDescriptor{}
	return &this
}

// GetIpv4Addr returns the Ipv4Addr field value if set, zero value otherwise.
func (o *DddTrafficDescriptor) GetIpv4Addr() string {
	if o == nil || IsNil(o.Ipv4Addr) {
		var ret string
		return ret
	}
	return *o.Ipv4Addr
}

// GetIpv4AddrOk returns a tuple with the Ipv4Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DddTrafficDescriptor) GetIpv4AddrOk() (*string, bool) {
	if o == nil || IsNil(o.Ipv4Addr) {
		return nil, false
	}
	return o.Ipv4Addr, true
}

// HasIpv4Addr returns a boolean if a field has been set.
func (o *DddTrafficDescriptor) HasIpv4Addr() bool {
	if o != nil && !IsNil(o.Ipv4Addr) {
		return true
	}

	return false
}

// SetIpv4Addr gets a reference to the given string and assigns it to the Ipv4Addr field.
func (o *DddTrafficDescriptor) SetIpv4Addr(v string) {
	o.Ipv4Addr = &v
}

// GetIpv6Addr returns the Ipv6Addr field value if set, zero value otherwise.
func (o *DddTrafficDescriptor) GetIpv6Addr() Ipv6Addr {
	if o == nil || IsNil(o.Ipv6Addr) {
		var ret Ipv6Addr
		return ret
	}
	return *o.Ipv6Addr
}

// GetIpv6AddrOk returns a tuple with the Ipv6Addr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DddTrafficDescriptor) GetIpv6AddrOk() (*Ipv6Addr, bool) {
	if o == nil || IsNil(o.Ipv6Addr) {
		return nil, false
	}
	return o.Ipv6Addr, true
}

// HasIpv6Addr returns a boolean if a field has been set.
func (o *DddTrafficDescriptor) HasIpv6Addr() bool {
	if o != nil && !IsNil(o.Ipv6Addr) {
		return true
	}

	return false
}

// SetIpv6Addr gets a reference to the given Ipv6Addr and assigns it to the Ipv6Addr field.
func (o *DddTrafficDescriptor) SetIpv6Addr(v Ipv6Addr) {
	o.Ipv6Addr = &v
}

// GetPortNumber returns the PortNumber field value if set, zero value otherwise.
func (o *DddTrafficDescriptor) GetPortNumber() int32 {
	if o == nil || IsNil(o.PortNumber) {
		var ret int32
		return ret
	}
	return *o.PortNumber
}

// GetPortNumberOk returns a tuple with the PortNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DddTrafficDescriptor) GetPortNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.PortNumber) {
		return nil, false
	}
	return o.PortNumber, true
}

// HasPortNumber returns a boolean if a field has been set.
func (o *DddTrafficDescriptor) HasPortNumber() bool {
	if o != nil && !IsNil(o.PortNumber) {
		return true
	}

	return false
}

// SetPortNumber gets a reference to the given int32 and assigns it to the PortNumber field.
func (o *DddTrafficDescriptor) SetPortNumber(v int32) {
	o.PortNumber = &v
}

// GetMacAddr returns the MacAddr field value if set, zero value otherwise.
func (o *DddTrafficDescriptor) GetMacAddr() string {
	if o == nil || IsNil(o.MacAddr) {
		var ret string
		return ret
	}
	return *o.MacAddr
}

// GetMacAddrOk returns a tuple with the MacAddr field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DddTrafficDescriptor) GetMacAddrOk() (*string, bool) {
	if o == nil || IsNil(o.MacAddr) {
		return nil, false
	}
	return o.MacAddr, true
}

// HasMacAddr returns a boolean if a field has been set.
func (o *DddTrafficDescriptor) HasMacAddr() bool {
	if o != nil && !IsNil(o.MacAddr) {
		return true
	}

	return false
}

// SetMacAddr gets a reference to the given string and assigns it to the MacAddr field.
func (o *DddTrafficDescriptor) SetMacAddr(v string) {
	o.MacAddr = &v
}

func (o DddTrafficDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DddTrafficDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Ipv4Addr) {
		toSerialize["ipv4Addr"] = o.Ipv4Addr
	}
	if !IsNil(o.Ipv6Addr) {
		toSerialize["ipv6Addr"] = o.Ipv6Addr
	}
	if !IsNil(o.PortNumber) {
		toSerialize["portNumber"] = o.PortNumber
	}
	if !IsNil(o.MacAddr) {
		toSerialize["macAddr"] = o.MacAddr
	}
	return toSerialize, nil
}

type NullableDddTrafficDescriptor struct {
	value *DddTrafficDescriptor
	isSet bool
}

func (v NullableDddTrafficDescriptor) Get() *DddTrafficDescriptor {
	return v.value
}

func (v *NullableDddTrafficDescriptor) Set(val *DddTrafficDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableDddTrafficDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableDddTrafficDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDddTrafficDescriptor(val *DddTrafficDescriptor) *NullableDddTrafficDescriptor {
	return &NullableDddTrafficDescriptor{value: val, isSet: true}
}

func (v NullableDddTrafficDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDddTrafficDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

