/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the NgKsi type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NgKsi{}

// NgKsi Represents the ngKSI
type NgKsi struct {
	Tsc ScType `json:"tsc"`
	Ksi int32 `json:"ksi"`
}

// NewNgKsi instantiates a new NgKsi object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNgKsi(tsc ScType, ksi int32) *NgKsi {
	this := NgKsi{}
	this.Tsc = tsc
	this.Ksi = ksi
	return &this
}

// NewNgKsiWithDefaults instantiates a new NgKsi object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNgKsiWithDefaults() *NgKsi {
	this := NgKsi{}
	return &this
}

// GetTsc returns the Tsc field value
func (o *NgKsi) GetTsc() ScType {
	if o == nil {
		var ret ScType
		return ret
	}

	return o.Tsc
}

// GetTscOk returns a tuple with the Tsc field value
// and a boolean to check if the value has been set.
func (o *NgKsi) GetTscOk() (*ScType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Tsc, true
}

// SetTsc sets field value
func (o *NgKsi) SetTsc(v ScType) {
	o.Tsc = v
}

// GetKsi returns the Ksi field value
func (o *NgKsi) GetKsi() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ksi
}

// GetKsiOk returns a tuple with the Ksi field value
// and a boolean to check if the value has been set.
func (o *NgKsi) GetKsiOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ksi, true
}

// SetKsi sets field value
func (o *NgKsi) SetKsi(v int32) {
	o.Ksi = v
}

func (o NgKsi) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NgKsi) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tsc"] = o.Tsc
	toSerialize["ksi"] = o.Ksi
	return toSerialize, nil
}

type NullableNgKsi struct {
	value *NgKsi
	isSet bool
}

func (v NullableNgKsi) Get() *NgKsi {
	return v.value
}

func (v *NullableNgKsi) Set(val *NgKsi) {
	v.value = val
	v.isSet = true
}

func (v NullableNgKsi) IsSet() bool {
	return v.isSet
}

func (v *NullableNgKsi) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNgKsi(val *NgKsi) *NullableNgKsi {
	return &NullableNgKsi{value: val, isSet: true}
}

func (v NullableNgKsi) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNgKsi) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


