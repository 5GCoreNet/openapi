/*
Namf_Location

AMF Location Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Location

import (
	"encoding/json"
)

// checks if the CancelPosInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CancelPosInfo{}

// CancelPosInfo Data within a Cancel Location Request
type CancelPosInfo struct {
	// String identifying a Supi that shall contain either an IMSI, a network specific identifier, a Global Cable Identifier (GCI) or a Global Line Identifier (GLI) as specified in clause  2.2A of 3GPP TS 23.003. It shall be formatted as follows  - for an IMSI \"imsi-<imsi>\", where <imsi> shall be formatted according to clause 2.2    of 3GPP TS 23.003 that describes an IMSI.  - for a network specific identifier \"nai-<nai>, where <nai> shall be formatted    according to clause 28.7.2 of 3GPP TS 23.003 that describes an NAI.  - for a GCI \"gci-<gci>\", where <gci> shall be formatted according to clause 28.15.2    of 3GPP TS 23.003.  - for a GLI \"gli-<gli>\", where <gli> shall be formatted according to clause 28.16.2 of    3GPP TS 23.003.To enable that the value is used as part of an URI, the string shall    only contain characters allowed according to the \"lower-with-hyphen\" naming convention    defined in 3GPP TS 29.501.
	Supi string `json:"supi"`
	// String providing an URI formatted according to RFC 3986.
	HgmlcCallBackURI string `json:"hgmlcCallBackURI"`
	// LDR Reference.
	LdrReference string `json:"ldrReference"`
	// LMF identification.
	ServingLMFIdentification *string `json:"servingLMFIdentification,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewCancelPosInfo instantiates a new CancelPosInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelPosInfo(supi string, hgmlcCallBackURI string, ldrReference string) *CancelPosInfo {
	this := CancelPosInfo{}
	this.Supi = supi
	this.HgmlcCallBackURI = hgmlcCallBackURI
	this.LdrReference = ldrReference
	return &this
}

// NewCancelPosInfoWithDefaults instantiates a new CancelPosInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelPosInfoWithDefaults() *CancelPosInfo {
	this := CancelPosInfo{}
	return &this
}

// GetSupi returns the Supi field value
func (o *CancelPosInfo) GetSupi() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Supi
}

// GetSupiOk returns a tuple with the Supi field value
// and a boolean to check if the value has been set.
func (o *CancelPosInfo) GetSupiOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Supi, true
}

// SetSupi sets field value
func (o *CancelPosInfo) SetSupi(v string) {
	o.Supi = v
}

// GetHgmlcCallBackURI returns the HgmlcCallBackURI field value
func (o *CancelPosInfo) GetHgmlcCallBackURI() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HgmlcCallBackURI
}

// GetHgmlcCallBackURIOk returns a tuple with the HgmlcCallBackURI field value
// and a boolean to check if the value has been set.
func (o *CancelPosInfo) GetHgmlcCallBackURIOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HgmlcCallBackURI, true
}

// SetHgmlcCallBackURI sets field value
func (o *CancelPosInfo) SetHgmlcCallBackURI(v string) {
	o.HgmlcCallBackURI = v
}

// GetLdrReference returns the LdrReference field value
func (o *CancelPosInfo) GetLdrReference() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LdrReference
}

// GetLdrReferenceOk returns a tuple with the LdrReference field value
// and a boolean to check if the value has been set.
func (o *CancelPosInfo) GetLdrReferenceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LdrReference, true
}

// SetLdrReference sets field value
func (o *CancelPosInfo) SetLdrReference(v string) {
	o.LdrReference = v
}

// GetServingLMFIdentification returns the ServingLMFIdentification field value if set, zero value otherwise.
func (o *CancelPosInfo) GetServingLMFIdentification() string {
	if o == nil || IsNil(o.ServingLMFIdentification) {
		var ret string
		return ret
	}
	return *o.ServingLMFIdentification
}

// GetServingLMFIdentificationOk returns a tuple with the ServingLMFIdentification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelPosInfo) GetServingLMFIdentificationOk() (*string, bool) {
	if o == nil || IsNil(o.ServingLMFIdentification) {
		return nil, false
	}
	return o.ServingLMFIdentification, true
}

// HasServingLMFIdentification returns a boolean if a field has been set.
func (o *CancelPosInfo) HasServingLMFIdentification() bool {
	if o != nil && !IsNil(o.ServingLMFIdentification) {
		return true
	}

	return false
}

// SetServingLMFIdentification gets a reference to the given string and assigns it to the ServingLMFIdentification field.
func (o *CancelPosInfo) SetServingLMFIdentification(v string) {
	o.ServingLMFIdentification = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *CancelPosInfo) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelPosInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *CancelPosInfo) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *CancelPosInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o CancelPosInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CancelPosInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["supi"] = o.Supi
	toSerialize["hgmlcCallBackURI"] = o.HgmlcCallBackURI
	toSerialize["ldrReference"] = o.LdrReference
	if !IsNil(o.ServingLMFIdentification) {
		toSerialize["servingLMFIdentification"] = o.ServingLMFIdentification
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableCancelPosInfo struct {
	value *CancelPosInfo
	isSet bool
}

func (v NullableCancelPosInfo) Get() *CancelPosInfo {
	return v.value
}

func (v *NullableCancelPosInfo) Set(val *CancelPosInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelPosInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelPosInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelPosInfo(val *CancelPosInfo) *NullableCancelPosInfo {
	return &NullableCancelPosInfo{value: val, isSet: true}
}

func (v NullableCancelPosInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelPosInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
