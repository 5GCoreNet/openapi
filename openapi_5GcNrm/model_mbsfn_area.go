/*
3GPP 5GC NRM

OAS 3.0.1 specification of the 5GC NRM © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC). All rights reserved.

API version: 18.2.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_5GcNrm

import (
	"encoding/json"
)

// checks if the MbsfnArea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsfnArea{}

// MbsfnArea struct for MbsfnArea
type MbsfnArea struct {
	MbsfnAreaId *int32 `json:"mbsfnAreaId,omitempty"`
	Earfcn      *int32 `json:"earfcn,omitempty"`
}

// NewMbsfnArea instantiates a new MbsfnArea object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMbsfnArea() *MbsfnArea {
	this := MbsfnArea{}
	return &this
}

// NewMbsfnAreaWithDefaults instantiates a new MbsfnArea object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMbsfnAreaWithDefaults() *MbsfnArea {
	this := MbsfnArea{}
	return &this
}

// GetMbsfnAreaId returns the MbsfnAreaId field value if set, zero value otherwise.
func (o *MbsfnArea) GetMbsfnAreaId() int32 {
	if o == nil || IsNil(o.MbsfnAreaId) {
		var ret int32
		return ret
	}
	return *o.MbsfnAreaId
}

// GetMbsfnAreaIdOk returns a tuple with the MbsfnAreaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsfnArea) GetMbsfnAreaIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MbsfnAreaId) {
		return nil, false
	}
	return o.MbsfnAreaId, true
}

// HasMbsfnAreaId returns a boolean if a field has been set.
func (o *MbsfnArea) HasMbsfnAreaId() bool {
	if o != nil && !IsNil(o.MbsfnAreaId) {
		return true
	}

	return false
}

// SetMbsfnAreaId gets a reference to the given int32 and assigns it to the MbsfnAreaId field.
func (o *MbsfnArea) SetMbsfnAreaId(v int32) {
	o.MbsfnAreaId = &v
}

// GetEarfcn returns the Earfcn field value if set, zero value otherwise.
func (o *MbsfnArea) GetEarfcn() int32 {
	if o == nil || IsNil(o.Earfcn) {
		var ret int32
		return ret
	}
	return *o.Earfcn
}

// GetEarfcnOk returns a tuple with the Earfcn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsfnArea) GetEarfcnOk() (*int32, bool) {
	if o == nil || IsNil(o.Earfcn) {
		return nil, false
	}
	return o.Earfcn, true
}

// HasEarfcn returns a boolean if a field has been set.
func (o *MbsfnArea) HasEarfcn() bool {
	if o != nil && !IsNil(o.Earfcn) {
		return true
	}

	return false
}

// SetEarfcn gets a reference to the given int32 and assigns it to the Earfcn field.
func (o *MbsfnArea) SetEarfcn(v int32) {
	o.Earfcn = &v
}

func (o MbsfnArea) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MbsfnArea) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MbsfnAreaId) {
		toSerialize["mbsfnAreaId"] = o.MbsfnAreaId
	}
	if !IsNil(o.Earfcn) {
		toSerialize["earfcn"] = o.Earfcn
	}
	return toSerialize, nil
}

type NullableMbsfnArea struct {
	value *MbsfnArea
	isSet bool
}

func (v NullableMbsfnArea) Get() *MbsfnArea {
	return v.value
}

func (v *NullableMbsfnArea) Set(val *MbsfnArea) {
	v.value = val
	v.isSet = true
}

func (v NullableMbsfnArea) IsSet() bool {
	return v.isSet
}

func (v *NullableMbsfnArea) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMbsfnArea(val *MbsfnArea) *NullableMbsfnArea {
	return &NullableMbsfnArea{value: val, isSet: true}
}

func (v NullableMbsfnArea) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMbsfnArea) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
