/*
3gpp-akma

API for AKMA.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.0.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_AKMA

import (
	"encoding/json"
	"time"
)

// checks if the AkmaAfKeyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AkmaAfKeyData{}

// AkmaAfKeyData Represents AKMA Application Key information data.
type AkmaAfKeyData struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
	// String identifying a Gpsi shall contain either an External Id or an MSISDN.  It shall be formatted as follows -External Identifier= \"extid-'extid', where 'extid'  shall be formatted according to clause 19.7.2 of 3GPP TS 23.003 that describes an  External Identifier.  
	Gpsi *string `json:"gpsi,omitempty"`
	// string with format \"date-time\" as defined in OpenAPI.
	Expiry time.Time `json:"expiry"`
	Kaf string `json:"kaf"`
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501. 
	Supi *string `json:"supi,omitempty"`
}

// NewAkmaAfKeyData instantiates a new AkmaAfKeyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAkmaAfKeyData(expiry time.Time, kaf string) *AkmaAfKeyData {
	this := AkmaAfKeyData{}
	this.Expiry = expiry
	this.Kaf = kaf
	return &this
}

// NewAkmaAfKeyDataWithDefaults instantiates a new AkmaAfKeyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAkmaAfKeyDataWithDefaults() *AkmaAfKeyData {
	this := AkmaAfKeyData{}
	return &this
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *AkmaAfKeyData) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyData) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *AkmaAfKeyData) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *AkmaAfKeyData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

// GetGpsi returns the Gpsi field value if set, zero value otherwise.
func (o *AkmaAfKeyData) GetGpsi() string {
	if o == nil || IsNil(o.Gpsi) {
		var ret string
		return ret
	}
	return *o.Gpsi
}

// GetGpsiOk returns a tuple with the Gpsi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyData) GetGpsiOk() (*string, bool) {
	if o == nil || IsNil(o.Gpsi) {
		return nil, false
	}
	return o.Gpsi, true
}

// HasGpsi returns a boolean if a field has been set.
func (o *AkmaAfKeyData) HasGpsi() bool {
	if o != nil && !IsNil(o.Gpsi) {
		return true
	}

	return false
}

// SetGpsi gets a reference to the given string and assigns it to the Gpsi field.
func (o *AkmaAfKeyData) SetGpsi(v string) {
	o.Gpsi = &v
}

// GetExpiry returns the Expiry field value
func (o *AkmaAfKeyData) GetExpiry() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Expiry
}

// GetExpiryOk returns a tuple with the Expiry field value
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyData) GetExpiryOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Expiry, true
}

// SetExpiry sets field value
func (o *AkmaAfKeyData) SetExpiry(v time.Time) {
	o.Expiry = v
}

// GetKaf returns the Kaf field value
func (o *AkmaAfKeyData) GetKaf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Kaf
}

// GetKafOk returns a tuple with the Kaf field value
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyData) GetKafOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kaf, true
}

// SetKaf sets field value
func (o *AkmaAfKeyData) SetKaf(v string) {
	o.Kaf = v
}

// GetSupi returns the Supi field value if set, zero value otherwise.
func (o *AkmaAfKeyData) GetSupi() string {
	if o == nil || IsNil(o.Supi) {
		var ret string
		return ret
	}
	return *o.Supi
}

// GetSupiOk returns a tuple with the Supi field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AkmaAfKeyData) GetSupiOk() (*string, bool) {
	if o == nil || IsNil(o.Supi) {
		return nil, false
	}
	return o.Supi, true
}

// HasSupi returns a boolean if a field has been set.
func (o *AkmaAfKeyData) HasSupi() bool {
	if o != nil && !IsNil(o.Supi) {
		return true
	}

	return false
}

// SetSupi gets a reference to the given string and assigns it to the Supi field.
func (o *AkmaAfKeyData) SetSupi(v string) {
	o.Supi = &v
}

func (o AkmaAfKeyData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AkmaAfKeyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	if !IsNil(o.Gpsi) {
		toSerialize["gpsi"] = o.Gpsi
	}
	toSerialize["expiry"] = o.Expiry
	toSerialize["kaf"] = o.Kaf
	if !IsNil(o.Supi) {
		toSerialize["supi"] = o.Supi
	}
	return toSerialize, nil
}

type NullableAkmaAfKeyData struct {
	value *AkmaAfKeyData
	isSet bool
}

func (v NullableAkmaAfKeyData) Get() *AkmaAfKeyData {
	return v.value
}

func (v *NullableAkmaAfKeyData) Set(val *AkmaAfKeyData) {
	v.value = val
	v.isSet = true
}

func (v NullableAkmaAfKeyData) IsSet() bool {
	return v.isSet
}

func (v *NullableAkmaAfKeyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAkmaAfKeyData(val *AkmaAfKeyData) *NullableAkmaAfKeyData {
	return &NullableAkmaAfKeyData{value: val, isSet: true}
}

func (v NullableAkmaAfKeyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAkmaAfKeyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


