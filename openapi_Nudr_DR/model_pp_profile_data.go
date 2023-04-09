/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the PpProfileData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PpProfileData{}

// PpProfileData struct for PpProfileData
type PpProfileData struct {
	// A map (list of key-value pairs where PpDataType serves as key) of AllowedMtcProviderInfo lists. In addition to defined PpDataType, the key value \"ALL\" may be used to identify a map entry which contains a list of AllowedMtcProviderInfo that are allowed to provision all types of the PP data for the user using UDM ParameterProvision service.
	AllowedMtcProviders *map[string][]AllowedMtcProviderInfo `json:"allowedMtcProviders,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewPpProfileData instantiates a new PpProfileData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPpProfileData() *PpProfileData {
	this := PpProfileData{}
	return &this
}

// NewPpProfileDataWithDefaults instantiates a new PpProfileData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPpProfileDataWithDefaults() *PpProfileData {
	this := PpProfileData{}
	return &this
}

// GetAllowedMtcProviders returns the AllowedMtcProviders field value if set, zero value otherwise.
func (o *PpProfileData) GetAllowedMtcProviders() map[string][]AllowedMtcProviderInfo {
	if o == nil || IsNil(o.AllowedMtcProviders) {
		var ret map[string][]AllowedMtcProviderInfo
		return ret
	}
	return *o.AllowedMtcProviders
}

// GetAllowedMtcProvidersOk returns a tuple with the AllowedMtcProviders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpProfileData) GetAllowedMtcProvidersOk() (*map[string][]AllowedMtcProviderInfo, bool) {
	if o == nil || IsNil(o.AllowedMtcProviders) {
		return nil, false
	}
	return o.AllowedMtcProviders, true
}

// HasAllowedMtcProviders returns a boolean if a field has been set.
func (o *PpProfileData) HasAllowedMtcProviders() bool {
	if o != nil && !IsNil(o.AllowedMtcProviders) {
		return true
	}

	return false
}

// SetAllowedMtcProviders gets a reference to the given map[string][]AllowedMtcProviderInfo and assigns it to the AllowedMtcProviders field.
func (o *PpProfileData) SetAllowedMtcProviders(v map[string][]AllowedMtcProviderInfo) {
	o.AllowedMtcProviders = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *PpProfileData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PpProfileData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *PpProfileData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *PpProfileData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o PpProfileData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PpProfileData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowedMtcProviders) {
		toSerialize["allowedMtcProviders"] = o.AllowedMtcProviders
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullablePpProfileData struct {
	value *PpProfileData
	isSet bool
}

func (v NullablePpProfileData) Get() *PpProfileData {
	return v.value
}

func (v *NullablePpProfileData) Set(val *PpProfileData) {
	v.value = val
	v.isSet = true
}

func (v NullablePpProfileData) IsSet() bool {
	return v.isSet
}

func (v *NullablePpProfileData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePpProfileData(val *PpProfileData) *NullablePpProfileData {
	return &NullablePpProfileData{value: val, isSet: true}
}

func (v NullablePpProfileData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePpProfileData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
