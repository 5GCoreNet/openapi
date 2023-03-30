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

// checks if the UeInAreaFilter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UeInAreaFilter{}

// UeInAreaFilter Additional filters for UE in Area Report event
type UeInAreaFilter struct {
	UeType *UeType `json:"ueType,omitempty"`
	AerialSrvDnnInd *bool `json:"aerialSrvDnnInd,omitempty"`
}

// NewUeInAreaFilter instantiates a new UeInAreaFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUeInAreaFilter() *UeInAreaFilter {
	this := UeInAreaFilter{}
	var aerialSrvDnnInd bool = false
	this.AerialSrvDnnInd = &aerialSrvDnnInd
	return &this
}

// NewUeInAreaFilterWithDefaults instantiates a new UeInAreaFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUeInAreaFilterWithDefaults() *UeInAreaFilter {
	this := UeInAreaFilter{}
	var aerialSrvDnnInd bool = false
	this.AerialSrvDnnInd = &aerialSrvDnnInd
	return &this
}

// GetUeType returns the UeType field value if set, zero value otherwise.
func (o *UeInAreaFilter) GetUeType() UeType {
	if o == nil || IsNil(o.UeType) {
		var ret UeType
		return ret
	}
	return *o.UeType
}

// GetUeTypeOk returns a tuple with the UeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeInAreaFilter) GetUeTypeOk() (*UeType, bool) {
	if o == nil || IsNil(o.UeType) {
		return nil, false
	}
	return o.UeType, true
}

// HasUeType returns a boolean if a field has been set.
func (o *UeInAreaFilter) HasUeType() bool {
	if o != nil && !IsNil(o.UeType) {
		return true
	}

	return false
}

// SetUeType gets a reference to the given UeType and assigns it to the UeType field.
func (o *UeInAreaFilter) SetUeType(v UeType) {
	o.UeType = &v
}

// GetAerialSrvDnnInd returns the AerialSrvDnnInd field value if set, zero value otherwise.
func (o *UeInAreaFilter) GetAerialSrvDnnInd() bool {
	if o == nil || IsNil(o.AerialSrvDnnInd) {
		var ret bool
		return ret
	}
	return *o.AerialSrvDnnInd
}

// GetAerialSrvDnnIndOk returns a tuple with the AerialSrvDnnInd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UeInAreaFilter) GetAerialSrvDnnIndOk() (*bool, bool) {
	if o == nil || IsNil(o.AerialSrvDnnInd) {
		return nil, false
	}
	return o.AerialSrvDnnInd, true
}

// HasAerialSrvDnnInd returns a boolean if a field has been set.
func (o *UeInAreaFilter) HasAerialSrvDnnInd() bool {
	if o != nil && !IsNil(o.AerialSrvDnnInd) {
		return true
	}

	return false
}

// SetAerialSrvDnnInd gets a reference to the given bool and assigns it to the AerialSrvDnnInd field.
func (o *UeInAreaFilter) SetAerialSrvDnnInd(v bool) {
	o.AerialSrvDnnInd = &v
}

func (o UeInAreaFilter) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UeInAreaFilter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.UeType) {
		toSerialize["ueType"] = o.UeType
	}
	if !IsNil(o.AerialSrvDnnInd) {
		toSerialize["aerialSrvDnnInd"] = o.AerialSrvDnnInd
	}
	return toSerialize, nil
}

type NullableUeInAreaFilter struct {
	value *UeInAreaFilter
	isSet bool
}

func (v NullableUeInAreaFilter) Get() *UeInAreaFilter {
	return v.value
}

func (v *NullableUeInAreaFilter) Set(val *UeInAreaFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableUeInAreaFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableUeInAreaFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUeInAreaFilter(val *UeInAreaFilter) *NullableUeInAreaFilter {
	return &NullableUeInAreaFilter{value: val, isSet: true}
}

func (v NullableUeInAreaFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUeInAreaFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


