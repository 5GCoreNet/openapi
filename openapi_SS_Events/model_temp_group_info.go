/*
SS_Events

API for SEAL Events management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_SS_Events

import (
	"encoding/json"
)

// checks if the TempGroupInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TempGroupInfo{}

// TempGroupInfo Represents the created temporary VAL group information.
type TempGroupInfo struct {
	ValGrpIds []string `json:"valGrpIds"`
	TempValGrpId string `json:"tempValGrpId"`
	ValServIds []string `json:"valServIds,omitempty"`
}

// NewTempGroupInfo instantiates a new TempGroupInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTempGroupInfo(valGrpIds []string, tempValGrpId string) *TempGroupInfo {
	this := TempGroupInfo{}
	this.ValGrpIds = valGrpIds
	this.TempValGrpId = tempValGrpId
	return &this
}

// NewTempGroupInfoWithDefaults instantiates a new TempGroupInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTempGroupInfoWithDefaults() *TempGroupInfo {
	this := TempGroupInfo{}
	return &this
}

// GetValGrpIds returns the ValGrpIds field value
func (o *TempGroupInfo) GetValGrpIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.ValGrpIds
}

// GetValGrpIdsOk returns a tuple with the ValGrpIds field value
// and a boolean to check if the value has been set.
func (o *TempGroupInfo) GetValGrpIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ValGrpIds, true
}

// SetValGrpIds sets field value
func (o *TempGroupInfo) SetValGrpIds(v []string) {
	o.ValGrpIds = v
}

// GetTempValGrpId returns the TempValGrpId field value
func (o *TempGroupInfo) GetTempValGrpId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TempValGrpId
}

// GetTempValGrpIdOk returns a tuple with the TempValGrpId field value
// and a boolean to check if the value has been set.
func (o *TempGroupInfo) GetTempValGrpIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TempValGrpId, true
}

// SetTempValGrpId sets field value
func (o *TempGroupInfo) SetTempValGrpId(v string) {
	o.TempValGrpId = v
}

// GetValServIds returns the ValServIds field value if set, zero value otherwise.
func (o *TempGroupInfo) GetValServIds() []string {
	if o == nil || IsNil(o.ValServIds) {
		var ret []string
		return ret
	}
	return o.ValServIds
}

// GetValServIdsOk returns a tuple with the ValServIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TempGroupInfo) GetValServIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.ValServIds) {
		return nil, false
	}
	return o.ValServIds, true
}

// HasValServIds returns a boolean if a field has been set.
func (o *TempGroupInfo) HasValServIds() bool {
	if o != nil && !IsNil(o.ValServIds) {
		return true
	}

	return false
}

// SetValServIds gets a reference to the given []string and assigns it to the ValServIds field.
func (o *TempGroupInfo) SetValServIds(v []string) {
	o.ValServIds = v
}

func (o TempGroupInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TempGroupInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["valGrpIds"] = o.ValGrpIds
	toSerialize["tempValGrpId"] = o.TempValGrpId
	if !IsNil(o.ValServIds) {
		toSerialize["valServIds"] = o.ValServIds
	}
	return toSerialize, nil
}

type NullableTempGroupInfo struct {
	value *TempGroupInfo
	isSet bool
}

func (v NullableTempGroupInfo) Get() *TempGroupInfo {
	return v.value
}

func (v *NullableTempGroupInfo) Set(val *TempGroupInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableTempGroupInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableTempGroupInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTempGroupInfo(val *TempGroupInfo) *NullableTempGroupInfo {
	return &NullableTempGroupInfo{value: val, isSet: true}
}

func (v NullableTempGroupInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTempGroupInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


