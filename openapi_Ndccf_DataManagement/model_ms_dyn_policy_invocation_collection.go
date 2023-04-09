/*
Ndccf_DataManagement

DCCF Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.1.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Ndccf_DataManagement

import (
	"encoding/json"
)

// checks if the MsDynPolicyInvocationCollection type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MsDynPolicyInvocationCollection{}

// MsDynPolicyInvocationCollection Contains the Media Streaming Dynamic Policy invocation collected for an UE Application via AF.
type MsDynPolicyInvocationCollection struct {
	MsDynPlyInvocs []DynamicPolicy `json:"msDynPlyInvocs"`
}

// NewMsDynPolicyInvocationCollection instantiates a new MsDynPolicyInvocationCollection object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMsDynPolicyInvocationCollection(msDynPlyInvocs []DynamicPolicy) *MsDynPolicyInvocationCollection {
	this := MsDynPolicyInvocationCollection{}
	this.MsDynPlyInvocs = msDynPlyInvocs
	return &this
}

// NewMsDynPolicyInvocationCollectionWithDefaults instantiates a new MsDynPolicyInvocationCollection object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMsDynPolicyInvocationCollectionWithDefaults() *MsDynPolicyInvocationCollection {
	this := MsDynPolicyInvocationCollection{}
	return &this
}

// GetMsDynPlyInvocs returns the MsDynPlyInvocs field value
func (o *MsDynPolicyInvocationCollection) GetMsDynPlyInvocs() []DynamicPolicy {
	if o == nil {
		var ret []DynamicPolicy
		return ret
	}

	return o.MsDynPlyInvocs
}

// GetMsDynPlyInvocsOk returns a tuple with the MsDynPlyInvocs field value
// and a boolean to check if the value has been set.
func (o *MsDynPolicyInvocationCollection) GetMsDynPlyInvocsOk() ([]DynamicPolicy, bool) {
	if o == nil {
		return nil, false
	}
	return o.MsDynPlyInvocs, true
}

// SetMsDynPlyInvocs sets field value
func (o *MsDynPolicyInvocationCollection) SetMsDynPlyInvocs(v []DynamicPolicy) {
	o.MsDynPlyInvocs = v
}

func (o MsDynPolicyInvocationCollection) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MsDynPolicyInvocationCollection) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["msDynPlyInvocs"] = o.MsDynPlyInvocs
	return toSerialize, nil
}

type NullableMsDynPolicyInvocationCollection struct {
	value *MsDynPolicyInvocationCollection
	isSet bool
}

func (v NullableMsDynPolicyInvocationCollection) Get() *MsDynPolicyInvocationCollection {
	return v.value
}

func (v *NullableMsDynPolicyInvocationCollection) Set(val *MsDynPolicyInvocationCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableMsDynPolicyInvocationCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableMsDynPolicyInvocationCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMsDynPolicyInvocationCollection(val *MsDynPolicyInvocationCollection) *NullableMsDynPolicyInvocationCollection {
	return &NullableMsDynPolicyInvocationCollection{value: val, isSet: true}
}

func (v NullableMsDynPolicyInvocationCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMsDynPolicyInvocationCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
