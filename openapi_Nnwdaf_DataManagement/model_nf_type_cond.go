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

// checks if the NfTypeCond type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NfTypeCond{}

// NfTypeCond Subscription to a set of NFs based on their NF Type
type NfTypeCond struct {
	NfType NFType `json:"nfType"`
}

// NewNfTypeCond instantiates a new NfTypeCond object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNfTypeCond(nfType NFType) *NfTypeCond {
	this := NfTypeCond{}
	this.NfType = nfType
	return &this
}

// NewNfTypeCondWithDefaults instantiates a new NfTypeCond object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNfTypeCondWithDefaults() *NfTypeCond {
	this := NfTypeCond{}
	return &this
}

// GetNfType returns the NfType field value
func (o *NfTypeCond) GetNfType() NFType {
	if o == nil {
		var ret NFType
		return ret
	}

	return o.NfType
}

// GetNfTypeOk returns a tuple with the NfType field value
// and a boolean to check if the value has been set.
func (o *NfTypeCond) GetNfTypeOk() (*NFType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NfType, true
}

// SetNfType sets field value
func (o *NfTypeCond) SetNfType(v NFType) {
	o.NfType = v
}

func (o NfTypeCond) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NfTypeCond) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nfType"] = o.NfType
	return toSerialize, nil
}

type NullableNfTypeCond struct {
	value *NfTypeCond
	isSet bool
}

func (v NullableNfTypeCond) Get() *NfTypeCond {
	return v.value
}

func (v *NullableNfTypeCond) Set(val *NfTypeCond) {
	v.value = val
	v.isSet = true
}

func (v NullableNfTypeCond) IsSet() bool {
	return v.isSet
}

func (v *NullableNfTypeCond) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNfTypeCond(val *NfTypeCond) *NullableNfTypeCond {
	return &NullableNfTypeCond{value: val, isSet: true}
}

func (v NullableNfTypeCond) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNfTypeCond) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


