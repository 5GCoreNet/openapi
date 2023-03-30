/*
Unified Data Repository Service API file for Application Data

The API version is defined in 3GPP TS 29.504   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: -
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Application_Data

import (
	"encoding/json"
)

// checks if the BdtPolicyDataPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BdtPolicyDataPatch{}

// BdtPolicyDataPatch Represents modification instructions to be performed on the applied BDT policy data. 
type BdtPolicyDataPatch struct {
	// string identifying a BDT Reference ID as defined in clause 5.3.3 of 3GPP TS 29.154.
	BdtRefId string `json:"bdtRefId"`
}

// NewBdtPolicyDataPatch instantiates a new BdtPolicyDataPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBdtPolicyDataPatch(bdtRefId string) *BdtPolicyDataPatch {
	this := BdtPolicyDataPatch{}
	this.BdtRefId = bdtRefId
	return &this
}

// NewBdtPolicyDataPatchWithDefaults instantiates a new BdtPolicyDataPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBdtPolicyDataPatchWithDefaults() *BdtPolicyDataPatch {
	this := BdtPolicyDataPatch{}
	return &this
}

// GetBdtRefId returns the BdtRefId field value
func (o *BdtPolicyDataPatch) GetBdtRefId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.BdtRefId
}

// GetBdtRefIdOk returns a tuple with the BdtRefId field value
// and a boolean to check if the value has been set.
func (o *BdtPolicyDataPatch) GetBdtRefIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.BdtRefId, true
}

// SetBdtRefId sets field value
func (o *BdtPolicyDataPatch) SetBdtRefId(v string) {
	o.BdtRefId = v
}

func (o BdtPolicyDataPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BdtPolicyDataPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["bdtRefId"] = o.BdtRefId
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


