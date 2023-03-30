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

// checks if the AllowedMtcProviderInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AllowedMtcProviderInfo{}

// AllowedMtcProviderInfo struct for AllowedMtcProviderInfo
type AllowedMtcProviderInfo struct {
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
	AfId *string `json:"afId,omitempty"`
}

// NewAllowedMtcProviderInfo instantiates a new AllowedMtcProviderInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAllowedMtcProviderInfo() *AllowedMtcProviderInfo {
	this := AllowedMtcProviderInfo{}
	return &this
}

// NewAllowedMtcProviderInfoWithDefaults instantiates a new AllowedMtcProviderInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAllowedMtcProviderInfoWithDefaults() *AllowedMtcProviderInfo {
	this := AllowedMtcProviderInfo{}
	return &this
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *AllowedMtcProviderInfo) GetMtcProviderInformation() string {
	if o == nil || IsNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedMtcProviderInfo) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderInformation) {
		return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *AllowedMtcProviderInfo) HasMtcProviderInformation() bool {
	if o != nil && !IsNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *AllowedMtcProviderInfo) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

// GetAfId returns the AfId field value if set, zero value otherwise.
func (o *AllowedMtcProviderInfo) GetAfId() string {
	if o == nil || IsNil(o.AfId) {
		var ret string
		return ret
	}
	return *o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AllowedMtcProviderInfo) GetAfIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfId) {
		return nil, false
	}
	return o.AfId, true
}

// HasAfId returns a boolean if a field has been set.
func (o *AllowedMtcProviderInfo) HasAfId() bool {
	if o != nil && !IsNil(o.AfId) {
		return true
	}

	return false
}

// SetAfId gets a reference to the given string and assigns it to the AfId field.
func (o *AllowedMtcProviderInfo) SetAfId(v string) {
	o.AfId = &v
}

func (o AllowedMtcProviderInfo) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AllowedMtcProviderInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	if !IsNil(o.AfId) {
		toSerialize["afId"] = o.AfId
	}
	return toSerialize, nil
}

type NullableAllowedMtcProviderInfo struct {
	value *AllowedMtcProviderInfo
	isSet bool
}

func (v NullableAllowedMtcProviderInfo) Get() *AllowedMtcProviderInfo {
	return v.value
}

func (v *NullableAllowedMtcProviderInfo) Set(val *AllowedMtcProviderInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableAllowedMtcProviderInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableAllowedMtcProviderInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAllowedMtcProviderInfo(val *AllowedMtcProviderInfo) *NullableAllowedMtcProviderInfo {
	return &NullableAllowedMtcProviderInfo{value: val, isSet: true}
}

func (v NullableAllowedMtcProviderInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAllowedMtcProviderInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

