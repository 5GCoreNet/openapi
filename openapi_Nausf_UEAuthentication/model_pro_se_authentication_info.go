/*
AUSF API

AUSF UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nausf_UEAuthentication

import (
	"encoding/json"
)

// checks if the ProSeAuthenticationInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProSeAuthenticationInfo{}

// ProSeAuthenticationInfo Contains the UE id (i.e. SUCI) or CP-PRUK ID, Relay Service Code and Nonce_1.
type ProSeAuthenticationInfo struct {
	// String identifying a SUPI or a SUCI.
	SupiOrSuci *string `json:"supiOrSuci,omitempty"`
	// A string carrying the CP-PRUK ID of the remote UE. The CP-PRUK ID is a string in NAI format as specified in clause 28.7.11 of 3GPP TS 23.003.
	Var5gPrukId *string `json:"5gPrukId,omitempty"`
	// Relay Service Code to identify a connectivity service provided by the UE-to-Network relay.
	RelayServiceCode int32 `json:"relayServiceCode"`
	// contains an Nonce1
	Nonce1 NullableString `json:"nonce1"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewProSeAuthenticationInfo instantiates a new ProSeAuthenticationInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProSeAuthenticationInfo(relayServiceCode int32, nonce1 NullableString) *ProSeAuthenticationInfo {
	this := ProSeAuthenticationInfo{}
	this.RelayServiceCode = relayServiceCode
	this.Nonce1 = nonce1
	return &this
}

// NewProSeAuthenticationInfoWithDefaults instantiates a new ProSeAuthenticationInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProSeAuthenticationInfoWithDefaults() *ProSeAuthenticationInfo {
	this := ProSeAuthenticationInfo{}
	return &this
}

// GetSupiOrSuci returns the SupiOrSuci field value if set, zero value otherwise.
func (o *ProSeAuthenticationInfo) GetSupiOrSuci() string {
	if o == nil || IsNil(o.SupiOrSuci) {
		var ret string
		return ret
	}
	return *o.SupiOrSuci
}

// GetSupiOrSuciOk returns a tuple with the SupiOrSuci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeAuthenticationInfo) GetSupiOrSuciOk() (*string, bool) {
	if o == nil || IsNil(o.SupiOrSuci) {
		return nil, false
	}
	return o.SupiOrSuci, true
}

// HasSupiOrSuci returns a boolean if a field has been set.
func (o *ProSeAuthenticationInfo) HasSupiOrSuci() bool {
	if o != nil && !IsNil(o.SupiOrSuci) {
		return true
	}

	return false
}

// SetSupiOrSuci gets a reference to the given string and assigns it to the SupiOrSuci field.
func (o *ProSeAuthenticationInfo) SetSupiOrSuci(v string) {
	o.SupiOrSuci = &v
}

// GetVar5gPrukId returns the Var5gPrukId field value if set, zero value otherwise.
func (o *ProSeAuthenticationInfo) GetVar5gPrukId() string {
	if o == nil || IsNil(o.Var5gPrukId) {
		var ret string
		return ret
	}
	return *o.Var5gPrukId
}

// GetVar5gPrukIdOk returns a tuple with the Var5gPrukId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeAuthenticationInfo) GetVar5gPrukIdOk() (*string, bool) {
	if o == nil || IsNil(o.Var5gPrukId) {
		return nil, false
	}
	return o.Var5gPrukId, true
}

// HasVar5gPrukId returns a boolean if a field has been set.
func (o *ProSeAuthenticationInfo) HasVar5gPrukId() bool {
	if o != nil && !IsNil(o.Var5gPrukId) {
		return true
	}

	return false
}

// SetVar5gPrukId gets a reference to the given string and assigns it to the Var5gPrukId field.
func (o *ProSeAuthenticationInfo) SetVar5gPrukId(v string) {
	o.Var5gPrukId = &v
}

// GetRelayServiceCode returns the RelayServiceCode field value
func (o *ProSeAuthenticationInfo) GetRelayServiceCode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.RelayServiceCode
}

// GetRelayServiceCodeOk returns a tuple with the RelayServiceCode field value
// and a boolean to check if the value has been set.
func (o *ProSeAuthenticationInfo) GetRelayServiceCodeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RelayServiceCode, true
}

// SetRelayServiceCode sets field value
func (o *ProSeAuthenticationInfo) SetRelayServiceCode(v int32) {
	o.RelayServiceCode = v
}

// GetNonce1 returns the Nonce1 field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProSeAuthenticationInfo) GetNonce1() string {
	if o == nil || o.Nonce1.Get() == nil {
		var ret string
		return ret
	}

	return *o.Nonce1.Get()
}

// GetNonce1Ok returns a tuple with the Nonce1 field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProSeAuthenticationInfo) GetNonce1Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nonce1.Get(), o.Nonce1.IsSet()
}

// SetNonce1 sets field value
func (o *ProSeAuthenticationInfo) SetNonce1(v string) {
	o.Nonce1.Set(&v)
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ProSeAuthenticationInfo) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeAuthenticationInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ProSeAuthenticationInfo) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ProSeAuthenticationInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o ProSeAuthenticationInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProSeAuthenticationInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupiOrSuci) {
		toSerialize["supiOrSuci"] = o.SupiOrSuci
	}
	if !IsNil(o.Var5gPrukId) {
		toSerialize["5gPrukId"] = o.Var5gPrukId
	}
	toSerialize["relayServiceCode"] = o.RelayServiceCode
	toSerialize["nonce1"] = o.Nonce1.Get()
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableProSeAuthenticationInfo struct {
	value *ProSeAuthenticationInfo
	isSet bool
}

func (v NullableProSeAuthenticationInfo) Get() *ProSeAuthenticationInfo {
	return v.value
}

func (v *NullableProSeAuthenticationInfo) Set(val *ProSeAuthenticationInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProSeAuthenticationInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProSeAuthenticationInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProSeAuthenticationInfo(val *ProSeAuthenticationInfo) *NullableProSeAuthenticationInfo {
	return &NullableProSeAuthenticationInfo{value: val, isSet: true}
}

func (v NullableProSeAuthenticationInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProSeAuthenticationInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
