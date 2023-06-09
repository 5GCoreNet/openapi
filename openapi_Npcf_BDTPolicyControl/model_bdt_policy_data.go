/*
Npcf_BDTPolicyControl Service API

PCF BDT Policy Control Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_BDTPolicyControl

import (
	"encoding/json"
)

// checks if the BdtPolicyData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BdtPolicyData{}

// BdtPolicyData Describes the authorization data of an Individual BDT policy resource.
type BdtPolicyData struct {
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`
	// Contains transfer policies.
	TransfPolicies []TransferPolicy `json:"transfPolicies"`
	// Contains an identity of the selected transfer policy.
	SelTransPolicyId *int32 `json:"selTransPolicyId,omitempty"`
	// A string used to indicate the features supported by an API that is used as defined in clause  6.6 in 3GPP TS 29.500. The string shall contain a bitmask indicating supported features in  hexadecimal representation Each character in the string shall take a value of \"0\" to \"9\",  \"a\" to \"f\" or \"A\" to \"F\" and shall represent the support of 4 features as described in  table 5.2.2-3. The most significant character representing the highest-numbered features shall  appear first in the string, and the character representing features 1 to 4 shall appear last  in the string. The list of features and their numbering (starting with 1) are defined  separately for each API. If the string contains a lower number of characters than there are  defined features for an API, all features that would be represented by characters that are not  present in the string are not supported.
	SuppFeat *string `json:"suppFeat,omitempty"`
}

// NewBdtPolicyData instantiates a new BdtPolicyData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBdtPolicyData(bdtRefId string, transfPolicies []TransferPolicy) *BdtPolicyData {
	this := BdtPolicyData{}
	this.BdtRefId = bdtRefId
	this.TransfPolicies = transfPolicies
	return &this
}

// NewBdtPolicyDataWithDefaults instantiates a new BdtPolicyData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBdtPolicyDataWithDefaults() *BdtPolicyData {
	this := BdtPolicyData{}
	return &this
}

// GetBdtRefId returns the BdtRefId field value
func (o *BdtPolicyData) GetBdtRefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BdtRefId
}

// GetBdtRefIdOk returns a tuple with the BdtRefId field value
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetBdtRefIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BdtRefId, true
}

// SetBdtRefId sets field value
func (o *BdtPolicyData) SetBdtRefId(v string) {
	o.BdtRefId = v
}

// GetTransfPolicies returns the TransfPolicies field value
func (o *BdtPolicyData) GetTransfPolicies() []TransferPolicy {
	if o == nil {
		var ret []TransferPolicy
		return ret
	}

	return o.TransfPolicies
}

// GetTransfPoliciesOk returns a tuple with the TransfPolicies field value
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetTransfPoliciesOk() ([]TransferPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.TransfPolicies, true
}

// SetTransfPolicies sets field value
func (o *BdtPolicyData) SetTransfPolicies(v []TransferPolicy) {
	o.TransfPolicies = v
}

// GetSelTransPolicyId returns the SelTransPolicyId field value if set, zero value otherwise.
func (o *BdtPolicyData) GetSelTransPolicyId() int32 {
	if o == nil || IsNil(o.SelTransPolicyId) {
		var ret int32
		return ret
	}
	return *o.SelTransPolicyId
}

// GetSelTransPolicyIdOk returns a tuple with the SelTransPolicyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetSelTransPolicyIdOk() (*int32, bool) {
	if o == nil || IsNil(o.SelTransPolicyId) {
		return nil, false
	}
	return o.SelTransPolicyId, true
}

// HasSelTransPolicyId returns a boolean if a field has been set.
func (o *BdtPolicyData) HasSelTransPolicyId() bool {
	if o != nil && !IsNil(o.SelTransPolicyId) {
		return true
	}

	return false
}

// SetSelTransPolicyId gets a reference to the given int32 and assigns it to the SelTransPolicyId field.
func (o *BdtPolicyData) SetSelTransPolicyId(v int32) {
	o.SelTransPolicyId = &v
}

// GetSuppFeat returns the SuppFeat field value if set, zero value otherwise.
func (o *BdtPolicyData) GetSuppFeat() string {
	if o == nil || IsNil(o.SuppFeat) {
		var ret string
		return ret
	}
	return *o.SuppFeat
}

// GetSuppFeatOk returns a tuple with the SuppFeat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BdtPolicyData) GetSuppFeatOk() (*string, bool) {
	if o == nil || IsNil(o.SuppFeat) {
		return nil, false
	}
	return o.SuppFeat, true
}

// HasSuppFeat returns a boolean if a field has been set.
func (o *BdtPolicyData) HasSuppFeat() bool {
	if o != nil && !IsNil(o.SuppFeat) {
		return true
	}

	return false
}

// SetSuppFeat gets a reference to the given string and assigns it to the SuppFeat field.
func (o *BdtPolicyData) SetSuppFeat(v string) {
	o.SuppFeat = &v
}

func (o BdtPolicyData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BdtPolicyData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bdtRefId"] = o.BdtRefId
	toSerialize["transfPolicies"] = o.TransfPolicies
	if !IsNil(o.SelTransPolicyId) {
		toSerialize["selTransPolicyId"] = o.SelTransPolicyId
	}
	if !IsNil(o.SuppFeat) {
		toSerialize["suppFeat"] = o.SuppFeat
	}
	return toSerialize, nil
}

type NullableBdtPolicyData struct {
	value *BdtPolicyData
	isSet bool
}

func (v NullableBdtPolicyData) Get() *BdtPolicyData {
	return v.value
}

func (v *NullableBdtPolicyData) Set(val *BdtPolicyData) {
	v.value = val
	v.isSet = true
}

func (v NullableBdtPolicyData) IsSet() bool {
	return v.isSet
}

func (v *NullableBdtPolicyData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBdtPolicyData(val *BdtPolicyData) *NullableBdtPolicyData {
	return &NullableBdtPolicyData{value: val, isSet: true}
}

func (v NullableBdtPolicyData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBdtPolicyData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
