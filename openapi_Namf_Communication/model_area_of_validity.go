/*
Namf_Communication

AMF Communication Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Namf_Communication

import (
	"encoding/json"
)

// checks if the AreaOfValidity type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AreaOfValidity{}

// AreaOfValidity Area of validity information for N2 information transfer
type AreaOfValidity struct {
	TaiList []Tai `json:"taiList"`
}

// NewAreaOfValidity instantiates a new AreaOfValidity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAreaOfValidity(taiList []Tai) *AreaOfValidity {
	this := AreaOfValidity{}
	this.TaiList = taiList
	return &this
}

// NewAreaOfValidityWithDefaults instantiates a new AreaOfValidity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAreaOfValidityWithDefaults() *AreaOfValidity {
	this := AreaOfValidity{}
	return &this
}

// GetTaiList returns the TaiList field value
func (o *AreaOfValidity) GetTaiList() []Tai {
	if o == nil {
		var ret []Tai
		return ret
	}

	return o.TaiList
}

// GetTaiListOk returns a tuple with the TaiList field value
// and a boolean to check if the value has been set.
func (o *AreaOfValidity) GetTaiListOk() ([]Tai, bool) {
	if o == nil {
		return nil, false
	}
	return o.TaiList, true
}

// SetTaiList sets field value
func (o *AreaOfValidity) SetTaiList(v []Tai) {
	o.TaiList = v
}

func (o AreaOfValidity) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AreaOfValidity) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["taiList"] = o.TaiList
	return toSerialize, nil
}

type NullableAreaOfValidity struct {
	value *AreaOfValidity
	isSet bool
}

func (v NullableAreaOfValidity) Get() *AreaOfValidity {
	return v.value
}

func (v *NullableAreaOfValidity) Set(val *AreaOfValidity) {
	v.value = val
	v.isSet = true
}

func (v NullableAreaOfValidity) IsSet() bool {
	return v.isSet
}

func (v *NullableAreaOfValidity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAreaOfValidity(val *AreaOfValidity) *NullableAreaOfValidity {
	return &NullableAreaOfValidity{value: val, isSet: true}
}

func (v NullableAreaOfValidity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAreaOfValidity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
