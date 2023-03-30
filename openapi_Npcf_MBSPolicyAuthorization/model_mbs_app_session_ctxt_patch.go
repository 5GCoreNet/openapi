/*
Npcf_MBSPolicyAuthorization API

MBS Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_MBSPolicyAuthorization

import (
	"encoding/json"
)

// checks if the MbsAppSessionCtxtPatch type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsAppSessionCtxtPatch{}

// MbsAppSessionCtxtPatch Represents the modifications to an existing MBS Application Session Context resource. 
type MbsAppSessionCtxtPatch struct {
	MbsServInfo *MbsServiceInfo `json:"mbsServInfo,omitempty"`
}

// NewMbsAppSessionCtxtPatch instantiates a new MbsAppSessionCtxtPatch object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsAppSessionCtxtPatch() *MbsAppSessionCtxtPatch {
	this := MbsAppSessionCtxtPatch{}
	return &this
}

// NewMbsAppSessionCtxtPatchWithDefaults instantiates a new MbsAppSessionCtxtPatch object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsAppSessionCtxtPatchWithDefaults() *MbsAppSessionCtxtPatch {
	this := MbsAppSessionCtxtPatch{}
	return &this
}

// GetMbsServInfo returns the MbsServInfo field value if set, zero value otherwise.
func (o *MbsAppSessionCtxtPatch) GetMbsServInfo() MbsServiceInfo {
	if o == nil || IsNil(o.MbsServInfo) {
		var ret MbsServiceInfo
		return ret
	}
	return *o.MbsServInfo
}

// GetMbsServInfoOk returns a tuple with the MbsServInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsAppSessionCtxtPatch) GetMbsServInfoOk() (*MbsServiceInfo, bool) {
	if o == nil || IsNil(o.MbsServInfo) {
		return nil, false
	}
	return o.MbsServInfo, true
}

// HasMbsServInfo returns a boolean if a field has been set.
func (o *MbsAppSessionCtxtPatch) HasMbsServInfo() bool {
	if o != nil && !IsNil(o.MbsServInfo) {
		return true
	}

	return false
}

// SetMbsServInfo gets a reference to the given MbsServiceInfo and assigns it to the MbsServInfo field.
func (o *MbsAppSessionCtxtPatch) SetMbsServInfo(v MbsServiceInfo) {
	o.MbsServInfo = &v
}

func (o MbsAppSessionCtxtPatch) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsAppSessionCtxtPatch) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsServInfo) {
		toSerialize["mbsServInfo"] = o.MbsServInfo
	}
	return toSerialize, nil
}

type NullableMbsAppSessionCtxtPatch struct {
	value *MbsAppSessionCtxtPatch
	isSet bool
}

func (v NullableMbsAppSessionCtxtPatch) Get() *MbsAppSessionCtxtPatch {
	return v.value
}

func (v *NullableMbsAppSessionCtxtPatch) Set(val *MbsAppSessionCtxtPatch) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsAppSessionCtxtPatch) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsAppSessionCtxtPatch) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsAppSessionCtxtPatch(val *MbsAppSessionCtxtPatch) *NullableMbsAppSessionCtxtPatch {
	return &NullableMbsAppSessionCtxtPatch{value: val, isSet: true}
}

func (v NullableMbsAppSessionCtxtPatch) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsAppSessionCtxtPatch) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


