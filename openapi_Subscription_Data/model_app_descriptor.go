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

// checks if the AppDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppDescriptor{}

// AppDescriptor struct for AppDescriptor
type AppDescriptor struct {
	// Represents the Operating System of the served UE.
	OsId *string `json:"osId,omitempty"`
	AppId *string `json:"appId,omitempty"`
}

// NewAppDescriptor instantiates a new AppDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppDescriptor() *AppDescriptor {
	this := AppDescriptor{}
	return &this
}

// NewAppDescriptorWithDefaults instantiates a new AppDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppDescriptorWithDefaults() *AppDescriptor {
	this := AppDescriptor{}
	return &this
}

// GetOsId returns the OsId field value if set, zero value otherwise.
func (o *AppDescriptor) GetOsId() string {
	if o == nil || isNil(o.OsId) {
		var ret string
		return ret
	}
	return *o.OsId
}

// GetOsIdOk returns a tuple with the OsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDescriptor) GetOsIdOk() (*string, bool) {
	if o == nil || isNil(o.OsId) {
		return nil, false
	}
	return o.OsId, true
}

// HasOsId returns a boolean if a field has been set.
func (o *AppDescriptor) HasOsId() bool {
	if o != nil && !isNil(o.OsId) {
		return true
	}

	return false
}

// SetOsId gets a reference to the given string and assigns it to the OsId field.
func (o *AppDescriptor) SetOsId(v string) {
	o.OsId = &v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *AppDescriptor) GetAppId() string {
	if o == nil || isNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppDescriptor) GetAppIdOk() (*string, bool) {
	if o == nil || isNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *AppDescriptor) HasAppId() bool {
	if o != nil && !isNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *AppDescriptor) SetAppId(v string) {
	o.AppId = &v
}

func (o AppDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.OsId) {
		toSerialize["osId"] = o.OsId
	}
	if !isNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	return toSerialize, nil
}

type NullableAppDescriptor struct {
	value *AppDescriptor
	isSet bool
}

func (v NullableAppDescriptor) Get() *AppDescriptor {
	return v.value
}

func (v *NullableAppDescriptor) Set(val *AppDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableAppDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableAppDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppDescriptor(val *AppDescriptor) *NullableAppDescriptor {
	return &NullableAppDescriptor{value: val, isSet: true}
}

func (v NullableAppDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


