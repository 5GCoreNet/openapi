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

// checks if the Model5GDdnmfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Model5GDdnmfInfo{}

// Model5GDdnmfInfo Information of an 5G DDNMF NF Instance
type Model5GDdnmfInfo struct {
	PlmnId PlmnId `json:"plmnId"`
}

// NewModel5GDdnmfInfo instantiates a new Model5GDdnmfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewModel5GDdnmfInfo(plmnId PlmnId) *Model5GDdnmfInfo {
	this := Model5GDdnmfInfo{}
	this.PlmnId = plmnId
	return &this
}

// NewModel5GDdnmfInfoWithDefaults instantiates a new Model5GDdnmfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewModel5GDdnmfInfoWithDefaults() *Model5GDdnmfInfo {
	this := Model5GDdnmfInfo{}
	return &this
}

// GetPlmnId returns the PlmnId field value
func (o *Model5GDdnmfInfo) GetPlmnId() PlmnId {
	if o == nil {
		var ret PlmnId
		return ret
	}

	return o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value
// and a boolean to check if the value has been set.
func (o *Model5GDdnmfInfo) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlmnId, true
}

// SetPlmnId sets field value
func (o *Model5GDdnmfInfo) SetPlmnId(v PlmnId) {
	o.PlmnId = v
}

func (o Model5GDdnmfInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Model5GDdnmfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["plmnId"] = o.PlmnId
	return toSerialize, nil
}

type NullableModel5GDdnmfInfo struct {
	value *Model5GDdnmfInfo
	isSet bool
}

func (v NullableModel5GDdnmfInfo) Get() *Model5GDdnmfInfo {
	return v.value
}

func (v *NullableModel5GDdnmfInfo) Set(val *Model5GDdnmfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableModel5GDdnmfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableModel5GDdnmfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableModel5GDdnmfInfo(val *Model5GDdnmfInfo) *NullableModel5GDdnmfInfo {
	return &NullableModel5GDdnmfInfo{value: val, isSet: true}
}

func (v NullableModel5GDdnmfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableModel5GDdnmfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
