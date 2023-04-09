/*
Nudm_PP

Nudm Parameter Provision Service.   © 2022, 3GPP Organizational Partners (ARIB, ATIS, CCSA, ETSI, TSDSI, TTA, TTC).   All rights reserved.

API version: 1.3.0-alpha.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi_Nudm_PP

import (
	"encoding/json"
)

// checks if the CommunicationCharacteristicsAF type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommunicationCharacteristicsAF{}

// CommunicationCharacteristicsAF struct for CommunicationCharacteristicsAF
type CommunicationCharacteristicsAF struct {
	PpDlPacketCount NullableInt32 `json:"ppDlPacketCount,omitempty"`
	// indicating a time in seconds.
	MaximumResponseTime *int32 `json:"maximumResponseTime,omitempty"`
	// indicating a time in seconds.
	MaximumLatency *int32 `json:"maximumLatency,omitempty"`
}

// NewCommunicationCharacteristicsAF instantiates a new CommunicationCharacteristicsAF object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommunicationCharacteristicsAF() *CommunicationCharacteristicsAF {
	this := CommunicationCharacteristicsAF{}
	return &this
}

// NewCommunicationCharacteristicsAFWithDefaults instantiates a new CommunicationCharacteristicsAF object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommunicationCharacteristicsAFWithDefaults() *CommunicationCharacteristicsAF {
	this := CommunicationCharacteristicsAF{}
	return &this
}

// GetPpDlPacketCount returns the PpDlPacketCount field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CommunicationCharacteristicsAF) GetPpDlPacketCount() int32 {
	if o == nil || IsNil(o.PpDlPacketCount.Get()) {
		var ret int32
		return ret
	}
	return *o.PpDlPacketCount.Get()
}

// GetPpDlPacketCountOk returns a tuple with the PpDlPacketCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CommunicationCharacteristicsAF) GetPpDlPacketCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.PpDlPacketCount.Get(), o.PpDlPacketCount.IsSet()
}

// HasPpDlPacketCount returns a boolean if a field has been set.
func (o *CommunicationCharacteristicsAF) HasPpDlPacketCount() bool {
	if o != nil && o.PpDlPacketCount.IsSet() {
		return true
	}

	return false
}

// SetPpDlPacketCount gets a reference to the given NullableInt32 and assigns it to the PpDlPacketCount field.
func (o *CommunicationCharacteristicsAF) SetPpDlPacketCount(v int32) {
	o.PpDlPacketCount.Set(&v)
}

// SetPpDlPacketCountNil sets the value for PpDlPacketCount to be an explicit nil
func (o *CommunicationCharacteristicsAF) SetPpDlPacketCountNil() {
	o.PpDlPacketCount.Set(nil)
}

// UnsetPpDlPacketCount ensures that no value is present for PpDlPacketCount, not even an explicit nil
func (o *CommunicationCharacteristicsAF) UnsetPpDlPacketCount() {
	o.PpDlPacketCount.Unset()
}

// GetMaximumResponseTime returns the MaximumResponseTime field value if set, zero value otherwise.
func (o *CommunicationCharacteristicsAF) GetMaximumResponseTime() int32 {
	if o == nil || IsNil(o.MaximumResponseTime) {
		var ret int32
		return ret
	}
	return *o.MaximumResponseTime
}

// GetMaximumResponseTimeOk returns a tuple with the MaximumResponseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationCharacteristicsAF) GetMaximumResponseTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumResponseTime) {
		return nil, false
	}
	return o.MaximumResponseTime, true
}

// HasMaximumResponseTime returns a boolean if a field has been set.
func (o *CommunicationCharacteristicsAF) HasMaximumResponseTime() bool {
	if o != nil && !IsNil(o.MaximumResponseTime) {
		return true
	}

	return false
}

// SetMaximumResponseTime gets a reference to the given int32 and assigns it to the MaximumResponseTime field.
func (o *CommunicationCharacteristicsAF) SetMaximumResponseTime(v int32) {
	o.MaximumResponseTime = &v
}

// GetMaximumLatency returns the MaximumLatency field value if set, zero value otherwise.
func (o *CommunicationCharacteristicsAF) GetMaximumLatency() int32 {
	if o == nil || IsNil(o.MaximumLatency) {
		var ret int32
		return ret
	}
	return *o.MaximumLatency
}

// GetMaximumLatencyOk returns a tuple with the MaximumLatency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommunicationCharacteristicsAF) GetMaximumLatencyOk() (*int32, bool) {
	if o == nil || IsNil(o.MaximumLatency) {
		return nil, false
	}
	return o.MaximumLatency, true
}

// HasMaximumLatency returns a boolean if a field has been set.
func (o *CommunicationCharacteristicsAF) HasMaximumLatency() bool {
	if o != nil && !IsNil(o.MaximumLatency) {
		return true
	}

	return false
}

// SetMaximumLatency gets a reference to the given int32 and assigns it to the MaximumLatency field.
func (o *CommunicationCharacteristicsAF) SetMaximumLatency(v int32) {
	o.MaximumLatency = &v
}

func (o CommunicationCharacteristicsAF) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommunicationCharacteristicsAF) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.PpDlPacketCount.IsSet() {
		toSerialize["ppDlPacketCount"] = o.PpDlPacketCount.Get()
	}
	if !IsNil(o.MaximumResponseTime) {
		toSerialize["maximumResponseTime"] = o.MaximumResponseTime
	}
	if !IsNil(o.MaximumLatency) {
		toSerialize["maximumLatency"] = o.MaximumLatency
	}
	return toSerialize, nil
}

type NullableCommunicationCharacteristicsAF struct {
	value *CommunicationCharacteristicsAF
	isSet bool
}

func (v NullableCommunicationCharacteristicsAF) Get() *CommunicationCharacteristicsAF {
	return v.value
}

func (v *NullableCommunicationCharacteristicsAF) Set(val *CommunicationCharacteristicsAF) {
	v.value = val
	v.isSet = true
}

func (v NullableCommunicationCharacteristicsAF) IsSet() bool {
	return v.isSet
}

func (v *NullableCommunicationCharacteristicsAF) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommunicationCharacteristicsAF(val *CommunicationCharacteristicsAF) *NullableCommunicationCharacteristicsAF {
	return &NullableCommunicationCharacteristicsAF{value: val, isSet: true}
}

func (v NullableCommunicationCharacteristicsAF) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommunicationCharacteristicsAF) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
