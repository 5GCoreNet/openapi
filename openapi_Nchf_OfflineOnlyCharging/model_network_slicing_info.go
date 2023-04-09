/*
Nchf_OfflineOnlyCharging

OfflineOnlyCharging Service © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_OfflineOnlyCharging

import (
	"encoding/json"
)

// checks if the NetworkSlicingInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkSlicingInfo{}

// NetworkSlicingInfo struct for NetworkSlicingInfo
type NetworkSlicingInfo struct {
	SNSSAI Snssai `json:"sNSSAI"`
}

// NewNetworkSlicingInfo instantiates a new NetworkSlicingInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkSlicingInfo(sNSSAI Snssai) *NetworkSlicingInfo {
	this := NetworkSlicingInfo{}
	this.SNSSAI = sNSSAI
	return &this
}

// NewNetworkSlicingInfoWithDefaults instantiates a new NetworkSlicingInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkSlicingInfoWithDefaults() *NetworkSlicingInfo {
	this := NetworkSlicingInfo{}
	return &this
}

// GetSNSSAI returns the SNSSAI field value
func (o *NetworkSlicingInfo) GetSNSSAI() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.SNSSAI
}

// GetSNSSAIOk returns a tuple with the SNSSAI field value
// and a boolean to check if the value has been set.
func (o *NetworkSlicingInfo) GetSNSSAIOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNSSAI, true
}

// SetSNSSAI sets field value
func (o *NetworkSlicingInfo) SetSNSSAI(v Snssai) {
	o.SNSSAI = v
}

func (o NetworkSlicingInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkSlicingInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNSSAI"] = o.SNSSAI
	return toSerialize, nil
}

type NullableNetworkSlicingInfo struct {
	value *NetworkSlicingInfo
	isSet bool
}

func (v NullableNetworkSlicingInfo) Get() *NetworkSlicingInfo {
	return v.value
}

func (v *NullableNetworkSlicingInfo) Set(val *NetworkSlicingInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNetworkSlicingInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNetworkSlicingInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNetworkSlicingInfo(val *NetworkSlicingInfo) *NullableNetworkSlicingInfo {
	return &NullableNetworkSlicingInfo{value: val, isSet: true}
}

func (v NullableNetworkSlicingInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNetworkSlicingInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
