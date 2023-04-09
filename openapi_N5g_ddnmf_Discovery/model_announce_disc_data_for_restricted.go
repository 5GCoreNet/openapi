/*
N5g-ddnmf_Discovery API

N5g-ddnmf_Discovery Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_N5g_ddnmf_Discovery

import (
	"encoding/json"
	"time"
)

// checks if the AnnounceDiscDataForRestricted type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AnnounceDiscDataForRestricted{}

// AnnounceDiscDataForRestricted Represents Data for restricted discovery used to request the authorization to announce for a UE
type AnnounceDiscDataForRestricted struct {
	// Contains the RPAUID.
	Rpauid string `json:"rpauid"`
	// Contains the Application ID.
	AppId string `json:"appId"`
	// string with format 'date-time' as defined in OpenAPI.
	ValidityTime time.Time `json:"validityTime"`
	// Contains the ProSe Restricted Code.
	ProseRestrictedCode *string `json:"proseRestrictedCode,omitempty"`
	// Contains the ProSe Restricted Code Prefix.
	ProseRestrictedPrefix *string                   `json:"proseRestrictedPrefix,omitempty"`
	CodeSuffixPool        *RestrictedCodeSuffixPool `json:"codeSuffixPool,omitempty"`
}

// NewAnnounceDiscDataForRestricted instantiates a new AnnounceDiscDataForRestricted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAnnounceDiscDataForRestricted(rpauid string, appId string, validityTime time.Time) *AnnounceDiscDataForRestricted {
	this := AnnounceDiscDataForRestricted{}
	this.Rpauid = rpauid
	this.AppId = appId
	this.ValidityTime = validityTime
	return &this
}

// NewAnnounceDiscDataForRestrictedWithDefaults instantiates a new AnnounceDiscDataForRestricted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAnnounceDiscDataForRestrictedWithDefaults() *AnnounceDiscDataForRestricted {
	this := AnnounceDiscDataForRestricted{}
	return &this
}

// GetRpauid returns the Rpauid field value
func (o *AnnounceDiscDataForRestricted) GetRpauid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Rpauid
}

// GetRpauidOk returns a tuple with the Rpauid field value
// and a boolean to check if the value has been set.
func (o *AnnounceDiscDataForRestricted) GetRpauidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rpauid, true
}

// SetRpauid sets field value
func (o *AnnounceDiscDataForRestricted) SetRpauid(v string) {
	o.Rpauid = v
}

// GetAppId returns the AppId field value
func (o *AnnounceDiscDataForRestricted) GetAppId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AppId
}

// GetAppIdOk returns a tuple with the AppId field value
// and a boolean to check if the value has been set.
func (o *AnnounceDiscDataForRestricted) GetAppIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AppId, true
}

// SetAppId sets field value
func (o *AnnounceDiscDataForRestricted) SetAppId(v string) {
	o.AppId = v
}

// GetValidityTime returns the ValidityTime field value
func (o *AnnounceDiscDataForRestricted) GetValidityTime() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ValidityTime
}

// GetValidityTimeOk returns a tuple with the ValidityTime field value
// and a boolean to check if the value has been set.
func (o *AnnounceDiscDataForRestricted) GetValidityTimeOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ValidityTime, true
}

// SetValidityTime sets field value
func (o *AnnounceDiscDataForRestricted) SetValidityTime(v time.Time) {
	o.ValidityTime = v
}

// GetProseRestrictedCode returns the ProseRestrictedCode field value if set, zero value otherwise.
func (o *AnnounceDiscDataForRestricted) GetProseRestrictedCode() string {
	if o == nil || IsNil(o.ProseRestrictedCode) {
		var ret string
		return ret
	}
	return *o.ProseRestrictedCode
}

// GetProseRestrictedCodeOk returns a tuple with the ProseRestrictedCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnounceDiscDataForRestricted) GetProseRestrictedCodeOk() (*string, bool) {
	if o == nil || IsNil(o.ProseRestrictedCode) {
		return nil, false
	}
	return o.ProseRestrictedCode, true
}

// HasProseRestrictedCode returns a boolean if a field has been set.
func (o *AnnounceDiscDataForRestricted) HasProseRestrictedCode() bool {
	if o != nil && !IsNil(o.ProseRestrictedCode) {
		return true
	}

	return false
}

// SetProseRestrictedCode gets a reference to the given string and assigns it to the ProseRestrictedCode field.
func (o *AnnounceDiscDataForRestricted) SetProseRestrictedCode(v string) {
	o.ProseRestrictedCode = &v
}

// GetProseRestrictedPrefix returns the ProseRestrictedPrefix field value if set, zero value otherwise.
func (o *AnnounceDiscDataForRestricted) GetProseRestrictedPrefix() string {
	if o == nil || IsNil(o.ProseRestrictedPrefix) {
		var ret string
		return ret
	}
	return *o.ProseRestrictedPrefix
}

// GetProseRestrictedPrefixOk returns a tuple with the ProseRestrictedPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnounceDiscDataForRestricted) GetProseRestrictedPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ProseRestrictedPrefix) {
		return nil, false
	}
	return o.ProseRestrictedPrefix, true
}

// HasProseRestrictedPrefix returns a boolean if a field has been set.
func (o *AnnounceDiscDataForRestricted) HasProseRestrictedPrefix() bool {
	if o != nil && !IsNil(o.ProseRestrictedPrefix) {
		return true
	}

	return false
}

// SetProseRestrictedPrefix gets a reference to the given string and assigns it to the ProseRestrictedPrefix field.
func (o *AnnounceDiscDataForRestricted) SetProseRestrictedPrefix(v string) {
	o.ProseRestrictedPrefix = &v
}

// GetCodeSuffixPool returns the CodeSuffixPool field value if set, zero value otherwise.
func (o *AnnounceDiscDataForRestricted) GetCodeSuffixPool() RestrictedCodeSuffixPool {
	if o == nil || IsNil(o.CodeSuffixPool) {
		var ret RestrictedCodeSuffixPool
		return ret
	}
	return *o.CodeSuffixPool
}

// GetCodeSuffixPoolOk returns a tuple with the CodeSuffixPool field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AnnounceDiscDataForRestricted) GetCodeSuffixPoolOk() (*RestrictedCodeSuffixPool, bool) {
	if o == nil || IsNil(o.CodeSuffixPool) {
		return nil, false
	}
	return o.CodeSuffixPool, true
}

// HasCodeSuffixPool returns a boolean if a field has been set.
func (o *AnnounceDiscDataForRestricted) HasCodeSuffixPool() bool {
	if o != nil && !IsNil(o.CodeSuffixPool) {
		return true
	}

	return false
}

// SetCodeSuffixPool gets a reference to the given RestrictedCodeSuffixPool and assigns it to the CodeSuffixPool field.
func (o *AnnounceDiscDataForRestricted) SetCodeSuffixPool(v RestrictedCodeSuffixPool) {
	o.CodeSuffixPool = &v
}

func (o AnnounceDiscDataForRestricted) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AnnounceDiscDataForRestricted) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rpauid"] = o.Rpauid
	toSerialize["appId"] = o.AppId
	toSerialize["validityTime"] = o.ValidityTime
	if !IsNil(o.ProseRestrictedCode) {
		toSerialize["proseRestrictedCode"] = o.ProseRestrictedCode
	}
	if !IsNil(o.ProseRestrictedPrefix) {
		toSerialize["proseRestrictedPrefix"] = o.ProseRestrictedPrefix
	}
	if !IsNil(o.CodeSuffixPool) {
		toSerialize["codeSuffixPool"] = o.CodeSuffixPool
	}
	return toSerialize, nil
}

type NullableAnnounceDiscDataForRestricted struct {
	value *AnnounceDiscDataForRestricted
	isSet bool
}

func (v NullableAnnounceDiscDataForRestricted) Get() *AnnounceDiscDataForRestricted {
	return v.value
}

func (v *NullableAnnounceDiscDataForRestricted) Set(val *AnnounceDiscDataForRestricted) {
	v.value = val
	v.isSet = true
}

func (v NullableAnnounceDiscDataForRestricted) IsSet() bool {
	return v.isSet
}

func (v *NullableAnnounceDiscDataForRestricted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAnnounceDiscDataForRestricted(val *AnnounceDiscDataForRestricted) *NullableAnnounceDiscDataForRestricted {
	return &NullableAnnounceDiscDataForRestricted{value: val, isSet: true}
}

func (v NullableAnnounceDiscDataForRestricted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAnnounceDiscDataForRestricted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
