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

// checks if the ProseServiceAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProseServiceAuth{}

// ProseServiceAuth Indicates whether the UE is authorized to use ProSe Direct Discovery, ProSe Direct Communication, or both. 
type ProseServiceAuth struct {
	ProseDirectDiscoveryAuth *UeAuth `json:"proseDirectDiscoveryAuth,omitempty"`
	ProseDirectCommunicationAuth *UeAuth `json:"proseDirectCommunicationAuth,omitempty"`
}

// NewProseServiceAuth instantiates a new ProseServiceAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProseServiceAuth() *ProseServiceAuth {
	this := ProseServiceAuth{}
	return &this
}

// NewProseServiceAuthWithDefaults instantiates a new ProseServiceAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProseServiceAuthWithDefaults() *ProseServiceAuth {
	this := ProseServiceAuth{}
	return &this
}

// GetProseDirectDiscoveryAuth returns the ProseDirectDiscoveryAuth field value if set, zero value otherwise.
func (o *ProseServiceAuth) GetProseDirectDiscoveryAuth() UeAuth {
	if o == nil || IsNil(o.ProseDirectDiscoveryAuth) {
		var ret UeAuth
		return ret
	}
	return *o.ProseDirectDiscoveryAuth
}

// GetProseDirectDiscoveryAuthOk returns a tuple with the ProseDirectDiscoveryAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseServiceAuth) GetProseDirectDiscoveryAuthOk() (*UeAuth, bool) {
	if o == nil || IsNil(o.ProseDirectDiscoveryAuth) {
		return nil, false
	}
	return o.ProseDirectDiscoveryAuth, true
}

// HasProseDirectDiscoveryAuth returns a boolean if a field has been set.
func (o *ProseServiceAuth) HasProseDirectDiscoveryAuth() bool {
	if o != nil && !IsNil(o.ProseDirectDiscoveryAuth) {
		return true
	}

	return false
}

// SetProseDirectDiscoveryAuth gets a reference to the given UeAuth and assigns it to the ProseDirectDiscoveryAuth field.
func (o *ProseServiceAuth) SetProseDirectDiscoveryAuth(v UeAuth) {
	o.ProseDirectDiscoveryAuth = &v
}

// GetProseDirectCommunicationAuth returns the ProseDirectCommunicationAuth field value if set, zero value otherwise.
func (o *ProseServiceAuth) GetProseDirectCommunicationAuth() UeAuth {
	if o == nil || IsNil(o.ProseDirectCommunicationAuth) {
		var ret UeAuth
		return ret
	}
	return *o.ProseDirectCommunicationAuth
}

// GetProseDirectCommunicationAuthOk returns a tuple with the ProseDirectCommunicationAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProseServiceAuth) GetProseDirectCommunicationAuthOk() (*UeAuth, bool) {
	if o == nil || IsNil(o.ProseDirectCommunicationAuth) {
		return nil, false
	}
	return o.ProseDirectCommunicationAuth, true
}

// HasProseDirectCommunicationAuth returns a boolean if a field has been set.
func (o *ProseServiceAuth) HasProseDirectCommunicationAuth() bool {
	if o != nil && !IsNil(o.ProseDirectCommunicationAuth) {
		return true
	}

	return false
}

// SetProseDirectCommunicationAuth gets a reference to the given UeAuth and assigns it to the ProseDirectCommunicationAuth field.
func (o *ProseServiceAuth) SetProseDirectCommunicationAuth(v UeAuth) {
	o.ProseDirectCommunicationAuth = &v
}

func (o ProseServiceAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProseServiceAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProseDirectDiscoveryAuth) {
		toSerialize["proseDirectDiscoveryAuth"] = o.ProseDirectDiscoveryAuth
	}
	if !IsNil(o.ProseDirectCommunicationAuth) {
		toSerialize["proseDirectCommunicationAuth"] = o.ProseDirectCommunicationAuth
	}
	return toSerialize, nil
}

type NullableProseServiceAuth struct {
	value *ProseServiceAuth
	isSet bool
}

func (v NullableProseServiceAuth) Get() *ProseServiceAuth {
	return v.value
}

func (v *NullableProseServiceAuth) Set(val *ProseServiceAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableProseServiceAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableProseServiceAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseServiceAuth(val *ProseServiceAuth) *NullableProseServiceAuth {
	return &NullableProseServiceAuth{value: val, isSet: true}
}

func (v NullableProseServiceAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseServiceAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


