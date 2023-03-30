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

// checks if the MbSmfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbSmfInfo{}

// MbSmfInfo Information of an MB-SMF NF Instance
type MbSmfInfo struct {
	// A map (list of key-value pairs) where a valid JSON string serves as key
	SNssaiInfoList *map[string]SnssaiMbSmfInfoItem `json:"sNssaiInfoList,omitempty"`
	// A map (list of key-value pairs) where a valid JSON string serves as key
	TmgiRangeList *map[string]TmgiRange `json:"tmgiRangeList,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
	TaiRangeList []TaiRange `json:"taiRangeList,omitempty"`
	// A map (list of key-value pairs) where a valid JSON string serves as key
	MbsSessionList *map[string]MbsSession `json:"mbsSessionList,omitempty"`
}

// NewMbSmfInfo instantiates a new MbSmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbSmfInfo() *MbSmfInfo {
	this := MbSmfInfo{}
	return &this
}

// NewMbSmfInfoWithDefaults instantiates a new MbSmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbSmfInfoWithDefaults() *MbSmfInfo {
	this := MbSmfInfo{}
	return &this
}

// GetSNssaiInfoList returns the SNssaiInfoList field value if set, zero value otherwise.
func (o *MbSmfInfo) GetSNssaiInfoList() map[string]SnssaiMbSmfInfoItem {
	if o == nil || IsNil(o.SNssaiInfoList) {
		var ret map[string]SnssaiMbSmfInfoItem
		return ret
	}
	return *o.SNssaiInfoList
}

// GetSNssaiInfoListOk returns a tuple with the SNssaiInfoList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbSmfInfo) GetSNssaiInfoListOk() (*map[string]SnssaiMbSmfInfoItem, bool) {
	if o == nil || IsNil(o.SNssaiInfoList) {
		return nil, false
	}
	return o.SNssaiInfoList, true
}

// HasSNssaiInfoList returns a boolean if a field has been set.
func (o *MbSmfInfo) HasSNssaiInfoList() bool {
	if o != nil && !IsNil(o.SNssaiInfoList) {
		return true
	}

	return false
}

// SetSNssaiInfoList gets a reference to the given map[string]SnssaiMbSmfInfoItem and assigns it to the SNssaiInfoList field.
func (o *MbSmfInfo) SetSNssaiInfoList(v map[string]SnssaiMbSmfInfoItem) {
	o.SNssaiInfoList = &v
}

// GetTmgiRangeList returns the TmgiRangeList field value if set, zero value otherwise.
func (o *MbSmfInfo) GetTmgiRangeList() map[string]TmgiRange {
	if o == nil || IsNil(o.TmgiRangeList) {
		var ret map[string]TmgiRange
		return ret
	}
	return *o.TmgiRangeList
}

// GetTmgiRangeListOk returns a tuple with the TmgiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbSmfInfo) GetTmgiRangeListOk() (*map[string]TmgiRange, bool) {
	if o == nil || IsNil(o.TmgiRangeList) {
		return nil, false
	}
	return o.TmgiRangeList, true
}

// HasTmgiRangeList returns a boolean if a field has been set.
func (o *MbSmfInfo) HasTmgiRangeList() bool {
	if o != nil && !IsNil(o.TmgiRangeList) {
		return true
	}

	return false
}

// SetTmgiRangeList gets a reference to the given map[string]TmgiRange and assigns it to the TmgiRangeList field.
func (o *MbSmfInfo) SetTmgiRangeList(v map[string]TmgiRange) {
	o.TmgiRangeList = &v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *MbSmfInfo) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbSmfInfo) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *MbSmfInfo) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *MbSmfInfo) SetTaiList(v []Tai) {
	o.TaiList = v
}

// GetTaiRangeList returns the TaiRangeList field value if set, zero value otherwise.
func (o *MbSmfInfo) GetTaiRangeList() []TaiRange {
	if o == nil || IsNil(o.TaiRangeList) {
		var ret []TaiRange
		return ret
	}
	return o.TaiRangeList
}

// GetTaiRangeListOk returns a tuple with the TaiRangeList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbSmfInfo) GetTaiRangeListOk() ([]TaiRange, bool) {
	if o == nil || IsNil(o.TaiRangeList) {
		return nil, false
	}
	return o.TaiRangeList, true
}

// HasTaiRangeList returns a boolean if a field has been set.
func (o *MbSmfInfo) HasTaiRangeList() bool {
	if o != nil && !IsNil(o.TaiRangeList) {
		return true
	}

	return false
}

// SetTaiRangeList gets a reference to the given []TaiRange and assigns it to the TaiRangeList field.
func (o *MbSmfInfo) SetTaiRangeList(v []TaiRange) {
	o.TaiRangeList = v
}

// GetMbsSessionList returns the MbsSessionList field value if set, zero value otherwise.
func (o *MbSmfInfo) GetMbsSessionList() map[string]MbsSession {
	if o == nil || IsNil(o.MbsSessionList) {
		var ret map[string]MbsSession
		return ret
	}
	return *o.MbsSessionList
}

// GetMbsSessionListOk returns a tuple with the MbsSessionList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbSmfInfo) GetMbsSessionListOk() (*map[string]MbsSession, bool) {
	if o == nil || IsNil(o.MbsSessionList) {
		return nil, false
	}
	return o.MbsSessionList, true
}

// HasMbsSessionList returns a boolean if a field has been set.
func (o *MbSmfInfo) HasMbsSessionList() bool {
	if o != nil && !IsNil(o.MbsSessionList) {
		return true
	}

	return false
}

// SetMbsSessionList gets a reference to the given map[string]MbsSession and assigns it to the MbsSessionList field.
func (o *MbSmfInfo) SetMbsSessionList(v map[string]MbsSession) {
	o.MbsSessionList = &v
}

func (o MbSmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbSmfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SNssaiInfoList) {
		toSerialize["sNssaiInfoList"] = o.SNssaiInfoList
	}
	if !IsNil(o.TmgiRangeList) {
		toSerialize["tmgiRangeList"] = o.TmgiRangeList
	}
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	if !IsNil(o.TaiRangeList) {
		toSerialize["taiRangeList"] = o.TaiRangeList
	}
	if !IsNil(o.MbsSessionList) {
		toSerialize["mbsSessionList"] = o.MbsSessionList
	}
	return toSerialize, nil
}

type NullableMbSmfInfo struct {
	value *MbSmfInfo
	isSet bool
}

func (v NullableMbSmfInfo) Get() *MbSmfInfo {
	return v.value
}

func (v *NullableMbSmfInfo) Set(val *MbSmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMbSmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMbSmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbSmfInfo(val *MbSmfInfo) *NullableMbSmfInfo {
	return &NullableMbSmfInfo{value: val, isSet: true}
}

func (v NullableMbSmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbSmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


