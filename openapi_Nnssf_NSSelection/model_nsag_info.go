/*
NSSF NS Selection

NSSF Network Slice Selection Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssf_NSSelection

import (
	"encoding/json"
)

// checks if the NsagInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NsagInfo{}

// NsagInfo Contains the association of NSAGs and S-NSSAI(s) along with the TA(s) within which the association is valid.
type NsagInfo struct {
	NsagIds      []int32    `json:"nsagIds"`
	SnssaiList   []Snssai   `json:"snssaiList"`
	TaiList      []Tai      `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
}

// NewNsagInfo instantiates a new NsagInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNsagInfo(nsagIds []int32, snssaiList []Snssai) *NsagInfo {
	this := NsagInfo{}
	this.NsagIds = nsagIds
	this.SnssaiList = snssaiList
	return &this
}

// NewNsagInfoWithDefaults instantiates a new NsagInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNsagInfoWithDefaults() *NsagInfo {
	this := NsagInfo{}
	return &this
}

// GetNsagIds returns the NsagIds field value
func (o *NsagInfo) GetNsagIds() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.NsagIds
}

// GetNsagIdsOk returns a tuple with the NsagIds field value
// and a boolean to check if the value has been set.
func (o *NsagInfo) GetNsagIdsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.NsagIds, true
}

// SetNsagIds sets field value
func (o *NsagInfo) SetNsagIds(v []int32) {
	o.NsagIds = v
}

// GetSnssaiList returns the SnssaiList field value
func (o *NsagInfo) GetSnssaiList() []Snssai {
	if o == nil {
		var ret []Snssai
		return ret
	}

	return o.SnssaiList
}

// GetSnssaiListOk returns a tuple with the SnssaiList field value
// and a boolean to check if the value has been set.
func (o *NsagInfo) GetSnssaiListOk() ([]Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return o.SnssaiList, true
}

// SetSnssaiList sets field value
func (o *NsagInfo) SetSnssaiList(v []Snssai) {
	o.SnssaiList = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *NsagInfo) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsagInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *NsagInfo) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *NsagInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *NsagInfo) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NsagInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *NsagInfo) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *NsagInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

func (o NsagInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NsagInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nsagIds"] = o.NsagIds
	toSerialize["snssaiList"] = o.SnssaiList
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	return toSerialize, nil
}

type NullableNsagInfo struct {
	value *NsagInfo
	isSet bool
}

func (v NullableNsagInfo) Get() *NsagInfo {
	return v.value
}

func (v *NullableNsagInfo) Set(val *NsagInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNsagInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNsagInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNsagInfo(val *NsagInfo) *NullableNsagInfo {
	return &NullableNsagInfo{value: val, isSet: true}
}

func (v NullableNsagInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNsagInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
