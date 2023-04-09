/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the NiddAuthorizationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NiddAuthorizationInfo{}

// NiddAuthorizationInfo Information related to active NIDD Authorizations
type NiddAuthorizationInfo struct {
	NiddAuthorizationList []AuthorizationInfo `json:"niddAuthorizationList"`
}

// NewNiddAuthorizationInfo instantiates a new NiddAuthorizationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNiddAuthorizationInfo(niddAuthorizationList []AuthorizationInfo) *NiddAuthorizationInfo {
	this := NiddAuthorizationInfo{}
	this.NiddAuthorizationList = niddAuthorizationList
	return &this
}

// NewNiddAuthorizationInfoWithDefaults instantiates a new NiddAuthorizationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNiddAuthorizationInfoWithDefaults() *NiddAuthorizationInfo {
	this := NiddAuthorizationInfo{}
	return &this
}

// GetNiddAuthorizationList returns the NiddAuthorizationList field value
func (o *NiddAuthorizationInfo) GetNiddAuthorizationList() []AuthorizationInfo {
	if o == nil {
		var ret []AuthorizationInfo
		return ret
	}

	return o.NiddAuthorizationList
}

// GetNiddAuthorizationListOk returns a tuple with the NiddAuthorizationList field value
// and a boolean to check if the value has been set.
func (o *NiddAuthorizationInfo) GetNiddAuthorizationListOk() ([]AuthorizationInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.NiddAuthorizationList, true
}

// SetNiddAuthorizationList sets field value
func (o *NiddAuthorizationInfo) SetNiddAuthorizationList(v []AuthorizationInfo) {
	o.NiddAuthorizationList = v
}

func (o NiddAuthorizationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NiddAuthorizationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["niddAuthorizationList"] = o.NiddAuthorizationList
	return toSerialize, nil
}

type NullableNiddAuthorizationInfo struct {
	value *NiddAuthorizationInfo
	isSet bool
}

func (v NullableNiddAuthorizationInfo) Get() *NiddAuthorizationInfo {
	return v.value
}

func (v *NullableNiddAuthorizationInfo) Set(val *NiddAuthorizationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableNiddAuthorizationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableNiddAuthorizationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNiddAuthorizationInfo(val *NiddAuthorizationInfo) *NullableNiddAuthorizationInfo {
	return &NullableNiddAuthorizationInfo{value: val, isSet: true}
}

func (v NullableNiddAuthorizationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNiddAuthorizationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
