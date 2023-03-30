/*
Nnwdaf_AnalyticsInfo

Nnwdaf_AnalyticsInfo Service API.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnwdaf_AnalyticsInfo

import (
	"encoding/json"
)

// checks if the DccfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DccfInfo{}

// DccfInfo Information of a DCCF NF Instance
type DccfInfo struct {
	ServingNfTypeList []NFType `json:"servingNfTypeList,omitempty"`
	ServingNfSetIdList []string `json:"servingNfSetIdList,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
}

// NewDccfInfo instantiates a new DccfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDccfInfo() *DccfInfo {
	this := DccfInfo{}
	return &this
}

// NewDccfInfoWithDefaults instantiates a new DccfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDccfInfoWithDefaults() *DccfInfo {
	this := DccfInfo{}
	return &this
}

// GetServingNfTypeList returns the ServingNfTypeList field value if set, zero value otherwise.
func (o *DccfInfo) GetServingNfTypeList() []NFType {
	if o == nil || IsNil(o.ServingNfTypeList) {
		var ret []NFType
		return ret
	}
	return o.ServingNfTypeList
}

// GetServingNfTypeListOk returns a tuple with the ServingNfTypeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DccfInfo) GetServingNfTypeListOk() ([]NFType, bool) {
	if o == nil || IsNil(o.ServingNfTypeList) {
		return nil, false
	}
	return o.ServingNfTypeList, true
}

// HasServingNfTypeList returns a boolean if a field has been set.
func (o *DccfInfo) HasServingNfTypeList() bool {
	if o != nil && !IsNil(o.ServingNfTypeList) {
		return true
	}

	return false
}

// SetServingNfTypeList gets a reference to the given []NFType and assigns it to the ServingNfTypeList field.
func (o *DccfInfo) SetServingNfTypeList(v []NFType) {
	o.ServingNfTypeList = v
}

// GetServingNfSetIdList returns the ServingNfSetIdList field value if set, zero value otherwise.
func (o *DccfInfo) GetServingNfSetIdList() []string {
	if o == nil || IsNil(o.ServingNfSetIdList) {
		var ret []string
		return ret
	}
	return o.ServingNfSetIdList
}

// GetServingNfSetIdListOk returns a tuple with the ServingNfSetIdList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DccfInfo) GetServingNfSetIdListOk() ([]string, bool) {
	if o == nil || IsNil(o.ServingNfSetIdList) {
		return nil, false
	}
	return o.ServingNfSetIdList, true
}

// HasServingNfSetIdList returns a boolean if a field has been set.
func (o *DccfInfo) HasServingNfSetIdList() bool {
	if o != nil && !IsNil(o.ServingNfSetIdList) {
		return true
	}

	return false
}

// SetServingNfSetIdList gets a reference to the given []string and assigns it to the ServingNfSetIdList field.
func (o *DccfInfo) SetServingNfSetIdList(v []string) {
	o.ServingNfSetIdList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *DccfInfo) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DccfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *DccfInfo) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *DccfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *DccfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DccfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *DccfInfo) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *DccfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

func (o DccfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DccfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ServingNfTypeList) {
		toSerialize["servingNfTypeList"] = o.ServingNfTypeList
	}
	if !IsNil(o.ServingNfSetIdList) {
		toSerialize["servingNfSetIdList"] = o.ServingNfSetIdList
	}
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	return toSerialize, nil
}

type NullableDccfInfo struct {
	value *DccfInfo
	isSet bool
}

func (v NullableDccfInfo) Get() *DccfInfo {
	return v.value
}

func (v *NullableDccfInfo) Set(val *DccfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDccfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDccfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDccfInfo(val *DccfInfo) *NullableDccfInfo {
	return &NullableDccfInfo{value: val, isSet: true}
}

func (v NullableDccfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDccfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

