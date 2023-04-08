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

// checks if the ProSeAuthenticationCtx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProSeAuthenticationCtx{}

// ProSeAuthenticationCtx Contains the information related to the resource generated to handle the ProSe authentication.
type ProSeAuthenticationCtx struct {
	AuthType AuthType `json:"authType"`
	// A map(list of key-value pairs) where member serves as key
	Links map[string]LinksValueSchema `json:"_links"`
	ProSeAuthData ProSeAuthData `json:"proSeAuthData"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewProSeAuthenticationCtx instantiates a new ProSeAuthenticationCtx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProSeAuthenticationCtx(authType AuthType, links map[string]LinksValueSchema, proSeAuthData ProSeAuthData) *ProSeAuthenticationCtx {
	this := ProSeAuthenticationCtx{}
	this.AuthType = authType
	this.Links = links
	this.ProSeAuthData = proSeAuthData
	return &this
}

// NewProSeAuthenticationCtxWithDefaults instantiates a new ProSeAuthenticationCtx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProSeAuthenticationCtxWithDefaults() *ProSeAuthenticationCtx {
	this := ProSeAuthenticationCtx{}
	return &this
}

// GetAuthType returns the AuthType field value
func (o *ProSeAuthenticationCtx) GetAuthType() AuthType {
	if o == nil {
		var ret AuthType
		return ret
	}

	return o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value
// and a boolean to check if the value has been set.
func (o *ProSeAuthenticationCtx) GetAuthTypeOk() (*AuthType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthType, true
}

// SetAuthType sets field value
func (o *ProSeAuthenticationCtx) SetAuthType(v AuthType) {
	o.AuthType = v
}

// GetLinks returns the Links field value
func (o *ProSeAuthenticationCtx) GetLinks() map[string]LinksValueSchema {
	if o == nil {
		var ret map[string]LinksValueSchema
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ProSeAuthenticationCtx) GetLinksOk() (*map[string]LinksValueSchema, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ProSeAuthenticationCtx) SetLinks(v map[string]LinksValueSchema) {
	o.Links = v
}

// GetProSeAuthData returns the ProSeAuthData field value
func (o *ProSeAuthenticationCtx) GetProSeAuthData() ProSeAuthData {
	if o == nil {
		var ret ProSeAuthData
		return ret
	}

	return o.ProSeAuthData
}

// GetProSeAuthDataOk returns a tuple with the ProSeAuthData field value
// and a boolean to check if the value has been set.
func (o *ProSeAuthenticationCtx) GetProSeAuthDataOk() (*ProSeAuthData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProSeAuthData, true
}

// SetProSeAuthData sets field value
func (o *ProSeAuthenticationCtx) SetProSeAuthData(v ProSeAuthData) {
	o.ProSeAuthData = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *ProSeAuthenticationCtx) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProSeAuthenticationCtx) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *ProSeAuthenticationCtx) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *ProSeAuthenticationCtx) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o ProSeAuthenticationCtx) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProSeAuthenticationCtx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authType"] = o.AuthType
	toSerialize["_links"] = o.Links
	toSerialize["proSeAuthData"] = o.ProSeAuthData
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableProSeAuthenticationCtx struct {
	value *ProSeAuthenticationCtx
	isSet bool
}

func (v NullableProSeAuthenticationCtx) Get() *ProSeAuthenticationCtx {
	return v.value
}

func (v *NullableProSeAuthenticationCtx) Set(val *ProSeAuthenticationCtx) {
	v.value = val
	v.isSet = true
}

func (v NullableProSeAuthenticationCtx) IsSet() bool {
	return v.isSet
}

func (v *NullableProSeAuthenticationCtx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProSeAuthenticationCtx(val *ProSeAuthenticationCtx) *NullableProSeAuthenticationCtx {
	return &NullableProSeAuthenticationCtx{value: val, isSet: true}
}

func (v NullableProSeAuthenticationCtx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProSeAuthenticationCtx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


