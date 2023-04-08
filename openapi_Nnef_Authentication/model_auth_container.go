/*
Nnef_Authentication

NEF Auth Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnef_Authentication

import (
	"encoding/json"
)

// checks if the AuthContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthContainer{}

// AuthContainer Authentication/Authorization data
type AuthContainer struct {
	// string with format 'bytes' as defined in OpenAPI
	AuthMsgType *string `json:"authMsgType,omitempty"`
	AuthMsgPayload *RefToBinaryData `json:"authMsgPayload,omitempty"`
	AuthResult *AuthResult `json:"authResult,omitempty"`
}

// NewAuthContainer instantiates a new AuthContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthContainer() *AuthContainer {
	this := AuthContainer{}
	return &this
}

// NewAuthContainerWithDefaults instantiates a new AuthContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthContainerWithDefaults() *AuthContainer {
	this := AuthContainer{}
	return &this
}

// GetAuthMsgType returns the AuthMsgType field value if set, zero value otherwise.
func (o *AuthContainer) GetAuthMsgType() string {
	if o == nil || isNil(o.AuthMsgType) {
		var ret string
		return ret
	}
	return *o.AuthMsgType
}

// GetAuthMsgTypeOk returns a tuple with the AuthMsgType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthContainer) GetAuthMsgTypeOk() (*string, bool) {
	if o == nil || isNil(o.AuthMsgType) {
		return nil, false
	}
	return o.AuthMsgType, true
}

// HasAuthMsgType returns a boolean if a field has been set.
func (o *AuthContainer) HasAuthMsgType() bool {
	if o != nil && !isNil(o.AuthMsgType) {
		return true
	}

	return false
}

// SetAuthMsgType gets a reference to the given string and assigns it to the AuthMsgType field.
func (o *AuthContainer) SetAuthMsgType(v string) {
	o.AuthMsgType = &v
}

// GetAuthMsgPayload returns the AuthMsgPayload field value if set, zero value otherwise.
func (o *AuthContainer) GetAuthMsgPayload() RefToBinaryData {
	if o == nil || isNil(o.AuthMsgPayload) {
		var ret RefToBinaryData
		return ret
	}
	return *o.AuthMsgPayload
}

// GetAuthMsgPayloadOk returns a tuple with the AuthMsgPayload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthContainer) GetAuthMsgPayloadOk() (*RefToBinaryData, bool) {
	if o == nil || isNil(o.AuthMsgPayload) {
		return nil, false
	}
	return o.AuthMsgPayload, true
}

// HasAuthMsgPayload returns a boolean if a field has been set.
func (o *AuthContainer) HasAuthMsgPayload() bool {
	if o != nil && !isNil(o.AuthMsgPayload) {
		return true
	}

	return false
}

// SetAuthMsgPayload gets a reference to the given RefToBinaryData and assigns it to the AuthMsgPayload field.
func (o *AuthContainer) SetAuthMsgPayload(v RefToBinaryData) {
	o.AuthMsgPayload = &v
}

// GetAuthResult returns the AuthResult field value if set, zero value otherwise.
func (o *AuthContainer) GetAuthResult() AuthResult {
	if o == nil || isNil(o.AuthResult) {
		var ret AuthResult
		return ret
	}
	return *o.AuthResult
}

// GetAuthResultOk returns a tuple with the AuthResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthContainer) GetAuthResultOk() (*AuthResult, bool) {
	if o == nil || isNil(o.AuthResult) {
		return nil, false
	}
	return o.AuthResult, true
}

// HasAuthResult returns a boolean if a field has been set.
func (o *AuthContainer) HasAuthResult() bool {
	if o != nil && !isNil(o.AuthResult) {
		return true
	}

	return false
}

// SetAuthResult gets a reference to the given AuthResult and assigns it to the AuthResult field.
func (o *AuthContainer) SetAuthResult(v AuthResult) {
	o.AuthResult = &v
}

func (o AuthContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AuthMsgType) {
		toSerialize["authMsgType"] = o.AuthMsgType
	}
	if !isNil(o.AuthMsgPayload) {
		toSerialize["authMsgPayload"] = o.AuthMsgPayload
	}
	if !isNil(o.AuthResult) {
		toSerialize["authResult"] = o.AuthResult
	}
	return toSerialize, nil
}

type NullableAuthContainer struct {
	value *AuthContainer
	isSet bool
}

func (v NullableAuthContainer) Get() *AuthContainer {
	return v.value
}

func (v *NullableAuthContainer) Set(val *AuthContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthContainer(val *AuthContainer) *NullableAuthContainer {
	return &NullableAuthContainer{value: val, isSet: true}
}

func (v NullableAuthContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


