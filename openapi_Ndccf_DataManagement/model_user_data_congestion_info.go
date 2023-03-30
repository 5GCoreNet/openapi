/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the UserDataCongestionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserDataCongestionInfo{}

// UserDataCongestionInfo Represents the user data congestion information.
type UserDataCongestionInfo struct {
	NetworkArea NetworkAreaInfo `json:"networkArea"`
	CongestionInfo CongestionInfo `json:"congestionInfo"`
	Snssai *Snssai `json:"snssai,omitempty"`
}

// NewUserDataCongestionInfo instantiates a new UserDataCongestionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserDataCongestionInfo(networkArea NetworkAreaInfo, congestionInfo CongestionInfo) *UserDataCongestionInfo {
	this := UserDataCongestionInfo{}
	this.NetworkArea = networkArea
	this.CongestionInfo = congestionInfo
	return &this
}

// NewUserDataCongestionInfoWithDefaults instantiates a new UserDataCongestionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserDataCongestionInfoWithDefaults() *UserDataCongestionInfo {
	this := UserDataCongestionInfo{}
	return &this
}

// GetNetworkArea returns the NetworkArea field value
func (o *UserDataCongestionInfo) GetNetworkArea() NetworkAreaInfo {
	if o == nil {
		var ret NetworkAreaInfo
		return ret
	}

	return o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value
// and a boolean to check if the value has been set.
func (o *UserDataCongestionInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NetworkArea, true
}

// SetNetworkArea sets field value
func (o *UserDataCongestionInfo) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = v
}

// GetCongestionInfo returns the CongestionInfo field value
func (o *UserDataCongestionInfo) GetCongestionInfo() CongestionInfo {
	if o == nil {
		var ret CongestionInfo
		return ret
	}

	return o.CongestionInfo
}

// GetCongestionInfoOk returns a tuple with the CongestionInfo field value
// and a boolean to check if the value has been set.
func (o *UserDataCongestionInfo) GetCongestionInfoOk() (*CongestionInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CongestionInfo, true
}

// SetCongestionInfo sets field value
func (o *UserDataCongestionInfo) SetCongestionInfo(v CongestionInfo) {
	o.CongestionInfo = v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *UserDataCongestionInfo) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserDataCongestionInfo) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *UserDataCongestionInfo) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *UserDataCongestionInfo) SetSnssai(v Snssai) {
	o.Snssai = &v
}

func (o UserDataCongestionInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserDataCongestionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["networkArea"] = o.NetworkArea
	toSerialize["congestionInfo"] = o.CongestionInfo
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	return toSerialize, nil
}

type NullableUserDataCongestionInfo struct {
	value *UserDataCongestionInfo
	isSet bool
}

func (v NullableUserDataCongestionInfo) Get() *UserDataCongestionInfo {
	return v.value
}

func (v *NullableUserDataCongestionInfo) Set(val *UserDataCongestionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUserDataCongestionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUserDataCongestionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserDataCongestionInfo(val *UserDataCongestionInfo) *NullableUserDataCongestionInfo {
	return &NullableUserDataCongestionInfo{value: val, isSet: true}
}

func (v NullableUserDataCongestionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserDataCongestionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


