/*
Nudm_SDM

Nudm Subscriber Data Management Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 2.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_SDM

import (
	"encoding/json"
)

// checks if the MbsfnArea type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MbsfnArea{}

// MbsfnArea Contains an MBSFN area information.
type MbsfnArea struct {
	// This IE shall contain the MBSFN Area ID.
	MbsfnAreaId *int32 `json:"mbsfnAreaId,omitempty"`
	// When present, this IE shall contain the Carrier Frequency (EARFCN).
	CarrierFrequency *int32 `json:"carrierFrequency,omitempty"`
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

// GetCarrierFrequency returns the CarrierFrequency field value if set, zero value otherwise.
func (o *MbsfnArea) GetCarrierFrequency() int32 {
	if o == nil || IsNil(o.CarrierFrequency) {
		var ret int32
		return ret
	}
	return *o.CarrierFrequency
}

// GetCarrierFrequencyOk returns a tuple with the CarrierFrequency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MbsfnArea) GetCarrierFrequencyOk() (*int32, bool) {
	if o == nil || IsNil(o.CarrierFrequency) {
		return nil, false
	}
	return o.CarrierFrequency, true
}

// HasCarrierFrequency returns a boolean if a field has been set.
func (o *MbsfnArea) HasCarrierFrequency() bool {
	if o != nil && !IsNil(o.CarrierFrequency) {
		return true
	}

	return false
}

// SetCarrierFrequency gets a reference to the given int32 and assigns it to the CarrierFrequency field.
func (o *MbsfnArea) SetCarrierFrequency(v int32) {
	o.CarrierFrequency = &v
}

func (o MbsfnArea) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
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
	if !IsNil(o.CarrierFrequency) {
		toSerialize["carrierFrequency"] = o.CarrierFrequency
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


