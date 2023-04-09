/*
NRF NFDiscovery Service

NRF NFDiscovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nnrf_NFDiscovery

import (
	"encoding/json"
)

// checks if the DefSubServiceInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DefSubServiceInfo{}

// DefSubServiceInfo Service Specific information for Default Notification Subscription.
type DefSubServiceInfo struct {
	Versions []string `json:"versions,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewDefSubServiceInfo instantiates a new DefSubServiceInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDefSubServiceInfo() *DefSubServiceInfo {
	this := DefSubServiceInfo{}
	return &this
}

// NewDefSubServiceInfoWithDefaults instantiates a new DefSubServiceInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDefSubServiceInfoWithDefaults() *DefSubServiceInfo {
	this := DefSubServiceInfo{}
	return &this
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *DefSubServiceInfo) GetVersions() []string {
	if o == nil || IsNil(o.Versions) {
		var ret []string
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefSubServiceInfo) GetVersionsOk() ([]string, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// HasVersions returns a boolean if a field has been set.
func (o *DefSubServiceInfo) HasVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []string and assigns it to the Versions field.
func (o *DefSubServiceInfo) SetVersions(v []string) {
	o.Versions = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *DefSubServiceInfo) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DefSubServiceInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *DefSubServiceInfo) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *DefSubServiceInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o DefSubServiceInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DefSubServiceInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableDefSubServiceInfo struct {
	value *DefSubServiceInfo
	isSet bool
}

func (v NullableDefSubServiceInfo) Get() *DefSubServiceInfo {
	return v.value
}

func (v *NullableDefSubServiceInfo) Set(val *DefSubServiceInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableDefSubServiceInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableDefSubServiceInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDefSubServiceInfo(val *DefSubServiceInfo) *NullableDefSubServiceInfo {
	return &NullableDefSubServiceInfo{value: val, isSet: true}
}

func (v NullableDefSubServiceInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDefSubServiceInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
