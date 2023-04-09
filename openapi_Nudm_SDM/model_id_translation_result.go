/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the IdTranslationResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdTranslationResult{}

// IdTranslationResult struct for IdTranslationResult
type IdTranslationResult struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.
	Gpsi            *string  `json:"gpsi,omitempty"`
	AdditionalSupis []string `json:"additionalSupis,omitempty"`
	AdditionalGpsis []string `json:"additionalGpsis,omitempty"`
}

// NewIdTranslationResult instantiates a new IdTranslationResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdTranslationResult(supi string) *IdTranslationResult {
	this := IdTranslationResult{}
	this.Supi = supi
	return &this
}

// NewIdTranslationResultWithDefaults instantiates a new IdTranslationResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdTranslationResultWithDefaults() *IdTranslationResult {
	this := IdTranslationResult{}
	return &this
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *IdTranslationResult) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdTranslationResult) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *IdTranslationResult) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *IdTranslationResult) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

// GetSupi returns the Supi field value
func (o *IdTranslationResult) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *IdTranslationResult) GetSupiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *IdTranslationResult) SetSupi(v string) {
	o.Supi = v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *IdTranslationResult) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdTranslationResult) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *IdTranslationResult) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *IdTranslationResult) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetAdditionalSupis returns the AdditionalSupis field value if set, zero value otherwise.
func (o *IdTranslationResult) GetAdditionalSupis() []string {
	if o == nil || IsNil(o.AdditionalSupis) {
		var ret []string
		return ret
	}
	return o.AdditionalSupis
}

// GetAdditionalSupisOk returns a tuple with the AdditionalSupis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdTranslationResult) GetAdditionalSupisOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalSupis) {
		return nil, false
	}
	return o.AdditionalSupis, true
}

// HasAdditionalSupis returns a boolean if a field has been set.
func (o *IdTranslationResult) HasAdditionalSupis() bool {
	if o != nil && !IsNil(o.AdditionalSupis) {
		return true
	}

	return false
}

// SetAdditionalSupis gets a reference to the given []string and assigns it to the AdditionalSupis field.
func (o *IdTranslationResult) SetAdditionalSupis(v []string) {
	o.AdditionalSupis = v
}

// GetAdditionalGpsis returns the AdditionalGpsis field value if set, zero value otherwise.
func (o *IdTranslationResult) GetAdditionalGpsis() []string {
	if o == nil || IsNil(o.AdditionalGpsis) {
		var ret []string
		return ret
	}
	return o.AdditionalGpsis
}

// GetAdditionalGpsisOk returns a tuple with the AdditionalGpsis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdTranslationResult) GetAdditionalGpsisOk() ([]string, bool) {
	if o == nil || IsNil(o.AdditionalGpsis) {
		return nil, false
	}
	return o.AdditionalGpsis, true
}

// HasAdditionalGpsis returns a boolean if a field has been set.
func (o *IdTranslationResult) HasAdditionalGpsis() bool {
	if o != nil && !IsNil(o.AdditionalGpsis) {
		return true
	}

	return false
}

// SetAdditionalGpsis gets a reference to the given []string and assigns it to the AdditionalGpsis field.
func (o *IdTranslationResult) SetAdditionalGpsis(v []string) {
	o.AdditionalGpsis = v
}

func (o IdTranslationResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdTranslationResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	toSerialize["supi"] = o.Supi
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	if !IsNil(o.AdditionalSupis) {
		toSerialize["additionalSupis"] = o.AdditionalSupis
	}
	if !IsNil(o.AdditionalGpsis) {
		toSerialize["additionalGpsis"] = o.AdditionalGpsis
	}
	return toSerialize, nil
}

type NullableIdTranslationResult struct {
	value *IdTranslationResult
	isSet bool
}

func (v NullableIdTranslationResult) Get() *IdTranslationResult {
	return v.value
}

func (v *NullableIdTranslationResult) Set(val *IdTranslationResult) {
	v.value = val
	v.isSet = true
}

func (v NullableIdTranslationResult) IsSet() bool {
	return v.isSet
}

func (v *NullableIdTranslationResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdTranslationResult(val *IdTranslationResult) *NullableIdTranslationResult {
	return &NullableIdTranslationResult{value: val, isSet: true}
}

func (v NullableIdTranslationResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdTranslationResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
