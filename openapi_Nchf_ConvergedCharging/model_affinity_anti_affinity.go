/*
Nchf_ConvergedCharging

ConvergedCharging Service    © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved. 

API version: 3.2.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nchf_ConvergedCharging

import (
	"encoding/json"
)

// checks if the AffinityAntiAffinity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AffinityAntiAffinity{}

// AffinityAntiAffinity struct for AffinityAntiAffinity
type AffinityAntiAffinity struct {
	AffinityEAS []string `json:"affinityEAS,omitempty"`
	AntiAffinityEAS *string `json:"antiAffinityEAS,omitempty"`
}

// NewAffinityAntiAffinity instantiates a new AffinityAntiAffinity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAffinityAntiAffinity() *AffinityAntiAffinity {
	this := AffinityAntiAffinity{}
	return &this
}

// NewAffinityAntiAffinityWithDefaults instantiates a new AffinityAntiAffinity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAffinityAntiAffinityWithDefaults() *AffinityAntiAffinity {
	this := AffinityAntiAffinity{}
	return &this
}

// GetAffinityEAS returns the AffinityEAS field value if set, zero value otherwise.
func (o *AffinityAntiAffinity) GetAffinityEAS() []string {
	if o == nil || IsNil(o.AffinityEAS) {
		var ret []string
		return ret
	}
	return o.AffinityEAS
}

// GetAffinityEASOk returns a tuple with the AffinityEAS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinityAntiAffinity) GetAffinityEASOk() ([]string, bool) {
	if o == nil || IsNil(o.AffinityEAS) {
		return nil, false
	}
	return o.AffinityEAS, true
}

// HasAffinityEAS returns a boolean if a field has been set.
func (o *AffinityAntiAffinity) HasAffinityEAS() bool {
	if o != nil && !IsNil(o.AffinityEAS) {
		return true
	}

	return false
}

// SetAffinityEAS gets a reference to the given []string and assigns it to the AffinityEAS field.
func (o *AffinityAntiAffinity) SetAffinityEAS(v []string) {
	o.AffinityEAS = v
}

// GetAntiAffinityEAS returns the AntiAffinityEAS field value if set, zero value otherwise.
func (o *AffinityAntiAffinity) GetAntiAffinityEAS() string {
	if o == nil || IsNil(o.AntiAffinityEAS) {
		var ret string
		return ret
	}
	return *o.AntiAffinityEAS
}

// GetAntiAffinityEASOk returns a tuple with the AntiAffinityEAS field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AffinityAntiAffinity) GetAntiAffinityEASOk() (*string, bool) {
	if o == nil || IsNil(o.AntiAffinityEAS) {
		return nil, false
	}
	return o.AntiAffinityEAS, true
}

// HasAntiAffinityEAS returns a boolean if a field has been set.
func (o *AffinityAntiAffinity) HasAntiAffinityEAS() bool {
	if o != nil && !IsNil(o.AntiAffinityEAS) {
		return true
	}

	return false
}

// SetAntiAffinityEAS gets a reference to the given string and assigns it to the AntiAffinityEAS field.
func (o *AffinityAntiAffinity) SetAntiAffinityEAS(v string) {
	o.AntiAffinityEAS = &v
}

func (o AffinityAntiAffinity) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AffinityAntiAffinity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AffinityEAS) {
		toSerialize["affinityEAS"] = o.AffinityEAS
	}
	if !IsNil(o.AntiAffinityEAS) {
		toSerialize["antiAffinityEAS"] = o.AntiAffinityEAS
	}
	return toSerialize, nil
}

type NullableAffinityAntiAffinity struct {
	value *AffinityAntiAffinity
	isSet bool
}

func (v NullableAffinityAntiAffinity) Get() *AffinityAntiAffinity {
	return v.value
}

func (v *NullableAffinityAntiAffinity) Set(val *AffinityAntiAffinity) {
	v.value = val
	v.isSet = true
}

func (v NullableAffinityAntiAffinity) IsSet() bool {
	return v.isSet
}

func (v *NullableAffinityAntiAffinity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAffinityAntiAffinity(val *AffinityAntiAffinity) *NullableAffinityAntiAffinity {
	return &NullableAffinityAntiAffinity{value: val, isSet: true}
}

func (v NullableAffinityAntiAffinity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAffinityAntiAffinity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


