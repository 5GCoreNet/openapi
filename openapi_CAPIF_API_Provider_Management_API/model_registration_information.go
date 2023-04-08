/*
CAPIF_API_Provider_Management_API

API for API provider domain functions management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_CAPIF_API_Provider_Management_API

import (
	"encoding/json"
)

// checks if the RegistrationInformation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RegistrationInformation{}

// RegistrationInformation Represents registration information of an individual API provider domain function. 
type RegistrationInformation struct {
	// Public Key of API Provider domain function.
	ApiProvPubKey string `json:"apiProvPubKey"`
	// API provider domain function's client certificate
	ApiProvCert *string `json:"apiProvCert,omitempty"`
}

// NewRegistrationInformation instantiates a new RegistrationInformation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRegistrationInformation(apiProvPubKey string) *RegistrationInformation {
	this := RegistrationInformation{}
	this.ApiProvPubKey = apiProvPubKey
	return &this
}

// NewRegistrationInformationWithDefaults instantiates a new RegistrationInformation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRegistrationInformationWithDefaults() *RegistrationInformation {
	this := RegistrationInformation{}
	return &this
}

// GetApiProvPubKey returns the ApiProvPubKey field value
func (o *RegistrationInformation) GetApiProvPubKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiProvPubKey
}

// GetApiProvPubKeyOk returns a tuple with the ApiProvPubKey field value
// and a boolean to check if the value has been set.
func (o *RegistrationInformation) GetApiProvPubKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiProvPubKey, true
}

// SetApiProvPubKey sets field value
func (o *RegistrationInformation) SetApiProvPubKey(v string) {
	o.ApiProvPubKey = v
}

// GetApiProvCert returns the ApiProvCert field value if set, zero value otherwise.
func (o *RegistrationInformation) GetApiProvCert() string {
	if o == nil || isNil(o.ApiProvCert) {
		var ret string
		return ret
	}
	return *o.ApiProvCert
}

// GetApiProvCertOk returns a tuple with the ApiProvCert field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RegistrationInformation) GetApiProvCertOk() (*string, bool) {
	if o == nil || isNil(o.ApiProvCert) {
		return nil, false
	}
	return o.ApiProvCert, true
}

// HasApiProvCert returns a boolean if a field has been set.
func (o *RegistrationInformation) HasApiProvCert() bool {
	if o != nil && !isNil(o.ApiProvCert) {
		return true
	}

	return false
}

// SetApiProvCert gets a reference to the given string and assigns it to the ApiProvCert field.
func (o *RegistrationInformation) SetApiProvCert(v string) {
	o.ApiProvCert = &v
}

func (o RegistrationInformation) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RegistrationInformation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiProvPubKey"] = o.ApiProvPubKey
	if !isNil(o.ApiProvCert) {
		toSerialize["apiProvCert"] = o.ApiProvCert
	}
	return toSerialize, nil
}

type NullableRegistrationInformation struct {
	value *RegistrationInformation
	isSet bool
}

func (v NullableRegistrationInformation) Get() *RegistrationInformation {
	return v.value
}

func (v *NullableRegistrationInformation) Set(val *RegistrationInformation) {
	v.value = val
	v.isSet = true
}

func (v NullableRegistrationInformation) IsSet() bool {
	return v.isSet
}

func (v *NullableRegistrationInformation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRegistrationInformation(val *RegistrationInformation) *NullableRegistrationInformation {
	return &NullableRegistrationInformation{value: val, isSet: true}
}

func (v NullableRegistrationInformation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRegistrationInformation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


