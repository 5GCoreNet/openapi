/*
JOSE Protected Message Forwarding API

N32-f Message Forwarding Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_JOSEProtectedMessageForwarding

import (
	"encoding/json"
)

// checks if the FlatJwsJson type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FlatJwsJson{}

// FlatJwsJson Contains the modification from IPXes on path
type FlatJwsJson struct {
	Payload   string                 `json:"payload"`
	Protected *string                `json:"protected,omitempty"`
	Header    map[string]interface{} `json:"header,omitempty"`
	Signature string                 `json:"signature"`
}

// NewFlatJwsJson instantiates a new FlatJwsJson object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFlatJwsJson(payload string, signature string) *FlatJwsJson {
	this := FlatJwsJson{}
	this.Payload = payload
	this.Signature = signature
	return &this
}

// NewFlatJwsJsonWithDefaults instantiates a new FlatJwsJson object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFlatJwsJsonWithDefaults() *FlatJwsJson {
	this := FlatJwsJson{}
	return &this
}

// GetPayload returns the Payload field value
func (o *FlatJwsJson) GetPayload() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value
// and a boolean to check if the value has been set.
func (o *FlatJwsJson) GetPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Payload, true
}

// SetPayload sets field value
func (o *FlatJwsJson) SetPayload(v string) {
	o.Payload = v
}

// GetProtected returns the Protected field value if set, zero value otherwise.
func (o *FlatJwsJson) GetProtected() string {
	if o == nil || IsNil(o.Protected) {
		var ret string
		return ret
	}
	return *o.Protected
}

// GetProtectedOk returns a tuple with the Protected field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJwsJson) GetProtectedOk() (*string, bool) {
	if o == nil || IsNil(o.Protected) {
		return nil, false
	}
	return o.Protected, true
}

// HasProtected returns a boolean if a field has been set.
func (o *FlatJwsJson) HasProtected() bool {
	if o != nil && !IsNil(o.Protected) {
		return true
	}

	return false
}

// SetProtected gets a reference to the given string and assigns it to the Protected field.
func (o *FlatJwsJson) SetProtected(v string) {
	o.Protected = &v
}

// GetHeader returns the Header field value if set, zero value otherwise.
func (o *FlatJwsJson) GetHeader() map[string]interface{} {
	if o == nil || IsNil(o.Header) {
		var ret map[string]interface{}
		return ret
	}
	return o.Header
}

// GetHeaderOk returns a tuple with the Header field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FlatJwsJson) GetHeaderOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Header) {
		return map[string]interface{}{}, false
	}
	return o.Header, true
}

// HasHeader returns a boolean if a field has been set.
func (o *FlatJwsJson) HasHeader() bool {
	if o != nil && !IsNil(o.Header) {
		return true
	}

	return false
}

// SetHeader gets a reference to the given map[string]interface{} and assigns it to the Header field.
func (o *FlatJwsJson) SetHeader(v map[string]interface{}) {
	o.Header = v
}

// GetSignature returns the Signature field value
func (o *FlatJwsJson) GetSignature() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Signature
}

// GetSignatureOk returns a tuple with the Signature field value
// and a boolean to check if the value has been set.
func (o *FlatJwsJson) GetSignatureOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Signature, true
}

// SetSignature sets field value
func (o *FlatJwsJson) SetSignature(v string) {
	o.Signature = v
}

func (o FlatJwsJson) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FlatJwsJson) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["payload"] = o.Payload
	if !IsNil(o.Protected) {
		toSerialize["protected"] = o.Protected
	}
	if !IsNil(o.Header) {
		toSerialize["header"] = o.Header
	}
	toSerialize["signature"] = o.Signature
	return toSerialize, nil
}

type NullableFlatJwsJson struct {
	value *FlatJwsJson
	isSet bool
}

func (v NullableFlatJwsJson) Get() *FlatJwsJson {
	return v.value
}

func (v *NullableFlatJwsJson) Set(val *FlatJwsJson) {
	v.value = val
	v.isSet = true
}

func (v NullableFlatJwsJson) IsSet() bool {
	return v.isSet
}

func (v *NullableFlatJwsJson) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFlatJwsJson(val *FlatJwsJson) *NullableFlatJwsJson {
	return &NullableFlatJwsJson{value: val, isSet: true}
}

func (v NullableFlatJwsJson) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFlatJwsJson) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
