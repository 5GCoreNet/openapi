/*
Nmfaf_3caDataManagement

MFAF 3GPP Consumer Adaptor (3CA) Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nmfaf_3caDataManagement

import (
	"encoding/json"
)

// checks if the WlanPerformanceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WlanPerformanceInfo{}

// WlanPerformanceInfo The WLAN performance related information.
type WlanPerformanceInfo struct {
	NetworkArea      *NetworkAreaInfo             `json:"networkArea,omitempty"`
	WlanPerSsidInfos []WlanPerSsIdPerformanceInfo `json:"wlanPerSsidInfos"`
}

// NewWlanPerformanceInfo instantiates a new WlanPerformanceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWlanPerformanceInfo(wlanPerSsidInfos []WlanPerSsIdPerformanceInfo) *WlanPerformanceInfo {
	this := WlanPerformanceInfo{}
	this.WlanPerSsidInfos = wlanPerSsidInfos
	return &this
}

// NewWlanPerformanceInfoWithDefaults instantiates a new WlanPerformanceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWlanPerformanceInfoWithDefaults() *WlanPerformanceInfo {
	this := WlanPerformanceInfo{}
	return &this
}

// GetNetworkArea returns the NetworkArea field value if set, zero value otherwise.
func (o *WlanPerformanceInfo) GetNetworkArea() NetworkAreaInfo {
	if o == nil || IsNil(o.NetworkArea) {
		var ret NetworkAreaInfo
		return ret
	}
	return *o.NetworkArea
}

// GetNetworkAreaOk returns a tuple with the NetworkArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WlanPerformanceInfo) GetNetworkAreaOk() (*NetworkAreaInfo, bool) {
	if o == nil || IsNil(o.NetworkArea) {
		return nil, false
	}
	return o.NetworkArea, true
}

// HasNetworkArea returns a boolean if a field has been set.
func (o *WlanPerformanceInfo) HasNetworkArea() bool {
	if o != nil && !IsNil(o.NetworkArea) {
		return true
	}

	return false
}

// SetNetworkArea gets a reference to the given NetworkAreaInfo and assigns it to the NetworkArea field.
func (o *WlanPerformanceInfo) SetNetworkArea(v NetworkAreaInfo) {
	o.NetworkArea = &v
}

// GetWlanPerSsidInfos returns the WlanPerSsidInfos field value
func (o *WlanPerformanceInfo) GetWlanPerSsidInfos() []WlanPerSsIdPerformanceInfo {
	if o == nil {
		var ret []WlanPerSsIdPerformanceInfo
		return ret
	}

	return o.WlanPerSsidInfos
}

// GetWlanPerSsidInfosOk returns a tuple with the WlanPerSsidInfos field value
// and a boolean to check if the value has been set.
func (o *WlanPerformanceInfo) GetWlanPerSsidInfosOk() ([]WlanPerSsIdPerformanceInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.WlanPerSsidInfos, true
}

// SetWlanPerSsidInfos sets field value
func (o *WlanPerformanceInfo) SetWlanPerSsidInfos(v []WlanPerSsIdPerformanceInfo) {
	o.WlanPerSsidInfos = v
}

func (o WlanPerformanceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WlanPerformanceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NetworkArea) {
		toSerialize["networkArea"] = o.NetworkArea
	}
	toSerialize["wlanPerSsidInfos"] = o.WlanPerSsidInfos
	return toSerialize, nil
}

type NullableWlanPerformanceInfo struct {
	value *WlanPerformanceInfo
	isSet bool
}

func (v NullableWlanPerformanceInfo) Get() *WlanPerformanceInfo {
	return v.value
}

func (v *NullableWlanPerformanceInfo) Set(val *WlanPerformanceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableWlanPerformanceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableWlanPerformanceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWlanPerformanceInfo(val *WlanPerformanceInfo) *NullableWlanPerformanceInfo {
	return &NullableWlanPerformanceInfo{value: val, isSet: true}
}

func (v NullableWlanPerformanceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWlanPerformanceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
