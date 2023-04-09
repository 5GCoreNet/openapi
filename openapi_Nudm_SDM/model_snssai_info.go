/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the SnssaiInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiInfo{}

// SnssaiInfo struct for SnssaiInfo
type SnssaiInfo struct {
	DnnInfos []DnnInfo `json:"dnnInfos"`
}

// NewSnssaiInfo instantiates a new SnssaiInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiInfo(dnnInfos []DnnInfo) *SnssaiInfo {
	this := SnssaiInfo{}
	this.DnnInfos = dnnInfos
	return &this
}

// NewSnssaiInfoWithDefaults instantiates a new SnssaiInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiInfoWithDefaults() *SnssaiInfo {
	this := SnssaiInfo{}
	return &this
}

// GetDnnInfos returns the DnnInfos field value
func (o *SnssaiInfo) GetDnnInfos() []DnnInfo {
	if o == nil {
		var ret []DnnInfo
		return ret
	}

	return o.DnnInfos
}

// GetDnnInfosOk returns a tuple with the DnnInfos field value
// and a boolean to check if the value has been set.
func (o *SnssaiInfo) GetDnnInfosOk() ([]DnnInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnInfos, true
}

// SetDnnInfos sets field value
func (o *SnssaiInfo) SetDnnInfos(v []DnnInfo) {
	o.DnnInfos = v
}

func (o SnssaiInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["dnnInfos"] = o.DnnInfos
	return toSerialize, nil
}

type NullableSnssaiInfo struct {
	value *SnssaiInfo
	isSet bool
}

func (v NullableSnssaiInfo) Get() *SnssaiInfo {
	return v.value
}

func (v *NullableSnssaiInfo) Set(val *SnssaiInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiInfo(val *SnssaiInfo) *NullableSnssaiInfo {
	return &NullableSnssaiInfo{value: val, isSet: true}
}

func (v NullableSnssaiInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
