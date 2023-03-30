/*
Unified Data Repository Service API file for subscription data

Unified Data Repository Service (subscription data).   The API version is defined in 3GPP TS 29.504.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Subscription_Data

import (
	"encoding/json"
)

// checks if the MtcProvider type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MtcProvider{}

// MtcProvider MTC provider information.
type MtcProvider struct {
	// String uniquely identifying MTC provider information.
	MtcProviderInformation *string `json:"mtcProviderInformation,omitempty"`
	AfId *string `json:"afId,omitempty"`
}

// NewMtcProvider instantiates a new MtcProvider object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMtcProvider() *MtcProvider {
	this := MtcProvider{}
	return &this
}

// NewMtcProviderWithDefaults instantiates a new MtcProvider object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMtcProviderWithDefaults() *MtcProvider {
	this := MtcProvider{}
	return &this
}

// GetMtcProviderInformation returns the MtcProviderInformation field value if set, zero value otherwise.
func (o *MtcProvider) GetMtcProviderInformation() string {
	if o == nil || IsNil(o.MtcProviderInformation) {
		var ret string
		return ret
	}
	return *o.MtcProviderInformation
}

// GetMtcProviderInformationOk returns a tuple with the MtcProviderInformation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtcProvider) GetMtcProviderInformationOk() (*string, bool) {
	if o == nil || IsNil(o.MtcProviderInformation) {
		return nil, false
	}
	return o.MtcProviderInformation, true
}

// HasMtcProviderInformation returns a boolean if a field has been set.
func (o *MtcProvider) HasMtcProviderInformation() bool {
	if o != nil && !IsNil(o.MtcProviderInformation) {
		return true
	}

	return false
}

// SetMtcProviderInformation gets a reference to the given string and assigns it to the MtcProviderInformation field.
func (o *MtcProvider) SetMtcProviderInformation(v string) {
	o.MtcProviderInformation = &v
}

// GetAfId returns the AfId field value if set, zero value otherwise.
func (o *MtcProvider) GetAfId() string {
	if o == nil || IsNil(o.AfId) {
		var ret string
		return ret
	}
	return *o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MtcProvider) GetAfIdOk() (*string, bool) {
	if o == nil || IsNil(o.AfId) {
		return nil, false
	}
	return o.AfId, true
}

// HasAfId returns a boolean if a field has been set.
func (o *MtcProvider) HasAfId() bool {
	if o != nil && !IsNil(o.AfId) {
		return true
	}

	return false
}

// SetAfId gets a reference to the given string and assigns it to the AfId field.
func (o *MtcProvider) SetAfId(v string) {
	o.AfId = &v
}

func (o MtcProvider) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MtcProvider) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MtcProviderInformation) {
		toSerialize["mtcProviderInformation"] = o.MtcProviderInformation
	}
	if !IsNil(o.AfId) {
		toSerialize["afId"] = o.AfId
	}
	return toSerialize, nil
}

type NullableMtcProvider struct {
	value *MtcProvider
	isSet bool
}

func (v NullableMtcProvider) Get() *MtcProvider {
	return v.value
}

func (v *NullableMtcProvider) Set(val *MtcProvider) {
	v.value = val
	v.isSet = true
}

func (v NullableMtcProvider) IsSet() bool {
	return v.isSet
}

func (v *NullableMtcProvider) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMtcProvider(val *MtcProvider) *NullableMtcProvider {
	return &NullableMtcProvider{value: val, isSet: true}
}

func (v NullableMtcProvider) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMtcProvider) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


