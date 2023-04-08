/*
Nnssaaf_AIW

AAA Interworking Authentication and Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnssaaf_AIW

import (
	"encoding/json"
)

// checks if the AuthConfirmationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AuthConfirmationData{}

// AuthConfirmationData struct for AuthConfirmationData
type AuthConfirmationData struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi string `json:"supi"`
	// contains an EAP packet
	EapMessage NullableString `json:"eapMessage"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewAuthConfirmationData instantiates a new AuthConfirmationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAuthConfirmationData(supi string, eapMessage NullableString) *AuthConfirmationData {
	this := AuthConfirmationData{}
	this.Supi = supi
	this.EapMessage = eapMessage
	return &this
}

// NewAuthConfirmationDataWithDefaults instantiates a new AuthConfirmationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAuthConfirmationDataWithDefaults() *AuthConfirmationData {
	this := AuthConfirmationData{}
	return &this
}

// GetSupi returns the Supi field value
func (o *AuthConfirmationData) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *AuthConfirmationData) GetSupiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *AuthConfirmationData) SetSupi(v string) {
	o.Supi = v
}

// GetEapMessage returns the EapMessage field value
// If the value is explicit nil, the zero value for string will be returned
func (o *AuthConfirmationData) GetEapMessage() string {
	if o == nil || o.EapMessage.Get() == nil {
		var ret string
		return ret
	}

	return *o.EapMessage.Get()
}

// GetEapMessageOk returns a tuple with the EapMessage field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *AuthConfirmationData) GetEapMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EapMessage.Get(), o.EapMessage.IsSet()
}

// SetEapMessage sets field value
func (o *AuthConfirmationData) SetEapMessage(v string) {
	o.EapMessage.Set(&v)
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *AuthConfirmationData) GetSupportedFeatures() string {
	if o == nil || isNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AuthConfirmationData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || isNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *AuthConfirmationData) HasSupportedFeatures() bool {
	if o != nil && !isNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *AuthConfirmationData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o AuthConfirmationData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AuthConfirmationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supi"] = o.Supi
	toSerialize["eapMessage"] = o.EapMessage.Get()
	if !isNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableAuthConfirmationData struct {
	value *AuthConfirmationData
	isSet bool
}

func (v NullableAuthConfirmationData) Get() *AuthConfirmationData {
	return v.value
}

func (v *NullableAuthConfirmationData) Set(val *AuthConfirmationData) {
	v.value = val
	v.isSet = true
}

func (v NullableAuthConfirmationData) IsSet() bool {
	return v.isSet
}

func (v *NullableAuthConfirmationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAuthConfirmationData(val *AuthConfirmationData) *NullableAuthConfirmationData {
	return &NullableAuthConfirmationData{value: val, isSet: true}
}

func (v NullableAuthConfirmationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAuthConfirmationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


