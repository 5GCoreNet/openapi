/*
Npcf_PolicyAuthorization Service API

PCF Policy Authorization Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved. 

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Npcf_PolicyAuthorization

import (
	"encoding/json"
)

// checks if the TsnQosContainerRm type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TsnQosContainerRm{}

// TsnQosContainerRm Indicates removable TSC Traffic QoS.
type TsnQosContainerRm struct {
	// This data type is defined in the same way as the 'ExtMaxDataBurstVol' data type, but with the OpenAPI 'nullable: true' property. 
	MaxTscBurstSize NullableInt32 `json:"maxTscBurstSize,omitempty"`
	// This data type is defined in the same way as the 'PacketDelBudget' data type, but with the OpenAPI 'nullable: true' property. 
	TscPackDelay NullableInt32 `json:"tscPackDelay,omitempty"`
	// This data type is defined in the same way as the TscPriorityLevel data type, but with the OpenAPI nullable property set to true. 
	TscPrioLevel NullableInt32 `json:"tscPrioLevel,omitempty"`
}

// NewTsnQosContainerRm instantiates a new TsnQosContainerRm object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTsnQosContainerRm() *TsnQosContainerRm {
	this := TsnQosContainerRm{}
	return &this
}

// NewTsnQosContainerRmWithDefaults instantiates a new TsnQosContainerRm object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTsnQosContainerRmWithDefaults() *TsnQosContainerRm {
	this := TsnQosContainerRm{}
	return &this
}

// GetMaxTscBurstSize returns the MaxTscBurstSize field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TsnQosContainerRm) GetMaxTscBurstSize() int32 {
	if o == nil || IsNil(o.MaxTscBurstSize.Get()) {
		var ret int32
		return ret
	}
	return *o.MaxTscBurstSize.Get()
}

// GetMaxTscBurstSizeOk returns a tuple with the MaxTscBurstSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TsnQosContainerRm) GetMaxTscBurstSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxTscBurstSize.Get(), o.MaxTscBurstSize.IsSet()
}

// HasMaxTscBurstSize returns a boolean if a field has been set.
func (o *TsnQosContainerRm) HasMaxTscBurstSize() bool {
	if o != nil && o.MaxTscBurstSize.IsSet() {
		return true
	}

	return false
}

// SetMaxTscBurstSize gets a reference to the given NullableInt32 and assigns it to the MaxTscBurstSize field.
func (o *TsnQosContainerRm) SetMaxTscBurstSize(v int32) {
	o.MaxTscBurstSize.Set(&v)
}
// SetMaxTscBurstSizeNil sets the value for MaxTscBurstSize to be an explicit nil
func (o *TsnQosContainerRm) SetMaxTscBurstSizeNil() {
	o.MaxTscBurstSize.Set(nil)
}

// UnsetMaxTscBurstSize ensures that no value is present for MaxTscBurstSize, not even an explicit nil
func (o *TsnQosContainerRm) UnsetMaxTscBurstSize() {
	o.MaxTscBurstSize.Unset()
}

// GetTscPackDelay returns the TscPackDelay field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TsnQosContainerRm) GetTscPackDelay() int32 {
	if o == nil || IsNil(o.TscPackDelay.Get()) {
		var ret int32
		return ret
	}
	return *o.TscPackDelay.Get()
}

// GetTscPackDelayOk returns a tuple with the TscPackDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TsnQosContainerRm) GetTscPackDelayOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TscPackDelay.Get(), o.TscPackDelay.IsSet()
}

// HasTscPackDelay returns a boolean if a field has been set.
func (o *TsnQosContainerRm) HasTscPackDelay() bool {
	if o != nil && o.TscPackDelay.IsSet() {
		return true
	}

	return false
}

// SetTscPackDelay gets a reference to the given NullableInt32 and assigns it to the TscPackDelay field.
func (o *TsnQosContainerRm) SetTscPackDelay(v int32) {
	o.TscPackDelay.Set(&v)
}
// SetTscPackDelayNil sets the value for TscPackDelay to be an explicit nil
func (o *TsnQosContainerRm) SetTscPackDelayNil() {
	o.TscPackDelay.Set(nil)
}

// UnsetTscPackDelay ensures that no value is present for TscPackDelay, not even an explicit nil
func (o *TsnQosContainerRm) UnsetTscPackDelay() {
	o.TscPackDelay.Unset()
}

// GetTscPrioLevel returns the TscPrioLevel field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TsnQosContainerRm) GetTscPrioLevel() int32 {
	if o == nil || IsNil(o.TscPrioLevel.Get()) {
		var ret int32
		return ret
	}
	return *o.TscPrioLevel.Get()
}

// GetTscPrioLevelOk returns a tuple with the TscPrioLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TsnQosContainerRm) GetTscPrioLevelOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.TscPrioLevel.Get(), o.TscPrioLevel.IsSet()
}

// HasTscPrioLevel returns a boolean if a field has been set.
func (o *TsnQosContainerRm) HasTscPrioLevel() bool {
	if o != nil && o.TscPrioLevel.IsSet() {
		return true
	}

	return false
}

// SetTscPrioLevel gets a reference to the given NullableInt32 and assigns it to the TscPrioLevel field.
func (o *TsnQosContainerRm) SetTscPrioLevel(v int32) {
	o.TscPrioLevel.Set(&v)
}
// SetTscPrioLevelNil sets the value for TscPrioLevel to be an explicit nil
func (o *TsnQosContainerRm) SetTscPrioLevelNil() {
	o.TscPrioLevel.Set(nil)
}

// UnsetTscPrioLevel ensures that no value is present for TscPrioLevel, not even an explicit nil
func (o *TsnQosContainerRm) UnsetTscPrioLevel() {
	o.TscPrioLevel.Unset()
}

func (o TsnQosContainerRm) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TsnQosContainerRm) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.MaxTscBurstSize.IsSet() {
		toSerialize["maxTscBurstSize"] = o.MaxTscBurstSize.Get()
	}
	if o.TscPackDelay.IsSet() {
		toSerialize["tscPackDelay"] = o.TscPackDelay.Get()
	}
	if o.TscPrioLevel.IsSet() {
		toSerialize["tscPrioLevel"] = o.TscPrioLevel.Get()
	}
	return toSerialize, nil
}

type NullableTsnQosContainerRm struct {
	value *TsnQosContainerRm
	isSet bool
}

func (v NullableTsnQosContainerRm) Get() *TsnQosContainerRm {
	return v.value
}

func (v *NullableTsnQosContainerRm) Set(val *TsnQosContainerRm) {
	v.value = val
	v.isSet = true
}

func (v NullableTsnQosContainerRm) IsSet() bool {
	return v.isSet
}

func (v *NullableTsnQosContainerRm) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTsnQosContainerRm(val *TsnQosContainerRm) *NullableTsnQosContainerRm {
	return &NullableTsnQosContainerRm{value: val, isSet: true}
}

func (v NullableTsnQosContainerRm) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTsnQosContainerRm) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

