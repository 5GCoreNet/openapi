/*
Nudm_UECM

Nudm Context Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UECM

import (
	"encoding/json"
)

// checks if the NwdafRegistrationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NwdafRegistrationInfo{}

// NwdafRegistrationInfo List of NwdafRegistration
type NwdafRegistrationInfo struct {
	NwdafRegistrationList []NwdafRegistration `json:"nwdafRegistrationList"`
}

// NewNwdafRegistrationInfo instantiates a new NwdafRegistrationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNwdafRegistrationInfo(nwdafRegistrationList []NwdafRegistration) *NwdafRegistrationInfo {
	this := NwdafRegistrationInfo{}
	this.NwdafRegistrationList = nwdafRegistrationList
	return &this
}

// NewNwdafRegistrationInfoWithDefaults instantiates a new NwdafRegistrationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNwdafRegistrationInfoWithDefaults() *NwdafRegistrationInfo {
	this := NwdafRegistrationInfo{}
	return &this
}

// GetNwdafRegistrationList returns the NwdafRegistrationList field value
func (o *NwdafRegistrationInfo) GetNwdafRegistrationList() []NwdafRegistration {
	if o == nil {
		var ret []NwdafRegistration
		return ret
	}

	return o.NwdafRegistrationList
}

// GetNwdafRegistrationListOk returns a tuple with the NwdafRegistrationList field value
// and a boolean to check if the value has been set.
func (o *NwdafRegistrationInfo) GetNwdafRegistrationListOk() ([]NwdafRegistration, bool) {
	if o == nil {
		return nil, false
	}
	return o.NwdafRegistrationList, true
}

// SetNwdafRegistrationList sets field value
func (o *NwdafRegistrationInfo) SetNwdafRegistrationList(v []NwdafRegistration) {
	o.NwdafRegistrationList = v
}

func (o NwdafRegistrationInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NwdafRegistrationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nwdafRegistrationList"] = o.NwdafRegistrationList
	return toSerialize, nil
}

type NullableNwdafRegistrationInfo struct {
	value *NwdafRegistrationInfo
	isSet bool
}

func (v NullableNwdafRegistrationInfo) Get() *NwdafRegistrationInfo {
	return v.value
}

func (v *NullableNwdafRegistrationInfo) Set(val *NwdafRegistrationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNwdafRegistrationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNwdafRegistrationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNwdafRegistrationInfo(val *NwdafRegistrationInfo) *NullableNwdafRegistrationInfo {
	return &NullableNwdafRegistrationInfo{value: val, isSet: true}
}

func (v NullableNwdafRegistrationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNwdafRegistrationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


