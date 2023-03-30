/*
MSGS_ASRegistration

API for MSGS AS Registration Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MSGS_ASRegistration

import (
	"encoding/json"
)

// checks if the ASRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ASRegistration{}

// ASRegistration AS registration data
type ASRegistration struct {
	AsSvcId string `json:"asSvcId"`
	AppId *string `json:"appId,omitempty"`
	// String providing an URI formatted according to RFC 3986.
	TargetUri *string `json:"targetUri,omitempty"`
	AsProf *ASProfile `json:"asProf,omitempty"`
}

// NewASRegistration instantiates a new ASRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewASRegistration(asSvcId string) *ASRegistration {
	this := ASRegistration{}
	this.AsSvcId = asSvcId
	return &this
}

// NewASRegistrationWithDefaults instantiates a new ASRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewASRegistrationWithDefaults() *ASRegistration {
	this := ASRegistration{}
	return &this
}

// GetAsSvcId returns the AsSvcId field value
func (o *ASRegistration) GetAsSvcId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AsSvcId
}

// GetAsSvcIdOk returns a tuple with the AsSvcId field value
// and a boolean to check if the value has been set.
func (o *ASRegistration) GetAsSvcIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AsSvcId, true
}

// SetAsSvcId sets field value
func (o *ASRegistration) SetAsSvcId(v string) {
	o.AsSvcId = v
}

// GetAppId returns the AppId field value if set, zero value otherwise.
func (o *ASRegistration) GetAppId() string {
	if o == nil || IsNil(o.AppId) {
		var ret string
		return ret
	}
	return *o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASRegistration) GetAppIdOk() (*string, bool) {
	if o == nil || IsNil(o.AppId) {
		return nil, false
	}
	return o.AppId, true
}

// HasAppId returns a boolean if a field has been set.
func (o *ASRegistration) HasAppId() bool {
	if o != nil && !IsNil(o.AppId) {
		return true
	}

	return false
}

// SetAppId gets a reference to the given string and assigns it to the AppId field.
func (o *ASRegistration) SetAppId(v string) {
	o.AppId = &v
}

// GetTargetUri returns the TargetUri field value if set, zero value otherwise.
func (o *ASRegistration) GetTargetUri() string {
	if o == nil || IsNil(o.TargetUri) {
		var ret string
		return ret
	}
	return *o.TargetUri
}

// GetTargetUriOk returns a tuple with the TargetUri field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASRegistration) GetTargetUriOk() (*string, bool) {
	if o == nil || IsNil(o.TargetUri) {
		return nil, false
	}
	return o.TargetUri, true
}

// HasTargetUri returns a boolean if a field has been set.
func (o *ASRegistration) HasTargetUri() bool {
	if o != nil && !IsNil(o.TargetUri) {
		return true
	}

	return false
}

// SetTargetUri gets a reference to the given string and assigns it to the TargetUri field.
func (o *ASRegistration) SetTargetUri(v string) {
	o.TargetUri = &v
}

// GetAsProf returns the AsProf field value if set, zero value otherwise.
func (o *ASRegistration) GetAsProf() ASProfile {
	if o == nil || IsNil(o.AsProf) {
		var ret ASProfile
		return ret
	}
	return *o.AsProf
}

// GetAsProfOk returns a tuple with the AsProf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ASRegistration) GetAsProfOk() (*ASProfile, bool) {
	if o == nil || IsNil(o.AsProf) {
		return nil, false
	}
	return o.AsProf, true
}

// HasAsProf returns a boolean if a field has been set.
func (o *ASRegistration) HasAsProf() bool {
	if o != nil && !IsNil(o.AsProf) {
		return true
	}

	return false
}

// SetAsProf gets a reference to the given ASProfile and assigns it to the AsProf field.
func (o *ASRegistration) SetAsProf(v ASProfile) {
	o.AsProf = &v
}

func (o ASRegistration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ASRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["asSvcId"] = o.AsSvcId
	if !IsNil(o.AppId) {
		toSerialize["appId"] = o.AppId
	}
	if !IsNil(o.TargetUri) {
		toSerialize["targetUri"] = o.TargetUri
	}
	if !IsNil(o.AsProf) {
		toSerialize["asProf"] = o.AsProf
	}
	return toSerialize, nil
}

type NullableASRegistration struct {
	value *ASRegistration
	isSet bool
}

func (v NullableASRegistration) Get() *ASRegistration {
	return v.value
}

func (v *NullableASRegistration) Set(val *ASRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableASRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableASRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableASRegistration(val *ASRegistration) *NullableASRegistration {
	return &NullableASRegistration{value: val, isSet: true}
}

func (v NullableASRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableASRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

