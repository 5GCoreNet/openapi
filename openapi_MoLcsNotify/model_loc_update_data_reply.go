/*
3gpp-mo-lcs-notify

API for UE updated location information notification.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_MoLcsNotify

import (
	"encoding/json"
)

// checks if the LocUpdateDataReply type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &LocUpdateDataReply{}

// LocUpdateDataReply Represents a reply to a MO LCS notification.
type LocUpdateDataReply struct {
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported. 
	SuppFeat string `json:"suppFeat"`
}

// NewLocUpdateDataReply instantiates a new LocUpdateDataReply object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocUpdateDataReply(suppFeat string) *LocUpdateDataReply {
	this := LocUpdateDataReply{}
	this.SuppFeat = suppFeat
	return &this
}

// NewLocUpdateDataReplyWithDefaults instantiates a new LocUpdateDataReply object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocUpdateDataReplyWithDefaults() *LocUpdateDataReply {
	this := LocUpdateDataReply{}
	return &this
}

// GetSuppFeat returns the SuppFeat field value
func (o *LocUpdateDataReply) GetSuppFeat() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value
// and a boolean to check if the value has been set.
func (o *LocUpdateDataReply) GetSuppFeatOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SuppFeat, true
}

// SetSuppFeat sets field value
func (o *LocUpdateDataReply) SetSuppFeat(v string) {
	o.SuppFeat = v
}

func (o LocUpdateDataReply) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LocUpdateDataReply) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["suppFeat"] = o.SuppFeat
	return toSerialize, nil
}

type NullableLocUpdateDataReply struct {
	value *LocUpdateDataReply
	isSet bool
}

func (v NullableLocUpdateDataReply) Get() *LocUpdateDataReply {
	return v.value
}

func (v *NullableLocUpdateDataReply) Set(val *LocUpdateDataReply) {
	v.value = val
	v.isSet = true
}

func (v NullableLocUpdateDataReply) IsSet() bool {
	return v.isSet
}

func (v *NullableLocUpdateDataReply) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocUpdateDataReply(val *LocUpdateDataReply) *NullableLocUpdateDataReply {
	return &NullableLocUpdateDataReply{value: val, isSet: true}
}

func (v NullableLocUpdateDataReply) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocUpdateDataReply) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

