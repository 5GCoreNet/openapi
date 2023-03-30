/*
nmbsf-mbs-ud-ingest

API for MBS User Data Ingest Session Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmbsf_MBSUserDataIngestSession

import (
	"encoding/json"
)

// checks if the PacketDistrMethInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PacketDistrMethInfo{}

// PacketDistrMethInfo Represents additional MBS Distribution Session parameters for the case of Packet  Distribution Method. 
type PacketDistrMethInfo struct {
	OperatingMode PktDistributionOperatingMode `json:"operatingMode"`
	PckIngMethod PktIngestMethod `json:"pckIngMethod"`
	IngEndpointAddrs MbStfIngestAddr `json:"ingEndpointAddrs"`
}

// NewPacketDistrMethInfo instantiates a new PacketDistrMethInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPacketDistrMethInfo(operatingMode PktDistributionOperatingMode, pckIngMethod PktIngestMethod, ingEndpointAddrs MbStfIngestAddr) *PacketDistrMethInfo {
	this := PacketDistrMethInfo{}
	this.OperatingMode = operatingMode
	this.PckIngMethod = pckIngMethod
	this.IngEndpointAddrs = ingEndpointAddrs
	return &this
}

// NewPacketDistrMethInfoWithDefaults instantiates a new PacketDistrMethInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPacketDistrMethInfoWithDefaults() *PacketDistrMethInfo {
	this := PacketDistrMethInfo{}
	return &this
}

// GetOperatingMode returns the OperatingMode field value
func (o *PacketDistrMethInfo) GetOperatingMode() PktDistributionOperatingMode {
	if o == nil {
		var ret PktDistributionOperatingMode
		return ret
	}

	return o.OperatingMode
}

// GetOperatingModeOk returns a tuple with the OperatingMode field value
// and a boolean to check if the value has been set.
func (o *PacketDistrMethInfo) GetOperatingModeOk() (*PktDistributionOperatingMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OperatingMode, true
}

// SetOperatingMode sets field value
func (o *PacketDistrMethInfo) SetOperatingMode(v PktDistributionOperatingMode) {
	o.OperatingMode = v
}

// GetPckIngMethod returns the PckIngMethod field value
func (o *PacketDistrMethInfo) GetPckIngMethod() PktIngestMethod {
	if o == nil {
		var ret PktIngestMethod
		return ret
	}

	return o.PckIngMethod
}

// GetPckIngMethodOk returns a tuple with the PckIngMethod field value
// and a boolean to check if the value has been set.
func (o *PacketDistrMethInfo) GetPckIngMethodOk() (*PktIngestMethod, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PckIngMethod, true
}

// SetPckIngMethod sets field value
func (o *PacketDistrMethInfo) SetPckIngMethod(v PktIngestMethod) {
	o.PckIngMethod = v
}

// GetIngEndpointAddrs returns the IngEndpointAddrs field value
func (o *PacketDistrMethInfo) GetIngEndpointAddrs() MbStfIngestAddr {
	if o == nil {
		var ret MbStfIngestAddr
		return ret
	}

	return o.IngEndpointAddrs
}

// GetIngEndpointAddrsOk returns a tuple with the IngEndpointAddrs field value
// and a boolean to check if the value has been set.
func (o *PacketDistrMethInfo) GetIngEndpointAddrsOk() (*MbStfIngestAddr, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IngEndpointAddrs, true
}

// SetIngEndpointAddrs sets field value
func (o *PacketDistrMethInfo) SetIngEndpointAddrs(v MbStfIngestAddr) {
	o.IngEndpointAddrs = v
}

func (o PacketDistrMethInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PacketDistrMethInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["operatingMode"] = o.OperatingMode
	toSerialize["pckIngMethod"] = o.PckIngMethod
	toSerialize["ingEndpointAddrs"] = o.IngEndpointAddrs
	return toSerialize, nil
}

type NullablePacketDistrMethInfo struct {
	value *PacketDistrMethInfo
	isSet bool
}

func (v NullablePacketDistrMethInfo) Get() *PacketDistrMethInfo {
	return v.value
}

func (v *NullablePacketDistrMethInfo) Set(val *PacketDistrMethInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePacketDistrMethInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePacketDistrMethInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePacketDistrMethInfo(val *PacketDistrMethInfo) *NullablePacketDistrMethInfo {
	return &NullablePacketDistrMethInfo{value: val, isSet: true}
}

func (v NullablePacketDistrMethInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePacketDistrMethInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

