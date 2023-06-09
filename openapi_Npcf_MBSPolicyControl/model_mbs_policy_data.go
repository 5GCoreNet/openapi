/*
Npcf_MBSPolicyControl API

MBS Policy Control Service   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_MBSPolicyControl

import (
	"encoding/json"
)

// checks if the MbsPolicyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsPolicyData{}

// MbsPolicyData Contains the MBS policy data provisioned as part of an MBS Policy Association.
type MbsPolicyData struct {
	MbsPolicyCtxtData MbsPolicyCtxtData  `json:"mbsPolicyCtxtData"`
	MbsPolicies       *MbsPolicyDecision `json:"mbsPolicies,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewMbsPolicyData instantiates a new MbsPolicyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsPolicyData(mbsPolicyCtxtData MbsPolicyCtxtData) *MbsPolicyData {
	this := MbsPolicyData{}
	this.MbsPolicyCtxtData = mbsPolicyCtxtData
	return &this
}

// NewMbsPolicyDataWithDefaults instantiates a new MbsPolicyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsPolicyDataWithDefaults() *MbsPolicyData {
	this := MbsPolicyData{}
	return &this
}

// GetMbsPolicyCtxtData returns the MbsPolicyCtxtData field value
func (o *MbsPolicyData) GetMbsPolicyCtxtData() MbsPolicyCtxtData {
	if o == nil {
		var ret MbsPolicyCtxtData
		return ret
	}

	return o.MbsPolicyCtxtData
}

// GetMbsPolicyCtxtDataOk returns a tuple with the MbsPolicyCtxtData field value
// and a boolean to check if the value has been set.
func (o *MbsPolicyData) GetMbsPolicyCtxtDataOk() (*MbsPolicyCtxtData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MbsPolicyCtxtData, true
}

// SetMbsPolicyCtxtData sets field value
func (o *MbsPolicyData) SetMbsPolicyCtxtData(v MbsPolicyCtxtData) {
	o.MbsPolicyCtxtData = v
}

// GetMbsPolicies returns the MbsPolicies field value if set, zero value otherwise.
func (o *MbsPolicyData) GetMbsPolicies() MbsPolicyDecision {
	if o == nil || IsNil(o.MbsPolicies) {
		var ret MbsPolicyDecision
		return ret
	}
	return *o.MbsPolicies
}

// GetMbsPoliciesOk returns a tuple with the MbsPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPolicyData) GetMbsPoliciesOk() (*MbsPolicyDecision, bool) {
	if o == nil || IsNil(o.MbsPolicies) {
		return nil, false
	}
	return o.MbsPolicies, true
}

// HasMbsPolicies returns a boolean if a field has been set.
func (o *MbsPolicyData) HasMbsPolicies() bool {
	if o != nil && !IsNil(o.MbsPolicies) {
		return true
	}

	return false
}

// SetMbsPolicies gets a reference to the given MbsPolicyDecision and assigns it to the MbsPolicies field.
func (o *MbsPolicyData) SetMbsPolicies(v MbsPolicyDecision) {
	o.MbsPolicies = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *MbsPolicyData) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsPolicyData) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *MbsPolicyData) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *MbsPolicyData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o MbsPolicyData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsPolicyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["mbsPolicyCtxtData"] = o.MbsPolicyCtxtData
	if !IsNil(o.MbsPolicies) {
		toSerialize["mbsPolicies"] = o.MbsPolicies
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableMbsPolicyData struct {
	value *MbsPolicyData
	isSet bool
}

func (v NullableMbsPolicyData) Get() *MbsPolicyData {
	return v.value
}

func (v *NullableMbsPolicyData) Set(val *MbsPolicyData) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsPolicyData) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsPolicyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsPolicyData(val *MbsPolicyData) *NullableMbsPolicyData {
	return &NullableMbsPolicyData{value: val, isSet: true}
}

func (v NullableMbsPolicyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsPolicyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
