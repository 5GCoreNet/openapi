/*
Nmbsmf-MBSSession

MB-SMF MBSSession Service. © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsmf_MBSSession

import (
	"encoding/json"
)

// checks if the IngressTunAddrInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IngressTunAddrInfo{}

// IngressTunAddrInfo Ingress Tunnel Address Information
type IngressTunAddrInfo struct {
	IngressTunAddr []TunnelAddress `json:"ingressTunAddr"`
}

// NewIngressTunAddrInfo instantiates a new IngressTunAddrInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIngressTunAddrInfo(ingressTunAddr []TunnelAddress) *IngressTunAddrInfo {
	this := IngressTunAddrInfo{}
	this.IngressTunAddr = ingressTunAddr
	return &this
}

// NewIngressTunAddrInfoWithDefaults instantiates a new IngressTunAddrInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIngressTunAddrInfoWithDefaults() *IngressTunAddrInfo {
	this := IngressTunAddrInfo{}
	return &this
}

// GetIngressTunAddr returns the IngressTunAddr field value
func (o *IngressTunAddrInfo) GetIngressTunAddr() []TunnelAddress {
	if o == nil {
		var ret []TunnelAddress
		return ret
	}

	return o.IngressTunAddr
}

// GetIngressTunAddrOk returns a tuple with the IngressTunAddr field value
// and a boolean to check if the value has been set.
func (o *IngressTunAddrInfo) GetIngressTunAddrOk() ([]TunnelAddress, bool) {
	if o == nil {
		return nil, false
	}
	return o.IngressTunAddr, true
}

// SetIngressTunAddr sets field value
func (o *IngressTunAddrInfo) SetIngressTunAddr(v []TunnelAddress) {
	o.IngressTunAddr = v
}

func (o IngressTunAddrInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IngressTunAddrInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["ingressTunAddr"] = o.IngressTunAddr
	return toSerialize, nil
}

type NullableIngressTunAddrInfo struct {
	value *IngressTunAddrInfo
	isSet bool
}

func (v NullableIngressTunAddrInfo) Get() *IngressTunAddrInfo {
	return v.value
}

func (v *NullableIngressTunAddrInfo) Set(val *IngressTunAddrInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableIngressTunAddrInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableIngressTunAddrInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIngressTunAddrInfo(val *IngressTunAddrInfo) *NullableIngressTunAddrInfo {
	return &NullableIngressTunAddrInfo{value: val, isSet: true}
}

func (v NullableIngressTunAddrInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIngressTunAddrInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
