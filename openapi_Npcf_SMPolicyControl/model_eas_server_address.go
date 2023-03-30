/*
Npcf_SMPolicyControl API

Session Management Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_SMPolicyControl

import (
	"encoding/json"
)

// checks if the EasServerAddress type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EasServerAddress{}

// EasServerAddress Represents the IP address and port of an EAS server.
type EasServerAddress struct {
	Ip IpAddr `json:"ip"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	Port int32 `json:"port"`
}

// NewEasServerAddress instantiates a new EasServerAddress object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEasServerAddress(ip IpAddr, port int32) *EasServerAddress {
	this := EasServerAddress{}
	this.Ip = ip
	this.Port = port
	return &this
}

// NewEasServerAddressWithDefaults instantiates a new EasServerAddress object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEasServerAddressWithDefaults() *EasServerAddress {
	this := EasServerAddress{}
	return &this
}

// GetIp returns the Ip field value
func (o *EasServerAddress) GetIp() IpAddr {
	if o == nil {
		var ret IpAddr
		return ret
	}

	return o.Ip
}

// GetIpOk returns a tuple with the Ip field value
// and a boolean to check if the value has been set.
func (o *EasServerAddress) GetIpOk() (*IpAddr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ip, true
}

// SetIp sets field value
func (o *EasServerAddress) SetIp(v IpAddr) {
	o.Ip = v
}

// GetPort returns the Port field value
func (o *EasServerAddress) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *EasServerAddress) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *EasServerAddress) SetPort(v int32) {
	o.Port = v
}

func (o EasServerAddress) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EasServerAddress) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ip"] = o.Ip
	toSerialize["port"] = o.Port
	return toSerialize, nil
}

type NullableEasServerAddress struct {
	value *EasServerAddress
	isSet bool
}

func (v NullableEasServerAddress) Get() *EasServerAddress {
	return v.value
}

func (v *NullableEasServerAddress) Set(val *EasServerAddress) {
	v.value = val
	v.isSet = true
}

func (v NullableEasServerAddress) IsSet() bool {
	return v.isSet
}

func (v *NullableEasServerAddress) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEasServerAddress(val *EasServerAddress) *NullableEasServerAddress {
	return &NullableEasServerAddress{value: val, isSet: true}
}

func (v NullableEasServerAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEasServerAddress) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


