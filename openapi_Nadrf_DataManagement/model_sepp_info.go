/*
Nadrf_DataManagement

ADRF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nadrf_DataManagement

import (
	"encoding/json"
)

// checks if the SeppInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SeppInfo{}

// SeppInfo Information of a SEPP Instance
type SeppInfo struct {
	SeppPrefix *string `json:"seppPrefix,omitempty"`
	// Port numbers for HTTP and HTTPS. The key of the map shall be \"http\" or \"https\". 
	SeppPorts *map[string]int32 `json:"seppPorts,omitempty"`
	RemotePlmnList []PlmnId `json:"remotePlmnList,omitempty"`
	RemoteSnpnList []PlmnIdNid `json:"remoteSnpnList,omitempty"`
	// N32 purposes supported by the SEPP
	N32Purposes []N32Purpose `json:"n32Purposes,omitempty"`
}

// NewSeppInfo instantiates a new SeppInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSeppInfo() *SeppInfo {
	this := SeppInfo{}
	return &this
}

// NewSeppInfoWithDefaults instantiates a new SeppInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSeppInfoWithDefaults() *SeppInfo {
	this := SeppInfo{}
	return &this
}

// GetSeppPrefix returns the SeppPrefix field value if set, zero value otherwise.
func (o *SeppInfo) GetSeppPrefix() string {
	if o == nil || isNil(o.SeppPrefix) {
		var ret string
		return ret
	}
	return *o.SeppPrefix
}

// GetSeppPrefixOk returns a tuple with the SeppPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppInfo) GetSeppPrefixOk() (*string, bool) {
	if o == nil || isNil(o.SeppPrefix) {
		return nil, false
	}
	return o.SeppPrefix, true
}

// HasSeppPrefix returns a boolean if a field has been set.
func (o *SeppInfo) HasSeppPrefix() bool {
	if o != nil && !isNil(o.SeppPrefix) {
		return true
	}

	return false
}

// SetSeppPrefix gets a reference to the given string and assigns it to the SeppPrefix field.
func (o *SeppInfo) SetSeppPrefix(v string) {
	o.SeppPrefix = &v
}

// GetSeppPorts returns the SeppPorts field value if set, zero value otherwise.
func (o *SeppInfo) GetSeppPorts() map[string]int32 {
	if o == nil || isNil(o.SeppPorts) {
		var ret map[string]int32
		return ret
	}
	return *o.SeppPorts
}

// GetSeppPortsOk returns a tuple with the SeppPorts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppInfo) GetSeppPortsOk() (*map[string]int32, bool) {
	if o == nil || isNil(o.SeppPorts) {
		return nil, false
	}
	return o.SeppPorts, true
}

// HasSeppPorts returns a boolean if a field has been set.
func (o *SeppInfo) HasSeppPorts() bool {
	if o != nil && !isNil(o.SeppPorts) {
		return true
	}

	return false
}

// SetSeppPorts gets a reference to the given map[string]int32 and assigns it to the SeppPorts field.
func (o *SeppInfo) SetSeppPorts(v map[string]int32) {
	o.SeppPorts = &v
}

// GetRemotePlmnList returns the RemotePlmnList field value if set, zero value otherwise.
func (o *SeppInfo) GetRemotePlmnList() []PlmnId {
	if o == nil || isNil(o.RemotePlmnList) {
		var ret []PlmnId
		return ret
	}
	return o.RemotePlmnList
}

// GetRemotePlmnListOk returns a tuple with the RemotePlmnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppInfo) GetRemotePlmnListOk() ([]PlmnId, bool) {
	if o == nil || isNil(o.RemotePlmnList) {
		return nil, false
	}
	return o.RemotePlmnList, true
}

// HasRemotePlmnList returns a boolean if a field has been set.
func (o *SeppInfo) HasRemotePlmnList() bool {
	if o != nil && !isNil(o.RemotePlmnList) {
		return true
	}

	return false
}

// SetRemotePlmnList gets a reference to the given []PlmnId and assigns it to the RemotePlmnList field.
func (o *SeppInfo) SetRemotePlmnList(v []PlmnId) {
	o.RemotePlmnList = v
}

// GetRemoteSnpnList returns the RemoteSnpnList field value if set, zero value otherwise.
func (o *SeppInfo) GetRemoteSnpnList() []PlmnIdNid {
	if o == nil || isNil(o.RemoteSnpnList) {
		var ret []PlmnIdNid
		return ret
	}
	return o.RemoteSnpnList
}

// GetRemoteSnpnListOk returns a tuple with the RemoteSnpnList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppInfo) GetRemoteSnpnListOk() ([]PlmnIdNid, bool) {
	if o == nil || isNil(o.RemoteSnpnList) {
		return nil, false
	}
	return o.RemoteSnpnList, true
}

// HasRemoteSnpnList returns a boolean if a field has been set.
func (o *SeppInfo) HasRemoteSnpnList() bool {
	if o != nil && !isNil(o.RemoteSnpnList) {
		return true
	}

	return false
}

// SetRemoteSnpnList gets a reference to the given []PlmnIdNid and assigns it to the RemoteSnpnList field.
func (o *SeppInfo) SetRemoteSnpnList(v []PlmnIdNid) {
	o.RemoteSnpnList = v
}

// GetN32Purposes returns the N32Purposes field value if set, zero value otherwise.
func (o *SeppInfo) GetN32Purposes() []N32Purpose {
	if o == nil || isNil(o.N32Purposes) {
		var ret []N32Purpose
		return ret
	}
	return o.N32Purposes
}

// GetN32PurposesOk returns a tuple with the N32Purposes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SeppInfo) GetN32PurposesOk() ([]N32Purpose, bool) {
	if o == nil || isNil(o.N32Purposes) {
		return nil, false
	}
	return o.N32Purposes, true
}

// HasN32Purposes returns a boolean if a field has been set.
func (o *SeppInfo) HasN32Purposes() bool {
	if o != nil && !isNil(o.N32Purposes) {
		return true
	}

	return false
}

// SetN32Purposes gets a reference to the given []N32Purpose and assigns it to the N32Purposes field.
func (o *SeppInfo) SetN32Purposes(v []N32Purpose) {
	o.N32Purposes = v
}

func (o SeppInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SeppInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.SeppPrefix) {
		toSerialize["seppPrefix"] = o.SeppPrefix
	}
	if !isNil(o.SeppPorts) {
		toSerialize["seppPorts"] = o.SeppPorts
	}
	if !isNil(o.RemotePlmnList) {
		toSerialize["remotePlmnList"] = o.RemotePlmnList
	}
	if !isNil(o.RemoteSnpnList) {
		toSerialize["remoteSnpnList"] = o.RemoteSnpnList
	}
	if !isNil(o.N32Purposes) {
		toSerialize["n32Purposes"] = o.N32Purposes
	}
	return toSerialize, nil
}

type NullableSeppInfo struct {
	value *SeppInfo
	isSet bool
}

func (v NullableSeppInfo) Get() *SeppInfo {
	return v.value
}

func (v *NullableSeppInfo) Set(val *SeppInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSeppInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSeppInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSeppInfo(val *SeppInfo) *NullableSeppInfo {
	return &NullableSeppInfo{value: val, isSet: true}
}

func (v NullableSeppInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSeppInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


