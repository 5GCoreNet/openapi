/*
Nhss_gbaUEAU

Nhss UE Authentication Service for GBA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nhss_gbaUEAU

import (
	"encoding/json"
)

// checks if the AuthenticationInfoResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthenticationInfoResult{}

// AuthenticationInfoResult Contains authentication information returned in the authentication response message (e.g. authentication vector, digest authentication parameters) 
type AuthenticationInfoResult struct {
	// IMS Private Identity of the UE
	Impi *string `json:"impi,omitempty"`
	Var3gAkaAv *Model3GAkaAv `json:"3gAkaAv,omitempty"`
	DigestAuth *DigestAuthentication `json:"digestAuth,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewAuthenticationInfoResult instantiates a new AuthenticationInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthenticationInfoResult() *AuthenticationInfoResult {
	this := AuthenticationInfoResult{}
	return &this
}

// NewAuthenticationInfoResultWithDefaults instantiates a new AuthenticationInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthenticationInfoResultWithDefaults() *AuthenticationInfoResult {
	this := AuthenticationInfoResult{}
	return &this
}

// GetImpi returns the Impi field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetImpi() string {
	if o == nil || isNil(o.Impi) {
		var ret string
		return ret
	}
	return *o.Impi
}

// GetImpiOk returns a tuple with the Impi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetImpiOk() (*string, bool) {
	if o == nil || isNil(o.Impi) {
		return nil, false
	}
	return o.Impi, true
}

// HasImpi returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasImpi() bool {
	if o != nil && !isNil(o.Impi) {
		return true
	}

	return false
}

// SetImpi gets a reference to the given string and assigns it to the Impi field.
func (o *AuthenticationInfoResult) SetImpi(v string) {
	o.Impi = &v
}

// GetVar3gAkaAv returns the Var3gAkaAv field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetVar3gAkaAv() Model3GAkaAv {
	if o == nil || isNil(o.Var3gAkaAv) {
		var ret Model3GAkaAv
		return ret
	}
	return *o.Var3gAkaAv
}

// GetVar3gAkaAvOk returns a tuple with the Var3gAkaAv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetVar3gAkaAvOk() (*Model3GAkaAv, bool) {
	if o == nil || isNil(o.Var3gAkaAv) {
		return nil, false
	}
	return o.Var3gAkaAv, true
}

// HasVar3gAkaAv returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasVar3gAkaAv() bool {
	if o != nil && !isNil(o.Var3gAkaAv) {
		return true
	}

	return false
}

// SetVar3gAkaAv gets a reference to the given Model3GAkaAv and assigns it to the Var3gAkaAv field.
func (o *AuthenticationInfoResult) SetVar3gAkaAv(v Model3GAkaAv) {
	o.Var3gAkaAv = &v
}

// GetDigestAuth returns the DigestAuth field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetDigestAuth() DigestAuthentication {
	if o == nil || isNil(o.DigestAuth) {
		var ret DigestAuthentication
		return ret
	}
	return *o.DigestAuth
}

// GetDigestAuthOk returns a tuple with the DigestAuth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetDigestAuthOk() (*DigestAuthentication, bool) {
	if o == nil || isNil(o.DigestAuth) {
		return nil, false
	}
	return o.DigestAuth, true
}

// HasDigestAuth returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasDigestAuth() bool {
	if o != nil && !isNil(o.DigestAuth) {
		return true
	}

	return false
}

// SetDigestAuth gets a reference to the given DigestAuthentication and assigns it to the DigestAuth field.
func (o *AuthenticationInfoResult) SetDigestAuth(v DigestAuthentication) {
	o.DigestAuth = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *AuthenticationInfoResult) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthenticationInfoResult) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *AuthenticationInfoResult) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *AuthenticationInfoResult) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o AuthenticationInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthenticationInfoResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Impi) {
		toSerialize["impi"] = o.Impi
	}
	if !isNil(o.Var3gAkaAv) {
		toSerialize["3gAkaAv"] = o.Var3gAkaAv
	}
	if !isNil(o.DigestAuth) {
		toSerialize["digestAuth"] = o.DigestAuth
	}
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableAuthenticationInfoResult struct {
	value *AuthenticationInfoResult
	isSet bool
}

func (v NullableAuthenticationInfoResult) Get() *AuthenticationInfoResult {
	return v.value
}

func (v *NullableAuthenticationInfoResult) Set(val *AuthenticationInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthenticationInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthenticationInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthenticationInfoResult(val *AuthenticationInfoResult) *NullableAuthenticationInfoResult {
	return &NullableAuthenticationInfoResult{value: val, isSet: true}
}

func (v NullableAuthenticationInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthenticationInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


