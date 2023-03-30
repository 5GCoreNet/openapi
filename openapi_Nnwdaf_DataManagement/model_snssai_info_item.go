/*
Nnwdaf_DataManagement

Nnwdaf_DataManagement API Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_DataManagement

import (
	"encoding/json"
)

// checks if the SnssaiInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiInfoItem{}

// SnssaiInfoItem Parameters supported by an NF for a given S-NSSAI Set of parameters supported by NF for a given S-NSSAI 
type SnssaiInfoItem struct {
	SNssai ExtSnssai `json:"sNssai"`
	DnnInfoList []DnnInfoItem `json:"dnnInfoList"`
}

// NewSnssaiInfoItem instantiates a new SnssaiInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiInfoItem(sNssai ExtSnssai, dnnInfoList []DnnInfoItem) *SnssaiInfoItem {
	this := SnssaiInfoItem{}
	this.SNssai = sNssai
	this.DnnInfoList = dnnInfoList
	return &this
}

// NewSnssaiInfoItemWithDefaults instantiates a new SnssaiInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiInfoItemWithDefaults() *SnssaiInfoItem {
	this := SnssaiInfoItem{}
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiInfoItem) GetSNssai() ExtSnssai {
	if o == nil {
		var ret ExtSnssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiInfoItem) GetSNssaiOk() (*ExtSnssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiInfoItem) SetSNssai(v ExtSnssai) {
	o.SNssai = v
}

// GetDnnInfoList returns the DnnInfoList field value
func (o *SnssaiInfoItem) GetDnnInfoList() []DnnInfoItem {
	if o == nil {
		var ret []DnnInfoItem
		return ret
	}

	return o.DnnInfoList
}

// GetDnnInfoListOk returns a tuple with the DnnInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiInfoItem) GetDnnInfoListOk() ([]DnnInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnInfoList, true
}

// SetDnnInfoList sets field value
func (o *SnssaiInfoItem) SetDnnInfoList(v []DnnInfoItem) {
	o.DnnInfoList = v
}

func (o SnssaiInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNssai"] = o.SNssai
	toSerialize["dnnInfoList"] = o.DnnInfoList
	return toSerialize, nil
}

type NullableSnssaiInfoItem struct {
	value *SnssaiInfoItem
	isSet bool
}

func (v NullableSnssaiInfoItem) Get() *SnssaiInfoItem {
	return v.value
}

func (v *NullableSnssaiInfoItem) Set(val *SnssaiInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiInfoItem(val *SnssaiInfoItem) *NullableSnssaiInfoItem {
	return &NullableSnssaiInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


