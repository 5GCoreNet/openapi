/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
)

// checks if the SnssaiEasdfInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiEasdfInfoItem{}

// SnssaiEasdfInfoItem Set of parameters supported by EASDF for a given S-NSSAI
type SnssaiEasdfInfoItem struct {
	SNssai           ExtSnssai          `json:"sNssai"`
	DnnEasdfInfoList []DnnEasdfInfoItem `json:"dnnEasdfInfoList"`
}

// NewSnssaiEasdfInfoItem instantiates a new SnssaiEasdfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiEasdfInfoItem(sNssai ExtSnssai, dnnEasdfInfoList []DnnEasdfInfoItem) *SnssaiEasdfInfoItem {
	this := SnssaiEasdfInfoItem{}
	this.SNssai = sNssai
	this.DnnEasdfInfoList = dnnEasdfInfoList
	return &this
}

// NewSnssaiEasdfInfoItemWithDefaults instantiates a new SnssaiEasdfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiEasdfInfoItemWithDefaults() *SnssaiEasdfInfoItem {
	this := SnssaiEasdfInfoItem{}
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiEasdfInfoItem) GetSNssai() ExtSnssai {
	if o == nil {
		var ret ExtSnssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiEasdfInfoItem) GetSNssaiOk() (*ExtSnssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiEasdfInfoItem) SetSNssai(v ExtSnssai) {
	o.SNssai = v
}

// GetDnnEasdfInfoList returns the DnnEasdfInfoList field value
func (o *SnssaiEasdfInfoItem) GetDnnEasdfInfoList() []DnnEasdfInfoItem {
	if o == nil {
		var ret []DnnEasdfInfoItem
		return ret
	}

	return o.DnnEasdfInfoList
}

// GetDnnEasdfInfoListOk returns a tuple with the DnnEasdfInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiEasdfInfoItem) GetDnnEasdfInfoListOk() ([]DnnEasdfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnEasdfInfoList, true
}

// SetDnnEasdfInfoList sets field value
func (o *SnssaiEasdfInfoItem) SetDnnEasdfInfoList(v []DnnEasdfInfoItem) {
	o.DnnEasdfInfoList = v
}

func (o SnssaiEasdfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiEasdfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNssai"] = o.SNssai
	toSerialize["dnnEasdfInfoList"] = o.DnnEasdfInfoList
	return toSerialize, nil
}

type NullableSnssaiEasdfInfoItem struct {
	value *SnssaiEasdfInfoItem
	isSet bool
}

func (v NullableSnssaiEasdfInfoItem) Get() *SnssaiEasdfInfoItem {
	return v.value
}

func (v *NullableSnssaiEasdfInfoItem) Set(val *SnssaiEasdfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiEasdfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiEasdfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiEasdfInfoItem(val *SnssaiEasdfInfoItem) *NullableSnssaiEasdfInfoItem {
	return &NullableSnssaiEasdfInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiEasdfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiEasdfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
