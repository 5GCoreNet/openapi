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

// checks if the ProSeEapSession type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProSeEapSession{}

// ProSeEapSession Contains information related to the EAP session. If present the 5gPrukId IE shall carry the CP-PRUK ID.
type ProSeEapSession struct {
	// contains an EAP packet
	EapPayload NullableString `json:"eapPayload"`
	// Contains the KNR_ProSe.
	KnrProSe *string `json:"knrProSe,omitempty"`
	// A map(list of key-value pairs) where member serves as key
	Links *map[string]LinksValueSchema `json:"_links,omitempty"`
	AuthResult *AuthResult `json:"authResult,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// contains an Nonce2
	Nonce2 NullableString `json:"nonce2,omitempty"`
	// A string carrying the CP-PRUK ID of the remote UE. The CP-PRUK ID is a string in NAI format as specified in clause 28.7.11 of 3GPP TS 23.003. 
	Var5gPrukId *string `json:"5gPrukId,omitempty"`
}

// NewProSeEapSession instantiates a new ProSeEapSession object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProSeEapSession(eapPayload NullableString) *ProSeEapSession {
	this := ProSeEapSession{}
	this.EapPayload = eapPayload
	return &this
}

// NewProSeEapSessionWithDefaults instantiates a new ProSeEapSession object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProSeEapSessionWithDefaults() *ProSeEapSession {
	this := ProSeEapSession{}
	return &this
}

// GetEapPayload returns the EapPayload field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ProSeEapSession) GetEapPayload() string {
	if o == nil || o.EapPayload.Get() == nil {
		var ret string
		return ret
	}

	return *o.EapPayload.Get()
}

// GetEapPayloadOk returns a tuple with the EapPayload field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProSeEapSession) GetEapPayloadOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EapPayload.Get(), o.EapPayload.IsSet()
}

// SetEapPayload sets field value
func (o *ProSeEapSession) SetEapPayload(v string) {
	o.EapPayload.Set(&v)
}

// GetKnrProSe returns the KnrProSe field value if set, zero value otherwise.
func (o *ProSeEapSession) GetKnrProSe() string {
	if o == nil || isNil(o.KnrProSe) {
		var ret string
		return ret
	}
	return *o.KnrProSe
}

// GetKnrProSeOk returns a tuple with the KnrProSe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeEapSession) GetKnrProSeOk() (*string, bool) {
	if o == nil || isNil(o.KnrProSe) {
		return nil, false
	}
	return o.KnrProSe, true
}

// HasKnrProSe returns a boolean if a field has been set.
func (o *ProSeEapSession) HasKnrProSe() bool {
	if o != nil && !isNil(o.KnrProSe) {
		return true
	}

	return false
}

// SetKnrProSe gets a reference to the given string and assigns it to the KnrProSe field.
func (o *ProSeEapSession) SetKnrProSe(v string) {
	o.KnrProSe = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *ProSeEapSession) GetLinks() map[string]LinksValueSchema {
	if o == nil || isNil(o.Links) {
		var ret map[string]LinksValueSchema
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeEapSession) GetLinksOk() (*map[string]LinksValueSchema, bool) {
	if o == nil || isNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *ProSeEapSession) HasLinks() bool {
	if o != nil && !isNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]LinksValueSchema and assigns it to the Links field.
func (o *ProSeEapSession) SetLinks(v map[string]LinksValueSchema) {
	o.Links = &v
}

// GetAuthResult returns the AuthResult field value if set, zero value otherwise.
func (o *ProSeEapSession) GetAuthResult() AuthResult {
	if o == nil || isNil(o.AuthResult) {
		var ret AuthResult
		return ret
	}
	return *o.AuthResult
}

// GetAuthResultOk returns a tuple with the AuthResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeEapSession) GetAuthResultOk() (*AuthResult, bool) {
	if o == nil || isNil(o.AuthResult) {
		return nil, false
	}
	return o.AuthResult, true
}

// HasAuthResult returns a boolean if a field has been set.
func (o *ProSeEapSession) HasAuthResult() bool {
	if o != nil && !isNil(o.AuthResult) {
		return true
	}

	return false
}

// SetAuthResult gets a reference to the given AuthResult and assigns it to the AuthResult field.
func (o *ProSeEapSession) SetAuthResult(v AuthResult) {
	o.AuthResult = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ProSeEapSession) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeEapSession) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ProSeEapSession) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ProSeEapSession) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetNonce2 returns the Nonce2 field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *ProSeEapSession) GetNonce2() string {
	if o == nil || isNil(o.Nonce2.Get()) {
		var ret string
		return ret
	}
	return *o.Nonce2.Get()
}

// GetNonce2Ok returns a tuple with the Nonce2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ProSeEapSession) GetNonce2Ok() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Nonce2.Get(), o.Nonce2.IsSet()
}

// HasNonce2 returns a boolean if a field has been set.
func (o *ProSeEapSession) HasNonce2() bool {
	if o != nil && o.Nonce2.IsSet() {
		return true
	}

	return false
}

// SetNonce2 gets a reference to the given NullableString and assigns it to the Nonce2 field.
func (o *ProSeEapSession) SetNonce2(v string) {
	o.Nonce2.Set(&v)
}
// SetNonce2Nil sets the value for Nonce2 to be an explicit nil
func (o *ProSeEapSession) SetNonce2Nil() {
	o.Nonce2.Set(nil)
}

// UnsetNonce2 ensures that no value is present for Nonce2, not even an explicit nil
func (o *ProSeEapSession) UnsetNonce2() {
	o.Nonce2.Unset()
}

// GetVar5gPrukId returns the Var5gPrukId field value if set, zero value otherwise.
func (o *ProSeEapSession) GetVar5gPrukId() string {
	if o == nil || isNil(o.Var5gPrukId) {
		var ret string
		return ret
	}
	return *o.Var5gPrukId
}

// GetVar5gPrukIdOk returns a tuple with the Var5gPrukId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeEapSession) GetVar5gPrukIdOk() (*string, bool) {
	if o == nil || isNil(o.Var5gPrukId) {
		return nil, false
	}
	return o.Var5gPrukId, true
}

// HasVar5gPrukId returns a boolean if a field has been set.
func (o *ProSeEapSession) HasVar5gPrukId() bool {
	if o != nil && !isNil(o.Var5gPrukId) {
		return true
	}

	return false
}

// SetVar5gPrukId gets a reference to the given string and assigns it to the Var5gPrukId field.
func (o *ProSeEapSession) SetVar5gPrukId(v string) {
	o.Var5gPrukId = &v
}

func (o ProSeEapSession) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProSeEapSession) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["eapPayload"] = o.EapPayload.Get()
	if !isNil(o.KnrProSe) {
		toSerialize["knrProSe"] = o.KnrProSe
	}
	if !isNil(o.Links) {
		toSerialize["_links"] = o.Links
	}
	if !isNil(o.AuthResult) {
		toSerialize["authResult"] = o.AuthResult
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	if o.Nonce2.IsSet() {
		toSerialize["nonce2"] = o.Nonce2.Get()
	}
	if !isNil(o.Var5gPrukId) {
		toSerialize["5gPrukId"] = o.Var5gPrukId
	}
	return toSerialize, nil
}

type NullableProSeEapSession struct {
	value *ProSeEapSession
	isSet bool
}

func (v NullableProSeEapSession) Get() *ProSeEapSession {
	return v.value
}

func (v *NullableProSeEapSession) Set(val *ProSeEapSession) {
	v.value = val
	v.isSet = true
}

func (v NullableProSeEapSession) IsSet() bool {
	return v.isSet
}

func (v *NullableProSeEapSession) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProSeEapSession(val *ProSeEapSession) *NullableProSeEapSession {
	return &NullableProSeEapSession{value: val, isSet: true}
}

func (v NullableProSeEapSession) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProSeEapSession) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


