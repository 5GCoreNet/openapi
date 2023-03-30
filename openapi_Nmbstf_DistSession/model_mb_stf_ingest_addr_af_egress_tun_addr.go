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

// checks if the MbStfIngestAddrAfEgressTunAddr type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbStfIngestAddrAfEgressTunAddr{}

// MbStfIngestAddrAfEgressTunAddr struct for MbStfIngestAddrAfEgressTunAddr
type MbStfIngestAddrAfEgressTunAddr struct {
}

// NewMbStfIngestAddrAfEgressTunAddr instantiates a new MbStfIngestAddrAfEgressTunAddr object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbStfIngestAddrAfEgressTunAddr() *MbStfIngestAddrAfEgressTunAddr {
	this := MbStfIngestAddrAfEgressTunAddr{}
	return &this
}

// NewMbStfIngestAddrAfEgressTunAddrWithDefaults instantiates a new MbStfIngestAddrAfEgressTunAddr object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbStfIngestAddrAfEgressTunAddrWithDefaults() *MbStfIngestAddrAfEgressTunAddr {
	this := MbStfIngestAddrAfEgressTunAddr{}
	return &this
}

func (o MbStfIngestAddrAfEgressTunAddr) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbStfIngestAddrAfEgressTunAddr) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	return toSerialize, nil
}

type NullableMbStfIngestAddrAfEgressTunAddr struct {
	value *MbStfIngestAddrAfEgressTunAddr
	isSet bool
}

func (v NullableMbStfIngestAddrAfEgressTunAddr) Get() *MbStfIngestAddrAfEgressTunAddr {
	return v.value
}

func (v *NullableMbStfIngestAddrAfEgressTunAddr) Set(val *MbStfIngestAddrAfEgressTunAddr) {
	v.value = val
	v.isSet = true
}

func (v NullableMbStfIngestAddrAfEgressTunAddr) IsSet() bool {
	return v.isSet
}

func (v *NullableMbStfIngestAddrAfEgressTunAddr) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbStfIngestAddrAfEgressTunAddr(val *MbStfIngestAddrAfEgressTunAddr) *NullableMbStfIngestAddrAfEgressTunAddr {
	return &NullableMbStfIngestAddrAfEgressTunAddr{value: val, isSet: true}
}

func (v NullableMbStfIngestAddrAfEgressTunAddr) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbStfIngestAddrAfEgressTunAddr) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


