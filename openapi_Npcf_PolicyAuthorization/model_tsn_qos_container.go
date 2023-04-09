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

// checks if the TsnQosContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TsnQosContainer{}

// TsnQosContainer Indicates TSC Traffic QoS.
type TsnQosContainer struct {
	// Unsigned integer indicating Maximum Data Burst Volume (see clauses 5.7.3.7 and 5.7.4 of 3GPP TS 23.501), expressed in Bytes.
	MaxTscBurstSize *int32 `json:"maxTscBurstSize,omitempty"`
	// Unsigned integer indicating Packet Delay Budget (see clauses 5.7.3.4 and 5.7.4 of 3GPP TS 23.501), expressed in milliseconds.
	TscPackDelay *int32 `json:"tscPackDelay,omitempty"`
	// Represents the priority level of TSC Flows.
	TscPrioLevel *int32 `json:"tscPrioLevel,omitempty"`
}

// NewTsnQosContainer instantiates a new TsnQosContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTsnQosContainer() *TsnQosContainer {
	this := TsnQosContainer{}
	return &this
}

// NewTsnQosContainerWithDefaults instantiates a new TsnQosContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTsnQosContainerWithDefaults() *TsnQosContainer {
	this := TsnQosContainer{}
	return &this
}

// GetMaxTscBurstSize returns the MaxTscBurstSize field value if set, zero value otherwise.
func (o *TsnQosContainer) GetMaxTscBurstSize() int32 {
	if o == nil || IsNil(o.MaxTscBurstSize) {
		var ret int32
		return ret
	}
	return *o.MaxTscBurstSize
}

// GetMaxTscBurstSizeOk returns a tuple with the MaxTscBurstSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnQosContainer) GetMaxTscBurstSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxTscBurstSize) {
		return nil, false
	}
	return o.MaxTscBurstSize, true
}

// HasMaxTscBurstSize returns a boolean if a field has been set.
func (o *TsnQosContainer) HasMaxTscBurstSize() bool {
	if o != nil && !IsNil(o.MaxTscBurstSize) {
		return true
	}

	return false
}

// SetMaxTscBurstSize gets a reference to the given int32 and assigns it to the MaxTscBurstSize field.
func (o *TsnQosContainer) SetMaxTscBurstSize(v int32) {
	o.MaxTscBurstSize = &v
}

// GetTscPackDelay returns the TscPackDelay field value if set, zero value otherwise.
func (o *TsnQosContainer) GetTscPackDelay() int32 {
	if o == nil || IsNil(o.TscPackDelay) {
		var ret int32
		return ret
	}
	return *o.TscPackDelay
}

// GetTscPackDelayOk returns a tuple with the TscPackDelay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnQosContainer) GetTscPackDelayOk() (*int32, bool) {
	if o == nil || IsNil(o.TscPackDelay) {
		return nil, false
	}
	return o.TscPackDelay, true
}

// HasTscPackDelay returns a boolean if a field has been set.
func (o *TsnQosContainer) HasTscPackDelay() bool {
	if o != nil && !IsNil(o.TscPackDelay) {
		return true
	}

	return false
}

// SetTscPackDelay gets a reference to the given int32 and assigns it to the TscPackDelay field.
func (o *TsnQosContainer) SetTscPackDelay(v int32) {
	o.TscPackDelay = &v
}

// GetTscPrioLevel returns the TscPrioLevel field value if set, zero value otherwise.
func (o *TsnQosContainer) GetTscPrioLevel() int32 {
	if o == nil || IsNil(o.TscPrioLevel) {
		var ret int32
		return ret
	}
	return *o.TscPrioLevel
}

// GetTscPrioLevelOk returns a tuple with the TscPrioLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TsnQosContainer) GetTscPrioLevelOk() (*int32, bool) {
	if o == nil || IsNil(o.TscPrioLevel) {
		return nil, false
	}
	return o.TscPrioLevel, true
}

// HasTscPrioLevel returns a boolean if a field has been set.
func (o *TsnQosContainer) HasTscPrioLevel() bool {
	if o != nil && !IsNil(o.TscPrioLevel) {
		return true
	}

	return false
}

// SetTscPrioLevel gets a reference to the given int32 and assigns it to the TscPrioLevel field.
func (o *TsnQosContainer) SetTscPrioLevel(v int32) {
	o.TscPrioLevel = &v
}

func (o TsnQosContainer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TsnQosContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxTscBurstSize) {
		toSerialize["maxTscBurstSize"] = o.MaxTscBurstSize
	}
	if !IsNil(o.TscPackDelay) {
		toSerialize["tscPackDelay"] = o.TscPackDelay
	}
	if !IsNil(o.TscPrioLevel) {
		toSerialize["tscPrioLevel"] = o.TscPrioLevel
	}
	return toSerialize, nil
}

type NullableTsnQosContainer struct {
	value *TsnQosContainer
	isSet bool
}

func (v NullableTsnQosContainer) Get() *TsnQosContainer {
	return v.value
}

func (v *NullableTsnQosContainer) Set(val *TsnQosContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableTsnQosContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableTsnQosContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTsnQosContainer(val *TsnQosContainer) *NullableTsnQosContainer {
	return &NullableTsnQosContainer{value: val, isSet: true}
}

func (v NullableTsnQosContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTsnQosContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
