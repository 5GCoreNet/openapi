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

// checks if the NsacfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsacfInfo{}

// NsacfInfo Information of a NSACF NF Instance
type NsacfInfo struct {
	NsacfCapability NsacfCapability `json:"nsacfCapability"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
}

// NewNsacfInfo instantiates a new NsacfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsacfInfo(nsacfCapability NsacfCapability) *NsacfInfo {
	this := NsacfInfo{}
	this.NsacfCapability = nsacfCapability
	return &this
}

// NewNsacfInfoWithDefaults instantiates a new NsacfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsacfInfoWithDefaults() *NsacfInfo {
	this := NsacfInfo{}
	return &this
}

// GetNsacfCapability returns the NsacfCapability field value
func (o *NsacfInfo) GetNsacfCapability() NsacfCapability {
	if o == nil {
		var ret NsacfCapability
		return ret
	}

	return o.NsacfCapability
}

// GetNsacfCapabilityOk returns a tuple with the NsacfCapability field value
// and a boolean to check if the value has been set.
func (o *NsacfInfo) GetNsacfCapabilityOk() (*NsacfCapability, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NsacfCapability, true
}

// SetNsacfCapability sets field value
func (o *NsacfInfo) SetNsacfCapability(v NsacfCapability) {
	o.NsacfCapability = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *NsacfInfo) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *NsacfInfo) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *NsacfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NsacfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsacfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NsacfInfo) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NsacfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

func (o NsacfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsacfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nsacfCapability"] = o.NsacfCapability
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	return toSerialize, nil
}

type NullableNsacfInfo struct {
	value *NsacfInfo
	isSet bool
}

func (v NullableNsacfInfo) Get() *NsacfInfo {
	return v.value
}

func (v *NullableNsacfInfo) Set(val *NsacfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNsacfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNsacfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsacfInfo(val *NsacfInfo) *NullableNsacfInfo {
	return &NullableNsacfInfo{value: val, isSet: true}
}

func (v NullableNsacfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsacfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


