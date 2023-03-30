/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g-ddnmf_Discovery

import (
	"encoding/json"
)

// checks if the DiscoveryAuthRespData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DiscoveryAuthRespData{}

// DiscoveryAuthRespData Represents the obtained authorization Data for a discoverer UE
type DiscoveryAuthRespData struct {
	AuthDataRestricted *AuthDataForRestricted `json:"authDataRestricted,omitempty"`
}

// NewDiscoveryAuthRespData instantiates a new DiscoveryAuthRespData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDiscoveryAuthRespData() *DiscoveryAuthRespData {
	this := DiscoveryAuthRespData{}
	return &this
}

// NewDiscoveryAuthRespDataWithDefaults instantiates a new DiscoveryAuthRespData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDiscoveryAuthRespDataWithDefaults() *DiscoveryAuthRespData {
	this := DiscoveryAuthRespData{}
	return &this
}

// GetAuthDataRestricted returns the AuthDataRestricted field value if set, zero value otherwise.
func (o *DiscoveryAuthRespData) GetAuthDataRestricted() AuthDataForRestricted {
	if o == nil || IsNil(o.AuthDataRestricted) {
		var ret AuthDataForRestricted
		return ret
	}
	return *o.AuthDataRestricted
}

// GetAuthDataRestrictedOk returns a tuple with the AuthDataRestricted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DiscoveryAuthRespData) GetAuthDataRestrictedOk() (*AuthDataForRestricted, bool) {
	if o == nil || IsNil(o.AuthDataRestricted) {
		return nil, false
	}
	return o.AuthDataRestricted, true
}

// HasAuthDataRestricted returns a boolean if a field has been set.
func (o *DiscoveryAuthRespData) HasAuthDataRestricted() bool {
	if o != nil && !IsNil(o.AuthDataRestricted) {
		return true
	}

	return false
}

// SetAuthDataRestricted gets a reference to the given AuthDataForRestricted and assigns it to the AuthDataRestricted field.
func (o *DiscoveryAuthRespData) SetAuthDataRestricted(v AuthDataForRestricted) {
	o.AuthDataRestricted = &v
}

func (o DiscoveryAuthRespData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DiscoveryAuthRespData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthDataRestricted) {
		toSerialize["authDataRestricted"] = o.AuthDataRestricted
	}
	return toSerialize, nil
}

type NullableDiscoveryAuthRespData struct {
	value *DiscoveryAuthRespData
	isSet bool
}

func (v NullableDiscoveryAuthRespData) Get() *DiscoveryAuthRespData {
	return v.value
}

func (v *NullableDiscoveryAuthRespData) Set(val *DiscoveryAuthRespData) {
	v.value = val
	v.isSet = true
}

func (v NullableDiscoveryAuthRespData) IsSet() bool {
	return v.isSet
}

func (v *NullableDiscoveryAuthRespData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDiscoveryAuthRespData(val *DiscoveryAuthRespData) *NullableDiscoveryAuthRespData {
	return &NullableDiscoveryAuthRespData{value: val, isSet: true}
}

func (v NullableDiscoveryAuthRespData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDiscoveryAuthRespData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


