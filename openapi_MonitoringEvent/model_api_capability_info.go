/*
3gpp-monitoring-event

API for Monitoring Event.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MonitoringEvent

import (
	"encoding/json"
)

// checks if the ApiCapabilityInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiCapabilityInfo{}

// ApiCapabilityInfo Represents the availability information of supported API.
type ApiCapabilityInfo struct {
	ApiName string `json:"apiName"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat string `json:"suppFeat"`
}

// NewApiCapabilityInfo instantiates a new ApiCapabilityInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiCapabilityInfo(apiName string, suppFeat string) *ApiCapabilityInfo {
	this := ApiCapabilityInfo{}
	this.ApiName = apiName
	this.SuppFeat = suppFeat
	return &this
}

// NewApiCapabilityInfoWithDefaults instantiates a new ApiCapabilityInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiCapabilityInfoWithDefaults() *ApiCapabilityInfo {
	this := ApiCapabilityInfo{}
	return &this
}

// GetApiName returns the ApiName field value
func (o *ApiCapabilityInfo) GetApiName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ApiName
}

// GetApiNameOk returns a tuple with the ApiName field value
// and a boolean to check if the value has been set.
func (o *ApiCapabilityInfo) GetApiNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ApiName, true
}

// SetApiName sets field value
func (o *ApiCapabilityInfo) SetApiName(v string) {
	o.ApiName = v
}

// GetSuppFeat returns the SuppFeat field value
func (o *ApiCapabilityInfo) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *ApiCapabilityInfo) GetSuppFeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *ApiCapabilityInfo) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o ApiCapabilityInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiCapabilityInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiName"] = o.ApiName
	toSerialize["suppFeat"] = o.SuppFeat
	return toSerialize, nil
}

type NullableApiCapabilityInfo struct {
	value *ApiCapabilityInfo
	isSet bool
}

func (v NullableApiCapabilityInfo) Get() *ApiCapabilityInfo {
	return v.value
}

func (v *NullableApiCapabilityInfo) Set(val *ApiCapabilityInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableApiCapabilityInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableApiCapabilityInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiCapabilityInfo(val *ApiCapabilityInfo) *NullableApiCapabilityInfo {
	return &NullableApiCapabilityInfo{value: val, isSet: true}
}

func (v NullableApiCapabilityInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiCapabilityInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
