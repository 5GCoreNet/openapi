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

// checks if the SnssaiTsctsfInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiTsctsfInfoItem{}

// SnssaiTsctsfInfoItem Set of parameters supported by TSCTSF for a given S-NSSAI
type SnssaiTsctsfInfoItem struct {
	SNssai ExtSnssai `json:"sNssai"`
	DnnInfoList []DnnTsctsfInfoItem `json:"dnnInfoList"`
}

// NewSnssaiTsctsfInfoItem instantiates a new SnssaiTsctsfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiTsctsfInfoItem(sNssai ExtSnssai, dnnInfoList []DnnTsctsfInfoItem) *SnssaiTsctsfInfoItem {
	this := SnssaiTsctsfInfoItem{}
	this.SNssai = sNssai
	this.DnnInfoList = dnnInfoList
	return &this
}

// NewSnssaiTsctsfInfoItemWithDefaults instantiates a new SnssaiTsctsfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiTsctsfInfoItemWithDefaults() *SnssaiTsctsfInfoItem {
	this := SnssaiTsctsfInfoItem{}
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiTsctsfInfoItem) GetSNssai() ExtSnssai {
	if o == nil {
		var ret ExtSnssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiTsctsfInfoItem) GetSNssaiOk() (*ExtSnssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiTsctsfInfoItem) SetSNssai(v ExtSnssai) {
	o.SNssai = v
}

// GetDnnInfoList returns the DnnInfoList field value
func (o *SnssaiTsctsfInfoItem) GetDnnInfoList() []DnnTsctsfInfoItem {
	if o == nil {
		var ret []DnnTsctsfInfoItem
		return ret
	}

	return o.DnnInfoList
}

// GetDnnInfoListOk returns a tuple with the DnnInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiTsctsfInfoItem) GetDnnInfoListOk() ([]DnnTsctsfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnInfoList, true
}

// SetDnnInfoList sets field value
func (o *SnssaiTsctsfInfoItem) SetDnnInfoList(v []DnnTsctsfInfoItem) {
	o.DnnInfoList = v
}

func (o SnssaiTsctsfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiTsctsfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNssai"] = o.SNssai
	toSerialize["dnnInfoList"] = o.DnnInfoList
	return toSerialize, nil
}

type NullableSnssaiTsctsfInfoItem struct {
	value *SnssaiTsctsfInfoItem
	isSet bool
}

func (v NullableSnssaiTsctsfInfoItem) Get() *SnssaiTsctsfInfoItem {
	return v.value
}

func (v *NullableSnssaiTsctsfInfoItem) Set(val *SnssaiTsctsfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiTsctsfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiTsctsfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiTsctsfInfoItem(val *SnssaiTsctsfInfoItem) *NullableSnssaiTsctsfInfoItem {
	return &NullableSnssaiTsctsfInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiTsctsfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiTsctsfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


