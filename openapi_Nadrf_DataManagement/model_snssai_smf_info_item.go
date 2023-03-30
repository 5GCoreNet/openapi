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

// checks if the SnssaiSmfInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiSmfInfoItem{}

// SnssaiSmfInfoItem Set of parameters supported by SMF for a given S-NSSAI
type SnssaiSmfInfoItem struct {
	SNssai ExtSnssai `json:"sNssai"`
	DnnSmfInfoList []DnnSmfInfoItem `json:"dnnSmfInfoList"`
}

// NewSnssaiSmfInfoItem instantiates a new SnssaiSmfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiSmfInfoItem(sNssai ExtSnssai, dnnSmfInfoList []DnnSmfInfoItem) *SnssaiSmfInfoItem {
	this := SnssaiSmfInfoItem{}
	this.SNssai = sNssai
	this.DnnSmfInfoList = dnnSmfInfoList
	return &this
}

// NewSnssaiSmfInfoItemWithDefaults instantiates a new SnssaiSmfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiSmfInfoItemWithDefaults() *SnssaiSmfInfoItem {
	this := SnssaiSmfInfoItem{}
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiSmfInfoItem) GetSNssai() ExtSnssai {
	if o == nil {
		var ret ExtSnssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiSmfInfoItem) GetSNssaiOk() (*ExtSnssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiSmfInfoItem) SetSNssai(v ExtSnssai) {
	o.SNssai = v
}

// GetDnnSmfInfoList returns the DnnSmfInfoList field value
func (o *SnssaiSmfInfoItem) GetDnnSmfInfoList() []DnnSmfInfoItem {
	if o == nil {
		var ret []DnnSmfInfoItem
		return ret
	}

	return o.DnnSmfInfoList
}

// GetDnnSmfInfoListOk returns a tuple with the DnnSmfInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiSmfInfoItem) GetDnnSmfInfoListOk() ([]DnnSmfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnSmfInfoList, true
}

// SetDnnSmfInfoList sets field value
func (o *SnssaiSmfInfoItem) SetDnnSmfInfoList(v []DnnSmfInfoItem) {
	o.DnnSmfInfoList = v
}

func (o SnssaiSmfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiSmfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNssai"] = o.SNssai
	toSerialize["dnnSmfInfoList"] = o.DnnSmfInfoList
	return toSerialize, nil
}

type NullableSnssaiSmfInfoItem struct {
	value *SnssaiSmfInfoItem
	isSet bool
}

func (v NullableSnssaiSmfInfoItem) Get() *SnssaiSmfInfoItem {
	return v.value
}

func (v *NullableSnssaiSmfInfoItem) Set(val *SnssaiSmfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiSmfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiSmfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiSmfInfoItem(val *SnssaiSmfInfoItem) *NullableSnssaiSmfInfoItem {
	return &NullableSnssaiSmfInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiSmfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiSmfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


