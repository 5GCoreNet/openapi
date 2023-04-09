/*
Nmbstf-distsession

MBSTF Distribution Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbstf_DistSession

import (
	"encoding/json"
)

// checks if the UpTrafficFlowInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpTrafficFlowInfo{}

// UpTrafficFlowInfo IP Multicast Address and Port Number
type UpTrafficFlowInfo struct {
	DestIpAddr IpAddr `json:"destIpAddr"`
	// Unsigned Integer, i.e. only value 0 and integers above 0 are permissible.
	PortNumber int32 `json:"portNumber"`
}

// NewUpTrafficFlowInfo instantiates a new UpTrafficFlowInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpTrafficFlowInfo(destIpAddr IpAddr, portNumber int32) *UpTrafficFlowInfo {
	this := UpTrafficFlowInfo{}
	this.DestIpAddr = destIpAddr
	this.PortNumber = portNumber
	return &this
}

// NewUpTrafficFlowInfoWithDefaults instantiates a new UpTrafficFlowInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpTrafficFlowInfoWithDefaults() *UpTrafficFlowInfo {
	this := UpTrafficFlowInfo{}
	return &this
}

// GetDestIpAddr returns the DestIpAddr field value
func (o *UpTrafficFlowInfo) GetDestIpAddr() IpAddr {
	if o == nil {
		var ret IpAddr
		return ret
	}

	return o.DestIpAddr
}

// GetDestIpAddrOk returns a tuple with the DestIpAddr field value
// and a boolean to check if the value has been set.
func (o *UpTrafficFlowInfo) GetDestIpAddrOk() (*IpAddr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DestIpAddr, true
}

// SetDestIpAddr sets field value
func (o *UpTrafficFlowInfo) SetDestIpAddr(v IpAddr) {
	o.DestIpAddr = v
}

// GetPortNumber returns the PortNumber field value
func (o *UpTrafficFlowInfo) GetPortNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PortNumber
}

// GetPortNumberOk returns a tuple with the PortNumber field value
// and a boolean to check if the value has been set.
func (o *UpTrafficFlowInfo) GetPortNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortNumber, true
}

// SetPortNumber sets field value
func (o *UpTrafficFlowInfo) SetPortNumber(v int32) {
	o.PortNumber = v
}

func (o UpTrafficFlowInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpTrafficFlowInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destIpAddr"] = o.DestIpAddr
	toSerialize["portNumber"] = o.PortNumber
	return toSerialize, nil
}

type NullableUpTrafficFlowInfo struct {
	value *UpTrafficFlowInfo
	isSet bool
}

func (v NullableUpTrafficFlowInfo) Get() *UpTrafficFlowInfo {
	return v.value
}

func (v *NullableUpTrafficFlowInfo) Set(val *UpTrafficFlowInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUpTrafficFlowInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUpTrafficFlowInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpTrafficFlowInfo(val *UpTrafficFlowInfo) *NullableUpTrafficFlowInfo {
	return &NullableUpTrafficFlowInfo{value: val, isSet: true}
}

func (v NullableUpTrafficFlowInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpTrafficFlowInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
