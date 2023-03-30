/*
Nudr_DataRepository API OpenAPI file

Unified Data Repository Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudr_DR

import (
	"encoding/json"
)

// checks if the UsageMonDataScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UsageMonDataScope{}

// UsageMonDataScope Contains a SNSSAI and DNN combinations to which the UsageMonData instance belongs to. 
type UsageMonDataScope struct {
	Snssai Snssai `json:"snssai"`
	Dnn []string `json:"dnn,omitempty"`
}

// NewUsageMonDataScope instantiates a new UsageMonDataScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUsageMonDataScope(snssai Snssai) *UsageMonDataScope {
	this := UsageMonDataScope{}
	this.Snssai = snssai
	return &this
}

// NewUsageMonDataScopeWithDefaults instantiates a new UsageMonDataScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUsageMonDataScopeWithDefaults() *UsageMonDataScope {
	this := UsageMonDataScope{}
	return &this
}

// GetSnssai returns the Snssai field value
func (o *UsageMonDataScope) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *UsageMonDataScope) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *UsageMonDataScope) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetDnn returns the Dnn field value if set, zero value otherwise.
func (o *UsageMonDataScope) GetDnn() []string {
	if o == nil || IsNil(o.Dnn) {
		var ret []string
		return ret
	}
	return o.Dnn
}

// GetDnnOk returns a tuple with the Dnn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UsageMonDataScope) GetDnnOk() ([]string, bool) {
	if o == nil || IsNil(o.Dnn) {
		return nil, false
	}
	return o.Dnn, true
}

// HasDnn returns a boolean if a field has been set.
func (o *UsageMonDataScope) HasDnn() bool {
	if o != nil && !IsNil(o.Dnn) {
		return true
	}

	return false
}

// SetDnn gets a reference to the given []string and assigns it to the Dnn field.
func (o *UsageMonDataScope) SetDnn(v []string) {
	o.Dnn = v
}

func (o UsageMonDataScope) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UsageMonDataScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snssai"] = o.Snssai
	if !IsNil(o.Dnn) {
		toSerialize["dnn"] = o.Dnn
	}
	return toSerialize, nil
}

type NullableUsageMonDataScope struct {
	value *UsageMonDataScope
	isSet bool
}

func (v NullableUsageMonDataScope) Get() *UsageMonDataScope {
	return v.value
}

func (v *NullableUsageMonDataScope) Set(val *UsageMonDataScope) {
	v.value = val
	v.isSet = true
}

func (v NullableUsageMonDataScope) IsSet() bool {
	return v.isSet
}

func (v *NullableUsageMonDataScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUsageMonDataScope(val *UsageMonDataScope) *NullableUsageMonDataScope {
	return &NullableUsageMonDataScope{value: val, isSet: true}
}

func (v NullableUsageMonDataScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUsageMonDataScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


