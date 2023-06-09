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

// checks if the AssuranceScope type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssuranceScope{}

// AssuranceScope struct for AssuranceScope
type AssuranceScope struct {
	TaiList []Tai `json:"taiList,omitempty"`
}

// NewAssuranceScope instantiates a new AssuranceScope object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssuranceScope() *AssuranceScope {
	this := AssuranceScope{}
	return &this
}

// NewAssuranceScopeWithDefaults instantiates a new AssuranceScope object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssuranceScopeWithDefaults() *AssuranceScope {
	this := AssuranceScope{}
	return &this
}

// GetTaiList returns the TaiList field value if set, zero value otherwise.
func (o *AssuranceScope) GetTaiList() []Tai {
	if o == nil || IsNil(o.TaiList) {
		var ret []Tai
		return ret
	}
	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssuranceScope) GetTaiListOk() ([]Tai, bool) {
	if o == nil || IsNil(o.TaiList) {
		return nil, false
	}
	return o.TaiList, true
}

// HasTaiList returns a boolean if a field has been set.
func (o *AssuranceScope) HasTaiList() bool {
	if o != nil && !IsNil(o.TaiList) {
		return true
	}

	return false
}

// SetTaiList gets a reference to the given []Tai and assigns it to the TaiList field.
func (o *AssuranceScope) SetTaiList(v []Tai) {
	o.TaiList = v
}

func (o AssuranceScope) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssuranceScope) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TaiList) {
		toSerialize["taiList"] = o.TaiList
	}
	return toSerialize, nil
}

type NullableAssuranceScope struct {
	value *AssuranceScope
	isSet bool
}

func (v NullableAssuranceScope) Get() *AssuranceScope {
	return v.value
}

func (v *NullableAssuranceScope) Set(val *AssuranceScope) {
	v.value = val
	v.isSet = true
}

func (v NullableAssuranceScope) IsSet() bool {
	return v.isSet
}

func (v *NullableAssuranceScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssuranceScope(val *AssuranceScope) *NullableAssuranceScope {
	return &NullableAssuranceScope{value: val, isSet: true}
}

func (v NullableAssuranceScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssuranceScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
