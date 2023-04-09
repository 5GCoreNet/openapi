/*
Ndccf_ContextManagement

DCCF Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_ContextManagement

import (
	"encoding/json"
)

// checks if the NfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NfInfo{}

// NfInfo Information of a generic NF Instance
type NfInfo struct {
	NfType *NFType `json:"nfType,omitempty"`
}

// NewNfInfo instantiates a new NfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfInfo() *NfInfo {
	this := NfInfo{}
	return &this
}

// NewNfInfoWithDefaults instantiates a new NfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfInfoWithDefaults() *NfInfo {
	this := NfInfo{}
	return &this
}

// GetNfType returns the NfType field value if set, zero value otherwise.
func (o *NfInfo) GetNfType() NFType {
	if o == nil || IsNil(o.NfType) {
		var ret NFType
		return ret
	}
	return *o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NfInfo) GetNfTypeOk() (*NFType, bool) {
	if o == nil || IsNil(o.NfType) {
		return nil, false
	}
	return o.NfType, true
}

// HasNfType returns a boolean if a field has been set.
func (o *NfInfo) HasNfType() bool {
	if o != nil && !IsNil(o.NfType) {
		return true
	}

	return false
}

// SetNfType gets a reference to the given NFType and assigns it to the NfType field.
func (o *NfInfo) SetNfType(v NFType) {
	o.NfType = &v
}

func (o NfInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.NfType) {
		toSerialize["nfType"] = o.NfType
	}
	return toSerialize, nil
}

type NullableNfInfo struct {
	value *NfInfo
	isSet bool
}

func (v NullableNfInfo) Get() *NfInfo {
	return v.value
}

func (v *NullableNfInfo) Set(val *NfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfInfo(val *NfInfo) *NullableNfInfo {
	return &NullableNfInfo{value: val, isSet: true}
}

func (v NullableNfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
