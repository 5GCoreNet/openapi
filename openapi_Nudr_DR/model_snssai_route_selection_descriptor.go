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

// checks if the SnssaiRouteSelectionDescriptor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SnssaiRouteSelectionDescriptor{}

// SnssaiRouteSelectionDescriptor Contains the route selector parameters (DNNs, PDU session types, SSC modes and ATSSS information) per SNSSAI 
type SnssaiRouteSelectionDescriptor struct {
	Snssai Snssai `json:"snssai"`
	DnnRouteSelDescs []DnnRouteSelectionDescriptor `json:"dnnRouteSelDescs,omitempty"`
}

// NewSnssaiRouteSelectionDescriptor instantiates a new SnssaiRouteSelectionDescriptor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSnssaiRouteSelectionDescriptor(snssai Snssai) *SnssaiRouteSelectionDescriptor {
	this := SnssaiRouteSelectionDescriptor{}
	this.Snssai = snssai
	return &this
}

// NewSnssaiRouteSelectionDescriptorWithDefaults instantiates a new SnssaiRouteSelectionDescriptor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSnssaiRouteSelectionDescriptorWithDefaults() *SnssaiRouteSelectionDescriptor {
	this := SnssaiRouteSelectionDescriptor{}
	return &this
}

// GetSnssai returns the Snssai field value
func (o *SnssaiRouteSelectionDescriptor) GetSnssai() Snssai {
	if o == nil {
		var ret Snssai
		return ret
	}

	return o.Snssai
}

// GetSnssaiOk returns a tuple with the Snssai field value
// and a boolean to check if the value has been set.
func (o *SnssaiRouteSelectionDescriptor) GetSnssaiOk() (*Snssai, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Snssai, true
}

// SetSnssai sets field value
func (o *SnssaiRouteSelectionDescriptor) SetSnssai(v Snssai) {
	o.Snssai = v
}

// GetDnnRouteSelDescs returns the DnnRouteSelDescs field value if set, zero value otherwise.
func (o *SnssaiRouteSelectionDescriptor) GetDnnRouteSelDescs() []DnnRouteSelectionDescriptor {
	if o == nil || IsNil(o.DnnRouteSelDescs) {
		var ret []DnnRouteSelectionDescriptor
		return ret
	}
	return o.DnnRouteSelDescs
}

// GetDnnRouteSelDescsOk returns a tuple with the DnnRouteSelDescs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SnssaiRouteSelectionDescriptor) GetDnnRouteSelDescsOk() ([]DnnRouteSelectionDescriptor, bool) {
	if o == nil || IsNil(o.DnnRouteSelDescs) {
		return nil, false
	}
	return o.DnnRouteSelDescs, true
}

// HasDnnRouteSelDescs returns a boolean if a field has been set.
func (o *SnssaiRouteSelectionDescriptor) HasDnnRouteSelDescs() bool {
	if o != nil && !IsNil(o.DnnRouteSelDescs) {
		return true
	}

	return false
}

// SetDnnRouteSelDescs gets a reference to the given []DnnRouteSelectionDescriptor and assigns it to the DnnRouteSelDescs field.
func (o *SnssaiRouteSelectionDescriptor) SetDnnRouteSelDescs(v []DnnRouteSelectionDescriptor) {
	o.DnnRouteSelDescs = v
}

func (o SnssaiRouteSelectionDescriptor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SnssaiRouteSelectionDescriptor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["snssai"] = o.Snssai
	if !IsNil(o.DnnRouteSelDescs) {
		toSerialize["dnnRouteSelDescs"] = o.DnnRouteSelDescs
	}
	return toSerialize, nil
}

type NullableSnssaiRouteSelectionDescriptor struct {
	value *SnssaiRouteSelectionDescriptor
	isSet bool
}

func (v NullableSnssaiRouteSelectionDescriptor) Get() *SnssaiRouteSelectionDescriptor {
	return v.value
}

func (v *NullableSnssaiRouteSelectionDescriptor) Set(val *SnssaiRouteSelectionDescriptor) {
	v.value = val
	v.isSet = true
}

func (v NullableSnssaiRouteSelectionDescriptor) IsSet() bool {
	return v.isSet
}

func (v *NullableSnssaiRouteSelectionDescriptor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSnssaiRouteSelectionDescriptor(val *SnssaiRouteSelectionDescriptor) *NullableSnssaiRouteSelectionDescriptor {
	return &NullableSnssaiRouteSelectionDescriptor{value: val, isSet: true}
}

func (v NullableSnssaiRouteSelectionDescriptor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSnssaiRouteSelectionDescriptor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


