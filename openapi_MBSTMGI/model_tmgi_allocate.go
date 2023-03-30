/*
3gpp-mbs-tmgi

API for the allocation, deallocation and management of TMGI(s) for MBS.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSTMGI

import (
	"encoding/json"
)

// checks if the TmgiAllocate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TmgiAllocate{}

// TmgiAllocate Data within TMGI Allocate Request
type TmgiAllocate struct {
	// The number of requested TMGIs
	TmgiNumber *int32 `json:"tmgiNumber,omitempty"`
	// The list of TMGIs to be refreshed
	TmgiList []Tmgi `json:"tmgiList,omitempty"`
}

// NewTmgiAllocate instantiates a new TmgiAllocate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTmgiAllocate() *TmgiAllocate {
	this := TmgiAllocate{}
	return &this
}

// NewTmgiAllocateWithDefaults instantiates a new TmgiAllocate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTmgiAllocateWithDefaults() *TmgiAllocate {
	this := TmgiAllocate{}
	return &this
}

// GetTmgiNumber returns the TmgiNumber field value if set, zero value otherwise.
func (o *TmgiAllocate) GetTmgiNumber() int32 {
	if o == nil || IsNil(o.TmgiNumber) {
		var ret int32
		return ret
	}
	return *o.TmgiNumber
}

// GetTmgiNumberOk returns a tuple with the TmgiNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmgiAllocate) GetTmgiNumberOk() (*int32, bool) {
	if o == nil || IsNil(o.TmgiNumber) {
		return nil, false
	}
	return o.TmgiNumber, true
}

// HasTmgiNumber returns a boolean if a field has been set.
func (o *TmgiAllocate) HasTmgiNumber() bool {
	if o != nil && !IsNil(o.TmgiNumber) {
		return true
	}

	return false
}

// SetTmgiNumber gets a reference to the given int32 and assigns it to the TmgiNumber field.
func (o *TmgiAllocate) SetTmgiNumber(v int32) {
	o.TmgiNumber = &v
}

// GetTmgiList returns the TmgiList field value if set, zero value otherwise.
func (o *TmgiAllocate) GetTmgiList() []Tmgi {
	if o == nil || IsNil(o.TmgiList) {
		var ret []Tmgi
		return ret
	}
	return o.TmgiList
}

// GetTmgiListOk returns a tuple with the TmgiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TmgiAllocate) GetTmgiListOk() ([]Tmgi, bool) {
	if o == nil || IsNil(o.TmgiList) {
		return nil, false
	}
	return o.TmgiList, true
}

// HasTmgiList returns a boolean if a field has been set.
func (o *TmgiAllocate) HasTmgiList() bool {
	if o != nil && !IsNil(o.TmgiList) {
		return true
	}

	return false
}

// SetTmgiList gets a reference to the given []Tmgi and assigns it to the TmgiList field.
func (o *TmgiAllocate) SetTmgiList(v []Tmgi) {
	o.TmgiList = v
}

func (o TmgiAllocate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TmgiAllocate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TmgiNumber) {
		toSerialize["tmgiNumber"] = o.TmgiNumber
	}
	if !IsNil(o.TmgiList) {
		toSerialize["tmgiList"] = o.TmgiList
	}
	return toSerialize, nil
}

type NullableTmgiAllocate struct {
	value *TmgiAllocate
	isSet bool
}

func (v NullableTmgiAllocate) Get() *TmgiAllocate {
	return v.value
}

func (v *NullableTmgiAllocate) Set(val *TmgiAllocate) {
	v.value = val
	v.isSet = true
}

func (v NullableTmgiAllocate) IsSet() bool {
	return v.isSet
}

func (v *NullableTmgiAllocate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTmgiAllocate(val *TmgiAllocate) *NullableTmgiAllocate {
	return &NullableTmgiAllocate{value: val, isSet: true}
}

func (v NullableTmgiAllocate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTmgiAllocate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


