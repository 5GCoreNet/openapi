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

// checks if the MnpfInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MnpfInfo{}

// MnpfInfo Information of an MNPF Instance
type MnpfInfo struct {
	MsisdnRanges []IdentityRange `json:"msisdnRanges"`
}

// NewMnpfInfo instantiates a new MnpfInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMnpfInfo(msisdnRanges []IdentityRange) *MnpfInfo {
	this := MnpfInfo{}
	this.MsisdnRanges = msisdnRanges
	return &this
}

// NewMnpfInfoWithDefaults instantiates a new MnpfInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMnpfInfoWithDefaults() *MnpfInfo {
	this := MnpfInfo{}
	return &this
}

// GetMsisdnRanges returns the MsisdnRanges field value
func (o *MnpfInfo) GetMsisdnRanges() []IdentityRange {
	if o == nil {
		var ret []IdentityRange
		return ret
	}

	return o.MsisdnRanges
}

// GetMsisdnRangesOk returns a tuple with the MsisdnRanges field value
// and a boolean to check if the value has been set.
func (o *MnpfInfo) GetMsisdnRangesOk() ([]IdentityRange, bool) {
	if o == nil {
		return nil, false
	}
	return o.MsisdnRanges, true
}

// SetMsisdnRanges sets field value
func (o *MnpfInfo) SetMsisdnRanges(v []IdentityRange) {
	o.MsisdnRanges = v
}

func (o MnpfInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MnpfInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msisdnRanges"] = o.MsisdnRanges
	return toSerialize, nil
}

type NullableMnpfInfo struct {
	value *MnpfInfo
	isSet bool
}

func (v NullableMnpfInfo) Get() *MnpfInfo {
	return v.value
}

func (v *NullableMnpfInfo) Set(val *MnpfInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableMnpfInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableMnpfInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMnpfInfo(val *MnpfInfo) *NullableMnpfInfo {
	return &NullableMnpfInfo{value: val, isSet: true}
}

func (v NullableMnpfInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMnpfInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


