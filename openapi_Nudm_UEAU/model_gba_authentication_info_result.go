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

// checks if the GbaAuthenticationInfoResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GbaAuthenticationInfoResult{}

// GbaAuthenticationInfoResult struct for GbaAuthenticationInfoResult
type GbaAuthenticationInfoResult struct {
	Var3gAkaAv *Model3GAkaAv `json:"3gAkaAv,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewGbaAuthenticationInfoResult instantiates a new GbaAuthenticationInfoResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGbaAuthenticationInfoResult() *GbaAuthenticationInfoResult {
	this := GbaAuthenticationInfoResult{}
	return &this
}

// NewGbaAuthenticationInfoResultWithDefaults instantiates a new GbaAuthenticationInfoResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGbaAuthenticationInfoResultWithDefaults() *GbaAuthenticationInfoResult {
	this := GbaAuthenticationInfoResult{}
	return &this
}

// GetVar3gAkaAv returns the Var3gAkaAv field value if set, zero value otherwise.
func (o *GbaAuthenticationInfoResult) GetVar3gAkaAv() Model3GAkaAv {
	if o == nil || IsNil(o.Var3gAkaAv) {
		var ret Model3GAkaAv
		return ret
	}
	return *o.Var3gAkaAv
}

// GetVar3gAkaAvOk returns a tuple with the Var3gAkaAv field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GbaAuthenticationInfoResult) GetVar3gAkaAvOk() (*Model3GAkaAv, bool) {
	if o == nil || IsNil(o.Var3gAkaAv) {
		return nil, false
	}
	return o.Var3gAkaAv, true
}

// HasVar3gAkaAv returns a boolean if a field has been set.
func (o *GbaAuthenticationInfoResult) HasVar3gAkaAv() bool {
	if o != nil && !IsNil(o.Var3gAkaAv) {
		return true
	}

	return false
}

// SetVar3gAkaAv gets a reference to the given Model3GAkaAv and assigns it to the Var3gAkaAv field.
func (o *GbaAuthenticationInfoResult) SetVar3gAkaAv(v Model3GAkaAv) {
	o.Var3gAkaAv = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *GbaAuthenticationInfoResult) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GbaAuthenticationInfoResult) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *GbaAuthenticationInfoResult) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *GbaAuthenticationInfoResult) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o GbaAuthenticationInfoResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GbaAuthenticationInfoResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Var3gAkaAv) {
		toSerialize["3gAkaAv"] = o.Var3gAkaAv
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableGbaAuthenticationInfoResult struct {
	value *GbaAuthenticationInfoResult
	isSet bool
}

func (v NullableGbaAuthenticationInfoResult) Get() *GbaAuthenticationInfoResult {
	return v.value
}

func (v *NullableGbaAuthenticationInfoResult) Set(val *GbaAuthenticationInfoResult) {
	v.value = val
	v.isSet = true
}

func (v NullableGbaAuthenticationInfoResult) IsSet() bool {
	return v.isSet
}

func (v *NullableGbaAuthenticationInfoResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGbaAuthenticationInfoResult(val *GbaAuthenticationInfoResult) *NullableGbaAuthenticationInfoResult {
	return &NullableGbaAuthenticationInfoResult{value: val, isSet: true}
}

func (v NullableGbaAuthenticationInfoResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGbaAuthenticationInfoResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
