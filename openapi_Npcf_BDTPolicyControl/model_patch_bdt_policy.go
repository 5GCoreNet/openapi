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

// checks if the PatchBdtPolicy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchBdtPolicy{}

// PatchBdtPolicy Describes the updates in authorization data of an Individual BDT Policy created by the PCF.
type PatchBdtPolicy struct {
	BdtPolData *BdtPolicyDataPatch `json:"bdtPolData,omitempty"`
	BdtReqData *BdtReqDataPatch    `json:"bdtReqData,omitempty"`
}

// NewPatchBdtPolicy instantiates a new PatchBdtPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchBdtPolicy() *PatchBdtPolicy {
	this := PatchBdtPolicy{}
	return &this
}

// NewPatchBdtPolicyWithDefaults instantiates a new PatchBdtPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchBdtPolicyWithDefaults() *PatchBdtPolicy {
	this := PatchBdtPolicy{}
	return &this
}

// GetBdtPolData returns the BdtPolData field value if set, zero value otherwise.
func (o *PatchBdtPolicy) GetBdtPolData() BdtPolicyDataPatch {
	if o == nil || IsNil(o.BdtPolData) {
		var ret BdtPolicyDataPatch
		return ret
	}
	return *o.BdtPolData
}

// GetBdtPolDataOk returns a tuple with the BdtPolData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBdtPolicy) GetBdtPolDataOk() (*BdtPolicyDataPatch, bool) {
	if o == nil || IsNil(o.BdtPolData) {
		return nil, false
	}
	return o.BdtPolData, true
}

// HasBdtPolData returns a boolean if a field has been set.
func (o *PatchBdtPolicy) HasBdtPolData() bool {
	if o != nil && !IsNil(o.BdtPolData) {
		return true
	}

	return false
}

// SetBdtPolData gets a reference to the given BdtPolicyDataPatch and assigns it to the BdtPolData field.
func (o *PatchBdtPolicy) SetBdtPolData(v BdtPolicyDataPatch) {
	o.BdtPolData = &v
}

// GetBdtReqData returns the BdtReqData field value if set, zero value otherwise.
func (o *PatchBdtPolicy) GetBdtReqData() BdtReqDataPatch {
	if o == nil || IsNil(o.BdtReqData) {
		var ret BdtReqDataPatch
		return ret
	}
	return *o.BdtReqData
}

// GetBdtReqDataOk returns a tuple with the BdtReqData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchBdtPolicy) GetBdtReqDataOk() (*BdtReqDataPatch, bool) {
	if o == nil || IsNil(o.BdtReqData) {
		return nil, false
	}
	return o.BdtReqData, true
}

// HasBdtReqData returns a boolean if a field has been set.
func (o *PatchBdtPolicy) HasBdtReqData() bool {
	if o != nil && !IsNil(o.BdtReqData) {
		return true
	}

	return false
}

// SetBdtReqData gets a reference to the given BdtReqDataPatch and assigns it to the BdtReqData field.
func (o *PatchBdtPolicy) SetBdtReqData(v BdtReqDataPatch) {
	o.BdtReqData = &v
}

func (o PatchBdtPolicy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchBdtPolicy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BdtPolData) {
		toSerialize["bdtPolData"] = o.BdtPolData
	}
	if !IsNil(o.BdtReqData) {
		toSerialize["bdtReqData"] = o.BdtReqData
	}
	return toSerialize, nil
}

type NullablePatchBdtPolicy struct {
	value *PatchBdtPolicy
	isSet bool
}

func (v NullablePatchBdtPolicy) Get() *PatchBdtPolicy {
	return v.value
}

func (v *NullablePatchBdtPolicy) Set(val *PatchBdtPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchBdtPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchBdtPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchBdtPolicy(val *PatchBdtPolicy) *NullablePatchBdtPolicy {
	return &NullablePatchBdtPolicy{value: val, isSet: true}
}

func (v NullablePatchBdtPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchBdtPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
