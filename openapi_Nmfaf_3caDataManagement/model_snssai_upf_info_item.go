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

// checks if the SnssaiUpfInfoItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiUpfInfoItem{}

// SnssaiUpfInfoItem Set of parameters supported by UPF for a given S-NSSAI
type SnssaiUpfInfoItem struct {
	SNssai             ExtSnssai        `json:"sNssai"`
	DnnUpfInfoList     []DnnUpfInfoItem `json:"dnnUpfInfoList"`
	RedundantTransport *bool            `json:"redundantTransport,omitempty"`
}

// NewSnssaiUpfInfoItem instantiates a new SnssaiUpfInfoItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiUpfInfoItem(sNssai ExtSnssai, dnnUpfInfoList []DnnUpfInfoItem) *SnssaiUpfInfoItem {
	this := SnssaiUpfInfoItem{}
	this.SNssai = sNssai
	this.DnnUpfInfoList = dnnUpfInfoList
	var redundantTransport bool = false
	this.RedundantTransport = &redundantTransport
	return &this
}

// NewSnssaiUpfInfoItemWithDefaults instantiates a new SnssaiUpfInfoItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiUpfInfoItemWithDefaults() *SnssaiUpfInfoItem {
	this := SnssaiUpfInfoItem{}
	var redundantTransport bool = false
	this.RedundantTransport = &redundantTransport
	return &this
}

// GetSNssai returns the SNssai field value
func (o *SnssaiUpfInfoItem) GetSNssai() ExtSnssai {
	if o == nil {
		var ret ExtSnssai
		return ret
	}

	return o.SNssai
}

// GetSNssaiOk returns a tuple with the SNssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiUpfInfoItem) GetSNssaiOk() (*ExtSnssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SNssai, true
}

// SetSNssai sets field value
func (o *SnssaiUpfInfoItem) SetSNssai(v ExtSnssai) {
	o.SNssai = v
}

// GetDnnUpfInfoList returns the DnnUpfInfoList field value
func (o *SnssaiUpfInfoItem) GetDnnUpfInfoList() []DnnUpfInfoItem {
	if o == nil {
		var ret []DnnUpfInfoItem
		return ret
	}

	return o.DnnUpfInfoList
}

// GetDnnUpfInfoListOk returns a tuple with the DnnUpfInfoList field value
// and a boolean to check if the value has been set.
func (o *SnssaiUpfInfoItem) GetDnnUpfInfoListOk() ([]DnnUpfInfoItem, bool) {
	if o == nil {
		return nil, false
	}
	return o.DnnUpfInfoList, true
}

// SetDnnUpfInfoList sets field value
func (o *SnssaiUpfInfoItem) SetDnnUpfInfoList(v []DnnUpfInfoItem) {
	o.DnnUpfInfoList = v
}

// GetRedundantTransport returns the RedundantTransport field value if set, zero value otherwise.
func (o *SnssaiUpfInfoItem) GetRedundantTransport() bool {
	if o == nil || IsNil(o.RedundantTransport) {
		var ret bool
		return ret
	}
	return *o.RedundantTransport
}

// GetRedundantTransportOk returns a tuple with the RedundantTransport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiUpfInfoItem) GetRedundantTransportOk() (*bool, bool) {
	if o == nil || IsNil(o.RedundantTransport) {
		return nil, false
	}
	return o.RedundantTransport, true
}

// HasRedundantTransport returns a boolean if a field has been set.
func (o *SnssaiUpfInfoItem) HasRedundantTransport() bool {
	if o != nil && !IsNil(o.RedundantTransport) {
		return true
	}

	return false
}

// SetRedundantTransport gets a reference to the given bool and assigns it to the RedundantTransport field.
func (o *SnssaiUpfInfoItem) SetRedundantTransport(v bool) {
	o.RedundantTransport = &v
}

func (o SnssaiUpfInfoItem) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiUpfInfoItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["sNssai"] = o.SNssai
	toSerialize["dnnUpfInfoList"] = o.DnnUpfInfoList
	if !IsNil(o.RedundantTransport) {
		toSerialize["redundantTransport"] = o.RedundantTransport
	}
	return toSerialize, nil
}

type NullableSnssaiUpfInfoItem struct {
	value *SnssaiUpfInfoItem
	isSet bool
}

func (v NullableSnssaiUpfInfoItem) Get() *SnssaiUpfInfoItem {
	return v.value
}

func (v *NullableSnssaiUpfInfoItem) Set(val *SnssaiUpfInfoItem) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiUpfInfoItem) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiUpfInfoItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiUpfInfoItem(val *SnssaiUpfInfoItem) *NullableSnssaiUpfInfoItem {
	return &NullableSnssaiUpfInfoItem{value: val, isSet: true}
}

func (v NullableSnssaiUpfInfoItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiUpfInfoItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
