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

// checks if the EcsOption type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EcsOption{}

// EcsOption ECS Option Information
type EcsOption struct {
	SourcePrefixLength int32 `json:"sourcePrefixLength"`
	ScopePrefixLength *int32 `json:"scopePrefixLength,omitempty"`
	IpAddr IpAddr `json:"ipAddr"`
}

// NewEcsOption instantiates a new EcsOption object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEcsOption(sourcePrefixLength int32, ipAddr IpAddr) *EcsOption {
	this := EcsOption{}
	this.SourcePrefixLength = sourcePrefixLength
	this.IpAddr = ipAddr
	return &this
}

// NewEcsOptionWithDefaults instantiates a new EcsOption object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEcsOptionWithDefaults() *EcsOption {
	this := EcsOption{}
	return &this
}

// GetSourcePrefixLength returns the SourcePrefixLength field value
func (o *EcsOption) GetSourcePrefixLength() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SourcePrefixLength
}

// GetSourcePrefixLengthOk returns a tuple with the SourcePrefixLength field value
// and a boolean to check if the value has been set.
func (o *EcsOption) GetSourcePrefixLengthOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourcePrefixLength, true
}

// SetSourcePrefixLength sets field value
func (o *EcsOption) SetSourcePrefixLength(v int32) {
	o.SourcePrefixLength = v
}

// GetScopePrefixLength returns the ScopePrefixLength field value if set, zero value otherwise.
func (o *EcsOption) GetScopePrefixLength() int32 {
	if o == nil || IsNil(o.ScopePrefixLength) {
		var ret int32
		return ret
	}
	return *o.ScopePrefixLength
}

// GetScopePrefixLengthOk returns a tuple with the ScopePrefixLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EcsOption) GetScopePrefixLengthOk() (*int32, bool) {
	if o == nil || IsNil(o.ScopePrefixLength) {
		return nil, false
	}
	return o.ScopePrefixLength, true
}

// HasScopePrefixLength returns a boolean if a field has been set.
func (o *EcsOption) HasScopePrefixLength() bool {
	if o != nil && !IsNil(o.ScopePrefixLength) {
		return true
	}

	return false
}

// SetScopePrefixLength gets a reference to the given int32 and assigns it to the ScopePrefixLength field.
func (o *EcsOption) SetScopePrefixLength(v int32) {
	o.ScopePrefixLength = &v
}

// GetIpAddr returns the IpAddr field value
func (o *EcsOption) GetIpAddr() IpAddr {
	if o == nil {
		var ret IpAddr
		return ret
	}

	return o.IpAddr
}

// GetIpAddrOk returns a tuple with the IpAddr field value
// and a boolean to check if the value has been set.
func (o *EcsOption) GetIpAddrOk() (*IpAddr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IpAddr, true
}

// SetIpAddr sets field value
func (o *EcsOption) SetIpAddr(v IpAddr) {
	o.IpAddr = v
}

func (o EcsOption) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EcsOption) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sourcePrefixLength"] = o.SourcePrefixLength
	if !IsNil(o.ScopePrefixLength) {
		toSerialize["scopePrefixLength"] = o.ScopePrefixLength
	}
	toSerialize["ipAddr"] = o.IpAddr
	return toSerialize, nil
}

type NullableEcsOption struct {
	value *EcsOption
	isSet bool
}

func (v NullableEcsOption) Get() *EcsOption {
	return v.value
}

func (v *NullableEcsOption) Set(val *EcsOption) {
	v.value = val
	v.isSet = true
}

func (v NullableEcsOption) IsSet() bool {
	return v.isSet
}

func (v *NullableEcsOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEcsOption(val *EcsOption) *NullableEcsOption {
	return &NullableEcsOption{value: val, isSet: true}
}

func (v NullableEcsOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEcsOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


