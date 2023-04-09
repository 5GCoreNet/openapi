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

// checks if the BdtPolicyDataPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BdtPolicyDataPatch{}

// BdtPolicyDataPatch A JSON Merge Patch body schema containing modification instruction to be performed on the bdtPolData attribute of the BdtPolicy data structure to select a transfer policy. Adds selTransPolicyId to BdtPolicyData data structure.
type BdtPolicyDataPatch struct {
	// Contains an identity (i.e. transPolicyId value) of the selected transfer policy. If the BdtNotification_5G feature is supported value 0 indicates that no transfer policy is selected.
	SelTransPolicyId int32 `json:"selTransPolicyId"`
}

// NewBdtPolicyDataPatch instantiates a new BdtPolicyDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBdtPolicyDataPatch(selTransPolicyId int32) *BdtPolicyDataPatch {
	this := BdtPolicyDataPatch{}
	this.SelTransPolicyId = selTransPolicyId
	return &this
}

// NewBdtPolicyDataPatchWithDefaults instantiates a new BdtPolicyDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBdtPolicyDataPatchWithDefaults() *BdtPolicyDataPatch {
	this := BdtPolicyDataPatch{}
	return &this
}

// GetSelTransPolicyId returns the SelTransPolicyId field value
func (o *BdtPolicyDataPatch) GetSelTransPolicyId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.SelTransPolicyId
}

// GetSelTransPolicyIdOk returns a tuple with the SelTransPolicyId field value
// and a boolean to check if the value has been set.
func (o *BdtPolicyDataPatch) GetSelTransPolicyIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelTransPolicyId, true
}

// SetSelTransPolicyId sets field value
func (o *BdtPolicyDataPatch) SetSelTransPolicyId(v int32) {
	o.SelTransPolicyId = v
}

func (o BdtPolicyDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BdtPolicyDataPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["selTransPolicyId"] = o.SelTransPolicyId
	return toSerialize, nil
}

type NullableBdtPolicyDataPatch struct {
	value *BdtPolicyDataPatch
	isSet bool
}

func (v NullableBdtPolicyDataPatch) Get() *BdtPolicyDataPatch {
	return v.value
}

func (v *NullableBdtPolicyDataPatch) Set(val *BdtPolicyDataPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableBdtPolicyDataPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableBdtPolicyDataPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBdtPolicyDataPatch(val *BdtPolicyDataPatch) *NullableBdtPolicyDataPatch {
	return &NullableBdtPolicyDataPatch{value: val, isSet: true}
}

func (v NullableBdtPolicyDataPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBdtPolicyDataPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
