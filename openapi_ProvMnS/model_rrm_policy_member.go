/*
Provisioning MnS

OAS 3.0.1 definition of the Provisioning MnS © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 17.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_ProvMnS

import (
	"encoding/json"
)

// checks if the RrmPolicyMember type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RrmPolicyMember{}

// RrmPolicyMember struct for RrmPolicyMember
type RrmPolicyMember struct {
	PlmnId *PlmnId `json:"plmnId,omitempty"`
	Snssai *Snssai `json:"snssai,omitempty"`
}

// NewRrmPolicyMember instantiates a new RrmPolicyMember object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRrmPolicyMember() *RrmPolicyMember {
	this := RrmPolicyMember{}
	return &this
}

// NewRrmPolicyMemberWithDefaults instantiates a new RrmPolicyMember object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRrmPolicyMemberWithDefaults() *RrmPolicyMember {
	this := RrmPolicyMember{}
	return &this
}

// GetPlmnId returns the PlmnId field value if set, zero value otherwise.
func (o *RrmPolicyMember) GetPlmnId() PlmnId {
	if o == nil || IsNil(o.PlmnId) {
		var ret PlmnId
		return ret
	}
	return *o.PlmnId
}

// GetPlmnIdOk returns a tuple with the PlmnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmPolicyMember) GetPlmnIdOk() (*PlmnId, bool) {
	if o == nil || IsNil(o.PlmnId) {
		return nil, false
	}
	return o.PlmnId, true
}

// HasPlmnId returns a boolean if a field has been set.
func (o *RrmPolicyMember) HasPlmnId() bool {
	if o != nil && !IsNil(o.PlmnId) {
		return true
	}

	return false
}

// SetPlmnId gets a reference to the given PlmnId and assigns it to the PlmnId field.
func (o *RrmPolicyMember) SetPlmnId(v PlmnId) {
	o.PlmnId = &v
}

// GetSnssai returns the Snssai field value if set, zero value otherwise.
func (o *RrmPolicyMember) GetSnssai() Snssai {
	if o == nil || IsNil(o.Snssai) {
		var ret Snssai
		return ret
	}
	return *o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RrmPolicyMember) GetSnssaiOk() (*Snssai, bool) {
	if o == nil || IsNil(o.Snssai) {
		return nil, false
	}
	return o.Snssai, true
}

// HasSnssai returns a boolean if a field has been set.
func (o *RrmPolicyMember) HasSnssai() bool {
	if o != nil && !IsNil(o.Snssai) {
		return true
	}

	return false
}

// SetSnssai gets a reference to the given Snssai and assigns it to the Snssai field.
func (o *RrmPolicyMember) SetSnssai(v Snssai) {
	o.Snssai = &v
}

func (o RrmPolicyMember) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RrmPolicyMember) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PlmnId) {
		toSerialize["plmnId"] = o.PlmnId
	}
	if !IsNil(o.Snssai) {
		toSerialize["snssai"] = o.Snssai
	}
	return toSerialize, nil
}

type NullableRrmPolicyMember struct {
	value *RrmPolicyMember
	isSet bool
}

func (v NullableRrmPolicyMember) Get() *RrmPolicyMember {
	return v.value
}

func (v *NullableRrmPolicyMember) Set(val *RrmPolicyMember) {
	v.value = val
	v.isSet = true
}

func (v NullableRrmPolicyMember) IsSet() bool {
	return v.isSet
}

func (v *NullableRrmPolicyMember) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRrmPolicyMember(val *RrmPolicyMember) *NullableRrmPolicyMember {
	return &NullableRrmPolicyMember{value: val, isSet: true}
}

func (v NullableRrmPolicyMember) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRrmPolicyMember) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
