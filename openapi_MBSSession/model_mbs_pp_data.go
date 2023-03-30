/*
3gpp-mbs-session

API for MBS Session Management.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MBSSession

import (
	"encoding/json"
)

// checks if the MbsPpData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsPpData{}

// MbsPpData Represents MBS Parameters Provisioning data.
type MbsPpData struct {
	AfId string `json:"afId"`
	MbsSessAuthData *MbsSessAuthData `json:"mbsSessAuthData,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewMbsPpData instantiates a new MbsPpData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsPpData(afId string) *MbsPpData {
	this := MbsPpData{}
	this.AfId = afId
	return &this
}

// NewMbsPpDataWithDefaults instantiates a new MbsPpData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsPpDataWithDefaults() *MbsPpData {
	this := MbsPpData{}
	return &this
}

// GetAfId returns the AfId field value
func (o *MbsPpData) GetAfId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AfId
}

// GetAfIdOk returns a tuple with the AfId field value
// and a boolean to check if the value has been set.
func (o *MbsPpData) GetAfIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AfId, true
}

// SetAfId sets field value
func (o *MbsPpData) SetAfId(v string) {
	o.AfId = v
}

// GetMbsSessAuthData returns the MbsSessAuthData field value if set, zero value otherwise.
func (o *MbsPpData) GetMbsSessAuthData() MbsSessAuthData {
	if o == nil || IsNil(o.MbsSessAuthData) {
		var ret MbsSessAuthData
		return ret
	}
	return *o.MbsSessAuthData
}

// GetMbsSessAuthDataOk returns a tuple with the MbsSessAuthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPpData) GetMbsSessAuthDataOk() (*MbsSessAuthData, bool) {
	if o == nil || IsNil(o.MbsSessAuthData) {
		return nil, false
	}
	return o.MbsSessAuthData, true
}

// HasMbsSessAuthData returns a boolean if a field has been set.
func (o *MbsPpData) HasMbsSessAuthData() bool {
	if o != nil && !IsNil(o.MbsSessAuthData) {
		return true
	}

	return false
}

// SetMbsSessAuthData gets a reference to the given MbsSessAuthData and assigns it to the MbsSessAuthData field.
func (o *MbsPpData) SetMbsSessAuthData(v MbsSessAuthData) {
	o.MbsSessAuthData = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *MbsPpData) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPpData) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *MbsPpData) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *MbsPpData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o MbsPpData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsPpData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["afId"] = o.AfId
	if !IsNil(o.MbsSessAuthData) {
		toSerialize["mbsSessAuthData"] = o.MbsSessAuthData
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableMbsPpData struct {
	value *MbsPpData
	isSet bool
}

func (v NullableMbsPpData) Get() *MbsPpData {
	return v.value
}

func (v *NullableMbsPpData) Set(val *MbsPpData) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsPpData) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsPpData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsPpData(val *MbsPpData) *NullableMbsPpData {
	return &NullableMbsPpData{value: val, isSet: true}
}

func (v NullableMbsPpData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsPpData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

