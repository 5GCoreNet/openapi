/*
Nudm_MT

UDM MT Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_MT

import (
	"encoding/json"
	"time"
)

// checks if the UeContextInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeContextInfo{}

// UeContextInfo UE Context Information
type UeContextInfo struct {
	SupportVoPS      *bool `json:"supportVoPS,omitempty"`
	SupportVoPSn3gpp *bool `json:"supportVoPSn3gpp,omitempty"`
	// string with format 'date-time' as defined in OpenAPI.
	LastActTime *time.Time  `json:"lastActTime,omitempty"`
	AccessType  *AccessType `json:"accessType,omitempty"`
	RatType     *RatType    `json:"ratType,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SupportedFeatures *string `json:"supportedFeatures,omitempty"`
}

// NewUeContextInfo instantiates a new UeContextInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeContextInfo() *UeContextInfo {
	this := UeContextInfo{}
	return &this
}

// NewUeContextInfoWithDefaults instantiates a new UeContextInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeContextInfoWithDefaults() *UeContextInfo {
	this := UeContextInfo{}
	return &this
}

// GetSupportVoPS returns the SupportVoPS field value if set, zero value otherwise.
func (o *UeContextInfo) GetSupportVoPS() bool {
	if o == nil || IsNil(o.SupportVoPS) {
		var ret bool
		return ret
	}
	return *o.SupportVoPS
}

// GetSupportVoPSOk returns a tuple with the SupportVoPS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInfo) GetSupportVoPSOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportVoPS) {
		return nil, false
	}
	return o.SupportVoPS, true
}

// HasSupportVoPS returns a boolean if a field has been set.
func (o *UeContextInfo) HasSupportVoPS() bool {
	if o != nil && !IsNil(o.SupportVoPS) {
		return true
	}

	return false
}

// SetSupportVoPS gets a reference to the given bool and assigns it to the SupportVoPS field.
func (o *UeContextInfo) SetSupportVoPS(v bool) {
	o.SupportVoPS = &v
}

// GetSupportVoPSn3gpp returns the SupportVoPSn3gpp field value if set, zero value otherwise.
func (o *UeContextInfo) GetSupportVoPSn3gpp() bool {
	if o == nil || IsNil(o.SupportVoPSn3gpp) {
		var ret bool
		return ret
	}
	return *o.SupportVoPSn3gpp
}

// GetSupportVoPSn3gppOk returns a tuple with the SupportVoPSn3gpp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInfo) GetSupportVoPSn3gppOk() (*bool, bool) {
	if o == nil || IsNil(o.SupportVoPSn3gpp) {
		return nil, false
	}
	return o.SupportVoPSn3gpp, true
}

// HasSupportVoPSn3gpp returns a boolean if a field has been set.
func (o *UeContextInfo) HasSupportVoPSn3gpp() bool {
	if o != nil && !IsNil(o.SupportVoPSn3gpp) {
		return true
	}

	return false
}

// SetSupportVoPSn3gpp gets a reference to the given bool and assigns it to the SupportVoPSn3gpp field.
func (o *UeContextInfo) SetSupportVoPSn3gpp(v bool) {
	o.SupportVoPSn3gpp = &v
}

// GetLastActTime returns the LastActTime field value if set, zero value otherwise.
func (o *UeContextInfo) GetLastActTime() time.Time {
	if o == nil || IsNil(o.LastActTime) {
		var ret time.Time
		return ret
	}
	return *o.LastActTime
}

// GetLastActTimeOk returns a tuple with the LastActTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInfo) GetLastActTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastActTime) {
		return nil, false
	}
	return o.LastActTime, true
}

// HasLastActTime returns a boolean if a field has been set.
func (o *UeContextInfo) HasLastActTime() bool {
	if o != nil && !IsNil(o.LastActTime) {
		return true
	}

	return false
}

// SetLastActTime gets a reference to the given time.Time and assigns it to the LastActTime field.
func (o *UeContextInfo) SetLastActTime(v time.Time) {
	o.LastActTime = &v
}

// GetAccessType returns the AccessType field value if set, zero value otherwise.
func (o *UeContextInfo) GetAccessType() AccessType {
	if o == nil || IsNil(o.AccessType) {
		var ret AccessType
		return ret
	}
	return *o.AccessType
}

// GetAccessTypeOk returns a tuple with the AccessType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInfo) GetAccessTypeOk() (*AccessType, bool) {
	if o == nil || IsNil(o.AccessType) {
		return nil, false
	}
	return o.AccessType, true
}

// HasAccessType returns a boolean if a field has been set.
func (o *UeContextInfo) HasAccessType() bool {
	if o != nil && !IsNil(o.AccessType) {
		return true
	}

	return false
}

// SetAccessType gets a reference to the given AccessType and assigns it to the AccessType field.
func (o *UeContextInfo) SetAccessType(v AccessType) {
	o.AccessType = &v
}

// GetRatType returns the RatType field value if set, zero value otherwise.
func (o *UeContextInfo) GetRatType() RatType {
	if o == nil || IsNil(o.RatType) {
		var ret RatType
		return ret
	}
	return *o.RatType
}

// GetRatTypeOk returns a tuple with the RatType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInfo) GetRatTypeOk() (*RatType, bool) {
	if o == nil || IsNil(o.RatType) {
		return nil, false
	}
	return o.RatType, true
}

// HasRatType returns a boolean if a field has been set.
func (o *UeContextInfo) HasRatType() bool {
	if o != nil && !IsNil(o.RatType) {
		return true
	}

	return false
}

// SetRatType gets a reference to the given RatType and assigns it to the RatType field.
func (o *UeContextInfo) SetRatType(v RatType) {
	o.RatType = &v
}

// GetSupportedFeatures returns the SupportedFeatures field value if set, zero value otherwise.
func (o *UeContextInfo) GetSupportedFeatures() string {
	if o == nil || IsNil(o.SupportedFeatures) {
		var ret string
		return ret
	}
	return *o.SupportedFeatures
}

// GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeContextInfo) GetSupportedFeaturesOk() (*string, bool) {
	if o == nil || IsNil(o.SupportedFeatures) {
		return nil, false
	}
	return o.SupportedFeatures, true
}

// HasSupportedFeatures returns a boolean if a field has been set.
func (o *UeContextInfo) HasSupportedFeatures() bool {
	if o != nil && !IsNil(o.SupportedFeatures) {
		return true
	}

	return false
}

// SetSupportedFeatures gets a reference to the given string and assigns it to the SupportedFeatures field.
func (o *UeContextInfo) SetSupportedFeatures(v string) {
	o.SupportedFeatures = &v
}

func (o UeContextInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeContextInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SupportVoPS) {
		toSerialize["supportVoPS"] = o.SupportVoPS
	}
	if !IsNil(o.SupportVoPSn3gpp) {
		toSerialize["supportVoPSn3gpp"] = o.SupportVoPSn3gpp
	}
	if !IsNil(o.LastActTime) {
		toSerialize["lastActTime"] = o.LastActTime
	}
	if !IsNil(o.AccessType) {
		toSerialize["accessType"] = o.AccessType
	}
	if !IsNil(o.RatType) {
		toSerialize["ratType"] = o.RatType
	}
	if !IsNil(o.SupportedFeatures) {
		toSerialize["supportedFeatures"] = o.SupportedFeatures
	}
	return toSerialize, nil
}

type NullableUeContextInfo struct {
	value *UeContextInfo
	isSet bool
}

func (v NullableUeContextInfo) Get() *UeContextInfo {
	return v.value
}

func (v *NullableUeContextInfo) Set(val *UeContextInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableUeContextInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableUeContextInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeContextInfo(val *UeContextInfo) *NullableUeContextInfo {
	return &NullableUeContextInfo{value: val, isSet: true}
}

func (v NullableUeContextInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeContextInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
