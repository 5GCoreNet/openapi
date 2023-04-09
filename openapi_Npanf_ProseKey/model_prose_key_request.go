/*
Npanf_ProseKey

PAnF ProseKey Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npanf_ProseKey

import (
	"encoding/json"
)

// checks if the ProseKeyRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProseKeyRequest{}

// ProseKeyRequest Prose Key Request.
type ProseKeyRequest struct {
	// A string carrying the CP-PRUK ID of the remote UE. The CP-PRUK ID is a string in NAI format as specified in clause 28.7.11 of 3GPP TS 23.003.
	Var5gPrukId string `json:"5gPrukId"`
	// Relay Service Code to identify a connectivity service provided by the UE-to-Network relay.
	RelayServiceCode int32 `json:"relayServiceCode"`
}

// NewProseKeyRequest instantiates a new ProseKeyRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProseKeyRequest(var5gPrukId string, relayServiceCode int32) *ProseKeyRequest {
	this := ProseKeyRequest{}
	this.Var5gPrukId = var5gPrukId
	this.RelayServiceCode = relayServiceCode
	return &this
}

// NewProseKeyRequestWithDefaults instantiates a new ProseKeyRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProseKeyRequestWithDefaults() *ProseKeyRequest {
	this := ProseKeyRequest{}
	return &this
}

// GetVar5gPrukId returns the Var5gPrukId field value
func (o *ProseKeyRequest) GetVar5gPrukId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Var5gPrukId
}

// GetVar5gPrukIdOk returns a tuple with the Var5gPrukId field value
// and a boolean to check if the value has been set.
func (o *ProseKeyRequest) GetVar5gPrukIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Var5gPrukId, true
}

// SetVar5gPrukId sets field value
func (o *ProseKeyRequest) SetVar5gPrukId(v string) {
	o.Var5gPrukId = v
}

// GetRelayServiceCode returns the RelayServiceCode field value
func (o *ProseKeyRequest) GetRelayServiceCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelayServiceCode
}

// GetRelayServiceCodeOk returns a tuple with the RelayServiceCode field value
// and a boolean to check if the value has been set.
func (o *ProseKeyRequest) GetRelayServiceCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelayServiceCode, true
}

// SetRelayServiceCode sets field value
func (o *ProseKeyRequest) SetRelayServiceCode(v int32) {
	o.RelayServiceCode = v
}

func (o ProseKeyRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProseKeyRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["5gPrukId"] = o.Var5gPrukId
	toSerialize["relayServiceCode"] = o.RelayServiceCode
	return toSerialize, nil
}

type NullableProseKeyRequest struct {
	value *ProseKeyRequest
	isSet bool
}

func (v NullableProseKeyRequest) Get() *ProseKeyRequest {
	return v.value
}

func (v *NullableProseKeyRequest) Set(val *ProseKeyRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableProseKeyRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableProseKeyRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProseKeyRequest(val *ProseKeyRequest) *NullableProseKeyRequest {
	return &NullableProseKeyRequest{value: val, isSet: true}
}

func (v NullableProseKeyRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProseKeyRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
