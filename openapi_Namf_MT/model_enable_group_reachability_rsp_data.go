/*
Namf_MT

AMF Mobile Terminated Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_MT

import (
	"encoding/json"
)

// checks if the EnableGroupReachabilityRspData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EnableGroupReachabilityRspData{}

// EnableGroupReachabilityRspData Data within the Enable Group Reachability Response
type EnableGroupReachabilityRspData struct {
	UeConnectedList []string `json:"ueConnectedList,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewEnableGroupReachabilityRspData instantiates a new EnableGroupReachabilityRspData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnableGroupReachabilityRspData() *EnableGroupReachabilityRspData {
	this := EnableGroupReachabilityRspData{}
	return &this
}

// NewEnableGroupReachabilityRspDataWithDefaults instantiates a new EnableGroupReachabilityRspData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnableGroupReachabilityRspDataWithDefaults() *EnableGroupReachabilityRspData {
	this := EnableGroupReachabilityRspData{}
	return &this
}

// GetUeConnectedList returns the UeConnectedList field value if set, zero value otherwise.
func (o *EnableGroupReachabilityRspData) GetUeConnectedList() []string {
	if o == nil || IsNil(o.UeConnectedList) {
		var ret []string
		return ret
	}
	return o.UeConnectedList
}

// GetUeConnectedListOk returns a tuple with the UeConnectedList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableGroupReachabilityRspData) GetUeConnectedListOk() ([]string, bool) {
	if o == nil || IsNil(o.UeConnectedList) {
		return nil, false
	}
	return o.UeConnectedList, true
}

// HasUeConnectedList returns a boolean if a field has been set.
func (o *EnableGroupReachabilityRspData) HasUeConnectedList() bool {
	if o != nil && !IsNil(o.UeConnectedList) {
		return true
	}

	return false
}

// SetUeConnectedList gets a reference to the given []string and assigns it to the UeConnectedList field.
func (o *EnableGroupReachabilityRspData) SetUeConnectedList(v []string) {
	o.UeConnectedList = v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *EnableGroupReachabilityRspData) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnableGroupReachabilityRspData) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *EnableGroupReachabilityRspData) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *EnableGroupReachabilityRspData) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o EnableGroupReachabilityRspData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EnableGroupReachabilityRspData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UeConnectedList) {
		toSerialize["ueConnectedList"] = o.UeConnectedList
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableEnableGroupReachabilityRspData struct {
	value *EnableGroupReachabilityRspData
	isSet bool
}

func (v NullableEnableGroupReachabilityRspData) Get() *EnableGroupReachabilityRspData {
	return v.value
}

func (v *NullableEnableGroupReachabilityRspData) Set(val *EnableGroupReachabilityRspData) {
	v.value = val
	v.isSet = true
}

func (v NullableEnableGroupReachabilityRspData) IsSet() bool {
	return v.isSet
}

func (v *NullableEnableGroupReachabilityRspData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnableGroupReachabilityRspData(val *EnableGroupReachabilityRspData) *NullableEnableGroupReachabilityRspData {
	return &NullableEnableGroupReachabilityRspData{value: val, isSet: true}
}

func (v NullableEnableGroupReachabilityRspData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnableGroupReachabilityRspData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
