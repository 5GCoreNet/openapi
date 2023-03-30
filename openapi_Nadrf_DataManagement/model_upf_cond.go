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

// checks if the UpfCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpfCond{}

// UpfCond Subscription to a set of NF Instances (UPFs), able to serve a certain service area (i.e. SMF serving area or TAI list) 
type UpfCond struct {
	ConditionType string `json:"conditionType"`
	SmfServingArea []string `json:"smfServingArea,omitempty"`
	TaiList []Tai `json:"taiList,omitempty"`
}

// NewUpfCond instantiates a new UpfCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpfCond(conditionType string) *UpfCond {
	this := UpfCond{}
	this.ConditionType = conditionType
	return &this
}

// NewUpfCondWithDefaults instantiates a new UpfCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpfCondWithDefaults() *UpfCond {
	this := UpfCond{}
	return &this
}

// GetConditionType returns the ConditionType field value
func (o *UpfCond) GetConditionType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ConditionType
}

// GetConditionTypeOk returns a tuple with the ConditionType field value
// and a boolean to check if the value has been set.
func (o *UpfCond) GetConditionTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ConditionType, true
}

// SetConditionType sets field value
func (o *UpfCond) SetConditionType(v string) {
	o.ConditionType = v
}

// GetSmfServingArea returns the SmfServingArea field value if set, zero value otherwise.
func (o *UpfCond) GetSmfServingArea() []string {
	if o == nil || IsNil(o.SmfServingArea) {
		var ret []string
		return ret
	}
	return o.SmfServingArea
}

// GetSmfServingAreaOk returns a tuple with the SmfServingArea field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfCond) GetSmfServingAreaOk() ([]string, bool) {
	if o == nil || IsNil(o.SmfServingArea) {
		return nil, false
	}
	return o.SmfServingArea, true
}

// HasSmfServingArea returns a boolean if a field has been set.
func (o *UpfCond) HasSmfServingArea() bool {
	if o != nil && !IsNil(o.SmfServingArea) {
		return true
	}

	return false
}

// SetSmfServingArea gets a reference to the given []string and assigns it to the SmfServingArea field.
func (o *UpfCond) SetSmfServingArea(v []string) {
	o.SmfServingArea = v
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *UpfCond) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpfCond) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *UpfCond) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *UpfCond) SetTaiList(v []Tai) {
	o.TaiList = v
}

func (o UpfCond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpfCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["conditionType"] = o.ConditionType
	if !IsNil(o.SmfServingArea) {
		toSerialize["smfServingArea"] = o.SmfServingArea
	}
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	return toSerialize, nil
}

type NullableUpfCond struct {
	value *UpfCond
	isSet bool
}

func (v NullableUpfCond) Get() *UpfCond {
	return v.value
}

func (v *NullableUpfCond) Set(val *UpfCond) {
	v.value = val
	v.isSet = true
}

func (v NullableUpfCond) IsSet() bool {
	return v.isSet
}

func (v *NullableUpfCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpfCond(val *UpfCond) *NullableUpfCond {
	return &NullableUpfCond{value: val, isSet: true}
}

func (v NullableUpfCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpfCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


