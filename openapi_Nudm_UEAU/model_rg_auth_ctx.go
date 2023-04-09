/*
Nudm_UEAU

UDM UE Authentication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_UEAU

import (
	"encoding/json"
)

// checks if the RgAuthCtx type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RgAuthCtx{}

// RgAuthCtx struct for RgAuthCtx
type RgAuthCtx struct {
	AuthInd bool `json:"authInd"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi *string `json:"supi,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewRgAuthCtx instantiates a new RgAuthCtx object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRgAuthCtx(authInd bool) *RgAuthCtx {
	this := RgAuthCtx{}
	this.AuthInd = authInd
	return &this
}

// NewRgAuthCtxWithDefaults instantiates a new RgAuthCtx object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRgAuthCtxWithDefaults() *RgAuthCtx {
	this := RgAuthCtx{}
	var authInd bool = false
	this.AuthInd = authInd
	return &this
}

// GetAuthInd returns the AuthInd field value
func (o *RgAuthCtx) GetAuthInd() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.AuthInd
}

// GetAuthIndOk returns a tuple with the AuthInd field value
// and a boolean to check if the value has been set.
func (o *RgAuthCtx) GetAuthIndOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AuthInd, true
}

// SetAuthInd sets field value
func (o *RgAuthCtx) SetAuthInd(v bool) {
	o.AuthInd = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *RgAuthCtx) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RgAuthCtx) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *RgAuthCtx) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *RgAuthCtx) SetSupi(v string) {
	o.Supi = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *RgAuthCtx) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RgAuthCtx) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *RgAuthCtx) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *RgAuthCtx) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o RgAuthCtx) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RgAuthCtx) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["authInd"] = o.AuthInd
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableRgAuthCtx struct {
	value *RgAuthCtx
	isSet bool
}

func (v NullableRgAuthCtx) Get() *RgAuthCtx {
	return v.value
}

func (v *NullableRgAuthCtx) Set(val *RgAuthCtx) {
	v.value = val
	v.isSet = true
}

func (v NullableRgAuthCtx) IsSet() bool {
	return v.isSet
}

func (v *NullableRgAuthCtx) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRgAuthCtx(val *RgAuthCtx) *NullableRgAuthCtx {
	return &NullableRgAuthCtx{value: val, isSet: true}
}

func (v NullableRgAuthCtx) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRgAuthCtx) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
